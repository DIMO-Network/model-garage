package vss_test

import (
	"testing"
	"time"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPackUnpackEvents_Roundtrip(t *testing.T) {
	t.Parallel()
	now := time.Now().UTC().Truncate(time.Millisecond)

	header := cloudevent.CloudEventHeader{
		SpecVersion: "1.0",
		ID:          "test-id",
		Source:      "test-source",
		Subject:     "test-subject",
		Producer:    "test-producer",
		Time:        now,
		Type:        cloudevent.TypeEvents,
		DataVersion: "1.0",
	}

	events := []vss.Event{
		{
			CloudEventHeader: header,
			Data: vss.EventData{
				Name:       "behavior.harshBraking",
				Timestamp:  now,
				DurationNs: 0,
				Metadata:   `{"counterValue":3}`,
			},
		},
		{
			CloudEventHeader: header,
			Data: vss.EventData{
				Name:       "behavior.extremeBraking",
				Timestamp:  now,
				DurationNs: 5000000000,
				Metadata:   `{"counterValue":1}`,
			},
		},
	}

	packed := vss.PackEvents(header, events)
	assert.Equal(t, "test-id", packed.ID)
	assert.Equal(t, "test-subject", packed.Subject)
	require.Len(t, packed.Data.Events, 2)
	assert.Equal(t, "behavior.harshBraking", packed.Data.Events[0].Name)
	assert.Equal(t, "behavior.extremeBraking", packed.Data.Events[1].Name)
	assert.Equal(t, cloudevent.TypeEvents, packed.Type)

	unpacked := vss.UnpackEvents(packed)
	require.Len(t, unpacked, 2)
	assert.Equal(t, "test-subject", unpacked[0].Subject)
	assert.Equal(t, "test-source", unpacked[0].Source)
	assert.Equal(t, "behavior.harshBraking", unpacked[0].Data.Name)
	assert.Equal(t, "behavior.extremeBraking", unpacked[1].Data.Name)
	assert.Equal(t, uint64(5000000000), unpacked[1].Data.DurationNs)
	assert.Equal(t, cloudevent.TypeEvent, unpacked[0].Type)
	assert.Equal(t, cloudevent.TypeEvent, unpacked[1].Type)
}

func TestEventToSlice_CloudEventID(t *testing.T) {
	t.Parallel()
	now := time.Now().UTC().Truncate(time.Millisecond)

	evt := vss.Event{
		CloudEventHeader: cloudevent.CloudEventHeader{
			ID:          "unique-event-id",
			Subject:     "did:erc721:1:0xABC:1",
			Source:      "test-source",
			Producer:    "test-producer",
			Type:        cloudevent.TypeEvent,
			DataVersion: "v1",
		},
		Data: vss.EventData{
			Name:         "behavior.harshBraking",
			Timestamp:    now,
			DurationNs:   1000000000,
			Metadata:     `{"counterValue":3}`,
			CloudEventID: "parent-event-id",
			Tags:         []string{"tag1"},
		},
	}

	slice := vss.EventToSlice(evt)

	// cloud_event_id column (index 3) should be Data.CloudEventID (parent), not header ID
	assert.Equal(t, "parent-event-id", slice[3], "cloud_event_id should come from Data.CloudEventID")
	assert.Equal(t, "did:erc721:1:0xABC:1", slice[0], "subject should come from header")
	assert.Equal(t, cloudevent.TypeEvent, slice[4], "type should come from header")
}

func TestPackUnpackEvents_CloudEventID(t *testing.T) {
	t.Parallel()
	now := time.Now().UTC().Truncate(time.Millisecond)

	header := cloudevent.CloudEventHeader{
		SpecVersion: "1.0",
		ID:          "envelope-id",
		Source:      "test-source",
		Subject:     "test-subject",
	}

	events := []vss.Event{
		{
			Data: vss.EventData{
				Name:         "behavior.harshBraking",
				Timestamp:    now,
				Metadata:     `{"counterValue":3}`,
				CloudEventID: "parent-event-id",
				Tags:         []string{},
			},
		},
	}

	packed := vss.PackEvents(header, events)
	assert.Equal(t, "parent-event-id", packed.Data.Events[0].CloudEventID)

	unpacked := vss.UnpackEvents(packed)
	require.Len(t, unpacked, 1)
	assert.Equal(t, "parent-event-id", unpacked[0].Data.CloudEventID, "CloudEventID should roundtrip through pack/unpack")
	assert.Equal(t, "envelope-id", unpacked[0].ID, "header ID should come from envelope")
}

func TestPackEvents_Empty(t *testing.T) {
	t.Parallel()
	header := cloudevent.CloudEventHeader{ID: "empty"}
	packed := vss.PackEvents(header, nil)
	assert.Empty(t, packed.Data.Events)

	unpacked := vss.UnpackEvents(packed)
	assert.Empty(t, unpacked)
}
