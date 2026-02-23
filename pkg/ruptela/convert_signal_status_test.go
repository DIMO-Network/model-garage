package ruptela_test

import (
	"cmp"
	"encoding/json"
	"slices"
	"testing"
	"time"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/ruptela"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/stretchr/testify/require"
)

func TestFullFromDataConversion(t *testing.T) {
	t.Parallel()
	var event cloudevent.RawEvent
	err := json.Unmarshal([]byte(fullInputJSON), &event)
	require.NoError(t, err)
	actualSignals, err := ruptela.SignalsFromV1Payload(event)
	require.NoErrorf(t, err, "error converting full input data: %v", err)

	// sort the signals so diffs are easier to read
	sortFunc := func(a, b vss.Signal) int {
		return cmp.Compare(a.Name, b.Name)
	}
	ts := time.Date(2024, 9, 27, 8, 33, 26, 0, time.UTC)

	expectedSignals := []vss.Signal{
		{TokenID: 33, Timestamp: ts, Name: vss.FieldLowVoltageBatteryCurrentVoltage, ValueNumber: 14.335, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldPowertrainFuelSystemAbsoluteLevel, ValueNumber: 5, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldCurrentLocationAltitude, ValueNumber: 104.8, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldDIMOAftermarketHDOP, ValueNumber: 0.6, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldCurrentLocationLatitude, ValueNumber: 52.2721466, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldCurrentLocationLongitude, ValueNumber: -0.9014316, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldPowertrainType, ValueString: "COMBUSTION", Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldPowertrainFuelSystemRelativeLevel, ValueNumber: 19.200000000000003, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldOBDDistanceWithMIL, ValueNumber: 0, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldPowertrainCombustionEngineTPS, ValueNumber: 18, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldPowertrainTransmissionTravelledDistance, ValueNumber: 8, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldSpeed, ValueNumber: 31.24609375, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldPowertrainTractionBatteryRange, ValueNumber: 31, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldChassisAxleRow1WheelLeftTirePressure, ValueNumber: 262.00088, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldChassisAxleRow1WheelRightTirePressure, ValueNumber: 310.2642, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldChassisAxleRow2WheelLeftTirePressure, ValueNumber: 282.68516, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldChassisAxleRow2WheelRightTirePressure, ValueNumber: 310.2642, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldPowertrainCombustionEngineEngineOilLevel, ValueString: "LOW", Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldPowertrainCombustionEngineEngineOilRelativeLevel, ValueNumber: 36.800000000000004, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldCurrentLocationHeading, ValueNumber: 73.7, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldOBDStatusDTCCount, ValueNumber: 18, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldPowertrainTractionBatteryStateOfHealth, ValueNumber: 98.5, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldPowertrainTractionBatteryChargingPower, ValueNumber: 24.400000000000002, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldPowertrainTractionBatteryChargingIsChargingCableConnected, ValueNumber: 1, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldIsIgnitionOn, ValueNumber: 1, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldOBDIsPluggedIn, ValueNumber: 1, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldOBDIsEngineBlocked, ValueNumber: 0, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldOBDFuelTypeName, ValueString: "GASOLINE", Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldBodyLightsIsAirbagWarningOn, ValueNumber: 1, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldChassisBrakeABSIsWarningOn, ValueNumber: 0, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldChassisTireSystemIsWarningOn, ValueNumber: 1, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldChassisBrakeCircuit1PressurePrimary, ValueNumber: 376.00, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldChassisBrakeCircuit2PressurePrimary, ValueNumber: 1016.00, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldPowertrainCombustionEngineSpeed, ValueNumber: 996.125, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldExteriorAirTemperature, ValueNumber: 103.59375, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldPowertrainTractionBatteryStateOfChargeCurrent, ValueNumber: 55, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldPowertrainCombustionEngineECT, ValueNumber: 85, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldPowertrainCombustionEngineDieselExhaustFluidLevel, ValueNumber: 24.400000000000002, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldPowertrainCombustionEngineEOP, ValueNumber: 368, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldOBDOilTemperature, ValueNumber: 103.59375, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldPowertrainTractionBatteryChargingIsCharging, ValueNumber: 1, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldChassisBrakePedalPosition, ValueNumber: 24, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldOBDEngineLoad, ValueNumber: 60, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldOBDIsPTOActive, ValueNumber: 1, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldChassisParkingBrakeIsEngaged, ValueNumber: 1, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldOBDFuelRate, ValueNumber: 15, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldChassisAxleRow3Weight, ValueNumber: 106, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldChassisAxleRow4Weight, ValueNumber: 22, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldChassisAxleRow5Weight, ValueNumber: 10, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldChassisBrakeIsPedalPressed, ValueNumber: 0, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldPowertrainFuelSystemAccumulatedConsumption, ValueNumber: 3990.5, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldPowertrainCombustionEngineTorquePercent, ValueNumber: 88, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldPowertrainTransmissionRetarderActualTorque, ValueNumber: 72, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldPowertrainTransmissionCurrentGear, ValueNumber: -1, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldPowertrainTransmissionSelectedGear, ValueNumber: 126, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldServiceTimeToService, ValueNumber: 2419200, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldCabinDoorRow1DriverSideIsOpen, ValueNumber: 1, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldCabinDoorRow1PassengerSideIsOpen, ValueNumber: 1, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldCabinDoorRow2DriverSideIsOpen, ValueNumber: 0, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldCabinDoorRow2PassengerSideIsOpen, ValueNumber: 0, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldBodyTrunkFrontIsOpen, ValueNumber: 0, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldBodyTrunkRearIsOpen, ValueNumber: 1, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldCabinSeatRow1DriverSideIsBelted, ValueNumber: 1, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldCabinSeatRow1PassengerSideIsBelted, ValueNumber: 1, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldCabinSeatRow2DriverSideIsBelted, ValueNumber: 1, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldCabinSeatRow2MiddleIsBelted, ValueNumber: 1, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldCabinSeatRow2PassengerSideIsBelted, ValueNumber: 1, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldCabinSeatRow3DriverSideIsBelted, ValueNumber: 1, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldCabinSeatRow3PassengerSideIsBelted, ValueNumber: 1, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldPowertrainTransmissionRetarderTorqueMode, ValueString: "HIGH SPEED GOVERNOR", Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldBodyLockIsLocked, ValueNumber: 1, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldPowertrainTransmissionIsClutchSwitchOperated, ValueNumber: 1, Source: "ruptela/TODO"},
		{TokenID: 33, Timestamp: ts, Name: vss.FieldConnectivityCellularIsJammingDetected, ValueNumber: 0, Source: "ruptela/TODO"},
	}

	slices.SortFunc(expectedSignals, sortFunc)
	slices.SortFunc(actualSignals, sortFunc)
	require.Equal(t, expectedSignals, actualSignals, "converted vehicle does not match expected vehicle")
}

var fullInputJSON = `
{
	"source": "ruptela/TODO",
	"data": {
		"pos": {
			"alt": 1048,
			"dir": 7370,
			"hdop": 6,
			"lat": 522721466,
			"lon": -9014316,
			"sat": 20,
			"spd": 65535
		},
		"prt": 0,
		"signals": {
			"35": "1",
			"88": "0", 
			"100": "012C",
			"101": "D5",
			"359": "9",
			"538": "1",
			"347": "3",
			"109": "7F",
			"520": "9D80",
			"360": "C5",
			"518": "23",
			"519": "FF",
			"36": "0",
			"92": "FFFFFFFF",
			"208": "1F2D",
			"313": "0",
			"316": "1",
			"367": "7C",
			"366": "FB",
			"102": "0",
			"103": "FF",
			"104": "53414C4C41414146",
			"105": "3341413534343438",
			"106": "3200000000000000",
			"107": "0",
			"108": "12",
			"645": "8",
			"135": "0",
			"136": "0",
			"137": "14",
			"173": "1",
			"205": "5",
			"207": "30",
			"29": "37FF",
			"30": "1080",
			"405": "1",
			"115": "7D",
			"355": "2F",
			"356": "7F",
			"516": "1F",
			"210": "1F3F",
			"206": "2D",
			"197": "1F21",
			"89": "2F13",
			"522": "1",
			"515": "FF",
			"722": "37",
			"409": "FF",
			"201": "FF",
			"357": "3D",
			"1187": "5C",
			"728": "5C",
			"727": "2F13",
			"517": "1",
			"482": "3C",
			"39": "3C",	
			"38": "5",
			"362": "1",
			"116": "FFFF", 
			"53": "D4",
			"54": "2C",
			"55": "14",
			"49": "FE",
			"50": "FA",
			"5005": "31",
			"5060": "6597",
			"51": "ED",
			"525": "A502A",
			"525_1": "A502A",
			"642": "FFFF",
			"645": "FFFFFFFF",
			"722": "FF",
			"754": "FB8F",
			"93": "0",
			"94": "FFFF",
			"950": "267A",
			"96": "FF",
			"95": "FF",
			"97": "FF",
			"98": "0",
			"964": "FF",
			"965": "26",
			"966": "2D",
			"967": "29",
			"968": "2D",
			"985": "0",
			"99": "1",
			"999": "0",
			"1190": "5F50",
			"1191": "1"
		},
		"trigger": 7
	},
	"ds": "r/v0/s",
	"signature": "0x6fb5849e21e66f3e0619f148bc032153aa4c90be4cd175e83c1f959e1bc551d940d516fe74f50aed380e432406675c583e75155bf1c77b9ec0761b1dbe1ab87e1c",
	"subject": "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:33",
	"time": "2024-09-27T08:33:26Z"
}`

func TestIgnoreUnplugged(t *testing.T) {
	t.Parallel()
	var event cloudevent.RawEvent
	err := json.Unmarshal([]byte(ignoreTestDoc), &event)
	require.NoError(t, err)
	actualSignals, err := ruptela.SignalsFromV1Payload(event)
	require.NoError(t, err)
	expectedSignals := []vss.Signal{
		{TokenID: 162682, Timestamp: time.Date(2025, 3, 28, 0, 51, 29, 0, time.UTC), Name: vss.FieldIsIgnitionOn, ValueNumber: 1, Source: "0xF26421509Efe92861a587482100c6d728aBf1CD0"},
		{TokenID: 162682, Timestamp: time.Date(2025, 3, 28, 0, 51, 29, 0, time.UTC), Name: vss.FieldOBDDistanceWithMIL, ValueNumber: 0, Source: "0xF26421509Efe92861a587482100c6d728aBf1CD0"},
		{TokenID: 162682, Timestamp: time.Date(2025, 3, 28, 0, 51, 29, 0, time.UTC), Name: vss.FieldOBDIsPluggedIn, ValueNumber: 0, Source: "0xF26421509Efe92861a587482100c6d728aBf1CD0"},
		{TokenID: 162682, Timestamp: time.Date(2025, 3, 28, 0, 51, 29, 0, time.UTC), Name: vss.FieldOBDStatusDTCCount, ValueNumber: 0, Source: "0xF26421509Efe92861a587482100c6d728aBf1CD0"},
		{TokenID: 162682, Timestamp: time.Date(2025, 3, 28, 0, 51, 29, 0, time.UTC), Name: vss.FieldPowertrainCombustionEngineTPS, ValueNumber: 0, Source: "0xF26421509Efe92861a587482100c6d728aBf1CD0"},
		{TokenID: 162682, Timestamp: time.Date(2025, 3, 28, 0, 51, 29, 0, time.UTC), Name: vss.FieldSpeed, ValueNumber: 5, Source: "0xF26421509Efe92861a587482100c6d728aBf1CD0"},
	}
	require.Equal(t, expectedSignals, actualSignals)
}

var ignoreTestDoc = `
{
  "id": "2uvJPjThhoJSwNulNpvfe6xwkpF",
  "source": "0xF26421509Efe92861a587482100c6d728aBf1CD0",
  "producer": "did:erc721:137:0x9c94C395cBcBDe662235E0A9d3bB87Ad708561BA:31648",
  "specversion": "1.0",
  "subject": "did:erc721:137:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:162682",
  "time": "2025-03-28T00:51:29Z",
  "type": "dimo.status",
  "datacontenttype": "application/json",
  "dataversion": "r/v0/s",
  "signature": "0x8dc9123f76361fadefa7753e3f299cc2597c26839f22d4f0983d2898e72fbdf733d1665a109a09ff5c81f0b214af29a97eb8e4d6383cb2035f543c449a75ab501b",
  "data": {
    "trigger": 7,
    "prt": 0,
    "pos": {
		"spd": 0
    },
    "signals": {
      "27": "0",
      "29": "186",
      "30": "102E",
      "93": "0",
      "94": "0",
      "95": "5",
      "96": "FF",
      "97": "FF",
      "98": "0",
      "99": "0",
      "102": "0",
      "103": "0",
      "104": "0",
      "105": "0",
      "106": "0",
      "107": "0",
      "108": "0",
      "134": "0",
      "135": "0",
      "136": "0",
      "137": "0",
      "169": "0",
      "173": "1",
      "402": "0",
      "403": "0",
      "404": "0",
      "409": "1",
      "418": "0",
      "525": "561AD",
      "642": "FFFF",
      "645": "FFFFFFFF",
      "722": "FF",
      "723": "FFFF",
      "754": "D30C3A",
      "762": "FF",
      "763": "0",
      "950": "0",
      "960": "FFFF",
      "961": "FFFF",
      "962": "FFFF",
      "963": "FFFF",
      "964": "FF",
      "985": "1",
      "999": "FF",
      "1148": "FF",
      "1149": "FF",
      "1150": "FF",
      "5005": "0",
      "5060": "0",
      "5114": "6666",
      "5115": "444",
      "5116": "6"
    }
  }
}`
