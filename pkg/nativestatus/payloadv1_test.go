package nativestatus_test

import (
	"cmp"
	"context"
	"slices"
	"testing"
	"time"

	"github.com/DIMO-Network/model-garage/pkg/nativestatus"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/stretchr/testify/require"
)

type tokenGetter struct{}

func (*tokenGetter) TokenIDFromSubject(context.Context, string) (uint32, error) {
	return 123, nil
}

func TestFullFromDataConversion(t *testing.T) {
	t.Parallel()
	getter := &tokenGetter{}
	actualSignals, err := nativestatus.SignalsFromPayload(context.Background(), getter, []byte(fullInputJSON))
	require.NoErrorf(t, err, "error converting full input data: %v", err)

	// sort the signals so diffs are easier to read
	sortFunc := func(a, b vss.Signal) int {
		return cmp.Compare(a.Name, b.Name)
	}
	slices.SortFunc(expectedSignals, sortFunc)
	slices.SortFunc(actualSignals, sortFunc)

	require.Equal(t, expectedSignals, actualSignals, "converted vehicle does not match expected vehicle")
}

var (
	fullInputJSON = `{
		"id": "randomIDnumber",
		"specversion": "1.0",
		"dataschema": "test.status/v1",
		"source": "dimo/integration/123",
		"subject": "Vehicle123",
		"time": "2022-01-01T12:34:56Z",
		"type": "DIMO",
		"data": {
			"tires": {
				"frontLeft": 30.5,
				"frontRight": 31.0,
				"backLeft": 32.2,
				"backRight": 33.1
			},
			"charger": {
				"power": 34.0
			},
			"altitude": 100.0,
			"latitude": 37.7749,
			"longitude": -122.4194,
			"timestamp": "2022-01-01T12:34:56Z",
			"definitionID": "123",
			"iD": "456",
			"ambientTemp": 25.0,
			"batteryVoltage": 12.5,
			"barometricPressure": 1013.25,
			"engineLoad": 0.75,
			"intakeTemp": 30.0,
			"runTime": 1200.0,
			"coolantTemp": 90.0,
			"oil": 0.10,
			"engineSpeed": 3000.0,
			"throttlePosition": 0.50,
			"fuelPercentRemaining": 0.6,
			"fuelType": "Gasoline",
			"range": 300.0,
			"chargeLimit": 0.8,
			"charging": true,
			"batteryCapacity": 60.0,
			"soc": 0.7,
			"odometer": 50000.0,
			"speed": 60.0,
			"make": "Toyota",
			"model": "Camry",
			"year": 2020,
			"vin": "1234567890",
			"isRedacted": true
		},
	}`
	ts = time.Date(2022, 1, 1, 12, 34, 56, 0, time.UTC)

	expectedSignals = []vss.Signal{
		{Subject: "Vehicle123", Timestamp: ts, Name: "chassisAxleRow1WheelLeftTirePressure", ValueNumber: 30.5, Source: "dimo/integration/123"},
		{Subject: "Vehicle123", Timestamp: ts, Name: "chassisAxleRow1WheelRightTirePressure", ValueNumber: 31, Source: "dimo/integration/123"},
		{Subject: "Vehicle123", Timestamp: ts, Name: "chassisAxleRow2WheelLeftTirePressure", ValueNumber: 32.2, Source: "dimo/integration/123"},
		{Subject: "Vehicle123", Timestamp: ts, Name: "chassisAxleRow2WheelRightTirePressure", ValueNumber: 33.1, Source: "dimo/integration/123"},
		{Subject: "Vehicle123", Timestamp: ts, Name: "powertrainTractionBatteryCurrentPower", ValueNumber: 34000.0, Source: "dimo/integration/123"},
		{Subject: "Vehicle123", Timestamp: ts, Name: "currentLocationAltitude", ValueNumber: 100, Source: "dimo/integration/123"},
		{Subject: "Vehicle123", Timestamp: ts, Name: "currentLocationLatitude", ValueNumber: 37.7749, Source: "dimo/integration/123"},
		{Subject: "Vehicle123", Timestamp: ts, Name: "currentLocationLongitude", ValueNumber: -122.4194, Source: "dimo/integration/123"},
		{Subject: "Vehicle123", Timestamp: ts, Name: "currentLocationIsRedacted", ValueNumber: 1, Source: "dimo/integration/123"},
		{Subject: "Vehicle123", Timestamp: ts, Name: "powertrainCombustionEngineECT", ValueNumber: 90, Source: "dimo/integration/123"},
		{Subject: "Vehicle123", Timestamp: ts, Name: "powertrainCombustionEngineEngineOilLevel", ValueString: "CRITICALLY_LOW", Source: "dimo/integration/123"},
		{Subject: "Vehicle123", Timestamp: ts, Name: "powertrainCombustionEngineEngineOilRelativeLevel", ValueNumber: 10, Source: "dimo/integration/123"},
		{Subject: "Vehicle123", Timestamp: ts, Name: "powertrainCombustionEngineSpeed", ValueNumber: 3000, Source: "dimo/integration/123"},
		{Subject: "Vehicle123", Timestamp: ts, Name: "powertrainCombustionEngineTPS", ValueNumber: 50, Source: "dimo/integration/123"},
		{Subject: "Vehicle123", Timestamp: ts, Name: "powertrainFuelSystemRelativeLevel", ValueNumber: 60, Source: "dimo/integration/123"},
		{Subject: "Vehicle123", Timestamp: ts, Name: "powertrainFuelSystemSupportedFuelTypes", ValueString: "GASOLINE", Source: "dimo/integration/123"},
		{Subject: "Vehicle123", Timestamp: ts, Name: "powertrainRange", ValueNumber: 300, Source: "dimo/integration/123"},
		{Subject: "Vehicle123", Timestamp: ts, Name: "powertrainType", ValueString: "COMBUSTION", Source: "dimo/integration/123"},
		{Subject: "Vehicle123", Timestamp: ts, Name: "powertrainTractionBatteryChargingChargeLimit", ValueNumber: 80, Source: "dimo/integration/123"},
		{Subject: "Vehicle123", Timestamp: ts, Name: "powertrainTractionBatteryChargingIsCharging", ValueNumber: 1, Source: "dimo/integration/123"},
		{Subject: "Vehicle123", Timestamp: ts, Name: "powertrainTractionBatteryGrossCapacity", ValueNumber: 60, Source: "dimo/integration/123"},
		{Subject: "Vehicle123", Timestamp: ts, Name: "powertrainTractionBatteryStateOfChargeCurrent", ValueNumber: 70, Source: "dimo/integration/123"},
		{Subject: "Vehicle123", Timestamp: ts, Name: "powertrainTransmissionTravelledDistance", ValueNumber: 50000, Source: "dimo/integration/123"},
		{Subject: "Vehicle123", Timestamp: ts, Name: "speed", ValueNumber: 60, Source: "dimo/integration/123"},
		{Subject: "Vehicle123", Timestamp: ts, Name: "exteriorAirTemperature", ValueNumber: 25, Source: "dimo/integration/123"},
		{Subject: "Vehicle123", Timestamp: ts, Name: "lowVoltageBatteryCurrentVoltage", ValueNumber: 12.5, Source: "dimo/integration/123"},
		{Subject: "Vehicle123", Timestamp: ts, Name: "obdBarometricPressure", ValueNumber: 1013.25, Source: "dimo/integration/123"},
		{Subject: "Vehicle123", Timestamp: ts, Name: "obdEngineLoad", ValueNumber: 75, Source: "dimo/integration/123"},
		{Subject: "Vehicle123", Timestamp: ts, Name: "obdIntakeTemp", ValueNumber: 30, Source: "dimo/integration/123"},
		{Subject: "Vehicle123", Timestamp: ts, Name: "obdRunTime", ValueNumber: 1200, Source: "dimo/integration/123"},
	}
)

func TestSkipNulls(t *testing.T) {
	t.Parallel()
	getter := &tokenGetter{}
	actualSignals, err := nativestatus.SignalsFromPayload(context.Background(), getter, []byte(inputJSONWithNull))
	require.NoErrorf(t, err, "error converting input data: %v", err)
	require.ElementsMatchf(t, expectedSignalsWithoutNull, actualSignals, "converted vehicle does not match expected vehicle")
}

var (
	inputJSONWithNull = `{
		"id": "randomIDnumber",
		"specversion": "1.0",
		"source": "dimo/integration/123",
		"subject": "Vehicle123",
		"time": "2022-01-01T12:34:56Z",
		"type": "DIMO",
		"data": {
			"odometer": 5024.0,
			"range": null,
			"speed": 25.0
		}
	}`

	expectedSignalsWithoutNull = []vss.Signal{
		{Subject: "Vehicle123", Timestamp: ts, Name: "speed", ValueNumber: 25.0, Source: "dimo/integration/123"},
		{Subject: "Vehicle123", Timestamp: ts, Name: "powertrainTransmissionTravelledDistance", ValueNumber: 5024, Source: "dimo/integration/123"},
	}
)

func TestWithTokenId(t *testing.T) {
	t.Parallel()
	actualSignals, err := nativestatus.SignalsFromPayload(context.Background(), nil, []byte(inputJSONWithTokenID))
	require.NoErrorf(t, err, "error converting input data: %v", err)
	require.ElementsMatchf(t, expectedSignalsWithFromTokenID, actualSignals, "converted vehicle does not match expected vehicle")
}

var (
	inputJSONWithTokenID = `{
		"id": "randomIDnumber",
		"specversion": "1.0",
		"source": "dimo/integration/123",
		"subject": "Vehicle123",
		"time": "2022-01-01T12:34:56Z",
		"type": "DIMO",
		"vehicleTokenId": 123,
		"data": {
			"odometer": 5024.0,
			"speed": 25.0
		}
	}`

	expectedSignalsWithFromTokenID = []vss.Signal{
		{Subject: "Vehicle123", Timestamp: ts, Name: "speed", ValueNumber: 25.0, Source: "dimo/integration/123"},
		{Subject: "Vehicle123", Timestamp: ts, Name: "powertrainTransmissionTravelledDistance", ValueNumber: 5024, Source: "dimo/integration/123"},
	}
)
