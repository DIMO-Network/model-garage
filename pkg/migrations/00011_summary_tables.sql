-- +goose Up

-- Per-name summaries (count, first seen, last seen) for the dataSummary API.
-- AggregatingMergeTree with SimpleAggregateFunction columns: the MV emits one
-- partially-aggregated row per insert block, merges combine them (sum/min/max),
-- and readers re-aggregate at query time, so results are exact regardless of
-- merge state.
--
-- Count semantics: counts are over raw inserted rows. The signal/event tables
-- are ReplacingMergeTree, so rows deduplicated there later are still counted
-- here. The previous count(*) reads (no FINAL) were equally approximate.

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS signal_summary
(
    `subject` String COMMENT 'Subject of the signal, typically a vehicle DID.',
    `name` LowCardinality(String) COMMENT 'Name of the signal.',
    `source` LowCardinality(String) COMMENT 'Source of the signal.',
    `count` SimpleAggregateFunction(sum, UInt64) COMMENT 'Number of rows inserted for this signal.',
    `first_seen` SimpleAggregateFunction(min, DateTime64(6, 'UTC')) COMMENT 'Earliest timestamp for this signal.',
    `last_seen` SimpleAggregateFunction(max, DateTime64(6, 'UTC')) COMMENT 'Latest timestamp for this signal.'
)
ENGINE = AggregatingMergeTree
ORDER BY (subject, name, source)
COMMENT 'Per-signal summary per (subject, name, source). Fed by signal_summary_mv.';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE MATERIALIZED VIEW IF NOT EXISTS signal_summary_mv
TO signal_summary
AS
SELECT
    subject,
    name,
    source,
    toUInt64(count()) AS count,
    min(timestamp) AS first_seen,
    max(timestamp) AS last_seen
FROM signal
GROUP BY subject, name, source;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS event_summary
(
    `subject` String COMMENT 'Subject of the event, typically a vehicle DID.',
    `name` LowCardinality(String) COMMENT 'Name of the event.',
    `source` LowCardinality(String) COMMENT 'Source of the event.',
    `count` SimpleAggregateFunction(sum, UInt64) COMMENT 'Number of rows inserted for this event name.',
    `first_seen` SimpleAggregateFunction(min, DateTime64(6, 'UTC')) COMMENT 'Earliest timestamp for this event name.',
    `last_seen` SimpleAggregateFunction(max, DateTime64(6, 'UTC')) COMMENT 'Latest timestamp for this event name.'
)
ENGINE = AggregatingMergeTree
ORDER BY (subject, name, source)
COMMENT 'Per-event summary per (subject, name, source). Fed by event_summary_mv.';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE MATERIALIZED VIEW IF NOT EXISTS event_summary_mv
TO event_summary
AS
SELECT
    subject,
    name,
    source,
    toUInt64(count()) AS count,
    min(timestamp) AS first_seen,
    max(timestamp) AS last_seen
FROM event
GROUP BY subject, name, source;
-- +goose StatementEnd

-- Operator backfill (run once, after applying, before telemetry-api switches):
--
--   INSERT INTO signal_summary
--   SELECT subject, name, source, toUInt64(count()), min(timestamp), max(timestamp)
--   FROM signal GROUP BY subject, name, source
--   SETTINGS max_execution_time = 0, max_memory_usage = 0,
--            max_bytes_before_external_group_by = 32000000000;
--
--   INSERT INTO event_summary
--   SELECT subject, name, source, toUInt64(count()), min(timestamp), max(timestamp)
--   FROM event GROUP BY subject, name, source
--   SETTINGS max_execution_time = 0, max_memory_usage = 0,
--            max_bytes_before_external_group_by = 32000000000;
--
-- IMPORTANT: backfill must run BEFORE meaningful MV data accumulates or counts
-- double. Practically: run both backfills immediately after the migration
-- applies, then correct the overlap window if needed (counts are approximate
-- by design; first/last_seen are idempotent under overlap).

-- +goose Down

-- +goose StatementBegin
DROP VIEW IF EXISTS event_summary_mv;
-- +goose StatementEnd

-- +goose StatementBegin
DROP VIEW IF EXISTS signal_summary_mv;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS event_summary;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS signal_summary;
-- +goose StatementEnd
