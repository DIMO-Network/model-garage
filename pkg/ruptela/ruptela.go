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

func fuelTypeNameConversion(val float64) (string, error) {
	switch val {
	case 1:
		return "GASOLINE", nil
	case 2:
		return "METHANOL", nil
	case 3:
		return "ETHANOL", nil
	case 4:
		return "DIESEL", nil
	case 5:
		return "LPG", nil
	case 6:
		return "CNG", nil
	case 7:
		return "PROPANE", nil
	case 8:
		return "ELECTRIC", nil
	case 9:
		return "BIFUEL RUNNING GASOLINE", nil
	case 10:
		return "BIFUEL RUNNING METHANOL", nil
	case 11:
		return "BIFUEL RUNNING ETHANOL", nil
	case 12:
		return "BIFUEL RUNNING LPG", nil
	case 13:
		return "BIFUEL RUNNING CNG", nil
	case 14:
		return "BIFUEL RUNNING PROPANE", nil
	case 15:
		return "BIFUEL RUNNING ELECTRICITY", nil
	case 16:
		return "BIFUEL RUNNING ELECTRIC AND COMBUSTION ENGINE", nil
	case 17:
		return "HYBRID GASOLINE", nil
	case 18:
		return "HYBRID ETHANOL", nil
	case 19:
		return "HYBRID DIESEL", nil
	case 20:
		return "HYBRID ELECTRIC", nil
	case 21:
		return "HYBRID RUNNING ELECTRIC AND COMBUSTION ENGINE", nil
	case 22:
		return "HYBRID REGENERATIVE", nil
	case 23:
		return "BIFUEL RUNNING DIESEL", nil
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
	result := gjson.GetBytes(originalDoc, "signals.409")
	if !result.Exists() || result.Type != gjson.String {
		return false
	}
	return result.Str == "0"
}

func unplugged(originalDoc []byte) bool {
	result := gjson.GetBytes(originalDoc, "signals.985")
	if !result.Exists() || result.Type != gjson.String {
		return false
	}
	return result.Str == "1"
}

// ConvertPSIToKPa converts a pressure value from psi to kPa.
func ConvertPSIToKPa(psi float64) float64 {
	return psi * 6.89476
}
