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
		goose.AddNamedMigrationContext(filename, upAddLocation, downAddLocation)
	}
	registerFuncs = append(registerFuncs, registerFunc)
}

func upAddLocation(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	upStatements := []string{
		"ALTER TABLE signal ADD COLUMN value_location Tuple(longitude Float64, latitude Float64) COMMENT 'Geographic point value, expressed in WGS-84 degrees, with longitude first in order to match most ClickHouse functions.'",
	}
	for _, upStatement := range upStatements {
		_, err := tx.ExecContext(ctx, upStatement)
		if err != nil {
			return err
		}
	}
	return nil
}

func downAddLocation(ctx context.Context, tx *sql.Tx) error {
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
