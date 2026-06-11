-- +goose Up

-- signal_latest holds the latest state per (subject, kind, name, source) so
-- latest-value / last-seen queries are point lookups instead of fan-out reads
-- across every part of the ~97B-row signal table. ReplacingMergeTree keeps the
-- row with the greatest `timestamp` (the version column) per sorting key at
-- merge time; readers still use argMax/max at query time so results are exact
-- before merges complete.
--
-- kind = 0: latest raw row for the signal, regardless of value.
-- kind = 1: latest row whose value_location is not (0, 0) — serves the
--           "latest valid GPS fix" semantics used by signalsLatest location
--           fields. Only location-bearing rows are written with kind = 1.
--           value_number/value_string on kind = 1 rows are not meaningful.

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS signal_latest
(
    `subject` String COMMENT 'Subject of the signal, typically a vehicle DID.',
    `kind` UInt8 COMMENT '0 = latest raw row; 1 = latest non-(0,0)-location row.',
    `name` LowCardinality(String) COMMENT 'Name of the signal.',
    `source` LowCardinality(String) COMMENT 'Source of the signal.',
    `timestamp` DateTime64(6, 'UTC') COMMENT 'Timestamp of the captured row; also the ReplacingMergeTree version.',
    `value_number` Float64 COMMENT 'Numeric value of the captured row.',
    `value_string` String COMMENT 'String value of the captured row.',
    `value_location` Tuple(
        latitude Float64,
        longitude Float64,
        hdop Float64,
        heading Float64) COMMENT 'Location value of the captured row.'
)
ENGINE = ReplacingMergeTree(timestamp)
ORDER BY (subject, kind, name, source)
COMMENT 'Latest signal state per (subject, kind, name, source). Fed by signal_latest_raw_mv and signal_latest_loc_mv.';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE MATERIALIZED VIEW IF NOT EXISTS signal_latest_raw_mv
TO signal_latest
AS
SELECT
    subject,
    0 AS kind,
    name,
    source,
    timestamp,
    value_number,
    value_string,
    value_location
FROM signal;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE MATERIALIZED VIEW IF NOT EXISTS signal_latest_loc_mv
TO signal_latest
AS
SELECT
    subject,
    1 AS kind,
    name,
    source,
    timestamp,
    value_number,
    value_string,
    value_location
FROM signal
WHERE (tupleElement(value_location, 'latitude') != 0) OR (tupleElement(value_location, 'longitude') != 0);
-- +goose StatementEnd

-- Backfill is NOT part of the migration: it aggregates the full signal table
-- and must be run by an operator once, after this migration is applied and
-- before telemetry-api switches to signal_latest. Overlap between the MVs and
-- the backfill is collapsed by the ReplacingMergeTree version. Operator SQL is
-- documented in the implementation plan (Phase 2) and duplicated here:
--
-- NOTE: do not alias the aggregates to source-column names (e.g.
-- max(timestamp) AS timestamp): the alias shadows the column inside the other
-- aggregates and ClickHouse rejects the query with "aggregate function found
-- inside another aggregate function" (error 184). The explicit INSERT column
-- list maps the SELECT positionally instead.
--
--   INSERT INTO signal_latest (subject, kind, name, source, timestamp, value_number, value_string, value_location)
--   SELECT subject, 0, name, source,
--          max(timestamp) AS ts,
--          argMax(value_number, timestamp),
--          argMax(value_string, timestamp),
--          argMax(value_location, timestamp)
--   FROM signal
--   GROUP BY subject, name, source
--   SETTINGS max_execution_time = 0, max_memory_usage = 0,
--            max_bytes_before_external_group_by = 32000000000;
--
--   INSERT INTO signal_latest (subject, kind, name, source, timestamp, value_number, value_string, value_location)
--   SELECT subject, 1, name, source,
--          maxIf(timestamp, (tupleElement(value_location, 'latitude') != 0) OR (tupleElement(value_location, 'longitude') != 0)) AS ts,
--          0, '',
--          argMaxIf(value_location, timestamp, (tupleElement(value_location, 'latitude') != 0) OR (tupleElement(value_location, 'longitude') != 0))
--   FROM signal
--   GROUP BY subject, name, source
--   HAVING ts > toDateTime64(0, 6)
--   SETTINGS max_execution_time = 0, max_memory_usage = 0,
--            max_bytes_before_external_group_by = 32000000000;

-- +goose Down

-- +goose StatementBegin
DROP VIEW IF EXISTS signal_latest_loc_mv;
-- +goose StatementEnd

-- +goose StatementBegin
DROP VIEW IF EXISTS signal_latest_raw_mv;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS signal_latest;
-- +goose StatementEnd
