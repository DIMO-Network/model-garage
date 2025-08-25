-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS event (
	-- original cloud event headers
    subject String COMMENT 'identifies the entity the event pertains to.',
    source String COMMENT 'the entity that identified and submitted the event (oracle).',
    producer String COMMENT 'the specific origin of the data used to determine the event (device).',
	cloud_event_id String COMMENT 'identifier for the cloudevent.',

	-- event infos
	name String COMMENT 'name of the event indicated by the oracle transmitting it.',
	timestamp DateTime64(6, 'UTC') COMMENT 'time at which the event described occurred, transmitted by oracle.',
	duration_ns UInt64 COMMENT 'duration in nanoseconds of the event.',
	metadata String COMMENT 'arbitrary JSON metadata provided by the user, containing additional event-related information.'
) ENGINE = ReplacingMergeTree
ORDER BY (subject, timestamp, name, source) SETTINGS index_granularity = 8192;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE event;
-- +goose StatementEnd
