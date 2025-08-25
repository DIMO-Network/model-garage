-- +goose Up
-- +goose StatementBegin
ALTER TABLE signal ADD COLUMN producer String COMMENT 'producer of the collected signal.' AFTER source;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE signal ADD COLUMN cloud_event_id String COMMENT 'Id of the Cloud Event that this signal was extracted from.' AFTER producer;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE signal DROP COLUMN producer;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE signal DROP COLUMN cloud_event_id;
-- +goose StatementEnd
