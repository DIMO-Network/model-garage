package tesla

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/DIMO-Network/cloudevent"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestModule_CloudEventConvert(t *testing.T) {
	tests := []struct {
		name          string
		input         []byte
		expectHeaders int
		expectErr     bool
	}{
		{
			name: "valid cloud event with vin",
			input: []byte(`{
				"id": "2pcYwspbaBFJ7NPGZ2kivkuJ12a",
				"source": "0xFFEE022fAb46610EAFe98b87377B42e366364a71",
				"producer": "did:nft:80003:0x78513c8CB4D6B6079f813850376bc9c7fc8aE67f_12",
				"specversion": "1.0",
				"subject": "did:nft:80003:0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8_15",
				"time": "2024-12-01T15:31:12.378075897Z",
				"type": "dimo.status",
				"signature": "0x22cca92bb6a16fed01def56d02541427633ff82552bc8c5c2da2fffd69c4436927b256ab0f1352e584deb5394fff2f979699f206691f73fffee547cee1431c",
				"data": {
					"id": 234234,
					"user_id": 32425456,
					"vehicle_id": 33,
					"vin": "VF33E1EB4K55F700D",
					"color": null,
					"access_type": "OWNER",
					"granular_access": {
						"hide_private": false
					}
				}
			}`),
			expectHeaders: 2, // Expects both original and fingerprint headers
			expectErr:     false,
		},
		{
			name: "valid cloud event without vin",
			input: []byte(`{
				"id": "2pcYwspbaBFJ7NPGZ2kivkuJ12a",
				"source": "0xFFEE022fAb46610EAFe98b87377B42e366364a71",
				"producer": "did:nft:80003:0x78513c8CB4D6B6079f813850376bc9c7fc8aE67f_12",
				"specversion": "1.0",
				"subject": "did:nft:80003:0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8_15",
				"time": "2024-12-01T15:31:12.378075897Z",
				"type": "dimo.status",
				"signature": "0x22cca92bb6a16fed01def56d02541427633ff82552bc8c5c2da2fffd69c4436927b256ab0f1352e584deb5394fff2f979699f206691f73fffee547cee1431c",
				"data": {
					"id": 1492932094674954,
					"user_id": 32425456,
					"vehicle_id": 33,
					"access_type": "OWNER",
					"granular_access": {
						"hide_private": false
					}
				}
			}`),
			expectHeaders: 1, // Expects only original header
			expectErr:     false,
		},
		{
			name:          "invalid json",
			input:         []byte(`{invalid`),
			expectHeaders: 0,
			expectErr:     true,
		},
		{
			name: "malformed data field",
			input: []byte(`{
				"id": "2pcYwspbaBFJ7NPGZ2kivkuJ12a",
				"source": "0xFFEE022fAb46610EAFe98b87377B42e366364a71",
				"producer": "did:nft:80003:0x78513c8CB4D6B6079f813850376bc9c7fc8aE67f_12",
				"specversion": "1.0",
				"type": "dimo.status",
				"data": "not an object"
			}`),
			expectHeaders: 1,
			expectErr:     false,
		},
		{
			name: "standard fleet telemetry",
			input: []byte(`{
				"id": "2pcYwspbaBFJ7NPGZ2kivkuJ12a",
				"source": "0xFFEE022fAb46610EAFe98b87377B42e366364a71",
				"producer": "did:nft:80003:0x78513c8CB4D6B6079f813850376bc9c7fc8aE67f_12",
				"specversion": "1.0",
				"type": "dimo.status",
				"dataversion": "fleet_telemetry/v1.0.0",
				"data": {
					"payloads": [
						"ChgIFRIUOhIJLEfIQJ4jSkARxeOiWkQU7b8SCwjD15i+BhD+w5tCGhFYUDdZSENFUjNTQjU4MjUwNg==",
						"ChgIFRIUOhIJLEfIQJ4jSkARru/DQUIU7b8SCwjF15i+BhDEiKZCGhFYUDdZSENFUjNTQjU4MjUwNg=="
					]
				}
			}`),
			expectHeaders: 2,
			expectErr:     false,
		},
		{
			name: "fleet telemetry without payloads",
			input: []byte(`{
				"id": "2pcYwspbaBFJ7NPGZ2kivkuJ12a",
				"source": "0xFFEE022fAb46610EAFe98b87377B42e366364a71",
				"producer": "did:nft:80003:0x78513c8CB4D6B6079f813850376bc9c7fc8aE67f_12",
				"specversion": "1.0",
				"type": "dimo.status",
				"dataversion": "fleet_telemetry/v1.0.0",
				"data": {
					"payloads": []
				}
			}`),
			expectHeaders: 1,
			expectErr:     false,
		},
	}

	var module Module
	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers, data, err := module.CloudEventConvert(ctx, tt.input)

			if tt.expectErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Len(t, headers, tt.expectHeaders)
			assert.NotNil(t, data)

			if tt.expectHeaders == 2 {
				// Verify the fingerprint header
				assert.Equal(t, cloudevent.TypeFingerprint, headers[1].Type)
				// Verify original header fields are preserved
				assert.Equal(t, headers[0].Source, headers[1].Source)
				assert.Equal(t, headers[0].Subject, headers[1].Subject)
				assert.Equal(t, headers[0].Producer, headers[1].Producer)
			}
		})
	}
}

func TestFullFromDataConversion(t *testing.T) {
	var module Module
	t.Parallel()
	event := cloudevent.CloudEvent[json.RawMessage]{}

	for _, test := range []struct {
		ExpectedVin   string
		CloudEvent    string
		ExpectedError error
	}{
		{
			ExpectedVin: "VF33E1EB4K55F700D",
			CloudEvent:  pollingPayload,
		},
		{
			ExpectedVin: "5YJ3E1EBXKF477800",
			CloudEvent:  streamingPayload,
		},
	} {
		err := json.Unmarshal([]byte(test.CloudEvent), &event)
		require.NoError(t, err, "error unmarshalling input JSON")
		fp, err := module.FingerprintConvert(t.Context(), event)
		require.Equal(t, test.ExpectedError, err, "error decoding fingerprint")
		require.Equal(t, test.ExpectedVin, fp.VIN, "decoded VIN does not match expected VIN")
	}
}

var streamingPayload = `{
        "data": {
            "payloads": [
                "CgcIMRIDCgEwCgcIJRIDCgEwCgoIAhIGCgRJZGxlCgcIVhIDCgEyEgwIw9DwvQYQqJL26QIaETVZSjNFMUVCWEtGNDc3ODAw"
            ]
        },
        "dataversion": "fleet_telemetry/v1.0.0",
        "id": "2tTmE6tswqo9t0XjreAdJU80K5Q",
        "producer": "did:nft:80002:0x78513c8CB4D6B6079f813850376bc9c7fc8aE67f_37",
        "signature": "0x2ecd6732e72b9a51f63e7182be2a15958dd5520379d3773725430db4c0a088e25b3127ed7de4de2cfe347b319bac54b64d3851355069178bf403f1dd3f9c2ffb1c",
        "source": "0xc4035Fecb1cc906130423EF05f9C20977F643722",
        "specversion": "1.0",
        "subject": "did:nft:80002:0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8_71",
        "time": "2025-02-24T08:05:23.786519144Z",
        "type": "dimo.status"
    }`

var pollingPayload = `{
	"id": "2pcYwspbaBFJ7NPGZ2kivkuJ12a",
	"source": "0xFFEE022fAb46610EAFe98b87377B42e366364a71",
	"producer": "did:nft:80003:0x78513c8CB4D6B6079f813850376bc9c7fc8aE67f_12",
	"specversion": "1.0",
	"subject": "did:nft:80003:0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8_15",
	"time": "2024-12-01T15:31:12.378075897Z",
	"type": "dimo.status",
	"signature": "0x22cca92bb6a16fed01def56d02541427633ff82552bc8c5c2da2fffd69c4436927b256ab0f1352e584deb5394fff2f979699f206691f73fffee547cee1431c",
	"data": {
		"id": 234234,
		"user_id": 32425456,
		"vehicle_id": 33,
		"vin": "VF33E1EB4K55F700D",
		"color": null,
		"access_type": "OWNER",
		"granular_access": {
			"hide_private": false
		}
	}
}`
