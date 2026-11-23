package ruptela_test

import (
	"encoding/json"
	"testing"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/ruptela"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/stretchr/testify/require"
)

// statusEventWithGPRS returns a status record with a GPS fix and the given element 418 (GPRS status:
// 0 – disconnected, 1 – connected) value. The R1 buffers records it generates while offline and
// sends them after it reconnects; the ingest lambda re-dates them to their capture time, so their
// position is converted like any other.
func statusEventWithGPRS(t *testing.T, gprsStatus string) cloudevent.RawEvent {
	t.Helper()
	doc := `{
  "id": "2vOfflineStatusTest00000000000",
  "source": "0xF26421509Efe92861a587482100c6d728aBf1CD0",
  "producer": "did:erc721:137:0x9c94C395cBcBDe662235E0A9d3bB87Ad708561BA:1",
  "specversion": "1.0",
  "subject": "did:erc721:137:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:33",
  "time": "2026-01-01T12:00:00Z",
  "type": "dimo.status",
  "datacontenttype": "application/json",
  "dataversion": "r/v0/s",
  "data": {
    "trigger": 8,
    "prt": 0,
    "pos": {"lat": 520000000, "lon": 130000000, "alt": 1180, "dir": 26500, "spd": 88, "sat": 23, "hdop": 14},
    "signals": {"95": "58", "409": "1", "418": "` + gprsStatus + `", "525": "AAAB7"}
  }
}`
	var event cloudevent.RawEvent
	require.NoError(t, json.Unmarshal([]byte(doc), &event))
	return event
}

func TestV1PayloadKeepsPositionWhenGeneratedOffline(t *testing.T) {
	t.Parallel()
	for _, gprs := range []string{"0", "1"} {
		t.Run("418="+gprs, func(t *testing.T) {
			t.Parallel()
			signals, err := ruptela.SignalsFromV1Payload(statusEventWithGPRS(t, gprs))
			require.NoError(t, err)

			byName := make(map[string]vss.Signal, len(signals))
			for _, s := range signals {
				byName[s.Data.Name] = s
			}
			for _, name := range []string{vss.FieldCurrentLocationCoordinates, vss.FieldCurrentLocationAltitude, vss.FieldCurrentLocationHeading} {
				require.Contains(t, byName, name)
			}
			require.InDelta(t, 1.4, byName[vss.FieldCurrentLocationCoordinates].Data.ValueLocation.HDOP, 1e-9)
		})
	}
}
