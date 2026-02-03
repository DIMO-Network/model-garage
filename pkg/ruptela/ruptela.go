package ruptela

import (
	"github.com/tidwall/gjson"
)

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

func TorqueModeConversion(val float64) (string, error) {
	switch val {
	case 0:
		return "LOW IDLE GOVERNOR/NO REQUEST", nil
	case 1:
		return "ACCELERATOR PEDAL/OPERATOR SELECTION", nil
	case 2:
		return "CRUISE CONTROL", nil
	case 3:
		return "PTO GOVERNOR", nil
	case 4:
		return "ROAD SPEED GOVERNOR", nil
	case 5:
		return "ASR CONTROL", nil
	case 6:
		return "TRANSMISSION CONTROL", nil
	case 7:
		return "ABS CONTROL", nil
	case 8:
		return "TORQUE LIMITING", nil
	case 9:
		return "HIGH SPEED GOVERNOR", nil
	case 10:
		return "BRAKING SYSTEM", nil
	case 11:
		return "REMOTE ACCELERATOR", nil
	case 12:
		return "SERVICE PROCEDURE", nil
	case 13:
		return "NOT DEFINED", nil
	case 14:
		return "OTHER", nil
	case 15:
		return "NOT AVAILABLE", nil
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

func CANStatusToBool(val float64) (float64, error) {
	switch {
	case val == 0:
		return 0, nil
	case val > 0 && val <= 3:
		return 1, nil
	case val == 7:
		return 0, errNotFound
	default:
		return 0, errNotFound
	}
}

func CANBitToBool(val float64, bit uint) (float64, error) {
	raw := int(val)
	if raw < 0 || raw > 255 {
		return 0, errNotFound
	}

	if raw == 0xFF {
		return 0, errNotFound
	}

	mask := 1 << bit
	if (raw & mask) != 0 {
		return 1.0, nil
	}
	return 0.0, nil
}
