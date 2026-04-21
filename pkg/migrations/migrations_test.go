package migrations_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/DIMO-Network/clickhouse-infra/pkg/connect"
	"github.com/DIMO-Network/clickhouse-infra/pkg/connect/config"
	"github.com/DIMO-Network/clickhouse-infra/pkg/container"
	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/migrations"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testCHSettings() config.Settings {
	return config.Settings{
		Password: "test-password",
	}
}

func TestSignalMigration(t *testing.T) {
	ctx := context.Background()
	chcontainer, err := container.CreateClickHouseContainer(ctx, testCHSettings())
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

func TestSignalInsertAfterMigrations(t *testing.T) {
	ctx := context.Background()
	chcontainer, err := container.CreateClickHouseContainer(ctx, testCHSettings())
	require.NoError(t, err, "Failed to create clickhouse container")

	defer chcontainer.Terminate(ctx)

	db, err := chcontainer.GetClickhouseAsDB()
	require.NoError(t, err, "Failed to get clickhouse db")

	err = migrations.RunGoose(ctx, []string{"up", "-v"}, db)
	require.NoError(t, err, "Failed to run migration")

	conn, err := chcontainer.GetClickHouseAsConn()
	require.NoError(t, err, "Failed to get clickhouse connection")

	signal := vss.Signal{
		CloudEventHeader: cloudevent.CloudEventHeader{
			Subject:  "did:erc721:137:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:22892",
			Source:   "did:ethr:137:0xcd445F4c6bDAD32b68a2939b912150Fe3C88803E",
			Producer: "did:key:z6MkiTBz1sZ4n4bipYw4v8M4y9Q7gL1Y9s9m4h7G9r2x3w1q",
		},
		Data: vss.SignalData{
			Timestamp:    time.Now().UTC().Truncate(time.Microsecond),
			Name:         vss.FieldSpeed,
			ValueNumber:  42,
			CloudEventID: "signal-insert-regression-test",
		},
	}

	batch, err := conn.PrepareBatch(ctx, "INSERT INTO "+vss.TableName)
	require.NoError(t, err, "Failed to prepare signal batch")
	defer batch.Close()

	err = batch.Append(vss.SignalToSlice(signal)...)
	require.NoError(t, err, "Failed to append signal")

	err = batch.Send()
	require.NoError(t, err, "Failed to send signal")

	rows, err := conn.Query(ctx, fmt.Sprintf("SELECT * FROM %s", vss.TableName))
	require.NoError(t, err, "Failed to select signal")
	defer rows.Close()

	signals := []vss.Signal{}
	for rows.Next() {
		var signal vss.Signal
		err = rows.Scan(
			&signal.Subject,
			&signal.Data.Timestamp,
			&signal.Data.Name,
			&signal.Source,
			&signal.Producer,
			&signal.Data.CloudEventID,
			&signal.Data.ValueNumber,
			&signal.Data.ValueString,
			&signal.Data.ValueLocation,
		)
		require.NoError(t, err, "Failed to scan signal")
		signals = append(signals, signal)
	}

	require.Equal(t, 1, len(signals), "Expected 1 signal")
	assert.Truef(t, signal.Data.Timestamp.Equal(signals[0].Data.Timestamp), "Signal timestamp mismatch: %v != %v", signal.Data.Timestamp, signals[0].Data.Timestamp)
	signal.Data.Timestamp = signals[0].Data.Timestamp
	assert.Equal(t, signal, signals[0], "Signal mismatch")

	err = db.Close()
	assert.NoError(t, err, "Failed to close DB connection")
	err = conn.Close()
	assert.NoError(t, err, "Failed to close clickhouse connection")
}

func TestSignalLocationInsertAfterMigrations(t *testing.T) {
	ctx := context.Background()
	chcontainer, err := container.CreateClickHouseContainer(ctx, testCHSettings())
	require.NoError(t, err, "Failed to create clickhouse container")

	defer chcontainer.Terminate(ctx)

	db, err := chcontainer.GetClickhouseAsDB()
	require.NoError(t, err, "Failed to get clickhouse db")

	err = migrations.RunGoose(ctx, []string{"up", "-v"}, db)
	require.NoError(t, err, "Failed to run migration")

	conn, err := chcontainer.GetClickHouseAsConn()
	require.NoError(t, err, "Failed to get clickhouse connection")

	signal := vss.Signal{
		CloudEventHeader: cloudevent.CloudEventHeader{
			Subject:  "did:erc721:137:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:22892",
			Source:   "did:ethr:137:0xcd445F4c6bDAD32b68a2939b912150Fe3C88803E",
			Producer: "did:key:z6MkiTBz1sZ4n4bipYw4v8M4y9Q7gL1Y9s9m4h7G9r2x3w1q",
		},
		Data: vss.SignalData{
			Timestamp:    time.Now().UTC().Truncate(time.Microsecond),
			Name:         "currentLocationCoordinates",
			CloudEventID: "signal-location-insert-regression-test",
			ValueLocation: vss.Location{
				Latitude:  42.3314,
				Longitude: -83.0458,
				HDOP:      0.7,
				Heading:   135.5,
			},
		},
	}

	batch, err := conn.PrepareBatch(ctx, "INSERT INTO "+vss.TableName)
	require.NoError(t, err, "Failed to prepare signal batch")
	defer batch.Close()

	err = batch.Append(vss.SignalToSlice(signal)...)
	require.NoError(t, err, "Failed to append signal")

	err = batch.Send()
	require.NoError(t, err, "Failed to send signal")

	rows, err := conn.Query(ctx, fmt.Sprintf("SELECT * FROM %s WHERE name = '%s'", vss.TableName, signal.Data.Name))
	require.NoError(t, err, "Failed to select signal")
	defer rows.Close()

	signals := []vss.Signal{}
	for rows.Next() {
		var signal vss.Signal
		err = rows.Scan(
			&signal.Subject,
			&signal.Data.Timestamp,
			&signal.Data.Name,
			&signal.Source,
			&signal.Producer,
			&signal.Data.CloudEventID,
			&signal.Data.ValueNumber,
			&signal.Data.ValueString,
			&signal.Data.ValueLocation,
		)
		require.NoError(t, err, "Failed to scan signal")
		signals = append(signals, signal)
	}

	require.Equal(t, 1, len(signals), "Expected 1 signal")
	assert.Truef(t, signal.Data.Timestamp.Equal(signals[0].Data.Timestamp), "Signal timestamp mismatch: %v != %v", signal.Data.Timestamp, signals[0].Data.Timestamp)
	signal.Data.Timestamp = signals[0].Data.Timestamp
	assert.Equal(t, signal, signals[0], "Signal mismatch")

	err = db.Close()
	assert.NoError(t, err, "Failed to close DB connection")
	err = conn.Close()
	assert.NoError(t, err, "Failed to close clickhouse connection")
}

func TestEventMigration(t *testing.T) {
	ctx := context.Background()
	chcontainer, err := container.CreateClickHouseContainer(ctx, testCHSettings())
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
		{Name: vss.EventTypeCol, Type: "String", Comment: "CloudEvent type of the event."},
		{Name: vss.EventDataVersionCol, Type: "String", Comment: "Version of the data schema."},
		{Name: vss.EventNameCol, Type: "String", Comment: "name of the event indicated by the oracle transmitting it."},
		{Name: vss.EventTimestampCol, Type: "DateTime64(6, 'UTC')", Comment: "time at which the event described occurred, transmitted by oracle."},
		{Name: vss.EventDurationNsCol, Type: "UInt64", Comment: "duration in nanoseconds of the event."},
		{Name: vss.EventMetadataCol, Type: "String", Comment: "arbitrary JSON metadata provided by the user, containing additional event-related information."},
		{Name: vss.EventTagsCol, Type: "Array(String)", Comment: "tags for the event."},
	}

	// Check if the actual columns match the expected columns
	require.Equal(t, expectedColumns, columns, "Unexpected table columns")

	event := vss.Event{
		CloudEventHeader: cloudevent.CloudEventHeader{
			Subject:     "subject",
			Source:      "source",
			Producer:    "producer",
			Type:        cloudevent.TypeEvent,
			DataVersion: "v1",
		},
		Data: vss.EventData{
			Name:         "name",
			Timestamp:    time.Now().Truncate(time.Microsecond),
			DurationNs:   1000000000,
			Metadata:     "metadata",
			CloudEventID: "parentCloudEventId",
			Tags:         []string{"tag1", "tag2"},
		},
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
		err = rows.Scan(
			&event.Subject, &event.Source, &event.Producer, &event.Data.CloudEventID,
			&event.Type, &event.DataVersion,
			&event.Data.Name, &event.Data.Timestamp, &event.Data.DurationNs,
			&event.Data.Metadata, &event.Data.Tags,
		)

		require.NoError(t, err, "Failed to scan event")
		events = append(events, event)
	}

	require.Equal(t, 1, len(events), "Expected 1 event")

	assert.Truef(t, event.Data.Timestamp.Equal(events[0].Data.Timestamp), "Event timestamp mismatch: %v != %v", event.Data.Timestamp, events[0].Data.Timestamp)
	event.Data.Timestamp = events[0].Data.Timestamp
	assert.Equal(t, event, events[0], "Event mismatch")

	// Close the DB connection
	err = db.Close()
	assert.NoError(t, err, "Failed to close DB connection")
	err = conn.Close()
	assert.NoError(t, err, "Failed to close clickhouse connection")
}
