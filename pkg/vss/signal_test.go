package vss_test

import (
	"testing"
	"time"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPackUnpackSignals_Roundtrip(t *testing.T) {
	t.Parallel()
	now := time.Now().UTC().Truncate(time.Millisecond)

	header := cloudevent.CloudEventHeader{
		SpecVersion: "1.0",
		ID:          "sig-id",
		Source:      "test-source",
		Subject:     "test-subject",
		Producer:    "test-producer",
		Time:        now,
		Type:        "dimo.status",
	}

	signals := []vss.Signal{
		{
			Data: vss.SignalData{
				Timestamp:   now,
				Name:        "speed",
				ValueNumber: 65.5,
			},
		},
		{
			Data: vss.SignalData{
				Timestamp:   now,
				Name:        "odometer",
				ValueNumber: 12345,
			},
		},
	}

	packed := vss.PackSignals(header, signals)
	assert.Equal(t, "sig-id", packed.ID)
	assert.Equal(t, "test-subject", packed.Subject)
	require.Len(t, packed.Data.Signals, 2)
	assert.Equal(t, "speed", packed.Data.Signals[0].Name)
	assert.Equal(t, float64(65.5), packed.Data.Signals[0].ValueNumber)

	unpacked := vss.UnpackSignals(packed)
	require.Len(t, unpacked, 2)
	assert.Equal(t, "test-subject", unpacked[0].Subject)
	assert.Equal(t, "test-source", unpacked[0].Source)
	assert.Equal(t, "sig-id", unpacked[0].ID)
	assert.Equal(t, "speed", unpacked[0].Data.Name)
	assert.Equal(t, float64(65.5), unpacked[0].Data.ValueNumber)
	assert.Equal(t, "odometer", unpacked[1].Data.Name)
	assert.Equal(t, cloudevent.TypeSignal, unpacked[0].Type)
}

func TestPackSignals_WithLocation(t *testing.T) {
	t.Parallel()
	now := time.Now().UTC().Truncate(time.Millisecond)
	header := cloudevent.CloudEventHeader{ID: "loc-id", Subject: "subj"}

	signals := []vss.Signal{
		{
			Data: vss.SignalData{
				Timestamp: now,
				Name:      "currentLocationCoordinates",
				ValueLocation: vss.Location{
					Latitude:  37.7749,
					Longitude: -122.4194,
					HDOP:      1.2,
					Heading:   180.5,
				},
			},
		},
	}

	packed := vss.PackSignals(header, signals)
	unpacked := vss.UnpackSignals(packed)
	require.Len(t, unpacked, 1)
	assert.Equal(t, 37.7749, unpacked[0].Data.ValueLocation.Latitude)
	assert.Equal(t, -122.4194, unpacked[0].Data.ValueLocation.Longitude)
	assert.Equal(t, 1.2, unpacked[0].Data.ValueLocation.HDOP)
	assert.Equal(t, 180.5, unpacked[0].Data.ValueLocation.Heading)
	assert.Equal(t, cloudevent.TypeSignal, unpacked[0].Type)
}

func TestSignalToSlice_CloudEventID(t *testing.T) {
	t.Parallel()
	now := time.Now().UTC().Truncate(time.Millisecond)

	sig := vss.Signal{
		CloudEventHeader: cloudevent.CloudEventHeader{
			ID:      "unique-signal-id",
			Subject: "did:erc721:1:0xABC:1",
			Source:  "test-source",
			Type:    cloudevent.TypeSignal,
		},
		Data: vss.SignalData{
			Timestamp:    now,
			Name:         "speed",
			ValueNumber:  65.5,
			CloudEventID: "parent-event-id",
		},
	}

	slice := vss.SignalToSlice(sig)

	// cloud_event_id column (index 5) should be Data.CloudEventID (parent), not header ID
	assert.Equal(t, "parent-event-id", slice[5], "cloud_event_id should come from Data.CloudEventID")
	assert.Equal(t, "did:erc721:1:0xABC:1", slice[0], "subject should come from header")
	assert.Equal(t, "speed", slice[2], "name should come from Data.Name")
}

func TestPackUnpackSignals_CloudEventID(t *testing.T) {
	t.Parallel()
	now := time.Now().UTC().Truncate(time.Millisecond)

	header := cloudevent.CloudEventHeader{
		SpecVersion: "1.0",
		ID:          "envelope-id",
		Source:      "test-source",
		Subject:     "test-subject",
	}

	signals := []vss.Signal{
		{
			Data: vss.SignalData{
				Timestamp:    now,
				Name:         "speed",
				ValueNumber:  65.5,
				CloudEventID: "parent-event-id",
			},
		},
	}

	packed := vss.PackSignals(header, signals)
	assert.Equal(t, "parent-event-id", packed.Data.Signals[0].CloudEventID)

	unpacked := vss.UnpackSignals(packed)
	require.Len(t, unpacked, 1)
	assert.Equal(t, "parent-event-id", unpacked[0].Data.CloudEventID, "CloudEventID should roundtrip through pack/unpack")
	assert.Equal(t, "envelope-id", unpacked[0].ID, "header ID should come from envelope")
}

func TestPackSignals_Empty(t *testing.T) {
	t.Parallel()
	header := cloudevent.CloudEventHeader{ID: "empty"}
	packed := vss.PackSignals(header, nil)
	assert.Empty(t, packed.Data.Signals)

	unpacked := vss.UnpackSignals(packed)
	assert.Empty(t, unpacked)
}
