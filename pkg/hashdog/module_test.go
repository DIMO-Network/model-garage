package hashdog

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/stretchr/testify/require"
)

func TestSignalConvert(t *testing.T) {
	ts := time.Unix(1727360340, 0).UTC()

	// Signal payload data
	signalData := `{
    "vehicle": {		
         "signals": [
			 {
                    "timestamp": 1727360340000,
                    "name": "longTermFuelTrim1",
                    "value": 25
                },
                {
                    "timestamp": 1727360340000,
                    "name": "coolantTemp",
                    "value": 107
                }
		]
    }}`

	const source = "dimo/integration/2ULfuC8U9dOqRshZBAi0lMM1Rrx"
	tests := []struct {
		name            string
		cloudEvent      cloudevent.CloudEvent[json.RawMessage]
		expectedSignals []vss.Signal
		expectedError   error
	}{
		{
			name: "Valid Signal Payload",
			cloudEvent: cloudevent.CloudEvent[json.RawMessage]{
				CloudEventHeader: cloudevent.CloudEventHeader{
					DataVersion: DataVersion,
					Type:        cloudevent.TypeStatus,
					Source:      source,
					Subject:     "did:erc721:1:0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8:33",
					Time:        ts,
				},
				Data: json.RawMessage(signalData),
			},
			expectedSignals: []vss.Signal{
				{TokenID: 33, Timestamp: ts, Name: vss.FieldOBDLongTermFuelTrim1, ValueNumber: 25, Source: source},
				{TokenID: 33, Timestamp: ts, Name: vss.FieldPowertrainCombustionEngineECT, ValueNumber: 107, Source: source},
			},
			expectedError: nil,
		},
		{
			name: "Device Status Payload",
			cloudEvent: cloudevent.CloudEvent[json.RawMessage]{
				CloudEventHeader: cloudevent.CloudEventHeader{
					DataVersion: DataVersion,
					Type:        cloudevent.TypeStatus,
					Source:      source,
					Subject:     "did:erc721:1:0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8:33",
					Producer:    "did:erc721:1:0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8:33",
					Time:        ts,
				},
				Data: json.RawMessage(signalData),
			},
			expectedSignals: nil,
			expectedError:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the SignalConvert function
			module := Module{}
			signals, err := module.SignalConvert(context.Background(), tt.cloudEvent)

			if tt.expectedError != nil {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expectedSignals, signals)
			}
		})
	}
}

func TestFingerprintConvert(t *testing.T) {
	fingerPrintData := `{"id":"9xYzA8bCdEf2GhIj3KlMnOpQ7rS","source":"0x4c674ddE8189aEF6e3b58F5a36d7438b2b1f6Bc2","producer":"did:erc721:137:0x8a92B34cDeFg1H2i3J4k5L6m7N8o9P0qRsTuV_45678","specversion":"1.0","subject":"did:nft:137:0xAb12Cd34Ef56Gh78Ij90Kl12Mn34Op56Qr78St_12345","type":"dimo.fingerprint","data":{"decodedPayload":{"data_base64":"Abc123XyZaBcDeF0987654321HiJkLmNoPqRsTuVwXyZ","header":1,"latitude":12.345678,"longitude":-98.765432,"nsat":1,"protocol":6,"signature":"0x9a8b7c6d5e4f3g2h1i0j9k8l7m6n5o4p3q2r1s0t9u8v7w6x5y4z3a2b1c0d9e8f7g6h5i4j3k2l1m","timestamp":"2025-03-05T12:46:32.000Z","vin":"1ABCD2EFGH3JKLMNO"},"device":{"id":"A1B2C3D4E5F6G7H8","name":"0x1a2B3c4D5e6F7g8H9i0J1k2L3m4N5o6P7q8R9s","protocol":"lora_helium","tags":{"env":"prod","label":"prod"}},"header":1,"id":"a1b2c3d4-e5f6-g7h8-i9j0-k1l2m3n4o5p6","metadata":{"devAddr":"123456ab","fPort":"2","fcnt":"2627"},"payload":"Abc123XyZaBcDeF0987654321HiJkLmNoPqRsTuVwXyZ9a8b7c6d5e4f3g2h1i0j9k8l7m6n5o4p3q2r1s0t9u8v7w6x5y4z3a2b1c0d9e8f7g6h5i4j3k2l1m","timestamp":1741178792000,"vehicle":{"signals":[{"name":"data_base64","timestamp":1741178792000,"value":"Abc123XyZaBcDeF0987654321HiJkLmNoPqRsTuVwXyZ"},{"name":"latitude","timestamp":1741178792000,"value":12.345678},{"name":"longitude","timestamp":1741178792000,"value":-98.765432},{"name":"nsat","timestamp":1741178792000,"value":1},{"name":"protocol","timestamp":1741178792000,"value":6},{"name":"signature","timestamp":1741178792000,"value":"0x9a8b7c6d5e4f3g2h1i0j9k8l7m6n5o4p3q2r1s0t9u8v7w6x5y4z3a2b1c0d9e8f7g6h5i4j3k2l1m"},{"name":"vin","timestamp":1741178792000,"value":"1ABCD2EFGH3JKLMNO"}]},"via":[{"id":"1a2B3c4D5e6F7g8H9i0J1k2L3m4N5o6P7q8R9s0T1u2V3w4X5y6Z","location":{"latitude":12.345678,"longitude":-98.765432,"ref":"1a2B3c4D5e6F7g8H9i0J1k2L3m4N5o6P7q8R9s0T1u2V3w4X5y6Z","rssi":-110,"snr":5.2},"metadata":{"gatewayId":"1a2B3c4D5e6F7g8H9i0J1k2L3m4N5o6P7q8R9s0T1u2V3w4X5y6Z","gatewayName":"random-scrambled-identifier"},"network":"helium_iot","protocol":"LORAWAN","timestamp":1741178794588,"txInfo":{"frequency":905100000,"modulation":{"lora":{"bandwidth":125000,"codeRate":"CR_4:5","spreadingFactor":7}}}}]}}`
	tests := []struct {
		name          string
		cloudEvent    []byte
		expectedVIN   string
		expectedError error
	}{
		{
			name:          "Fingerprint Payload",
			cloudEvent:    []byte(fingerPrintData),
			expectedVIN:   "1ABCD2EFGH3JKLMNO",
			expectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			module := Module{}
			var event cloudevent.RawEvent
			err := json.Unmarshal(tt.cloudEvent, &event)
			require.NoError(t, err)
			fp, err := module.FingerprintConvert(context.Background(), event)

			if tt.expectedError != nil {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expectedVIN, fp.VIN)
			}
		})
	}
}
