-- +goose Up
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
  INDEX idx_token_name_ts (token_id, signal_name, timestamp) TYPE minmax GRANULARITY 4
)
ENGINE = MergeTree()
ORDER BY (token_id, signal_name, timestamp);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE MATERIALIZED VIEW IF NOT EXISTS signal_state_changes_mv
TO signal_state_changes
AS
SELECT * FROM (
  SELECT
    token_id,
    name as signal_name,
    timestamp,
    value_number as new_state,
    lagInFrame(value_number, 1, -1) OVER (
      PARTITION BY token_id, name 
      ORDER BY timestamp
    ) as prev_state,
    dateDiff('second',
      lagInFrame(timestamp, 1, timestamp) OVER (
        PARTITION BY token_id, name 
        ORDER BY timestamp
      ),
      timestamp
    ) as time_since_prev_seconds,
    source,
    producer,
    cloud_event_id
  FROM signal
  WHERE name IN ('isIgnitionOn')
) WHERE prev_state != new_state;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP VIEW IF EXISTS signal_state_changes_mv;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS signal_state_changes;
-- +goose StatementEnd

