package autopi

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/DIMO-Network/model-garage/pkg/cloudevent"
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

	const source = "dimo/integration/27qftVRWQYpVDcO5DltO5Ojbjxk"
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
					Subject:     "did:nft:1:0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8_33",
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
					Subject:     "did:nft:1:0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8_33",
					Producer:    "did:nft:1:0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8_33",
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
	ts := time.Unix(1727360340, 0).UTC()

	fingerPrintData := `{"timestamp":1638316800000,"device":{"rpiUptimeSecs":3600,"batteryVoltage":12.6,"softwareVersion":"1.0.0","hwVersion":"v1","imei":"123456789012345","serial":"unit123"},"vin":"1HGCM82633A123456","protocol":"ISO9141","odometer":12345.67}`

	tests := []struct {
		name          string
		cloudEvent    cloudevent.CloudEvent[json.RawMessage]
		expectedVIN   string
		expectedError error
	}{
		{
			name: "Fingerprint Payload",
			cloudEvent: cloudevent.CloudEvent[json.RawMessage]{
				CloudEventHeader: cloudevent.CloudEventHeader{
					DataVersion: DataVersion,
					Type:        cloudevent.TypeFingerprint,
					Subject:     "did:nft:1:0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8_33",
					Time:        ts,
				},
				Data: json.RawMessage(fingerPrintData),
			},
			expectedVIN:   "1HGCM82633A123456",
			expectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			module := Module{}
			fp, err := module.FingerprintConvert(context.Background(), tt.cloudEvent)

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
