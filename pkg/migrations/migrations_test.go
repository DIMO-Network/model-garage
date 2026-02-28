package migrations_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/DIMO-Network/clickhouse-infra/pkg/connect"
	"github.com/DIMO-Network/clickhouse-infra/pkg/connect/config"
	"github.com/DIMO-Network/clickhouse-infra/pkg/container"
	"github.com/DIMO-Network/model-garage/pkg/migrations"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSignalMigration(t *testing.T) {
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
		{Name: vss.SubjectCol, Type: "String", Comment: "Subject of the signal, typically a W3C DID."},
		{Name: vss.TimestampCol, Type: "DateTime64(6, 'UTC')", Comment: "Timestamp, ideally from when the signal was emitted."},
		{Name: vss.NameCol, Type: "LowCardinality(String)", Comment: "Name of the signal. The set of meaningful values for name depends on subject. The name also determines which of the value_ columns is expected to be populated."},
		{Name: vss.SourceCol, Type: "LowCardinality(String)", Comment: "Source of the signal. This is typically a checksummed connection address."},
		{Name: vss.ProducerCol, Type: "String", Comment: "Producer of the collected signal, typically another W3C DID."},
		{Name: vss.CloudEventIDCol, Type: "String", Comment: "Id of the CloudEvent from which this signal was extracted."},
		{Name: vss.ValueNumberCol, Type: "Float64", Comment: "The value for numeric (float64) signals."},
		{Name: vss.ValueStringCol, Type: "String", Comment: "The value for string signals."},
		{Name: vss.ValueLocationCol, Type: "Tuple(latitude Float64, longitude Float64, hdop Float64, heading Float64)", Comment: "The value for location signals. Some entries may be empty."},
	}

	// Check if the actual columns match the expected columns
	require.Equal(t, expectedColumns, columns, "Unexpected table columns")

	// Close the DB connection
	err = db.Close()
	assert.NoError(t, err, "Failed to close DB connection")
	err = conn.Close()
	assert.NoError(t, err, "Failed to close clickhouse connection")
}

func TestEventMigration(t *testing.T) {
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
	columns, err := connect.GetTableCols(ctx, conn, vss.EventTableName)
	require.NoError(t, err, "Failed to get current columns")

	expectedColumns := []connect.ColInfo{
		{Name: vss.EventSubjectCol, Type: "String", Comment: "identifies the entity the event pertains to."},
		{Name: vss.EventSourceCol, Type: "String", Comment: "the entity that identified and submitted the event (oracle)."},
		{Name: vss.EventProducerCol, Type: "String", Comment: "the specific origin of the data used to determine the event (device)."},
		{Name: vss.EventCloudEventIDCol, Type: "String", Comment: "identifier for the cloudevent."},
		{Name: vss.EventNameCol, Type: "String", Comment: "name of the event indicated by the oracle transmitting it."},
		{Name: vss.EventTimestampCol, Type: "DateTime64(6, 'UTC')", Comment: "time at which the event described occurred, transmitted by oracle."},
		{Name: vss.EventDurationNsCol, Type: "UInt64", Comment: "duration in nanoseconds of the event."},
		{Name: vss.EventMetadataCol, Type: "String", Comment: "arbitrary JSON metadata provided by the user, containing additional event-related information."},
		{Name: vss.EventTagsCol, Type: "Array(String)", Comment: "tags for the event."},
	}

	// Check if the actual columns match the expected columns
	require.Equal(t, expectedColumns, columns, "Unexpected table columns")

	event := vss.Event{
		Subject:      "subject",
		Source:       "source",
		Producer:     "producer",
		CloudEventID: "cloudEventId",
		Name:         "name",
		Timestamp:    time.Now().Truncate(time.Microsecond),
		DurationNs:   1000000000,
		Metadata:     "metadata",
		Tags:         []string{"tag1", "tag2"},
	}
	batch, err := conn.PrepareBatch(context.Background(), "INSERT INTO "+vss.EventTableName)
	require.NoError(t, err, "Failed to prepare batch")
	defer batch.Close()

	err = batch.Append(vss.EventToSlice(event)...)
	require.NoError(t, err, "Failed to append event")

	err = batch.Send()
	require.NoError(t, err, "Failed to send batch")

	rows, err := conn.Query(ctx, fmt.Sprintf("SELECT * FROM %s", vss.EventTableName))
	require.NoError(t, err, "Failed to select event")
	defer rows.Close()

	events := []vss.Event{}
	for rows.Next() {
		var event vss.Event
		err = rows.Scan(&event.Subject, &event.Source, &event.Producer, &event.CloudEventID, &event.Name, &event.Timestamp, &event.DurationNs, &event.Metadata, &event.Tags)

		require.NoError(t, err, "Failed to scan event")
		events = append(events, event)
	}

	require.Equal(t, 1, len(events), "Expected 1 event")

	assert.Truef(t, event.Timestamp.Equal(events[0].Timestamp), "Event timestamp mismatch: %v != %v", event.Timestamp, events[0].Timestamp)
	event.Timestamp = events[0].Timestamp
	assert.Equal(t, event, events[0], "Event mismatch")

	// Close the DB connection
	err = db.Close()
	assert.NoError(t, err, "Failed to close DB connection")
	err = conn.Close()
	assert.NoError(t, err, "Failed to close clickhouse connection")
}
