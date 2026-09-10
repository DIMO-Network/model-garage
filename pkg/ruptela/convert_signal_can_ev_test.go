package ruptela_test

import (
	"encoding/json"
	"testing"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/ruptela"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/stretchr/testify/require"
)

// canEVPayload builds a minimal r/v0/s status event with the given signals map.
// The shape mirrors a real R1 record: the OBD EV SoC element (722) is only
// answered every other record and reads FF in between, while a CAN profile that
// does not decode the vehicle fills the CAN EV elements (515/516/517) with a
// literal 0 instead of the FF "not available" sentinel.
func canEVPayload(t *testing.T, signals map[string]string) cloudevent.RawEvent {
	t.Helper()
	sigJSON, err := json.Marshal(signals)
	require.NoError(t, err)
	doc := `{
  "id": "2vAbCdEfGhIjKlMnOpQrStUvWxY",
  "source": "0xF26421509Efe92861a587482100c6d728aBf1CD0",
  "producer": "did:erc721:137:0x9c94C395cBcBDe662235E0A9d3bB87Ad708561BA:36798",
  "specversion": "1.0",
  "subject": "did:erc721:137:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:192097",
  "time": "2026-09-03T22:45:05Z",
  "type": "dimo.status",
  "datacontenttype": "application/json",
  "dataversion": "r/v0/s",
  "data": {"trigger": 7, "prt": 0, "signals": ` + string(sigJSON) + `}
}`
	var event cloudevent.RawEvent
	require.NoError(t, json.Unmarshal([]byte(doc), &event))
	return event
}

func signalValues(signals []vss.Signal) map[string]float64 {
	out := make(map[string]float64, len(signals))
	for _, s := range signals {
		out[s.Data.Name] = s.Data.ValueNumber
	}
	return out
}

func TestCANEVZeroIsNotAvailable(t *testing.T) {
	t.Parallel()
	event := canEVPayload(t, map[string]string{
		"409": "1",  // ignition on
		"722": "FF", // OBD EV SoC not answered in this record
		"515": "0",  // CAN EV SoC: dead CAN profile, literal 0
		"516": "0",  // CAN EV distance until recharge: literal 0
		"517": "0",  // CAN EV charging state: false
	})

	signals, err := ruptela.SignalsFromV1Payload(event)
	require.NoError(t, err)
	got := signalValues(signals)

	require.NotContains(t, got, vss.FieldPowertrainTractionBatteryStateOfChargeCurrent,
		"a literal 0 from the CAN EV SoC element must be treated as not available, not 0%")
	require.NotContains(t, got, vss.FieldPowertrainTractionBatteryRange,
		"a literal 0 from the CAN EV range element must be treated as not available, not 0 km")
	// Booleans keep their zero: false is a real value for the charging state.
	require.Contains(t, got, vss.FieldPowertrainTractionBatteryChargingIsCharging)
	require.Equal(t, float64(0), got[vss.FieldPowertrainTractionBatteryChargingIsCharging])
}

func TestCANEVNonZeroStillConverts(t *testing.T) {
	t.Parallel()
	event := canEVPayload(t, map[string]string{
		"409": "1",
		"722": "FF", // OBD EV SoC not answered, resolver must fall through to CAN
		"515": "24", // 0x24 = 36 * 0.5 = 18 %
		"516": "1F", // 0x1F = 31 km
	})

	signals, err := ruptela.SignalsFromV1Payload(event)
	require.NoError(t, err)
	got := signalValues(signals)

	require.Equal(t, float64(18), got[vss.FieldPowertrainTractionBatteryStateOfChargeCurrent])
	require.Equal(t, float64(31), got[vss.FieldPowertrainTractionBatteryRange])
}

func TestCANEVBattery2ZeroIsNotAvailable(t *testing.T) {
	t.Parallel()
	event := canEVPayload(t, map[string]string{
		"409": "1",
		"722": "FF",
		"515": "FF", // battery 1 CAN SoC genuinely unavailable
		"720": "0",  // battery 2 CAN SoC: literal 0
	})

	signals, err := ruptela.SignalsFromV1Payload(event)
	require.NoError(t, err)
	got := signalValues(signals)

	require.NotContains(t, got, vss.FieldPowertrainTractionBatteryStateOfChargeCurrent)
}
