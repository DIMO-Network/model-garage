-- +goose Up
-- +goose StatementBegin
ALTER TABLE event ADD COLUMN tags Array(String) COMMENT 'tags for the event.';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE event DROP COLUMN tags;
-- +goose StatementEnd
