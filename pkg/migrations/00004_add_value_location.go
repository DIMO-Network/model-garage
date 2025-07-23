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
		goose.AddNamedMigrationContext(filename, upAddValueLocation, downAddValueLocation)
	}
	registerFuncs = append(registerFuncs, registerFunc)
}

func upAddValueLocation(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	upStatements := []string{
		"ALTER TABLE signal ADD COLUMN value_location Tuple(latitude Float64, longitude Float64, hdop Float64) COMMENT 'Geographic point value, expressed in WGS-84 degrees.'",
	}
	for _, upStatement := range upStatements {
		_, err := tx.ExecContext(ctx, upStatement)
		if err != nil {
			return err
		}
	}
	return nil
}

func downAddValueLocation(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	downStatements := []string{
		"ALTER TABLE signal DROP COLUMN value_location",
	}
	for _, downStatement := range downStatements {
		_, err := tx.ExecContext(ctx, downStatement)
		if err != nil {
			return err
		}
	}
	return nil
}
