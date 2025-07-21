package ruptela_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/ruptela"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/stretchr/testify/require"
)

func TestDecodeEventWithAllEventTypes(t *testing.T) {
	t.Parallel()

	// Test data with all event types - braking, acceleration, and cornering
	allEventsInputJSON := `{
		"id": "test-cloud-event-id",
		"source": "ruptela/test",
		"producer": "test-producer",
		"specversion": "1.0",
		"subject": "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:33",
		"time": "2024-09-27T08:33:26Z",
		"type": "dimo.event",
		"data": {
			"signals": {
				"135": "35", 
				"136": "A",
				"143": "7"
			}
		}
	}`

	var event cloudevent.RawEvent
	err := json.Unmarshal([]byte(allEventsInputJSON), &event)
	require.NoError(t, err)

	actualEvents, err := ruptela.DecodeEvent(event)
	require.NoError(t, err)
	require.Len(t, actualEvents, 4, "should have 4 events total")

	// Verify common fields are set for all events
	for _, evt := range actualEvents {
		require.Equal(t, "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:33", evt.Subject)
		require.Equal(t, "ruptela/test", evt.Source)
		require.Equal(t, "test-producer", evt.Producer)
		require.Equal(t, "test-cloud-event-id", evt.CloudEventID)
		require.Equal(t, time.Date(2024, 9, 27, 8, 33, 26, 0, time.UTC), evt.Timestamp)
	}

	// Check specific events
	var harshBrakingEvents []vss.Event
	var extremeBrakingEvents []vss.Event
	var accelerationEvents []vss.Event
	var corneringEvents []vss.Event

	for _, evt := range actualEvents {
		switch evt.Name {
		case ruptela.EventNameHarshBraking:
			harshBrakingEvents = append(harshBrakingEvents, evt)
		case ruptela.EventNameExtremeBraking:
			extremeBrakingEvents = append(extremeBrakingEvents, evt)
		case ruptela.EventNameAcceleration:
			accelerationEvents = append(accelerationEvents, evt)
		case ruptela.EventNameCornering:
			corneringEvents = append(corneringEvents, evt)
		}
	}

	require.Len(t, harshBrakingEvents, 1, "should have 1 harsh braking event (LSB)")
	require.Len(t, extremeBrakingEvents, 1, "should have 1 extreme braking event (MSB)")
	require.Len(t, accelerationEvents, 1, "should have 1 acceleration event")
	require.Len(t, corneringEvents, 1, "should have 1 cornering event")

	// Verify braking event metadata
	require.JSONEq(t, `{"counterValue":5}`, harshBrakingEvents[0].Metadata)
	require.JSONEq(t, `{"counterValue":3}`, extremeBrakingEvents[0].Metadata)

	// Verify acceleration event metadata
	require.JSONEq(t, `{"counterValue":10}`, accelerationEvents[0].Metadata)

	// Verify cornering event metadata
	require.JSONEq(t, `{"counterValue":7}`, corneringEvents[0].Metadata)
}

func TestDecodeEventBrakingOnly(t *testing.T) {
	t.Parallel()

	// Test data with only braking event (LSB only)
	brakingOnlyInputJSON := `{
		"id": "test-cloud-event-id",
		"source": "ruptela/test",
		"producer": "test-producer",
		"specversion": "1.0",
		"subject": "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:33",
		"time": "2024-09-27T08:33:26Z",
		"type": "dimo.event",
		"data": {
			"signals": {
				"135": "8",
				"136": "0",
				"143": "0"
			}
		}
	}`

	var event cloudevent.RawEvent
	err := json.Unmarshal([]byte(brakingOnlyInputJSON), &event)
	require.NoError(t, err)

	actualEvents, err := ruptela.DecodeEvent(event)
	require.NoError(t, err)
	require.Len(t, actualEvents, 1, "should have 1 braking event")

	evt := actualEvents[0]
	require.Equal(t, ruptela.EventNameHarshBraking, evt.Name)
	require.JSONEq(t, `{"counterValue":8}`, evt.Metadata)
}

func TestDecodeEventAccelerationOnly(t *testing.T) {
	t.Parallel()

	// Test data with only acceleration event
	accelerationOnlyInputJSON := `{
		"id": "test-cloud-event-id",
		"source": "ruptela/test",
		"producer": "test-producer",
		"specversion": "1.0",
		"subject": "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:33",
		"time": "2024-09-27T08:33:26Z",
		"type": "dimo.event",
		"data": {
			"signals": {
				"135": "0",
				"136": "F",
				"143": "0"
			}
		}
	}`

	var event cloudevent.RawEvent
	err := json.Unmarshal([]byte(accelerationOnlyInputJSON), &event)
	require.NoError(t, err)

	actualEvents, err := ruptela.DecodeEvent(event)
	require.NoError(t, err)
	require.Len(t, actualEvents, 1, "should have 1 acceleration event")

	evt := actualEvents[0]
	require.Equal(t, ruptela.EventNameAcceleration, evt.Name)
	require.JSONEq(t, `{"counterValue":15}`, evt.Metadata)
}

func TestDecodeEventCorneringOnly(t *testing.T) {
	t.Parallel()

	// Test data with only cornering event
	corneringOnlyInputJSON := `{
		"id": "test-cloud-event-id",
		"source": "ruptela/test",
		"producer": "test-producer",
		"specversion": "1.0",
		"subject": "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:33",
		"time": "2024-09-27T08:33:26Z",
		"type": "dimo.event",
		"data": {
			"signals": {
				"135": "0",
				"136": "0",
				"143": "C"
			}
		}
	}`

	var event cloudevent.RawEvent
	err := json.Unmarshal([]byte(corneringOnlyInputJSON), &event)
	require.NoError(t, err)

	actualEvents, err := ruptela.DecodeEvent(event)
	require.NoError(t, err)
	require.Len(t, actualEvents, 1, "should have 1 cornering event")

	evt := actualEvents[0]
	require.Equal(t, ruptela.EventNameCornering, evt.Name)
	require.JSONEq(t, `{"counterValue":12}`, evt.Metadata)
}

func TestDecodeEventNoEvents(t *testing.T) {
	t.Parallel()

	// Test data with no events (all zeros)
	noEventsInputJSON := `{
		"id": "test-cloud-event-id",
		"source": "ruptela/test",
		"producer": "test-producer",
		"specversion": "1.0",
		"subject": "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:33",
		"time": "2024-09-27T08:33:26Z",
		"type": "dimo.event",
		"data": {
			"signals": {
				"135": "0",
				"136": "0",
				"143": "0"
			}
		}
	}`

	var event cloudevent.RawEvent
	err := json.Unmarshal([]byte(noEventsInputJSON), &event)
	require.NoError(t, err)

	actualEvents, err := ruptela.DecodeEvent(event)
	require.NoError(t, err)
	require.Len(t, actualEvents, 0, "should have no events when all values are zero")
}

func TestDecodeEventBrakingBothNibbles(t *testing.T) {
	t.Parallel()

	// Test data with braking event in both nibbles (LSB and MSB)
	brakingBothNibblesInputJSON := `{
		"id": "test-cloud-event-id",
		"source": "ruptela/test",
		"producer": "test-producer",
		"specversion": "1.0",
		"subject": "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:33",
		"time": "2024-09-27T08:33:26Z",
		"type": "dimo.event",
		"data": {
			"signals": {
				"135": "A5",
				"136": "0",
				"143": "0"
			}
		}
	}`

	var event cloudevent.RawEvent
	err := json.Unmarshal([]byte(brakingBothNibblesInputJSON), &event)
	require.NoError(t, err)

	actualEvents, err := ruptela.DecodeEvent(event)
	require.NoError(t, err)
	require.Len(t, actualEvents, 2, "should have 2 braking events for both nibbles")

	// Separate the events by type
	var harshBrakingEvents []vss.Event
	var extremeBrakingEvents []vss.Event

	for _, evt := range actualEvents {
		switch evt.Name {
		case ruptela.EventNameHarshBraking:
			harshBrakingEvents = append(harshBrakingEvents, evt)
		case ruptela.EventNameExtremeBraking:
			extremeBrakingEvents = append(extremeBrakingEvents, evt)
		}
	}

	require.Len(t, harshBrakingEvents, 1, "should have 1 harsh braking event (LSB)")
	require.Len(t, extremeBrakingEvents, 1, "should have 1 extreme braking event (MSB)")

	// Check the counter values - one from LSB (bits 0-3) and one from MSB (bits 4-7)
	// For value "A5" (hex) = 165 (decimal) = 10100101 (binary)
	// LSB (bits 0-3): 0101 = 5
	// MSB (bits 4-7): 1010 = 10
	var harshMetadata ruptela.CounterMetadata
	err = json.Unmarshal([]byte(harshBrakingEvents[0].Metadata), &harshMetadata)
	require.NoError(t, err)
	require.Equal(t, uint(5), harshMetadata.CounterValue, "harsh braking should have counter value 5 from LSB")

	var extremeMetadata ruptela.CounterMetadata
	err = json.Unmarshal([]byte(extremeBrakingEvents[0].Metadata), &extremeMetadata)
	require.NoError(t, err)
	require.Equal(t, uint(10), extremeMetadata.CounterValue, "extreme braking should have counter value 10 from MSB")
}

func TestDecodeEventInvalidJSON(t *testing.T) {
	t.Parallel()

	// Test data with invalid JSON structure
	invalidJSONInputJSON := `{
		"id": "test-cloud-event-id",
		"source": "ruptela/test",
		"producer": "test-producer",
		"specversion": "1.0",
		"subject": "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:33",
		"time": "2024-09-27T08:33:26Z",
		"type": "dimo.event",
		"data": "invalid json structure"
	}`

	var event cloudevent.RawEvent
	err := json.Unmarshal([]byte(invalidJSONInputJSON), &event)
	require.NoError(t, err)

	_, err = ruptela.DecodeEvent(event)
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to unmarshal event data")
}

func TestDecodeEventHarshBrakingOnly(t *testing.T) {
	t.Parallel()

	// Test data with only harsh braking event (LSB only, MSB is 0)
	harshBrakingOnlyInputJSON := `{
		"id": "test-cloud-event-id",
		"source": "ruptela/test",
		"producer": "test-producer",
		"specversion": "1.0",
		"subject": "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:33",
		"time": "2024-09-27T08:33:26Z",
		"type": "dimo.event",
		"data": {
			"signals": {
				"135": "5",
				"136": "0",
				"143": "0"
			}
		}
	}`

	var event cloudevent.RawEvent
	err := json.Unmarshal([]byte(harshBrakingOnlyInputJSON), &event)
	require.NoError(t, err)

	actualEvents, err := ruptela.DecodeEvent(event)
	require.NoError(t, err)
	require.Len(t, actualEvents, 1, "should have 1 harsh braking event")

	evt := actualEvents[0]
	require.Equal(t, ruptela.EventNameHarshBraking, evt.Name)
	require.JSONEq(t, `{"counterValue":5}`, evt.Metadata)
}

func TestDecodeEventExtremeBrakingOnly(t *testing.T) {
	t.Parallel()

	// Test data with only extreme braking event (MSB only, LSB is 0)
	// Using "50" hex = 80 decimal = 01010000 binary (MSB=5, LSB=0)
	extremeBrakingOnlyInputJSON := `{
		"id": "test-cloud-event-id",
		"source": "ruptela/test",
		"producer": "test-producer",
		"specversion": "1.0",
		"subject": "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:33",
		"time": "2024-09-27T08:33:26Z",
		"type": "dimo.event",
		"data": {
			"signals": {
				"135": "50",
				"136": "0",
				"143": "0"
			}
		}
	}`

	var event cloudevent.RawEvent
	err := json.Unmarshal([]byte(extremeBrakingOnlyInputJSON), &event)
	require.NoError(t, err)

	actualEvents, err := ruptela.DecodeEvent(event)
	require.NoError(t, err)
	require.Len(t, actualEvents, 1, "should have 1 extreme braking event")

	evt := actualEvents[0]
	require.Equal(t, ruptela.EventNameExtremeBraking, evt.Name)
	require.JSONEq(t, `{"counterValue":5}`, evt.Metadata)
}

func TestDecodeEventPartialErrors(t *testing.T) {
	t.Parallel()

	// Test data with invalid hex value for braking but valid values for acceleration and cornering
	partialErrorsInputJSON := `{
		"id": "test-cloud-event-id",
		"source": "ruptela/test",
		"producer": "test-producer",
		"specversion": "1.0",
		"subject": "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:33",
		"time": "2024-09-27T08:33:26Z",
		"type": "dimo.event",
		"data": {
			"signals": {
				"135": "ZZ",
				"136": "A",
				"143": "7"
			}
		}
	}`

	var event cloudevent.RawEvent
	err := json.Unmarshal([]byte(partialErrorsInputJSON), &event)
	require.NoError(t, err)

	actualEvents, err := ruptela.DecodeEvent(event)

	// Should get events for acceleration and cornering, but error for braking
	require.Error(t, err)
	require.Contains(t, err.Error(), "could not parse uint")
	require.Len(t, actualEvents, 2, "should have 2 events despite braking error")

	// Verify the successful events
	var accelerationEvents []vss.Event
	var corneringEvents []vss.Event

	for _, evt := range actualEvents {
		switch evt.Name {
		case ruptela.EventNameAcceleration:
			accelerationEvents = append(accelerationEvents, evt)
		case ruptela.EventNameCornering:
			corneringEvents = append(corneringEvents, evt)
		}
	}

	require.Len(t, accelerationEvents, 1, "should have 1 acceleration event")
	require.Len(t, corneringEvents, 1, "should have 1 cornering event")

	require.JSONEq(t, `{"counterValue":10}`, accelerationEvents[0].Metadata)
	require.JSONEq(t, `{"counterValue":7}`, corneringEvents[0].Metadata)
}
