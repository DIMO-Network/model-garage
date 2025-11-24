-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS signal_window_aggregates (
    token_id UInt32,
    window_start DateTime64(6, 'UTC'),
    window_size_seconds UInt16,
    signal_count SimpleAggregateFunction(sum, UInt64),
    distinct_signals AggregateFunction(uniq, UInt64)
)
ENGINE = AggregatingMergeTree()
ORDER BY (token_id, window_start, window_size_seconds)
PARTITION BY toYYYYMM(window_start)
-- +goose StatementEnd

-- +goose StatementBegin
CREATE MATERIALIZED VIEW IF NOT EXISTS signal_window_aggregates_mv
TO signal_window_aggregates
AS
SELECT
    token_id,
    toStartOfMinute(timestamp) AS window_start,
    60 AS window_size_seconds,
    sum(1) AS signal_count,
    uniqState(cityHash64(name)) AS distinct_signals
FROM signal
GROUP BY token_id, window_start, window_size_seconds;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP VIEW IF EXISTS signal_window_aggregates_mv;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS signal_window_aggregates;
-- +goose StatementEnd

