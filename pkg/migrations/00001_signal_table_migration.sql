-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS signal
(
	token_id UInt32 COMMENT 'token_id of this device data.',
	timestamp DateTime64(6, 'UTC') COMMENT 'timestamp of when this data was collected.',
	name LowCardinality(String) COMMENT 'name of the signal collected.',
	source String COMMENT 'source of the signal collected.',
	value_number Float64 COMMENT 'float64 value of the signal collected.',
	value_string String COMMENT 'string value of the signal collected.'
)
ENGINE = ReplacingMergeTree
ORDER BY (token_id, timestamp, name)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE signal
-- +goose StatementEnd
