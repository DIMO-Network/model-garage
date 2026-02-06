package ruptela_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/ruptela"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCloudEventConvert(t *testing.T) {
	module := ruptela.Module{
		ChainID:                 1,
		AftermarketContractAddr: common.HexToAddress("0x06012c8cf97BEaD5deAe237070F9587f8E7A266d"),
		VehicleContractAddr:     common.HexToAddress("0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF"),
	}
	tests := []struct {
		name             string
		input            []byte
		expectError      bool
		length           int
		expectedSubject  string
		expectedProducer string
	}{
		{
			name:             "Status payload with VIN and event",
			input:            []byte(`{"ds":"r/v0/s","signature":"test","time":"2022-01-01T00:00:00Z","data":{"signals":{"104":"4148544241334344","105":"3930363235323539","106":"3300000000000000", "135":"1"}},"subject":"test","vehicleTokenId":1, "deviceTokenId":2}`),
			expectError:      false,
			length:           3,
			expectedSubject:  "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:1",
			expectedProducer: "did:erc721:1:0x06012c8cf97BEaD5deAe237070F9587f8E7A266d:2",
		},
		{
			name:             "Status payload with VIN",
			input:            []byte(`{"ds":"r/v0/s","signature":"test","time":"2022-01-01T00:00:00Z","data":{"signals":{"104":"4148544241334344","105":"3930363235323539","106":"3300000000000000"}},"subject":"test","vehicleTokenId":1, "deviceTokenId":2}`),
			expectError:      false,
			length:           2,
			expectedSubject:  "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:1",
			expectedProducer: "did:erc721:1:0x06012c8cf97BEaD5deAe237070F9587f8E7A266d:2",
		},
		{
			name:             "Status payload with no VIN",
			input:            []byte(`{"ds":"r/v0/s","signature":"test","time":"2022-01-01T00:00:00Z", "vehicleTokenId":1, "deviceTokenId":2, "data":{"trigger":409,"prt":1,"signals":{"104":"0","105":"0","106":"0"}}}`),
			expectError:      false,
			length:           1,
			expectedSubject:  "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:1",
			expectedProducer: "did:erc721:1:0x06012c8cf97BEaD5deAe237070F9587f8E7A266d:2",
		},
		{
			name:             "Status payload with 0s for the VIN",
			input:            []byte(`{"ds":"r/v0/s","signature":"test","time":"2022-01-01T00:00:00Z", "vehicleTokenId":1, "deviceTokenId":2, "data":{"trigger":409,"prt":1,"signals":{"104":"0000000000000000","105":"0000000000000000","106":"0000000000000000"}}}`),
			expectError:      false,
			length:           1,
			expectedSubject:  "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:1",
			expectedProducer: "did:erc721:1:0x06012c8cf97BEaD5deAe237070F9587f8E7A266d:2",
		},
		{
			name:             "Status payload with 0s for the CAN VIN",
			input:            []byte(`{"ds":"r/v0/s","signature":"test","time":"2022-01-01T00:00:00Z", "vehicleTokenId":1, "deviceTokenId":2, "data":{"trigger":409,"prt":1,"signals":{"123":"0000000000000000","125":"0000000000000000","124":"0000000000000000"}}}`),
			expectError:      false,
			length:           1,
			expectedSubject:  "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:1",
			expectedProducer: "did:erc721:1:0x06012c8cf97BEaD5deAe237070F9587f8E7A266d:2",
		},
		{
			name:             "Status payload with CAN VIN (smart5)",
			input:            []byte(`{"ds":"r/v0/s","signature":"test","time":"2022-01-01T00:00:00Z", "vehicleTokenId":1, "deviceTokenId":2, "data":{"trigger":409,"prt":1,"signals":{"123":"4148544241334344","124":"3930363235323539","125":"3300000000000000"}}}`),
			expectError:      false,
			length:           2,
			expectedSubject:  "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:1",
			expectedProducer: "did:erc721:1:0x06012c8cf97BEaD5deAe237070F9587f8E7A266d:2",
		},
		{
			name:             "Status payload with no vehicleTokenId",
			input:            []byte(`{"ds":"r/v0/s","signature":"test","time":"2022-01-01T00:00:00Z", "deviceTokenId":2, "data":{"trigger":409,"prt":1,"signals":{"104":"0","105":"0","106":"0"}}}`),
			expectError:      false,
			length:           1,
			expectedSubject:  "",
			expectedProducer: "did:erc721:1:0x06012c8cf97BEaD5deAe237070F9587f8E7A266d:2",
		},
		{
			name:        "Status payload with no deviceTokenId",
			input:       []byte(`{"ds":"r/v0/s","signature":"test","time":"2022-01-01T00:00:00Z", "data":{"trigger":409,"prt":1,"signals":{"104":"0","105":"0","106":"0"}}}`),
			expectError: true,
			length:      1,
		},
		{
			name:             "Location payload",
			input:            []byte(`{"ds":"r/v0/s","signature":"test","time":"2022-01-01T00:00:00Z","vehicleTokenId":1, "deviceTokenId":2,"data":{"location":[{"ts":1727712225,"lat":522784033,"lon":-9085750,"alt":1049,"dir":22390,"hdop":50},{"ts":1727712226,"lat":522783650,"lon":-9086150,"alt":1044,"dir":20100,"hdop":50}]}}`),
			expectError:      false,
			length:           1,
			expectedSubject:  "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:1",
			expectedProducer: "did:erc721:1:0x06012c8cf97BEaD5deAe237070F9587f8E7A266d:2",
		},
		{
			name:             "Dev status payload",
			input:            []byte(`{"ds":"r/v0/dev","signature":"test","time":"2022-01-01T00:00:00Z","vehicleTokenId":1, "deviceTokenId":2,"data":{"sn":"869267077308554","battVolt":"12420","hwVersion":"FTX-04-12231","imei":"869267077308554","fwVersion":"00.06.56.45","sigStrength":"14","accessTech":"0","operator":"23415","locAreaCode":"13888","cellId":"29443"}}`),
			expectError:      false,
			length:           1,
			expectedSubject:  "did:erc721:1:0x06012c8cf97BEaD5deAe237070F9587f8E7A266d:2",
			expectedProducer: "did:erc721:1:0x06012c8cf97BEaD5deAe237070F9587f8E7A266d:2",
		},
		{
			name:             "Batt payload",
			input:            []byte(`{"ds":"r/v0/batt","signature":"test","time":"2022-01-01T00:00:00Z","vehicleTokenId":1,"deviceTokenId":2,"data":{"voltage":4200,"soc":85}}`),
			expectError:      false,
			length:           1,
			expectedSubject:  "did:erc721:1:0x06012c8cf97BEaD5deAe237070F9587f8E7A266d:2",
			expectedProducer: "did:erc721:1:0x06012c8cf97BEaD5deAe237070F9587f8E7A266d:2",
		},
		{
			name:             "DTC status payload",
			input:            []byte(`{ "ds":"r/v0/dtc", "signature":"test","time":"2024-09-26T14:19:14Z", "vehicleTokenId":1, "deviceTokenId":2,"data":{"dtc_codes":[{"id":"7E8","code":"P0101"},{"id":"7E8","code":"P0202"}]}}`),
			expectError:      false,
			length:           1,
			expectedSubject:  "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:1",
			expectedProducer: "did:erc721:1:0x06012c8cf97BEaD5deAe237070F9587f8E7A266d:2",
		},
		{
			name:        "Invalid time format",
			input:       []byte(`{"ds":"r/v0/loc","signature":"test","time":"1727712275"}`),
			expectError: true,
		},
		{
			name:        "Invalid input",
			input:       []byte(`invalid`),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hdrs, _, err := module.CloudEventConvert(context.Background(), tt.input)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Len(t, hdrs, tt.length)

				cloudEvent := hdrs[0]
				assert.Equal(t, tt.expectedSubject, cloudEvent.Subject)
				assert.Equal(t, tt.expectedProducer, cloudEvent.Producer)
			}
		})
	}
}

func TestSignalConvert(t *testing.T) {
	ts := time.Unix(1727360340, 0).UTC()

	// Signal payload data
	signalData := `{
		"signals": {
			"96": "8",
			"97": "8"
		}
	}`

	// Location payload data
	locationData := `{
		"location": [
			{
				"alt": 1232,
				"ts": 1727360340
			},
			{
				"alt": 12,
				"ts": 1727360341
			}
		]
	}`

	// Location payload data
	dtcData := `{
	 "dtc_codes":[
         {
            "id":"7E8",
            "code":"P0101"
         },
         {
            "id":"7E8",
            "code":"P0202"
         }
      ]
	}`

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
					DataVersion: ruptela.StatusEventDS,
					Type:        cloudevent.TypeStatus,
					Source:      "ruptela/TODO",
					Subject:     "did:erc721:1:0x06012c8cf97BEaD5deAe237070F9587f8E7A266d:33",
					Time:        ts,
				},
				Data: json.RawMessage(signalData),
			},
			expectedSignals: []vss.Signal{
				{TokenID: 33, Timestamp: ts, Name: vss.FieldExteriorAirTemperature, ValueNumber: -32, Source: "ruptela/TODO"},
				{TokenID: 33, Timestamp: ts, Name: vss.FieldPowertrainCombustionEngineECT, ValueNumber: -32, Source: "ruptela/TODO"},
			},
			expectedError: nil,
		},
		{
			name: "Valid Location Payload",
			cloudEvent: cloudevent.CloudEvent[json.RawMessage]{
				CloudEventHeader: cloudevent.CloudEventHeader{
					DataVersion: ruptela.LocationEventDS,
					Type:        cloudevent.TypeStatus,
					Source:      "ruptela/TODO",
					Subject:     "did:erc721:1:0x06012c8cf97BEaD5deAe237070F9587f8E7A266d:33",
					Time:        ts,
				},
				Data: json.RawMessage(locationData),
			},
			expectedSignals: []vss.Signal{
				{TokenID: 33, Timestamp: ts, Name: vss.FieldCurrentLocationAltitude, ValueNumber: 123.2, Source: "ruptela/TODO"},
				{TokenID: 33, Timestamp: ts.Add(time.Second), Name: vss.FieldCurrentLocationAltitude, ValueNumber: 1.2, Source: "ruptela/TODO"},
			},
			expectedError: nil,
		},
		{
			name: "Valid DTC Payload",
			cloudEvent: cloudevent.CloudEvent[json.RawMessage]{
				CloudEventHeader: cloudevent.CloudEventHeader{
					DataVersion: ruptela.DTCEventDS,
					Type:        cloudevent.TypeStatus,
					Source:      "ruptela/TODO",
					Subject:     "did:erc721:1:0x06012c8cf97BEaD5deAe237070F9587f8E7A266d:33",
					Time:        ts,
				},
				Data: json.RawMessage(dtcData),
			},
			expectedSignals: []vss.Signal{
				{TokenID: 33, Timestamp: ts, Name: "obdDTCList", ValueString: "[\"P0101\",\"P0202\"]", Source: "ruptela/TODO"},
			},
			expectedError: nil,
		},
		{
			name: "Invalid DTC Payload",
			cloudEvent: cloudevent.CloudEvent[json.RawMessage]{
				CloudEventHeader: cloudevent.CloudEventHeader{
					DataVersion: ruptela.DTCEventDS,
					Type:        cloudevent.TypeStatus,
					Source:      "ruptela/TODO",
					Subject:     "did:erc721:1:0x06012c8cf97BEaD5deAe237070F9587f8E7A266d:33",
					Time:        ts,
				},
				Data: json.RawMessage("{}"),
			},
			expectedError: errors.New("error getting obdDTCList: field not found 'OBDDTCList'"),
		},
		{
			name: "Invalid Event DataVersion",
			cloudEvent: cloudevent.CloudEvent[json.RawMessage]{
				CloudEventHeader: cloudevent.CloudEventHeader{
					DataVersion: "unknownVersion",
					Type:        cloudevent.TypeStatus,
				},
				Data: json.RawMessage(signalData),
			},
			expectedError: errors.New("unknown data version: unknownVersion"),
		},
		{
			name: "Non-Status Event Type",
			cloudEvent: cloudevent.CloudEvent[json.RawMessage]{
				CloudEventHeader: cloudevent.CloudEventHeader{
					DataVersion: ruptela.StatusEventDS,
					Type:        "fingerprint",
				},
				Data: json.RawMessage(signalData),
			},
			expectedError: nil, // Should skip non-status events without error
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Marshal CloudEvent to JSON

			// Call the SignalConvert function
			module := ruptela.Module{}
			signals, err := module.SignalConvert(t.Context(), tt.cloudEvent)

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
