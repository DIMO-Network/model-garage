// Code generated by github.com/DIMO-Network/model-garage.
package compass

import "strconv"

// This file is automatically populated with conversion functions for each field of the model struct.
// any conversion functions already defined in this package will be coppied through.
// note: DO NOT mutate the orginalDoc parameter which is shared between all conversion functions.

// ToPowertrainTransmissionTravelledDistance0 converts data from field 'labels.odometer.value' of type string to 'Vehicle.Powertrain.Transmission.TravelledDistance' of type float64.
// Vehicle.Powertrain.Transmission.TravelledDistance: Odometer reading, total distance travelled during the lifetime of the transmission.
// Unit: 'km'
func ToPowertrainTransmissionTravelledDistance0(originalDoc []byte, val string) (float64, error) {
	return strconv.ParseFloat(val, 64)
}

// ToSpeed0 converts data from field 'speed.value' of type string to 'Vehicle.Speed' of type float64.
// Vehicle.Speed: Vehicle speed.
// Unit: 'km/h'
func ToSpeed0(originalDoc []byte, val string) (float64, error) {
	return strconv.ParseFloat(val, 64)
}
