-- +goose Up
RENAME TABLE signal TO signal_backup;

CREATE TABLE IF NOT EXISTS signal;
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

-- +goose Down
DROP TABLE IF EXISTS signal;

RENAME TABLE signal_backup TO signal;
