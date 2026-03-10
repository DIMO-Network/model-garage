package autopi

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func assertAndNormalizeSignals(t *testing.T, signals []vss.Signal, expectedCloudEventID string) {
	t.Helper()
	for i := range signals {
		assert.Equal(t, cloudevent.TypeSignal, signals[i].Type, "signal %d should have TypeSignal", i)
		assert.False(t, signals[i].Time.IsZero(), "signal %d should have a non-zero Time", i)
		assert.Equal(t, expectedCloudEventID, signals[i].Data.CloudEventID, "signal %d should reference original event ID", i)
		signals[i].Type = ""
		signals[i].Time = time.Time{}
		signals[i].Data.CloudEventID = ""
	}
}

func normalizeExpectedSignals(signals []vss.Signal) {
	for i := range signals {
		signals[i].Type = ""
		signals[i].Time = time.Time{}
		signals[i].Data.CloudEventID = ""
	}
}

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

	const (
		source  = "dimo/integration/27qftVRWQYpVDcO5DltO5Ojbjxk"
		subject = "did:erc721:1:0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8:33"
	)

	validHeader := cloudevent.CloudEventHeader{
		DataVersion: DataVersion,
		Type:        cloudevent.TypeStatus,
		Source:      source,
		Subject:     subject,
		Time:        ts,
	}

	tests := []struct {
		name            string
		cloudEvent      cloudevent.RawEvent
		expectedSignals []vss.Signal
		expectedError   error
	}{
		{
			name: "Valid Signal Payload",
			cloudEvent: cloudevent.RawEvent{
				CloudEventHeader: validHeader,
				Data:             json.RawMessage(signalData),
			},
			expectedSignals: []vss.Signal{
				{CloudEventHeader: validHeader, Data: vss.SignalData{Timestamp: ts, Name: vss.FieldOBDLongTermFuelTrim1, ValueNumber: 25}},
				{CloudEventHeader: validHeader, Data: vss.SignalData{Timestamp: ts, Name: vss.FieldPowertrainCombustionEngineECT, ValueNumber: 107}},
			},
			expectedError: nil,
		},
		{
			name: "Device Status Payload",
			cloudEvent: cloudevent.RawEvent{
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
				if tt.expectedSignals != nil {
					assertAndNormalizeSignals(t, signals, tt.cloudEvent.ID)
					normalizeExpectedSignals(tt.expectedSignals)
				}
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
		cloudEvent    cloudevent.RawEvent
		expectedVIN   string
		expectedError error
	}{
		{
			name: "Fingerprint Payload",
			cloudEvent: cloudevent.RawEvent{
				CloudEventHeader: cloudevent.CloudEventHeader{
					DataVersion: DataVersion,
					Type:        cloudevent.TypeFingerprint,
					Subject:     "did:erc721:1:0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8:33",
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
