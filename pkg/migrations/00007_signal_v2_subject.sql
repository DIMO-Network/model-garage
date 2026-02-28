-- +goose Up

-- Drop these state change views first, since they reference signal.
DROP VIEW IF EXISTS signal_state_changes_mv;
DROP VIEW IF EXISTS signal_last_state_mv;

-- Back up existing tables.
RENAME TABLE signal TO signal_backup;
RENAME TABLE signal_last_state TO signal_last_state_backup;
RENAME TABLE signal_state_changes TO signal_state_changes_backup;

-- Recreate signal with subject instead of token_id.
CREATE TABLE IF NOT EXISTS signal
(
    `subject` String COMMENT 'Subject of the signal, typically a W3C DID.',
    `timestamp` DateTime64(6, 'UTC') COMMENT 'Timestamp, ideally from when the signal was emitted.' CODEC(Delta, ZSTD),
    `name` LowCardinality(String) COMMENT 'Name of the signal. The set of meaningful values for name depends on subject. The name also determines which of the value_ columns is expected to be populated.',
    `source` LowCardinality(String) COMMENT 'Source of the signal. This is typically a checksummed connection address.',
    `producer` String COMMENT 'Producer of the collected signal, typically another W3C DID.',
    `cloud_event_id` String COMMENT 'Id of the CloudEvent from which this signal was extracted.',
    `value_number` Float64 COMMENT 'The value for numeric (float64) signals.',
    `value_string` String COMMENT 'The value for string signals.',
    `value_location` Tuple(
        latitude Float64,
        longitude Float64,
        hdop Float64,
        heading Float64) COMMENT 'The value for location signals. Some entries may be empty.'
)
ENGINE = ReplacingMergeTree
PARTITION BY toYYYYMM(timestamp)
ORDER BY (subject, timestamp, name)
COMMENT 'Contains signals extracted from incoming CloudEvents. Most column names refer to CloudEvent concepts.';

-- Recreate signal_last_state and the associated view, based on the new signal table.
CREATE TABLE IF NOT EXISTS signal_last_state (
  subject String,
  signal_name LowCardinality(String),
  last_timestamp DateTime64(6, 'UTC'),
  last_value Float64,
  last_source String DEFAULT '',
  last_producer String DEFAULT '',
  version UInt64 DEFAULT now64()
)
ENGINE = ReplacingMergeTree(version)
ORDER BY (subject, signal_name);

CREATE MATERIALIZED VIEW IF NOT EXISTS signal_last_state_mv
TO signal_last_state
AS
SELECT
  subject,
  name as signal_name,
  timestamp as last_timestamp,
  value_number as last_value,
  source as last_source,
  producer as last_producer,
  now64() as version
FROM signal
WHERE name IN ('isIgnitionOn');

-- Recreate signal_state_changes similarly.
CREATE TABLE IF NOT EXISTS signal_state_changes (
  subject String,
  signal_name LowCardinality(String),
  timestamp DateTime64(6, 'UTC'),
  new_state Float64,
  prev_state Float64,
  time_since_prev_seconds Int32,
  source LowCardinality(String),
  producer LowCardinality(String),
  cloud_event_id String DEFAULT '',
  version UInt64 DEFAULT now64(),
  INDEX idx_subject_name_ts (subject, signal_name, timestamp) TYPE minmax GRANULARITY 4
)
ENGINE = ReplacingMergeTree(version)
ORDER BY (subject, signal_name, timestamp)
PARTITION BY toYYYYMM(timestamp);

CREATE MATERIALIZED VIEW IF NOT EXISTS signal_state_changes_mv
TO signal_state_changes
AS
SELECT
  s.subject,
  s.name as signal_name,
  s.timestamp,
  s.value_number as new_state,
  ifNull(ls.last_value, -1) as prev_state,
  dateDiff('second', ifNull(ls.last_timestamp, s.timestamp), s.timestamp) as time_since_prev_seconds,
  s.source,
  s.producer,
  s.cloud_event_id,
  now64() as version
FROM signal s
LEFT JOIN (
  SELECT * FROM signal_last_state FINAL
) ls ON
  s.subject = ls.subject AND
  s.name = ls.signal_name
WHERE
  s.name IN ('isIgnitionOn') AND
  (ls.last_value IS NULL OR s.value_number != ls.last_value) AND
  (ls.last_timestamp IS NULL OR s.timestamp > ls.last_timestamp);

-- +goose Down

-- Drop the new views and tables.
DROP VIEW IF EXISTS signal_state_changes_mv;
DROP VIEW IF EXISTS signal_last_state_mv;

DROP TABLE IF EXISTS signal_state_changes;
DROP TABLE IF EXISTS signal_last_state;

DROP TABLE IF EXISTS signal;

-- Restore backups.
RENAME TABLE signal_backup TO signal;
RENAME TABLE signal_last_state_backup TO signal_last_state;
RENAME TABLE signal_state_changes_backup TO signal_state_changes;

-- Recreate the original materialized views with token_id.
CREATE MATERIALIZED VIEW IF NOT EXISTS signal_last_state_mv
TO signal_last_state
AS
SELECT
  token_id,
  name as signal_name,
  timestamp as last_timestamp,
  value_number as last_value,
  source as last_source,
  producer as last_producer,
  now64() as version
FROM signal
WHERE name IN ('isIgnitionOn');

CREATE MATERIALIZED VIEW IF NOT EXISTS signal_state_changes_mv
TO signal_state_changes
AS
SELECT
  s.token_id,
  s.name as signal_name,
  s.timestamp,
  s.value_number as new_state,
  ifNull(ls.last_value, -1) as prev_state,
  dateDiff('second', ifNull(ls.last_timestamp, s.timestamp), s.timestamp) as time_since_prev_seconds,
  s.source,
  s.producer,
  s.cloud_event_id,
  now64() as version
FROM signal s
LEFT JOIN (
  SELECT * FROM signal_last_state FINAL
) ls ON
  s.token_id = ls.token_id AND
  s.name = ls.signal_name
WHERE
  s.name IN ('isIgnitionOn') AND
  (ls.last_value IS NULL OR s.value_number != ls.last_value) AND
  (ls.last_timestamp IS NULL OR s.timestamp > ls.last_timestamp);
-- +goose StatementEnd
