-- +goose Up

-- ReplacingMergeTree tables reject projections by default because replaced
-- rows would leave the projection stale. Switching to 'rebuild' tells
-- ClickHouse to recompute affected projection parts during a deduplicating
-- merge so the pre-aggregated values stay correct.

-- +goose StatementBegin
ALTER TABLE signal MODIFY SETTING deduplicate_merge_projection_mode = 'rebuild';
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE event MODIFY SETTING deduplicate_merge_projection_mode = 'rebuild';
-- +goose StatementEnd

-- Aggregating projection that answers "latest value per signal name",
-- "last seen", and per-name summary (count + min/max timestamp) queries
-- from a few rows per (subject, source, name) rather than a full-history
-- scan of the signal table. The argMaxIf/maxIf aggregates on value_location
-- exclude (0, 0) points so the projection can serve the location-latest
-- query directly. min(timestamp) + count() serve data summary use cases.

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
ALTER TABLE signal MATERIALIZE PROJECTION IF EXISTS signal_latest_by_subject_source_name;
-- +goose StatementEnd

-- Same shape for the event table (no value_ columns on events).

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

-- +goose StatementBegin
ALTER TABLE event MATERIALIZE PROJECTION IF EXISTS event_latest_by_subject_source_name;
-- +goose StatementEnd

-- +goose Down

-- +goose StatementBegin
ALTER TABLE signal DROP PROJECTION IF EXISTS signal_latest_by_subject_source_name;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE event DROP PROJECTION IF EXISTS event_latest_by_subject_source_name;
-- +goose StatementEnd

-- Drop the per-table override and fall back to the server default
-- (currently 'ignore' on DIMO's cluster), matching the pre-migration state.

-- +goose StatementBegin
ALTER TABLE signal RESET SETTING deduplicate_merge_projection_mode;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE event RESET SETTING deduplicate_merge_projection_mode;
-- +goose StatementEnd
