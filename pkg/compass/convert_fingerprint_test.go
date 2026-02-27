package compass_test

import (
	"encoding/json"
	"testing"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/compass"
	"github.com/stretchr/testify/require"
)

func TestFullFromDataConversion(t *testing.T) {
	t.Parallel()
	expectedVIN := "1C4SJSBP8RS133747"
	event := cloudevent.RawEvent{}

	err := json.Unmarshal([]byte(fullInputJSON), &event)
	require.NoError(t, err, "error unmarshalling JSON")

	fp, err := compass.DecodeFingerprint(event)
	require.NoError(t, err, "error decoding fingerprint")
	require.Equal(t, expectedVIN, fp.VIN, "decoded VIN does not match expected VIN")
}

var fullInputJSON = `
{
   "subject":"did:erc721:137:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:37",
   "source":"0x55BF1c27d468314Ea119CF74979E2b59F962295c",
   "time": "2024-09-27T08:33:26Z",
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
         "speed.value":"40.2336",
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
}`
