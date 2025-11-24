-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS signal_last_state (
  token_id UInt32,
  signal_name LowCardinality(String),
  last_timestamp DateTime64(6, 'UTC'),
  last_value Float64,
  last_source String DEFAULT '',
  last_producer String DEFAULT '',
  version UInt64 DEFAULT now64()
)
ENGINE = ReplacingMergeTree(version)
ORDER BY (token_id, signal_name);
-- +goose StatementEnd

-- +goose StatementBegin
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
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS signal_state_changes (
  token_id UInt32,
  signal_name LowCardinality(String),
  timestamp DateTime64(6, 'UTC'),
  new_state Float64,
  prev_state Float64,
  time_since_prev_seconds Int32,
  source LowCardinality(String),
  producer LowCardinality(String),
  cloud_event_id String DEFAULT '',
  version UInt64 DEFAULT now64(),
  INDEX idx_token_name_ts (token_id, signal_name, timestamp)
)
ENGINE = ReplacingMergeTree(version)
ORDER BY (token_id, signal_name, timestamp)
PARTITION BY toYYYYMM(timestamp);
-- +goose StatementEnd

-- +goose StatementBegin
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

-- +goose Down
-- +goose StatementBegin
DROP VIEW IF EXISTS signal_state_changes_mv;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS signal_state_changes;
-- +goose StatementEnd

-- +goose StatementBegin
DROP VIEW IF EXISTS signal_last_state_mv;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS signal_last_state;
-- +goose StatementEnd

