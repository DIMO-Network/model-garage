package migrations_test

import (
	"context"
	"testing"

	"github.com/DIMO-Network/clickhouse-infra/pkg/connect"
	"github.com/DIMO-Network/clickhouse-infra/pkg/connect/config"
	"github.com/DIMO-Network/clickhouse-infra/pkg/container"
	"github.com/DIMO-Network/model-garage/pkg/migrations"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMigration(t *testing.T) {
	ctx := context.Background()
	chcontainer, err := container.CreateClickHouseContainer(ctx, config.Settings{})
	require.NoError(t, err, "Failed to create clickhouse container")

	defer chcontainer.Terminate(ctx)

	db, err := chcontainer.GetClickhouseAsDB()
	require.NoError(t, err, "Failed to get clickhouse db")

	err = migrations.RunGoose(ctx, []string{"up", "-v"}, db)
	require.NoError(t, err, "Failed to run migration")

	conn, err := chcontainer.GetClickHouseAsConn()
	require.NoError(t, err, "Failed to get clickhouse connection")

	// Iterate over the rows and check the column names
	columns, err := connect.GetTableCols(ctx, conn, vss.TableName)
	require.NoError(t, err, "Failed to get current columns")

	expectedColumns := []connect.ColInfo{
		{Name: vss.TokenIDCol, Type: "UInt32", Comment: "token_id of this device data."},
		{Name: vss.TimestampCol, Type: "DateTime64(6, 'UTC')", Comment: "timestamp of when this data was collected."},
		{Name: vss.NameCol, Type: "LowCardinality(String)", Comment: "name of the signal collected."},
		{Name: vss.SourceCol, Type: "String", Comment: "source of the signal collected."},
		{Name: vss.ProducerCol, Type: "String", Comment: "producer of the collected signal."},
		{Name: vss.CloudEventIDCol, Type: "String", Comment: "Id of the Cloud Event that this signal was extracted from."},
		{Name: vss.ValueNumberCol, Type: "Float64", Comment: "float64 value of the signal collected."},
		{Name: vss.ValueStringCol, Type: "String", Comment: "string value of the signal collected."},
	}

	// Check if the actual columns match the expected columns
	require.Equal(t, expectedColumns, columns, "Unexpected table columns")

	// Close the DB connection
	err = db.Close()
	assert.NoError(t, err, "Failed to close DB connection")
	err = conn.Close()
	assert.NoError(t, err, "Failed to close clickhouse connection")
}
