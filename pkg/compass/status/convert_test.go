package status

import (
	"testing"
	"time"

	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var baseDoc = []byte(`
{
   "subject":"did:nft:137:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF_37",
   "source":"0x983110309620D911731Ac0932219af06091b6744",
   "data":{
      "id":"Ktd6b9VETWCtWr07joo74Q==",
      "vehicle_id":"1C4SJSBP8RS133747",
      "timestamp":{
         "seconds":1737988801
      },
      "transport_type":0,
      "vehicle_type":0,
      "position":{
         "latlng":{
            "lat":34.821394,
            "lng":-82.29141
         },
         "speed":109.6
      },
      "acceleration":{
         "x":0.34,
         "y":0.22
      },
      "labels":{
         "engine.oil.pressure.unit":"bar",
         "tirePressure.front.left.unit":"psi",
         "engine.status":"off",
         "geolocation.altitude.unit":"m",
         "engine.oil.pressure.value":"0.04",
         "engine.oil.lifeLeft.percentage":"90",
         "datetime":"2025-02-06T15:18:22.243Z",
         "odometer.value":"20446",
         "vin":"1C4SJSBP8RS133747",
         "geolocation.latitude":"34.878016",
         "moving":"false",
         "engine.speed.unit":"rpm",
         "fuel.residualAutonomy.value":"733",
         "transmissionGear.state":"p",
         "seatbelt.passenger.front":"false",
         "geolocation.longitude":"-82.223566",
         "doors.open.passenger.front":"false",
         "adas.abs":"false",
         "engine.battery.voltage.value":"13",
         "doors.open.driver":"false",
         "fuel.level.percentage":"99",
         "engine.speed.value":"0",
         "lights.fog.front":"false",
         "tirePressure.front.right.unit":"psi",
         "speed.unit":"km/h",
         "engine.ignition":"false",
         "engine.contact":"false",
         "_id":"6786afd0ff54c000078aa67a",
         "tirePressure.rear.right.unit":"psi",
         "status":"halted",
         "engine.coolant.temperature.value":"92",
         "engine.battery.charging":"false",
         "speed.value":"0",
         "fuel.averageConsumption.value":"6.4",
         "odometer.unit":"km",
         "fuel.level.unit":"L",
         "engine.battery.voltage.unit":"V",
         "engine.oil.temperature.value":"89",
         "tirePressure.front.right.value":"41",
         "engine.oil.temperature.unit":"°C",
         "fuel.averageConsumption.unit":"L/100 km",
         "seatbelt.driver":"false",
         "tirePressure.front.left.value":"41",
         "fuel.residualAutonomy.unit":"km",
         "doors.open.passenger.rear.left":"false",
         "tirePressure.rear.left.unit":"psi",
         "heading":"16",
         "lights.fog.rear":"false",
         "doors.open.passenger.rear.right":"false",
         "tirePressure.rear.left.value":"42",
         "fuel.level.value":"113.85",
         "engine.coolant.temperature.unit":"°C",
         "datetimeSending":"2025-02-06T15:18:22.243Z",
         "crash.autoEcall":"false",
         "tirePressure.rear.right.value":"41",
         "geolocation.altitude.value":"277.100006"
      },
      "ingested_at":{
         "seconds":1737988851,
         "nanos":384695000
      }
   }
}
`)

// TODO change this to the correct connection once available
const compassConnection = "0x983110309620D911731Ac0932219af06091b6744"

var expSignals = []vss.Signal{
	{TokenID: 37, Timestamp: time.UnixMilli(1730728805000), Name: "chassisAxleRow1WheelLeftTirePressure", ValueNumber: 312, Source: compassConnection},
	{TokenID: 37, Timestamp: time.UnixMilli(1730728805000), Name: "chassisAxleRow1WheelRightTirePressure", ValueNumber: 309, Source: compassConnection},
	{TokenID: 37, Timestamp: time.UnixMilli(1730728805000), Name: "chassisAxleRow2WheelLeftTirePressure", ValueNumber: 298, Source: compassConnection},
	{TokenID: 37, Timestamp: time.UnixMilli(1730728805000), Name: "chassisAxleRow2WheelRightTirePressure", ValueNumber: 299, Source: compassConnection},
	{TokenID: 37, Timestamp: time.UnixMilli(1730738800000), Name: "currentLocationLatitude", ValueNumber: 38.89, Source: compassConnection},
	{TokenID: 37, Timestamp: time.UnixMilli(1730738800000), Name: "currentLocationLongitude", ValueNumber: 77.03, Source: compassConnection},
	{TokenID: 37, Timestamp: time.UnixMilli(1730728802000), Name: "exteriorAirTemperature", ValueNumber: 19, Source: compassConnection},
	{TokenID: 37, Timestamp: time.UnixMilli(1730728800000), Name: "powertrainRange", ValueNumber: 548.7863040000001, Source: compassConnection},
	{TokenID: 37, Timestamp: time.UnixMilli(1730728800000), Name: "powertrainTractionBatteryChargingAddedEnergy", ValueNumber: 42, Source: compassConnection},
	{TokenID: 37, Timestamp: time.UnixMilli(1730728800000), Name: "powertrainTractionBatteryChargingChargeLimit", ValueNumber: 80, Source: compassConnection},
	{TokenID: 37, Timestamp: time.UnixMilli(1730728800000), Name: "powertrainTractionBatteryChargingIsCharging", ValueNumber: 1, Source: compassConnection},
	{TokenID: 37, Timestamp: time.UnixMilli(1730738800000), Name: "powertrainTractionBatteryCurrentPower", ValueNumber: 7000, Source: compassConnection},
	{TokenID: 37, Timestamp: time.UnixMilli(1730728800000), Name: "powertrainTractionBatteryStateOfChargeCurrent", ValueNumber: 23, Source: compassConnection},
	{TokenID: 37, Timestamp: time.UnixMilli(1730728805000), Name: "powertrainTransmissionTravelledDistance", ValueNumber: 9065.434752000001, Source: compassConnection},
	{TokenID: 37, Timestamp: time.UnixMilli(1730738800000), Name: "speed", ValueNumber: 40.2336, Source: compassConnection},
}

func TestSignalsFromCompass(t *testing.T) {
	computedSignals, err := Decode(baseDoc)
	require.Empty(t, err, "Expected no errors.")
	assert.ElementsMatch(t, computedSignals, expSignals)
}
