package ruptela

import "github.com/tidwall/gjson"

const (
	// StatusEventDS is the data version for status events.
	StatusEventDS = "r/v0/s"
	// DevStatusDS is the data version for device status events.
	DevStatusDS = "r/v0/dev"
	// LocationEventDS is the data version for location events.
	LocationEventDS = "r/v0/loc"
	// DTCEventDS is the data version for DTC events.
	DTCEventDS = "r/v0/dtc"
)

// fuelTypeConversion Encodings taken from https://en.wikipedia.org/wiki/OBD-II_PIDs#Fuel_Type_Coding
func fuelTypeConversion(val float64) (string, error) {
	switch val {
	case 1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13, 14, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32:
		return "COMBUSTION", nil
	case 8, 15:
		return "ELECTRIC", nil
	case 16, 17, 18, 19, 20, 21, 22:
		return "HYBRID", nil
	default:
		return "", errNotFound
	}
}

func ignoreZero(val float64, err error) (float64, error) {
	if err != nil {
		return 0, err
	}
	if val == 0 {
		return 0, errNotFound
	}
	return val, err
}

func ignitionOff(originalDoc []byte) bool {
	result := gjson.GetBytes(originalDoc, "data.signals.409")
	if !result.Exists() || result.Type != gjson.String {
		return false
	}
	return result.Str == "0"
}

func unplugged(originalDoc []byte) bool {
	result := gjson.GetBytes(originalDoc, "data.signals.985")
	if !result.Exists() || result.Type != gjson.String {
		return false
	}
	return result.Str == "1"
}
