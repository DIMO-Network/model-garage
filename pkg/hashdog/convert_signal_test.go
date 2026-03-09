package hashdog

import (
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

func TestFullFromV2DataConversion(t *testing.T) {
	t.Parallel()
	var event cloudevent.RawEvent
	err := json.Unmarshal([]byte(fullV2InputJSON), &event)
	require.NoErrorf(t, err, "error unmarshalling full input data: %v", err)
	originalID := event.ID
	baseHeader := event.CloudEventHeader
	actualSignals, err := SignalsFromV2Payload(event)

	expectedSignals := expectedV2Signals(baseHeader)
	require.NoErrorf(t, err, "error converting full input data: %v", err)
	require.Len(t, actualSignals, len(expectedSignals), "actual signals length does not match expected")
	assertAndNormalizeSignals(t, actualSignals, originalID)
	normalizeExpectedSignals(expectedSignals)
	require.Equalf(t, expectedSignals, actualSignals, "converted vehicle does not match expected vehicle")
}

var fullV2InputJSON = `{
    "id": "2fHbFXPWzrVActDb7WqWCfqeiYe",
    "source": "dimo/integration/123",
    "specversion": "1.0",
	"dataversion": "v2",
    "subject": "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:33",
    "time": "2024-04-18T17:20:46.436008782Z",
    "type": "dimo.status",
    "signature": "0x72208df3282c890ec72afe22abbcffb76ec73dc3e1ce8becd158469126f10c35245289e02ad41782e07376f9b9092a2fec96477a6e453fed1ca3860639e776f31b",
    "data": {
        "timestamp": 1713460846435,
        "device": {
            "rpiUptimeSecs": 218,
            "batteryVoltage": 12.28
        },
        "vehicle": {
            "signals": [
                {
                    "timestamp": 1713460823243,
                    "name": "longTermFuelTrim1",
                    "value": 25
                },
                {
                    "timestamp": 1713460826633,
                    "name": "coolantTemp",
                    "value": 107
                },
                {
                    "timestamp": 1713460827173,
                    "name": "maf",
                    "value": 475.79
                },
                {
                    "timestamp": 1713460829314,
                    "name": "engineLoad",
                    "value": 12.54912
                },
                {
                    "timestamp": 1713460829844,
                    "name": "throttlePosition",
                    "value": 23.529600000000002
                },
                {
                    "timestamp": 1713460830382,
                    "name": "shortTermFuelTrim1",
                    "value": 12.5
                },
                {
                    "timestamp": 1713460837235,
                    "name": "throttlePosition",
                    "value": 23.529600000000002
                },
                {
                    "timestamp": 1713460842256,
                    "name": "maf",
                    "value": 475.79
                },
                {
                    "timestamp": 1713460844422,
                    "name": "engineLoad",
                    "value": 12.54912
                },
                {
                    "timestamp": 1713460844962,
                    "name": "throttlePosition",
                    "value": 23.529600000000002
                },
                {
                    "timestamp": 1713460845497,
                    "name": "shortTermFuelTrim1",
                    "value": 12.5
                },
                {
                    "timestamp": 1713460846435,
                    "name": "isRedacted",
                    "value": false
                },
                {
                    "timestamp": 1713460846435,
                    "name": "longitude",
                    "value": -56.50151833333334
                },
                {
                    "timestamp": 1713460846435,
                    "name": "latitude",
                    "value": 56.27014
                },
                {
                    "timestamp": 1713460846435,
                    "name": "hdop",
                    "value": 1.4
                },
                {
                    "timestamp": 1713460846435,
                    "name": "nsat",
                    "value": 6
                },
                {
                    "timestamp": 1713460846435,
                    "name": "wpa_state",
                    "value": "COMPLETED"
                },
                {
                    "timestamp": 1713460846435,
                    "name": "ssid",
                    "value": "foo"
                },
                {
                    "timestamp": 1713460846435,
                    "name": "vehicleSpeed",
                    "value": 39
                },
                {
                    "timestamp": 1713460846435,
                    "name": "rpm",
                    "value": 2000
                },
                {
                    "timestamp": 1713460846435,
                    "name": "fuelLevel",
                    "value": 50
                }
            ]
        }
    },
    "vehicleTokenId": 123,
    "make": "",
    "model": "",
    "year": 0
}`

func expectedV2Signals(hdr cloudevent.CloudEventHeader) []vss.Signal {
	return []vss.Signal{
		{CloudEventHeader: hdr, Data: vss.SignalData{Timestamp: time.Date(2024, time.April, 18, 17, 20, 23, 243000000, time.UTC), Name: "obdLongTermFuelTrim1", ValueNumber: 25, ValueString: ""}},
		{CloudEventHeader: hdr, Data: vss.SignalData{Timestamp: time.Date(2024, time.April, 18, 17, 20, 26, 633000000, time.UTC), Name: "powertrainCombustionEngineECT", ValueNumber: 107, ValueString: ""}},
		{CloudEventHeader: hdr, Data: vss.SignalData{Timestamp: time.Date(2024, time.April, 18, 17, 20, 27, 173000000, time.UTC), Name: "powertrainCombustionEngineMAF", ValueNumber: 475.79, ValueString: ""}},
		{CloudEventHeader: hdr, Data: vss.SignalData{Timestamp: time.Date(2024, time.April, 18, 17, 20, 29, 314000000, time.UTC), Name: "obdEngineLoad", ValueNumber: 12.54912, ValueString: ""}},
		{CloudEventHeader: hdr, Data: vss.SignalData{Timestamp: time.Date(2024, time.April, 18, 17, 20, 29, 844000000, time.UTC), Name: "powertrainCombustionEngineTPS", ValueNumber: 23.529600000000002, ValueString: ""}},
		{CloudEventHeader: hdr, Data: vss.SignalData{Timestamp: time.Date(2024, time.April, 18, 17, 20, 30, 382000000, time.UTC), Name: "obdShortTermFuelTrim1", ValueNumber: 12.5, ValueString: ""}},
		{CloudEventHeader: hdr, Data: vss.SignalData{Timestamp: time.Date(2024, time.April, 18, 17, 20, 37, 235000000, time.UTC), Name: "powertrainCombustionEngineTPS", ValueNumber: 23.529600000000002, ValueString: ""}},
		{CloudEventHeader: hdr, Data: vss.SignalData{Timestamp: time.Date(2024, time.April, 18, 17, 20, 42, 256000000, time.UTC), Name: "powertrainCombustionEngineMAF", ValueNumber: 475.79, ValueString: ""}},
		{CloudEventHeader: hdr, Data: vss.SignalData{Timestamp: time.Date(2024, time.April, 18, 17, 20, 44, 422000000, time.UTC), Name: "obdEngineLoad", ValueNumber: 12.54912, ValueString: ""}},
		{CloudEventHeader: hdr, Data: vss.SignalData{Timestamp: time.Date(2024, time.April, 18, 17, 20, 44, 962000000, time.UTC), Name: "powertrainCombustionEngineTPS", ValueNumber: 23.529600000000002, ValueString: ""}},
		{CloudEventHeader: hdr, Data: vss.SignalData{Timestamp: time.Date(2024, time.April, 18, 17, 20, 45, 497000000, time.UTC), Name: "obdShortTermFuelTrim1", ValueNumber: 12.5, ValueString: ""}},
		{CloudEventHeader: hdr, Data: vss.SignalData{Timestamp: time.Date(2024, time.April, 18, 17, 20, 46, 435000000, time.UTC), Name: "currentLocationIsRedacted", ValueNumber: 0, ValueString: ""}},
		{CloudEventHeader: hdr, Data: vss.SignalData{Timestamp: time.Date(2024, time.April, 18, 17, 20, 46, 435000000, time.UTC), Name: "currentLocationLongitude", ValueNumber: -56.50151833333334, ValueString: ""}},
		{CloudEventHeader: hdr, Data: vss.SignalData{Timestamp: time.Date(2024, time.April, 18, 17, 20, 46, 435000000, time.UTC), Name: "currentLocationLatitude", ValueNumber: 56.27014, ValueString: ""}},
		{CloudEventHeader: hdr, Data: vss.SignalData{Timestamp: time.Date(2024, time.April, 18, 17, 20, 46, 435000000, time.UTC), Name: "dimoAftermarketHDOP", ValueNumber: 1.4, ValueString: ""}},
		{CloudEventHeader: hdr, Data: vss.SignalData{Timestamp: time.Date(2024, time.April, 18, 17, 20, 46, 435000000, time.UTC), Name: "speed", ValueNumber: 39, ValueString: ""}},
		{CloudEventHeader: hdr, Data: vss.SignalData{Timestamp: time.Date(2024, time.April, 18, 17, 20, 46, 435000000, time.UTC), Name: "powertrainCombustionEngineSpeed", ValueNumber: 2000, ValueString: ""}},
		{CloudEventHeader: hdr, Data: vss.SignalData{Timestamp: time.Date(2024, time.April, 18, 17, 20, 46, 435000000, time.UTC), Name: "powertrainFuelSystemRelativeLevel", ValueNumber: 50, ValueString: ""}},
	}
}

func TestNullSignals(t *testing.T) {
	t.Parallel()
	var event cloudevent.RawEvent
	err := json.Unmarshal([]byte(nilSignalsJSON), &event)
	require.NoError(t, err)
	actualSignals, err := SignalsFromV2Payload(event)
	require.NoErrorf(t, err, "error converting full input data: %v", err)
	require.Equalf(t, []vss.Signal{}, actualSignals, "converted vehicle does not match expected vehicle")
}

var nilSignalsJSON = `{
    "id": "2fHbFXPWzrVActDb7WqWCfqeiYe",
    "source": "dimo/integration/123",
    "specversion": "1.0",
    "subject": "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:33",
    "time": "2024-04-18T17:20:46.436008782Z",
    "type": "com.dimo.device.status",
    "vehicleTokenId": 123,
    "data": {
        "timestamp": 1713460846435,
        "device": {
            "rpiUptimeSecs": 218,
            "batteryVoltage": 12.28
        },
        "vehicle": {
            "signals": null
        }
    }
}`

func expectedDTCErrorsSignals(hdr cloudevent.CloudEventHeader) []vss.Signal {
	return []vss.Signal{
		{CloudEventHeader: hdr, Data: vss.SignalData{Timestamp: time.UnixMilli(1753387941604).UTC(), Name: vss.FieldOBDDTCList, ValueString: `["P1234","P1235"]`}},
		{CloudEventHeader: hdr, Data: vss.SignalData{Timestamp: time.UnixMilli(1753387941604).UTC(), Name: vss.FieldOBDStatusDTCCount, ValueNumber: 2}},
		{CloudEventHeader: hdr, Data: vss.SignalData{Timestamp: time.UnixMilli(1753387941605).UTC(), Name: vss.FieldOBDStatusDTCCount, ValueNumber: 0}},
	}
}

func TestDTCErrorCodesConversion(t *testing.T) {
	t.Parallel()
	var event cloudevent.RawEvent
	err := json.Unmarshal([]byte(dtcErrorsAPInputJSON), &event)
	require.NoError(t, err)
	originalID := event.ID
	baseHeader := event.CloudEventHeader
	actualSignals, err := SignalsFromV2Payload(event)

	expectedSignals := expectedDTCErrorsSignals(baseHeader)
	require.NoErrorf(t, err, "error converting full input data: %v", err)
	actualSignals = actualSignals[:len(expectedSignals)]
	require.Len(t, actualSignals, len(expectedSignals), "actual signals length does not match expected")
	assertAndNormalizeSignals(t, actualSignals, originalID)
	normalizeExpectedSignals(expectedSignals)
	require.Equalf(t, expectedSignals, actualSignals, "converted vehicle does not match expected vehicle")
}

var dtcErrorsAPInputJSON = `{
    "id": "30Kt9eBsiyIZEsPTxnW5ldsenCb",
    "subject": "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:33",
    "specversion": "1.0",
    "time": "2025-07-24T20:12:21.604Z",
    "datacontenttype": "application/json",
    "type": "com.dimo.device.status.v2",
    "dataschema": "dimo.zone.status/v2.0",
    "source": "dimo/integration/123",
    "data": {
        "vehicle": {
            "signals": [
                {
                    "timestamp": 1753387941604,
                    "name": "obdDTCList",
                    "value": "P1234,P1235"
                },
                {
                    "timestamp": 1753387941604,
                    "name": "obdStatusDTCCount",
                    "value": 2
                },
                {
                    "timestamp": 1753387941605,
                    "name": "obdStatusDTCCount",
                    "value": 0
                }
            ]
        }
    }
}`
