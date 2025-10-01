-- +goose Up
-- +goose StatementBegin
ALTER TABLE signal ADD COLUMN value_location Tuple(latitude Float64, longitude Float64, hdop Float64) COMMENT 'Location value of the signal collected.';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE signal DROP COLUMN value_location;
-- +goose StatementEnd
