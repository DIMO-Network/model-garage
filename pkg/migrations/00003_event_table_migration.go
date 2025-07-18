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
ORDER BY (subject, timestamp, name, source) SETTINGS index_granularity = 8192;`
