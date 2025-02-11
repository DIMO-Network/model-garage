// Code generated by github.com/DIMO-Network/model-garage.
package compass

import "strconv"

// This file is automatically populated with conversion functions for each field of the model struct.
// any conversion functions already defined in this package will be coppied through.
// note: DO NOT mutate the orginalDoc parameter which is shared between all conversion functions.

// ToCurrentLocationAltitude0 converts data from field 'labels.geolocation.altitude.value' of type string to 'Vehicle.CurrentLocation.Altitude' of type float64.
// Vehicle.CurrentLocation.Altitude: Current altitude relative to WGS 84 reference ellipsoid, as measured at the position of GNSS receiver antenna.
// Unit: 'm'
func ToCurrentLocationAltitude0(originalDoc []byte, val string) (float64, error) {
	return strconv.ParseFloat(val, 64)
}

// ToCurrentLocationLatitude0 converts data from field 'labels.geolocation.latitude' of type string to 'Vehicle.CurrentLocation.Latitude' of type float64.
// Vehicle.CurrentLocation.Latitude: Current latitude of vehicle in WGS 84 geodetic coordinates, as measured at the position of GNSS receiver antenna.
// Unit: 'degrees' Min: '-90' Max: '90'
func ToCurrentLocationLatitude0(originalDoc []byte, val string) (float64, error) {
	return strconv.ParseFloat(val, 64)
}

// ToCurrentLocationLongitude0 converts data from field 'labels.geolocation.longitude' of type string to 'Vehicle.CurrentLocation.Longitude' of type float64.
// Vehicle.CurrentLocation.Longitude: Current longitude of vehicle in WGS 84 geodetic coordinates, as measured at the position of GNSS receiver antenna.
// Unit: 'degrees' Min: '-180' Max: '180'
func ToCurrentLocationLongitude0(originalDoc []byte, val string) (float64, error) {
	return strconv.ParseFloat(val, 64)
}

// ToLowVoltageBatteryCurrentVoltage0 converts data from field 'labels.engine.battery.voltage.value' of type string to 'Vehicle.LowVoltageBattery.CurrentVoltage' of type float64.
// Vehicle.LowVoltageBattery.CurrentVoltage: Current Voltage of the low voltage battery.
// Unit: 'V'
func ToLowVoltageBatteryCurrentVoltage0(originalDoc []byte, val string) (float64, error) {
	return strconv.ParseFloat(val, 64)
}

// ToPowertrainFuelSystemAbsoluteLevel0 converts data from field 'labels.fuel.level.value' of type string to 'Vehicle.Powertrain.FuelSystem.AbsoluteLevel' of type float64.
// Vehicle.Powertrain.FuelSystem.AbsoluteLevel: Current available fuel in the fuel tank expressed in liters.
// Unit: 'l'
func ToPowertrainFuelSystemAbsoluteLevel0(originalDoc []byte, val string) (float64, error) {
	return strconv.ParseFloat(val, 64)
}

// ToPowertrainFuelSystemRelativeLevel0 converts data from field 'labels.fuel.level.percentage' of type string to 'Vehicle.Powertrain.FuelSystem.RelativeLevel' of type float64.
// Vehicle.Powertrain.FuelSystem.RelativeLevel: Level in fuel tank as percent of capacity. 0 = empty. 100 = full.
// Unit: 'percent' Min: '0' Max: '100'
func ToPowertrainFuelSystemRelativeLevel0(originalDoc []byte, val string) (float64, error) {
	return strconv.ParseFloat(val, 64)
}

// ToPowertrainTransmissionTravelledDistance0 converts data from field 'labels.odometer.value' of type string to 'Vehicle.Powertrain.Transmission.TravelledDistance' of type float64.
// Vehicle.Powertrain.Transmission.TravelledDistance: Odometer reading, total distance travelled during the lifetime of the transmission.
// Unit: 'km'
func ToPowertrainTransmissionTravelledDistance0(originalDoc []byte, val string) (float64, error) {
	return strconv.ParseFloat(val, 64)
}

// ToSpeed0 converts data from field 'labels.speed.value' of type string to 'Vehicle.Speed' of type float64.
// Vehicle.Speed: Vehicle speed.
// Unit: 'km/h'
func ToSpeed0(originalDoc []byte, val string) (float64, error) {
	return strconv.ParseFloat(val, 64)
}
