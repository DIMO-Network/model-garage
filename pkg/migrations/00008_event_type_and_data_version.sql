-- +goose Up
-- +goose StatementBegin
ALTER TABLE event ADD COLUMN type String DEFAULT '' COMMENT 'CloudEvent type of the event.' AFTER cloud_event_id;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE event ADD COLUMN data_version String DEFAULT '' COMMENT 'Version of the data schema.' AFTER type;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE event DROP COLUMN type;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE event DROP COLUMN data_version;
-- +goose StatementEnd
