// Package migrations provides the functionality to run goose migrations on a clickhouse dimo database.
package migrations

import (
	"context"
	"database/sql"
	"embed"

	"github.com/DIMO-Network/clickhouse-infra/pkg/migrate"
)

// BaseFS is the embed.FS for the migrations.
//
//go:embed *.sql
var BaseFS embed.FS

func RunGoose(ctx context.Context, gooseArgs []string, db *sql.DB) error {
	return migrate.RunGoose(ctx, gooseArgs, BaseFS, db)
}
