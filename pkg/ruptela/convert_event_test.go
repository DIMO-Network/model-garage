package ruptela_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/convert"
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
		"type": "dimo.status",
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
	seenIDs := map[string]bool{}
	for _, evt := range actualEvents {
		require.Equal(t, "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:33", evt.Subject)
		require.Equal(t, "ruptela/test", evt.Source)
		require.Equal(t, "test-producer", evt.Producer)
		require.NotEmpty(t, evt.ID, "event ID should be a non-empty ksuid")
		require.False(t, seenIDs[evt.ID], "event IDs should be unique")
		seenIDs[evt.ID] = true
		require.Equal(t, cloudevent.TypeEvent, evt.Type)
		require.False(t, evt.Time.IsZero(), "event Time should be set to a recent time")
		require.Equal(t, "test-cloud-event-id", evt.Data.CloudEventID)
		require.Equal(t, time.Date(2024, 9, 27, 8, 33, 26, 0, time.UTC), evt.Data.Timestamp)
	}

	// Check specific events
	var harshBrakingEvents []vss.Event
	var extremeBrakingEvents []vss.Event
	var accelerationEvents []vss.Event
	var corneringEvents []vss.Event

	for _, evt := range actualEvents {
		switch evt.Data.Name {
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
	require.JSONEq(t, `{"counterValue":5}`, harshBrakingEvents[0].Data.Metadata)
	require.JSONEq(t, `{"counterValue":3}`, extremeBrakingEvents[0].Data.Metadata)

	// Verify acceleration event metadata
	require.JSONEq(t, `{"counterValue":10}`, accelerationEvents[0].Data.Metadata)

	// Verify cornering event metadata
	require.JSONEq(t, `{"counterValue":7}`, corneringEvents[0].Data.Metadata)
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
		"type": "dimo.status",
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
	require.NotEmpty(t, evt.ID, "event ID should be a non-empty ksuid")
	require.Equal(t, cloudevent.TypeEvent, evt.Type)
	require.False(t, evt.Time.IsZero(), "event Time should be set to a recent time")
	require.Equal(t, "test-cloud-event-id", evt.Data.CloudEventID)
	require.Equal(t, time.Date(2024, 9, 27, 8, 33, 26, 0, time.UTC), evt.Data.Timestamp)
	require.Equal(t, ruptela.EventNameHarshBraking, evt.Data.Name)
	require.JSONEq(t, `{"counterValue":8}`, evt.Data.Metadata)
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
		"type": "dimo.status",
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
	require.NotEmpty(t, evt.ID, "event ID should be a non-empty ksuid")
	require.Equal(t, cloudevent.TypeEvent, evt.Type)
	require.False(t, evt.Time.IsZero(), "event Time should be set to a recent time")
	require.Equal(t, "test-cloud-event-id", evt.Data.CloudEventID)
	require.Equal(t, time.Date(2024, 9, 27, 8, 33, 26, 0, time.UTC), evt.Data.Timestamp)
	require.Equal(t, ruptela.EventNameAcceleration, evt.Data.Name)
	require.JSONEq(t, `{"counterValue":15}`, evt.Data.Metadata)
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
		"type": "dimo.status",
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
	require.NotEmpty(t, evt.ID, "event ID should be a non-empty ksuid")
	require.Equal(t, cloudevent.TypeEvent, evt.Type)
	require.False(t, evt.Time.IsZero(), "event Time should be set to a recent time")
	require.Equal(t, "test-cloud-event-id", evt.Data.CloudEventID)
	require.Equal(t, time.Date(2024, 9, 27, 8, 33, 26, 0, time.UTC), evt.Data.Timestamp)
	require.Equal(t, ruptela.EventNameCornering, evt.Data.Name)
	require.JSONEq(t, `{"counterValue":12}`, evt.Data.Metadata)
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
		"type": "dimo.status",
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
		"type": "dimo.status",
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
		"type": "dimo.status",
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
	require.NotEmpty(t, evt.ID, "event ID should be a non-empty ksuid")
	require.Equal(t, cloudevent.TypeEvent, evt.Type)
	require.False(t, evt.Time.IsZero(), "event Time should be set to a recent time")
	require.Equal(t, "test-cloud-event-id", evt.Data.CloudEventID)
	require.Equal(t, time.Date(2024, 9, 27, 8, 33, 26, 0, time.UTC), evt.Data.Timestamp)
	require.Equal(t, ruptela.EventNameHarshBraking, evt.Data.Name)
	require.JSONEq(t, `{"counterValue":5}`, evt.Data.Metadata)
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
		"type": "dimo.status",
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
	require.NotEmpty(t, evt.ID, "event ID should be a non-empty ksuid")
	require.Equal(t, cloudevent.TypeEvent, evt.Type)
	require.False(t, evt.Time.IsZero(), "event Time should be set to a recent time")
	require.Equal(t, "test-cloud-event-id", evt.Data.CloudEventID)
	require.Equal(t, time.Date(2024, 9, 27, 8, 33, 26, 0, time.UTC), evt.Data.Timestamp)
	require.Equal(t, ruptela.EventNameExtremeBraking, evt.Data.Name)
	require.JSONEq(t, `{"counterValue":5}`, evt.Data.Metadata)
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
		"type": "dimo.status",
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
	require.Len(t, actualEvents, 0, "should not have any events")

	var conversionError convert.ConversionError
	require.ErrorAs(t, err, &conversionError)
	require.Equal(t, "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:33", conversionError.Subject)
	require.Equal(t, "ruptela/test", conversionError.Source)
	require.Len(t, conversionError.DecodedEvents, 2, "should have 2 events")

	// Verify the successful events
	var accelerationEvents []vss.Event
	var corneringEvents []vss.Event

	for _, evt := range conversionError.DecodedEvents {
		switch evt.Data.Name {
		case ruptela.EventNameAcceleration:
			accelerationEvents = append(accelerationEvents, evt)
		case ruptela.EventNameCornering:
			corneringEvents = append(corneringEvents, evt)
		}
	}

	require.Len(t, accelerationEvents, 1, "should have 1 acceleration event")
	require.Len(t, corneringEvents, 1, "should have 1 cornering event")

	// Verify dynamic fields on decoded events from the conversion error
	for _, evt := range conversionError.DecodedEvents {
		require.NotEmpty(t, evt.ID, "event ID should be a non-empty ksuid")
		require.Equal(t, cloudevent.TypeEvent, evt.Type)
		require.False(t, evt.Time.IsZero(), "event Time should be set to a recent time")
		require.Equal(t, "test-cloud-event-id", evt.Data.CloudEventID)
		require.Equal(t, time.Date(2024, 9, 27, 8, 33, 26, 0, time.UTC), evt.Data.Timestamp)
	}

	require.JSONEq(t, `{"counterValue":10}`, accelerationEvents[0].Data.Metadata)
	require.JSONEq(t, `{"counterValue":7}`, corneringEvents[0].Data.Metadata)
}

func TestDecodeEventMissingSignals(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		inputJSON      string
		expectedEvents int
		expectedNames  []string
	}{
		{
			name: "missing signal 135 (braking)",
			inputJSON: `{
				"id": "test-cloud-event-id",
				"source": "ruptela/test",
				"producer": "test-producer",
				"specversion": "1.0",
				"subject": "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:33",
				"time": "2024-09-27T08:33:26Z",
				"type": "dimo.status",
				"data": {
					"signals": {
						"136": "0A",
						"143": "07"
					}
				}
			}`,
			expectedEvents: 2,
			expectedNames:  []string{ruptela.EventNameAcceleration, ruptela.EventNameCornering},
		},
		{
			name: "missing signals 136 and 143 (acceleration and cornering)",
			inputJSON: `{
				"id": "test-cloud-event-id",
				"source": "ruptela/test",
				"producer": "test-producer",
				"specversion": "1.0",
				"subject": "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:33",
				"time": "2024-09-27T08:33:26Z",
				"type": "dimo.status",
				"data": {
					"signals": {
						"135": "35"
					}
				}
			}`,
			expectedEvents: 2,
			expectedNames:  []string{ruptela.EventNameHarshBraking, ruptela.EventNameExtremeBraking},
		},
		{
			name: "all signals missing",
			inputJSON: `{
				"id": "test-cloud-event-id",
				"source": "ruptela/test",
				"producer": "test-producer",
				"specversion": "1.0",
				"subject": "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:33",
				"time": "2024-09-27T08:33:26Z",
				"type": "dimo.status",
				"data": {
					"signals": {
						"100": "0000",
						"101": "00"
					}
				}
			}`,
			expectedEvents: 0,
			expectedNames:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var event cloudevent.RawEvent
			err := json.Unmarshal([]byte(tt.inputJSON), &event)
			require.NoError(t, err)

			actualEvents, err := ruptela.DecodeEvent(event)
			require.NoError(t, err, "should not error when signals are missing")
			require.Len(t, actualEvents, tt.expectedEvents)

			if tt.expectedNames != nil {
				actualNames := make([]string, len(actualEvents))
				for i, evt := range actualEvents {
					actualNames[i] = evt.Data.Name
				}
				require.ElementsMatch(t, tt.expectedNames, actualNames)
			}

			// Verify dynamic fields on all returned events
			for _, evt := range actualEvents {
				require.NotEmpty(t, evt.ID, "event ID should be a non-empty ksuid")
				require.Equal(t, cloudevent.TypeEvent, evt.Type)
				require.False(t, evt.Time.IsZero(), "event Time should be set to a recent time")
				require.Equal(t, "test-cloud-event-id", evt.Data.CloudEventID)
				require.Equal(t, time.Date(2024, 9, 27, 8, 33, 26, 0, time.UTC), evt.Data.Timestamp)
			}
		})
	}
}

func TestDecodeEventEngineSecurityEvents(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		engineValue  string
		expectedName string
		expectedTags []string
	}{
		{
			name:         "engine block",
			engineValue:  "1",
			expectedName: ruptela.EventNameEngineBlock,
			expectedTags: []string{},
		},
		{
			name:         "engine unblock",
			engineValue:  "0",
			expectedName: ruptela.EventNameEngineUnblock,
			expectedTags: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			inputJSON := `{
				"id": "test-cloud-event-id",
				"source": "ruptela/test",
				"producer": "test-producer",
				"specversion": "1.0",
				"subject": "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:33",
				"time": "2024-09-27T08:33:26Z",
				"type": "dimo.status",
				"dataversion": "r/v0/cmd",
				"data": {
					"signals": {
						"135": "0",
						"136": "0",
						"143": "0",
						"405": "` + tt.engineValue + `"
					}
				}
			}`

			var event cloudevent.RawEvent
			err := json.Unmarshal([]byte(inputJSON), &event)
			require.NoError(t, err)

			actualEvents, err := ruptela.DecodeEvent(event)
			require.NoError(t, err)
			require.Len(t, actualEvents, 1)

			evt := actualEvents[0]
			require.Equal(t, tt.expectedName, evt.Data.Name)
			require.Equal(t, tt.expectedTags, evt.Data.Tags)
			require.Equal(t, "test-cloud-event-id", evt.Data.CloudEventID)
			require.Equal(t, time.Date(2024, 9, 27, 8, 33, 26, 0, time.UTC), evt.Data.Timestamp)
		})
	}
}

func TestDecodeEventEngineSecuritySignalEmptyIsIgnored(t *testing.T) {
	t.Parallel()

	inputJSON := `{
		"id": "test-cloud-event-id",
		"source": "ruptela/test",
		"producer": "test-producer",
		"specversion": "1.0",
		"subject": "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:33",
		"time": "2024-09-27T08:33:26Z",
		"type": "dimo.status",
		"dataversion": "r/v0/cmd",
		"data": {
			"signals": {
				"135": "0",
				"136": "0",
				"143": "0",
				"405": ""
			}
		}
	}`

	var event cloudevent.RawEvent
	err := json.Unmarshal([]byte(inputJSON), &event)
	require.NoError(t, err)

	actualEvents, err := ruptela.DecodeEvent(event)
	require.NoError(t, err)
	require.Len(t, actualEvents, 0)
}

func TestDecodeEventEngineSecurityIgnoredOnNonCmdTopic(t *testing.T) {
	t.Parallel()

	inputJSON := `{
		"id": "test-cloud-event-id",
		"source": "ruptela/test",
		"producer": "test-producer",
		"specversion": "1.0",
		"subject": "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:33",
		"time": "2024-09-27T08:33:26Z",
		"type": "dimo.status",
		"dataversion": "r/v0/s",
		"data": {
			"signals": {
				"135": "0",
				"136": "0",
				"143": "0",
				"405": "1"
			}
		}
	}`

	var event cloudevent.RawEvent
	err := json.Unmarshal([]byte(inputJSON), &event)
	require.NoError(t, err)

	actualEvents, err := ruptela.DecodeEvent(event)
	require.NoError(t, err)
	require.Len(t, actualEvents, 0, "engine block signal should be ignored on non-cmd topics")
}

func TestDecodeEventEngineSecurityPartialError(t *testing.T) {
	t.Parallel()

	inputJSON := `{
		"id": "test-cloud-event-id",
		"source": "ruptela/test",
		"producer": "test-producer",
		"specversion": "1.0",
		"subject": "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:33",
		"time": "2024-09-27T08:33:26Z",
		"type": "dimo.status",
		"dataversion": "r/v0/cmd",
		"data": {
			"signals": {
				"135": "0",
				"136": "A",
				"143": "0",
				"405": "ZZ"
			}
		}
	}`

	var event cloudevent.RawEvent
	err := json.Unmarshal([]byte(inputJSON), &event)
	require.NoError(t, err)

	actualEvents, err := ruptela.DecodeEvent(event)
	require.Error(t, err)
	require.Contains(t, err.Error(), "could not parse uint")
	require.Len(t, actualEvents, 0)

	var conversionError convert.ConversionError
	require.ErrorAs(t, err, &conversionError)
	require.Len(t, conversionError.DecodedEvents, 1)
	require.Equal(t, ruptela.EventNameAcceleration, conversionError.DecodedEvents[0].Data.Name)
	require.JSONEq(t, `{"counterValue":10}`, conversionError.DecodedEvents[0].Data.Metadata)
}

func TestToEngineSecurityEvent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		rawValue     string
		expectedName string
		expectedTags []string
		expectErr    bool
	}{
		{
			name:         "zero value maps to unblock",
			rawValue:     "0",
			expectedName: ruptela.EventNameEngineUnblock,
			expectedTags: nil,
			expectErr:    false,
		},
		{
			name:         "non-zero value maps to block",
			rawValue:     "A",
			expectedName: ruptela.EventNameEngineBlock,
			expectedTags: nil,
			expectErr:    false,
		},
		{
			name:      "invalid value returns error",
			rawValue:  "ZZ",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual, err := ruptela.ToEngineSecurityEvent(tt.rawValue)
			if tt.expectErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), "could not parse uint from engine security event")
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.expectedName, actual.Name)
			require.Equal(t, tt.expectedTags, actual.Tags)
		})
	}
}
