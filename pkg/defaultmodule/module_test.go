package defaultmodule_test

import (
	"context"
	"encoding/json"
	"strconv"
	"testing"
	"time"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/defaultmodule"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestModule_CloudEventConvert(t *testing.T) {
	tests := []struct {
		name              string
		input             []byte
		expectHeaderTypes []string
		expectErr         bool
	}{
		{
			name: "valid cloud event with vin",
			input: []byte(`{
				"id": "2pcYwspbaBFJ7NPGZ2kivkuJ12a",
				"source": "0x1234567890123456789012345678901234567890",
				"producer": "did:erc721:80003:0x78513c8CB4D6B6079f813850376bc9c7fc8aE67f:12",
				"specversion": "1.0",
				"subject": "did:erc721:80003:0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8:15",
				"time": "2024-12-01T15:31:12.378075897Z",
				"type": "dimo.status",
				"signature": "0x22cca92bb6a16fed01def56d02541427633ff82552bc8c5c2da2fffd69c4436927b256ab0f1352e584deb5394fff2f979699f206691f73fffee547cee1431c",
				"data": {
					"signals": [
						{
							"name": "speed",
							"value": 100,
							"timestamp": "2024-12-01T15:31:12.378075897Z"
						}
					],
					"vin": "HYBRID"
				}
			}`),
			expectHeaderTypes: []string{cloudevent.TypeFingerprint, cloudevent.TypeStatus},
			expectErr:         false,
		},
		{
			name: "valid cloud event without vin",
			input: []byte(`{
				"id": "2pcYwspbaBFJ7NPGZ2kivkuJ12a",
				"source": "0x1234567890123456789012345678901234567890",
				"producer": "did:erc721:80003:0x78513c8CB4D6B6079f813850376bc9c7fc8aE67f:12",
				"specversion": "1.0",
				"subject": "did:erc721:80003:0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8:15",
				"time": "2024-12-01T15:31:12.378075897Z",
				"type": "dimo.status",
				"data": {
					"signals": [
						{
							"name": "speed",
							"value": 100,
							"timestamp": "2024-12-01T15:31:12.378075897Z"
						}
					]
				}
			}`),
			expectHeaderTypes: []string{cloudevent.TypeStatus},
			expectErr:         false,
		},
		{
			name:              "invalid json",
			input:             []byte(`{invalid`),
			expectHeaderTypes: nil,
			expectErr:         true,
		},
		{
			name: "unexpected data field",
			input: []byte(`{
				"id": "2pcYwspbaBFJ7NPGZ2kivkuJ12a",
				"source": "0xFFEE022fAb46610EAFe98b87377B42e366364a71",
				"producer": "did:erc721:80003:0x78513c8CB4D6B6079f813850376bc9c7fc8aE67f:12",
				"specversion": "1.0",
				"type": "dimo.status",
				"data": "not an object"
			}`),
			expectHeaderTypes: []string{cloudevent.TypeUnknown},
			expectErr:         false,
		},
	}

	module := defaultmodule.Module{}
	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers, data, err := module.CloudEventConvert(ctx, tt.input)

			if tt.expectErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Len(t, headers, len(tt.expectHeaderTypes))
			for i, header := range headers {
				assert.Equal(t, tt.expectHeaderTypes[i], header.Type)
			}
			assert.NotNil(t, data)

			if len(tt.expectHeaderTypes) > 1 {
				first := headers[0]
				for i := 1; i < len(headers); i++ {
					// Verify original header fields are preserved
					assert.Equal(t, first.Source, headers[i].Source)
					assert.Equal(t, first.Subject, headers[i].Subject)
					assert.Equal(t, first.Producer, headers[i].Producer)
				}
			}
		})
	}
}

func TestModule_SignalConvert(t *testing.T) {
	tests := []struct {
		name            string
		input           []byte
		expectedSignals []vss.Signal
		expectError     bool
	}{
		{
			name: "valid signal data with numeric value",
			input: []byte(`{
				"id": "2pcYwspbaBFJ7NPGZ2kivkuJ12a",
				"source": "0x1234567890123456789012345678901234567890",
				"producer": "did:erc721:80003:0x78513c8CB4D6B6079f813850376bc9c7fc8aE67f:12",
				"specversion": "1.0",
				"subject": "did:erc721:80003:0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8:15",
				"time": "2024-12-01T15:31:12.378075897Z",
				"type": "dimo.status",
				"data": {
					"signals": [
						{
							"timestamp": "2024-12-01T15:31:12.378075897Z",
							"name": "speed",
							"value": 100.5
						},
						{
							"timestamp": "2024-12-01T15:31:12.378075897Z",
							"name": "speed",
							"value": "102"
						},
						{
							"timestamp": "2024-12-01T15:31:12.378075897Z",
							"name": "speed",
							"value": 100
						},
						{
							"timestamp": "2024-12-01T15:31:12.378075897Z",
							"name": "speed",
							"value": 0
						}
					]
				}
			}`),
			expectedSignals: []vss.Signal{
				{
					Name:        "speed",
					ValueNumber: 100.5,
					Timestamp:   time.Date(2024, 12, 1, 15, 31, 12, 378075897, time.UTC),
				},
				{
					Name:        "speed",
					ValueNumber: 102,
					Timestamp:   time.Date(2024, 12, 1, 15, 31, 12, 378075897, time.UTC),
				},
				{
					Name:        "speed",
					ValueNumber: 100,
					Timestamp:   time.Date(2024, 12, 1, 15, 31, 12, 378075897, time.UTC),
				},
				{
					Name:        "speed",
					ValueNumber: 0,
					Timestamp:   time.Date(2024, 12, 1, 15, 31, 12, 378075897, time.UTC),
				},
			},
			expectError: false,
		},
		{
			name: "valid signal data with string value",
			input: []byte(`{
				"id": "2pcYwspbaBFJ7NPGZ2kivkuJ12a",
				"source": "0x1234567890123456789012345678901234567890",
				"producer": "did:erc721:80003:0x78513c8CB4D6B6079f813850376bc9c7fc8aE67f:12",
				"specversion": "1.0",
				"subject": "did:erc721:80003:0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8:15",
				"time": "2024-12-01T15:31:12.378075897Z",
				"type": "dimo.status",
				"data": {
					"signals": [
						{
							"timestamp": "2024-12-01T15:31:12.378075897Z",
							"name": "powertrainType",
							"value": "HYBRID"
						}
					]
				}
			}`),
			expectedSignals: []vss.Signal{
				{
					Name:        "powertrainType",
					ValueString: "HYBRID",
					Timestamp:   time.Date(2024, 12, 1, 15, 31, 12, 378075897, time.UTC),
				},
			},
			expectError: false,
		},
		{
			name: "multiple valid signals",
			input: []byte(`{
				"id": "2pcYwspbaBFJ7NPGZ2kivkuJ12a",
				"source": "0x1234567890123456789012345678901234567890",
				"producer": "did:erc721:80003:0x78513c8CB4D6B6079f813850376bc9c7fc8aE67f:12",
				"specversion": "1.0",
				"subject": "did:erc721:80003:0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8:15",
				"time": "2024-12-01T15:31:12.378075897Z",
				"type": "dimo.status",
				"data": {
					"signals": [
						{
							"timestamp": "2024-12-01T15:31:12.378075897Z",
							"name": "speed",
							"value": 100.5
						},
						{
							"timestamp": "2024-12-01T15:31:12.378075897Z",
							"name": "powertrainType",
							"value": "HYBRID"
						},
						{
							"timestamp": "2024-12-01T15:31:12.378075897Z",
							"name": "currentLocationCoordinates",
							"value": {
								"latitude": 37.7749,
								"longitude": -23.45833,
								"hdop": 3
							}
						}
					]
				}
			}`),
			expectedSignals: []vss.Signal{
				{
					Name:        "speed",
					ValueNumber: 100.5,
					Timestamp:   time.Date(2024, 12, 1, 15, 31, 12, 378075897, time.UTC),
				},
				{
					Name:        "powertrainType",
					ValueString: "HYBRID",
					Timestamp:   time.Date(2024, 12, 1, 15, 31, 12, 378075897, time.UTC),
				},
				{
					Name: "currentLocationCoordinates",
					ValueLocation: vss.Location{
						Latitude:  37.7749,
						Longitude: -23.45833,
						HDOP:      3,
					},
					Timestamp: time.Date(2024, 12, 1, 15, 31, 12, 378075897, time.UTC),
				},
			},
			expectError: false,
		},
		{
			name: "undefined signal name",
			input: []byte(`{
				"id": "2pcYwspbaBFJ7NPGZ2kivkuJ12a",
				"source": "0x1234567890123456789012345678901234567890",
				"producer": "did:erc721:80003:0x78513c8CB4D6B6079f813850376bc9c7fc8aE67f:12",
				"specversion": "1.0",
				"subject": "did:erc721:80003:0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8:15",
				"time": "2024-12-01T15:31:12.378075897Z",
				"type": "dimo.status",
				"data": {
					"signals": [
						{
							"timestamp": "2024-12-01T15:31:12.378075897Z",
							"name": "UndefinedSignal",
							"value": 100.5
						}
					]
				}
			}`),
			expectedSignals: []vss.Signal{},
			expectError:     true,
		},
		{
			name: "invalid number for numeric type",
			input: []byte(`{
				"id": "2pcYwspbaBFJ7NPGZ2kivkuJ12a",
				"source": "0x1234567890123456789012345678901234567890",
				"producer": "did:erc721:80003:0x78513c8CB4D6B6079f813850376bc9c7fc8aE67f:12",
				"specversion": "1.0",
				"subject": "did:erc721:80003:0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8:15",
				"time": "2024-12-01T15:31:12.378075897Z",
				"type": "dimo.status",
				"data": {
					"signals": [
						{
							"timestamp": "2024-12-01T15:31:12.378075897Z",
							"name": "speed",
							"value": "not-a-number"
						}
					]
				}
			}`),
			expectedSignals: []vss.Signal{},
			expectError:     true,
		},
		{
			name: "missing value",
			input: []byte(`{
				"id": "2pcYwspbaBFJ7NPGZ2kivkuJ12a",
				"source": "0x1234567890123456789012345678901234567890",
				"producer": "did:erc721:80003:0x78513c8CB4D6B6079f813850376bc9c7fc8aE67f:12",
				"specversion": "1.0",
				"subject": "did:erc721:80003:0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8:15",
				"time": "2024-12-01T15:31:12.378075897Z",
				"type": "dimo.status",
				"data": {
					"signals": [
						{
							"timestamp": "2024-12-01T15:31:12.378075897Z",
							"name": "speed"
						}
					]
				}
			}`),
			expectedSignals: []vss.Signal{},
			expectError:     true,
		},
		{
			name: "mixed valid and invalid signals",
			input: []byte(`{
				"id": "2pcYwspbaBFJ7NPGZ2kivkuJ12a",
				"source": "0x1234567890123456789012345678901234567890",
				"producer": "did:erc721:80003:0x78513c8CB4D6B6079f813850376bc9c7fc8aE67f:12",
				"specversion": "1.0",
				"subject": "did:erc721:80003:0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8:15",
				"time": "2024-12-01T15:31:12.378075897Z",
				"type": "dimo.status",
				"data": {
					"signals": [
						{
							"timestamp": "2024-12-01T15:31:12.378075897Z",
							"name": "speed",
							"value": "100.5"
						},
						{
							"timestamp": "2024-12-01T15:31:12.378075897Z",
							"name": "UndefinedSignal",
							"value": "value"
						}
					]
				}
			}`),
			expectedSignals: []vss.Signal{
				{
					Name:        "speed",
					ValueNumber: 100.5,
					Timestamp:   time.Date(2024, 12, 1, 15, 31, 12, 378075897, time.UTC),
				},
			},
			expectError: true,
		},
	}

	module := defaultmodule.Module{}
	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var event cloudevent.RawEvent
			err := json.Unmarshal(tt.input, &event)
			require.NoError(t, err, "Failed to unmarshal input")
			signals, err := module.SignalConvert(ctx, event)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// Check signal count
			assert.Len(t, signals, len(tt.expectedSignals))

			// Check signal details
			for i, expectedSignal := range tt.expectedSignals {
				if i < len(signals) {
					assert.Equal(t, expectedSignal.Name, signals[i].Name)

					// Check the appropriate value based on type
					if expectedSignal.ValueNumber != 0 {
						assert.Equal(t, expectedSignal.ValueNumber, signals[i].ValueNumber)
					} else if expectedSignal.ValueString != "" {
						assert.Equal(t, expectedSignal.ValueString, signals[i].ValueString)
					}

					// Check timestamp - using WithinDuration to avoid any potential timestamp precision issues
					assert.WithinDuration(t, expectedSignal.Timestamp, signals[i].Timestamp, time.Millisecond)
				}
			}
		})
	}
}

func TestModule_EventConvert(t *testing.T) {
	module := defaultmodule.Module{}
	ctx := context.Background()

	eventTs := time.Now().Truncate(time.Millisecond).UTC()
	cloudEventTs := eventTs.Add(-time.Minute)

	tests := []struct {
		name           string
		inputJSON      string
		expectedEvents []vss.Event
		expectError    bool
	}{
		{
			name: "valid events",
			inputJSON: `{
				"id": "evt-1",
				"source": "test-source",
				"producer": "test-producer",
				"subject": "test-subject",
				"specversion": "1.0",
				"time": "` + cloudEventTs.Format(time.RFC3339Nano) + `",
				"type": "dimo.event",
				"data": {
					"events": [
						{
							"name": "harsh_braking",
							"timestamp": "` + eventTs.Format(time.RFC3339Nano) + `",
							"durationNs": 0,
							"metadata": "{\"side\":\"left\"}",
							"tags": ["` + vss.TagBehaviorHarshAcceleration + `"]
						},
						{
							"name": "charging_stopped",
							"timestamp": "` + eventTs.Format(time.RFC3339Nano) + `",
							"durationNs": ` + strconv.Itoa(int((time.Second * 5).Nanoseconds())) + `,
							"metadata": "{\"temp\":72}"
						}
					]
				}
			}`,
			expectedEvents: []vss.Event{
				{
					Name:       "harsh_braking",
					Timestamp:  eventTs,
					DurationNs: 0,
					Metadata:   `{"side":"left"}`,
				},
				{
					Name:       "charging_stopped",
					Timestamp:  eventTs,
					DurationNs: uint64(5 * time.Second.Nanoseconds()),
					Metadata:   `{"temp":72}`,
				},
			},
			expectError: false,
		},
		{
			name: "invalid metadata",
			inputJSON: `{
				"id": "evt-2",
				"source": "test-source",
				"producer": "test-producer",
				"subject": "test-subject",
				"specversion": "1.0",
				"time": "` + cloudEventTs.Format(time.RFC3339Nano) + `",
				"type": "dimo.event",
				"data": {
					"events": [
						{
							"name": "bad_event",
							"timestamp": "` + eventTs.Format(time.RFC3339Nano) + `",
							"durationNs": 0,
							"metadata": "{\"bad\":}"
						}
					]
				}
			}`,
			expectedEvents: nil,
			expectError:    true,
		},
		{
			name: "empty events",
			inputJSON: `{
				"id": "evt-3",
				"source": "test-source",
				"producer": "test-producer",
				"subject": "test-subject",
				"specversion": "1.0",
				"time": "` + cloudEventTs.Format(time.RFC3339Nano) + `",
				"type": "dimo.event",
				"data": {
					"events": []
				}
			}`,
			expectedEvents: nil,
			expectError:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rawEvent cloudevent.RawEvent
			require.NoError(t, json.Unmarshal([]byte(tt.inputJSON), &rawEvent))
			events, err := module.EventConvert(ctx, rawEvent)
			if tt.expectError {
				assert.Error(t, err)
				assert.Len(t, events, 0)
			} else {
				assert.NoError(t, err)
				assert.Len(t, events, len(tt.expectedEvents))
				for i, expected := range tt.expectedEvents {
					assert.Equal(t, expected.Name, events[i].Name)
					assert.Equal(t, expected.Timestamp, events[i].Timestamp)
					assert.Equal(t, expected.DurationNs, events[i].DurationNs)
					assert.Equal(t, expected.Metadata, events[i].Metadata)
				}
			}
		})
	}
}
