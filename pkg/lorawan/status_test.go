package lorawan

import (
	"encoding/json"
	"testing"

	"github.com/DIMO-Network/model-garage/pkg/cloudevent"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConvertToCloudEvents(t *testing.T) {
	tests := []struct {
		name             string
		input            []byte
		expectError      bool
		length           int
		expectedSubject  string
		expectedProducer string
	}{
		{
			name:             "Status payload",
			input:            []byte(`{"data":{"header": 2,"device":{"serial":"60d4af69-86e8-b790-02d3-c0a9dc4d6c8a","softwareVersion":"v1.0.0"},"timestamp":1732224181876,"vehicle":{"make":"MINI","model":"Countryman","signals":[{"name":"batteryVoltage","timestamp":1732224181876,"value":12.95}],"year":2018}},"signature":"0x67bdfbfce03ef7c6577a4a64de037db97d882ef158ee6d1b3adc96e0e58599b2508bb74f8780e102e0c50b7b30385ed6160aa8218c9793cb00fc8f8b355a966c1b","time":"2024-11-21T21:23:01.876617869Z","type":"com.dimo.device.status.v2","vehicleTokenId":1, "deviceTokenId": 2222}`),
			expectError:      false,
			length:           2,
			expectedSubject:  "did:nft:2:0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8_1",
			expectedProducer: "did:nft:2:0x325b45949C833986bC98e98a49F3CA5C5c4643B5_2222",
		},
		{
			name:             "Status payload with no vehicleTokenId",
			input:            []byte(`{"data":{"header": 2,"device":{"softwareVersion":"v1.0.0"},"timestamp":1732224181876,"vehicle":{"make":"MINI","model":"Countryman","year":2018}},"time":"2024-11-21T21:23:01.876617869Z","type":"com.dimo.device.status.v2", "deviceTokenId": 2222}`),
			expectError:      false,
			length:           2,
			expectedSubject:  "",
			expectedProducer: "did:nft:2:0x325b45949C833986bC98e98a49F3CA5C5c4643B5_2222",
		},
		{
			name:        "Status payload, device token id is missing",
			input:       []byte(`{"type":"com.dimo.device.status.v2"}`),
			expectError: true,
		},
		{
			name:             "Fingerprint payload",
			input:            []byte(`{"data":{"device":{"id": "F4CE368CC0DF1D1D","name": "0x3216049F6D65A414E785D1012F70D8944AA1EC44","protocol": "lora_helium","nonce": 28},"timestamp": 1732224181876,"vehicle": {"make": "MINI","model": "Countryman","signals":[{"name": "batteryVoltage","timestamp": 1732224181876,"value": 12.95}],"year": 2018},"header": 1},"signature": "0xd085bac8eb4dbbc76f413c763e9910b3501e883437c658fe5ac33d1a5a380440392c2dc93b5dcb180e47e062784d746d1a6643d7419d8a6c76b8c39ef9ba12601c", "time": "2024-11-21T21:23:01.876617869Z", "type": "com.dimo.device.status.v2","vehicleTokenId": 1, "deviceTokenId": 2222}`),
			expectError:      false,
			length:           3,
			expectedSubject:  "did:nft:2:0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8_1",
			expectedProducer: "did:nft:2:0x325b45949C833986bC98e98a49F3CA5C5c4643B5_2222",
		},
		{
			name:        "Unknown payload type",
			input:       []byte(`{"subject":"0x1234567890abcdef1234567890abcdef12345678","time":"2023-10-31T12:34:56Z","type":"some","vehicleTokenId":1, "deviceTokenId": 2222}`),
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
			events, err := ConvertToCloudEvents(tt.input, 2, "0x325b45949C833986bC98e98a49F3CA5C5c4643B5", "0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8")
			if tt.expectError {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Len(t, events, tt.length)

				var cloudEvent cloudevent.CloudEvent[json.RawMessage]
				errUnmarshal := json.Unmarshal(events[0], &cloudEvent)
				if errUnmarshal != nil {
					t.Fatalf("Failed to unmarshal cloud event: %v", errUnmarshal)
				}

				assert.Equal(t, tt.expectedSubject, cloudEvent.Subject)
				assert.Equal(t, tt.expectedProducer, cloudEvent.Producer)
			}
		})
	}
}
