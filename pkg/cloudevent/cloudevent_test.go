package cloudevent_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/DIMO-Network/cloudevent"
	modelce "github.com/DIMO-Network/model-garage/pkg/cloudevent"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFingerprintEvent_Roundtrip(t *testing.T) {
	t.Parallel()
	now := time.Now().UTC().Truncate(time.Millisecond)
	event := modelce.FingerprintEvent{
		CloudEventHeader: cloudevent.CloudEventHeader{
			ID:       "fp-1",
			Source:   "test-source",
			Producer: "test-producer",
			Subject:  "test-subject",
			Time:     now,
			Type:     cloudevent.TypeFingerprint,
		},
		Data: modelce.Fingerprint{VIN: "1ABCD2EFGH3JKLMNO"},
	}

	data, err := json.Marshal(event)
	require.NoError(t, err)

	var obj map[string]any
	require.NoError(t, json.Unmarshal(data, &obj))

	dataField, ok := obj["data"].(map[string]any)
	require.True(t, ok, "data field must be a JSON object")
	assert.Equal(t, "1ABCD2EFGH3JKLMNO", dataField["vin"], "VIN must appear in data.vin")

	var roundtrip modelce.FingerprintEvent
	require.NoError(t, json.Unmarshal(data, &roundtrip))
	assert.Equal(t, "1ABCD2EFGH3JKLMNO", roundtrip.Data.VIN, "VIN must survive roundtrip")
}
