-- +goose Up

-- The per-part latest projections inherited the signal table's part fan-out
-- (projections are stored per part) with fatter aggregate-state granules, and
-- measured slower than raw primary-key reads in prod (signalsLatest p50 ~3s
-- via projection vs ~110ms raw, 2026-06-11). signal_latest (00010) and the
-- summary tables (00011) replace them. Apply only after telemetry-api reads
-- the new tables (this migration ships in a later model-garage release to
-- enforce that ordering).

-- +goose StatementBegin
ALTER TABLE signal DROP PROJECTION IF EXISTS signal_latest_by_subject_source_name;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE event DROP PROJECTION IF EXISTS event_latest_by_subject_source_name;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE signal RESET SETTING deduplicate_merge_projection_mode;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE event RESET SETTING deduplicate_merge_projection_mode;
-- +goose StatementEnd

-- +goose Down

-- Recreate the projections as defined in migration 00009. Re-materializing is
-- an operator decision (heavy mutation); without it the projections only
-- cover parts written after this rollback.

-- +goose StatementBegin
ALTER TABLE signal MODIFY SETTING deduplicate_merge_projection_mode = 'rebuild';
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE event MODIFY SETTING deduplicate_merge_projection_mode = 'rebuild';
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE signal ADD PROJECTION IF NOT EXISTS signal_latest_by_subject_source_name (
    SELECT
        subject,
        source,
        name,
        max(timestamp),
        min(timestamp),
        argMax(value_number, timestamp),
        argMax(value_string, timestamp),
        argMaxIf(value_location, timestamp, (tupleElement(value_location, 'latitude') != 0) OR (tupleElement(value_location, 'longitude') != 0)),
        maxIf(timestamp, (tupleElement(value_location, 'latitude') != 0) OR (tupleElement(value_location, 'longitude') != 0)),
        argMax(producer, timestamp),
        argMax(cloud_event_id, timestamp),
        count()
    GROUP BY subject, source, name
);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE event ADD PROJECTION IF NOT EXISTS event_latest_by_subject_source_name (
    SELECT
        subject,
        source,
        name,
        max(timestamp),
        min(timestamp),
        argMax(producer, timestamp),
        argMax(cloud_event_id, timestamp),
        count()
    GROUP BY subject, source, name
);
-- +goose StatementEnd
