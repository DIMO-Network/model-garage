package defaultmodule_test

import (
	"context"
	"testing"

	"github.com/DIMO-Network/model-garage/pkg/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/defaultmodule"
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
				"producer": "did:nft:80003:0x78513c8CB4D6B6079f813850376bc9c7fc8aE67f_12",
				"specversion": "1.0",
				"subject": "did:nft:80003:0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8_15",
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
				"producer": "did:nft:80003:0x78513c8CB4D6B6079f813850376bc9c7fc8aE67f_12",
				"specversion": "1.0",
				"subject": "did:nft:80003:0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8_15",
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
				"producer": "did:nft:80003:0x78513c8CB4D6B6079f813850376bc9c7fc8aE67f_12",
				"specversion": "1.0",
				"type": "dimo.status",
				"data": "not an object"
			}`),
			expectHeaderTypes: []string{cloudevent.TypeUnknown},
			expectErr:         false,
		},
	}

	module, _ := defaultmodule.New()
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
