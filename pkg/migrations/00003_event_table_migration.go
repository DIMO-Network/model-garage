// Package migrations provides the functionality for managing database migrations for the vss tables.
package migrations

import (
	"context"
	"database/sql"
	"runtime"

	"github.com/pressly/goose/v3"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	registerFunc := func() { goose.AddNamedMigrationContext(filename, upEventTable, downEventTable) }
	registerFuncs = append(registerFuncs, registerFunc)
}

func upEventTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	upStatements := []string{
		createVehicleEventStmt,
	}
	for _, upStatement := range upStatements {
		_, err := tx.ExecContext(ctx, upStatement)
		if err != nil {
			return err
		}
	}
	return nil
}

func downEventTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	downStatements := []string{
		"DROP TABLE event",
	}
	for _, downStatement := range downStatements {
		_, err := tx.ExecContext(ctx, downStatement)
		if err != nil {
			return err
		}
	}
	return nil
}

const createVehicleEventStmt = `
CREATE TABLE IF NOT EXISTS event
(
	id String COMMENT 'unique event id',

	-- original cloud event headers
    subject String COMMENT 'identifying the subject of the event within the context of the event producer',
    cloudevent_time DateTime64(3, 'UTC') COMMENT 'Time at which the cloudevent occurred.',
    cloudevent_id String COMMENT 'Identifier for the cloudevent.',
    source String COMMENT 'Entity that is responsible for providing this cloud event',
    producer String COMMENT 'specific instance, process or device that creates the data structure describing the cloud event.',
    data_content_type String COMMENT 'Type of data of this object.',
    data_version String COMMENT 'Version of the data stored for this cloud event.',
	extras String COMMENT 'Extra metadata for the cloud event',
    index_key String COMMENT 'Key of the backing object for this cloud event',

	-- event infos
	event_name String COMMENT 'name of the event indicated by the oracle transmitting it',
	event_time DateTime64(3, 'UTC') COMMENT 'optional field denoting time at which the event described occurred, transmitted by oracle',
	event_duration String COMMENT 'optional event duration field transmitted by oracle',
	event_metadata JSON COMMENT 'freeform json containing infos relevant to event transmitted'
)
ENGINE = ReplacingMergeTree
ORDER BY (subject, event_time, event_name, source, cloudevent_id) SETTINGS index_granularity = 8192;`
