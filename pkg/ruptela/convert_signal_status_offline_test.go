package ruptela_test

import (
	"encoding/json"
	"testing"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/ruptela"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/stretchr/testify/require"
)

// The R1 buffers records it generates while it has no data connection and sends them after it
// reconnects, stamped with the send time instead of the capture time. Element 418 (GPRS status)
// is captured with the record, so 418 = 0 marks a record that was generated offline: its position
// is from minutes or hours before its timestamp. These fixtures are shaped like a real
// reconnect burst, where a position captured hours earlier was sent stamped with the reconnect time.
func offlineStatusEvent(t *testing.T, gprsStatus string) cloudevent.RawEvent {
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

func signalNames(signals []vss.Signal) map[string]bool {
	names := make(map[string]bool, len(signals))
	for _, s := range signals {
		names[s.Data.Name] = true
	}
	return names
}

func TestV1PayloadGeneratedOfflineOmitsPosition(t *testing.T) {
	t.Parallel()
	signals, err := ruptela.SignalsFromV1Payload(offlineStatusEvent(t, "0"))
	require.NoError(t, err)
	names := signalNames(signals)

	for _, name := range []string{vss.FieldCurrentLocationCoordinates, vss.FieldCurrentLocationAltitude, vss.FieldCurrentLocationHeading} {
		require.False(t, names[name], "%s from a record generated offline must not be emitted at its send time", name)
	}
	require.True(t, names[vss.FieldSpeed], "non-position signals are still emitted")
}

func TestV1PayloadGeneratedOnlineKeepsPosition(t *testing.T) {
	t.Parallel()
	signals, err := ruptela.SignalsFromV1Payload(offlineStatusEvent(t, "1"))
	require.NoError(t, err)
	names := signalNames(signals)

	for _, name := range []string{vss.FieldCurrentLocationCoordinates, vss.FieldCurrentLocationAltitude, vss.FieldCurrentLocationHeading} {
		require.True(t, names[name], "%s must be emitted for a record generated while connected", name)
	}
}
