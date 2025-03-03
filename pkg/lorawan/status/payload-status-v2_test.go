package status

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/DIMO-Network/model-garage/pkg/vss"
)

func TestFullFromV2DataConversion(t *testing.T) {
	t.Parallel()
	actualSignals, err := SignalsFromV2Payload([]byte(fullV2InputJSON))

	require.NoErrorf(t, err, "error converting full input data: %v", err)
	require.Len(t, actualSignals, len(expectedV2Signals), "actual signals length does not match expected")
	require.Equalf(t, expectedV2Signals, actualSignals, "converted vehicle does not match expected vehicle")
}

var fullV2InputJSON = `{
    "id": "2fHbFXPWzrVActDb7WqWCfqeiYe",
    "source": "dimo/integration/123",
    "specversion": "1.0",
	"dataversion": "v2",
    "subject": "did:nft:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF_33",
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
                },
            ]
        }
    },
    "vehicleTokenId": 123,
    "make": "",
    "model": "",
    "year": 0
}`

var (
	tokenID           = uint32(33)
	expectedV2Signals = []vss.Signal{
		{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 23, 243000000, time.UTC), Name: "obdLongTermFuelTrim1", ValueNumber: 25, ValueString: "", Source: "dimo/integration/123"},
		{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 26, 633000000, time.UTC), Name: "powertrainCombustionEngineECT", ValueNumber: 107, ValueString: "", Source: "dimo/integration/123"},
		{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 27, 173000000, time.UTC), Name: "powertrainCombustionEngineMAF", ValueNumber: 475.79, ValueString: "", Source: "dimo/integration/123"},
		{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 29, 314000000, time.UTC), Name: "obdEngineLoad", ValueNumber: 12.54912, ValueString: "", Source: "dimo/integration/123"},
		{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 29, 844000000, time.UTC), Name: "powertrainCombustionEngineTPS", ValueNumber: 23.529600000000002, ValueString: "", Source: "dimo/integration/123"},
		{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 30, 382000000, time.UTC), Name: "obdShortTermFuelTrim1", ValueNumber: 12.5, ValueString: "", Source: "dimo/integration/123"},
		{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 37, 235000000, time.UTC), Name: "powertrainCombustionEngineTPS", ValueNumber: 23.529600000000002, ValueString: "", Source: "dimo/integration/123"},
		{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 42, 256000000, time.UTC), Name: "powertrainCombustionEngineMAF", ValueNumber: 475.79, ValueString: "", Source: "dimo/integration/123"},
		{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 44, 422000000, time.UTC), Name: "obdEngineLoad", ValueNumber: 12.54912, ValueString: "", Source: "dimo/integration/123"},
		{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 44, 962000000, time.UTC), Name: "powertrainCombustionEngineTPS", ValueNumber: 23.529600000000002, ValueString: "", Source: "dimo/integration/123"},
		{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 45, 497000000, time.UTC), Name: "obdShortTermFuelTrim1", ValueNumber: 12.5, ValueString: "", Source: "dimo/integration/123"},
		{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 46, 435000000, time.UTC), Name: "currentLocationIsRedacted", ValueNumber: 0, ValueString: "", Source: "dimo/integration/123"},
		{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 46, 435000000, time.UTC), Name: "currentLocationLongitude", ValueNumber: -56.50151833333334, ValueString: "", Source: "dimo/integration/123"},
		{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 46, 435000000, time.UTC), Name: "currentLocationLatitude", ValueNumber: 56.27014, ValueString: "", Source: "dimo/integration/123"},
		{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 46, 435000000, time.UTC), Name: "dimoAftermarketHDOP", ValueNumber: 1.4, ValueString: "", Source: "dimo/integration/123"},
		{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 46, 435000000, time.UTC), Name: "dimoAftermarketNSAT", ValueNumber: 6, ValueString: "", Source: "dimo/integration/123"},
		{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 46, 435000000, time.UTC), Name: "dimoAftermarketWPAState", ValueNumber: 0, ValueString: "COMPLETED", Source: "dimo/integration/123"},
		{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 46, 435000000, time.UTC), Name: "dimoAftermarketSSID", ValueNumber: 0, ValueString: "foo", Source: "dimo/integration/123"},
		{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 46, 435000000, time.UTC), Name: "speed", ValueNumber: 39, ValueString: "", Source: "dimo/integration/123"},
		{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 46, 435000000, time.UTC), Name: "powertrainCombustionEngineSpeed", ValueNumber: 2000, ValueString: "", Source: "dimo/integration/123"},
		{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 46, 435000000, time.UTC), Name: "powertrainFuelSystemRelativeLevel", ValueNumber: 50, ValueString: "", Source: "dimo/integration/123"},
	}
)

func TestNullSignals(t *testing.T) {
	t.Parallel()
	actualSignals, err := SignalsFromV2Payload([]byte(nilSignalsJSON))
	require.NoErrorf(t, err, "error converting full input data: %v", err)
	require.Equalf(t, []vss.Signal{}, actualSignals, "converted vehicle does not match expected vehicle")
}

var nilSignalsJSON = `{
    "id": "2fHbFXPWzrVActDb7WqWCfqeiYe",
    "source": "dimo/integration/123",
    "specversion": "1.0",
    "subject": "did:nft:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF_33",
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
    },
}`
