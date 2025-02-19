package migrations

import (
	"context"
	"database/sql"
	"runtime"

	"github.com/pressly/goose/v3"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	registerFunc := func() {
		goose.AddNamedMigrationContext(filename, upAddProducerAndCloudEventID, downAddProducerAndCloudEventID)
	}
	registerFuncs = append(registerFuncs, registerFunc)
}

func upAddProducerAndCloudEventID(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	upStatements := []string{
		"ALTER TABLE signal ADD COLUMN producer String COMMENT 'producer of the collected signal.' AFTER source",
		"ALTER TABLE signal ADD COLUMN cloud_event_id String COMMENT 'Id of the Cloud Event that this signal was extracted from.' AFTER producer",
	}
	for _, upStatement := range upStatements {
		_, err := tx.ExecContext(ctx, upStatement)
		if err != nil {
			return err
		}
	}
	return nil
}

func downAddProducerAndCloudEventID(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	downStatements := []string{
		"ALTER TABLE signal DROP COLUMN producer",
		"ALTER TABLE signal DROP COLUMN cloud_event_id",
	}
	for _, downStatement := range downStatements {
		_, err := tx.ExecContext(ctx, downStatement)
		if err != nil {
			return err
		}
	}
	return nil
}
