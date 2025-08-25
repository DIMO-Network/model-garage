// Package main is the entry point to run goose CLI with model-garage migrations.
package main

import (
	"context"
	"log"

	"github.com/DIMO-Network/clickhouse-infra/pkg/migrate"
	"github.com/DIMO-Network/model-garage/pkg/migrations"
)

func main() {
	err := migrate.RunGooseCmd(context.Background(), migrations.BaseFS)
	if err != nil {
		log.Fatal(err)
	}
}
