// Code generated by github.com/DIMO-Network/model-garage DO NOT EDIT.
package ruptela

import (
	"errors"
	"fmt"

	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/tidwall/gjson"
)

var errNotFound = errors.New("field not found")

// SignalsFromV1Data creates a slice of vss.Signal from the given v1 status JSON data.
// On error, partial results may be returned.
func SignalsFromV1Data(baseSignal vss.Signal, jsonData []byte) ([]vss.Signal, []error) {
	var retSignals []vss.Signal

	var val any
	var err error
	var errs []error

	val, err = ChassisAxleRow1WheelLeftTirePressureFromV1Data(jsonData)
	if err != nil {
		if !errors.Is(err, errNotFound) {
			errs = append(errs, fmt.Errorf("failed to get 'ChassisAxleRow1WheelLeftTirePressure': %w", err))
		}
	} else {
		sig := vss.Signal{
			Name:      "chassisAxleRow1WheelLeftTirePressure",
			TokenID:   baseSignal.TokenID,
			Timestamp: baseSignal.Timestamp,
			Source:    baseSignal.Source,
		}
		sig.SetValue(val)
		retSignals = append(retSignals, sig)
	}

	val, err = ChassisAxleRow1WheelRightTirePressureFromV1Data(jsonData)
	if err != nil {
		if !errors.Is(err, errNotFound) {
			errs = append(errs, fmt.Errorf("failed to get 'ChassisAxleRow1WheelRightTirePressure': %w", err))
		}
	} else {
		sig := vss.Signal{
			Name:      "chassisAxleRow1WheelRightTirePressure",
			TokenID:   baseSignal.TokenID,
			Timestamp: baseSignal.Timestamp,
			Source:    baseSignal.Source,
		}
		sig.SetValue(val)
		retSignals = append(retSignals, sig)
	}

	val, err = ChassisAxleRow2WheelLeftTirePressureFromV1Data(jsonData)
	if err != nil {
		if !errors.Is(err, errNotFound) {
			errs = append(errs, fmt.Errorf("failed to get 'ChassisAxleRow2WheelLeftTirePressure': %w", err))
		}
	} else {
		sig := vss.Signal{
			Name:      "chassisAxleRow2WheelLeftTirePressure",
			TokenID:   baseSignal.TokenID,
			Timestamp: baseSignal.Timestamp,
			Source:    baseSignal.Source,
		}
		sig.SetValue(val)
		retSignals = append(retSignals, sig)
	}

	val, err = ChassisAxleRow2WheelRightTirePressureFromV1Data(jsonData)
	if err != nil {
		if !errors.Is(err, errNotFound) {
			errs = append(errs, fmt.Errorf("failed to get 'ChassisAxleRow2WheelRightTirePressure': %w", err))
		}
	} else {
		sig := vss.Signal{
			Name:      "chassisAxleRow2WheelRightTirePressure",
			TokenID:   baseSignal.TokenID,
			Timestamp: baseSignal.Timestamp,
			Source:    baseSignal.Source,
		}
		sig.SetValue(val)
		retSignals = append(retSignals, sig)
	}

	val, err = CurrentLocationAltitudeFromV1Data(jsonData)
	if err != nil {
		if !errors.Is(err, errNotFound) {
			errs = append(errs, fmt.Errorf("failed to get 'CurrentLocationAltitude': %w", err))
		}
	} else {
		sig := vss.Signal{
			Name:      "currentLocationAltitude",
			TokenID:   baseSignal.TokenID,
			Timestamp: baseSignal.Timestamp,
			Source:    baseSignal.Source,
		}
		sig.SetValue(val)
		retSignals = append(retSignals, sig)
	}

	val, err = CurrentLocationLatitudeFromV1Data(jsonData)
	if err != nil {
		if !errors.Is(err, errNotFound) {
			errs = append(errs, fmt.Errorf("failed to get 'CurrentLocationLatitude': %w", err))
		}
	} else {
		sig := vss.Signal{
			Name:      "currentLocationLatitude",
			TokenID:   baseSignal.TokenID,
			Timestamp: baseSignal.Timestamp,
			Source:    baseSignal.Source,
		}
		sig.SetValue(val)
		retSignals = append(retSignals, sig)
	}

	val, err = CurrentLocationLongitudeFromV1Data(jsonData)
	if err != nil {
		if !errors.Is(err, errNotFound) {
			errs = append(errs, fmt.Errorf("failed to get 'CurrentLocationLongitude': %w", err))
		}
	} else {
		sig := vss.Signal{
			Name:      "currentLocationLongitude",
			TokenID:   baseSignal.TokenID,
			Timestamp: baseSignal.Timestamp,
			Source:    baseSignal.Source,
		}
		sig.SetValue(val)
		retSignals = append(retSignals, sig)
	}

	val, err = DIMOAftermarketHDOPFromV1Data(jsonData)
	if err != nil {
		if !errors.Is(err, errNotFound) {
			errs = append(errs, fmt.Errorf("failed to get 'DIMOAftermarketHDOP': %w", err))
		}
	} else {
		sig := vss.Signal{
			Name:      "dimoAftermarketHDOP",
			TokenID:   baseSignal.TokenID,
			Timestamp: baseSignal.Timestamp,
			Source:    baseSignal.Source,
		}
		sig.SetValue(val)
		retSignals = append(retSignals, sig)
	}

	val, err = DIMOAftermarketNSATFromV1Data(jsonData)
	if err != nil {
		if !errors.Is(err, errNotFound) {
			errs = append(errs, fmt.Errorf("failed to get 'DIMOAftermarketNSAT': %w", err))
		}
	} else {
		sig := vss.Signal{
			Name:      "dimoAftermarketNSAT",
			TokenID:   baseSignal.TokenID,
			Timestamp: baseSignal.Timestamp,
			Source:    baseSignal.Source,
		}
		sig.SetValue(val)
		retSignals = append(retSignals, sig)
	}

	val, err = ExteriorAirTemperatureFromV1Data(jsonData)
	if err != nil {
		if !errors.Is(err, errNotFound) {
			errs = append(errs, fmt.Errorf("failed to get 'ExteriorAirTemperature': %w", err))
		}
	} else {
		sig := vss.Signal{
			Name:      "exteriorAirTemperature",
			TokenID:   baseSignal.TokenID,
			Timestamp: baseSignal.Timestamp,
			Source:    baseSignal.Source,
		}
		sig.SetValue(val)
		retSignals = append(retSignals, sig)
	}

	val, err = LowVoltageBatteryCurrentVoltageFromV1Data(jsonData)
	if err != nil {
		if !errors.Is(err, errNotFound) {
			errs = append(errs, fmt.Errorf("failed to get 'LowVoltageBatteryCurrentVoltage': %w", err))
		}
	} else {
		sig := vss.Signal{
			Name:      "lowVoltageBatteryCurrentVoltage",
			TokenID:   baseSignal.TokenID,
			Timestamp: baseSignal.Timestamp,
			Source:    baseSignal.Source,
		}
		sig.SetValue(val)
		retSignals = append(retSignals, sig)
	}

	val, err = OBDDistanceWithMILFromV1Data(jsonData)
	if err != nil {
		if !errors.Is(err, errNotFound) {
			errs = append(errs, fmt.Errorf("failed to get 'OBDDistanceWithMIL': %w", err))
		}
	} else {
		sig := vss.Signal{
			Name:      "obdDistanceWithMIL",
			TokenID:   baseSignal.TokenID,
			Timestamp: baseSignal.Timestamp,
			Source:    baseSignal.Source,
		}
		sig.SetValue(val)
		retSignals = append(retSignals, sig)
	}

	val, err = OBDRunTimeFromV1Data(jsonData)
	if err != nil {
		if !errors.Is(err, errNotFound) {
			errs = append(errs, fmt.Errorf("failed to get 'OBDRunTime': %w", err))
		}
	} else {
		sig := vss.Signal{
			Name:      "obdRunTime",
			TokenID:   baseSignal.TokenID,
			Timestamp: baseSignal.Timestamp,
			Source:    baseSignal.Source,
		}
		sig.SetValue(val)
		retSignals = append(retSignals, sig)
	}

	val, err = PowertrainCombustionEngineECTFromV1Data(jsonData)
	if err != nil {
		if !errors.Is(err, errNotFound) {
			errs = append(errs, fmt.Errorf("failed to get 'PowertrainCombustionEngineECT': %w", err))
		}
	} else {
		sig := vss.Signal{
			Name:      "powertrainCombustionEngineECT",
			TokenID:   baseSignal.TokenID,
			Timestamp: baseSignal.Timestamp,
			Source:    baseSignal.Source,
		}
		sig.SetValue(val)
		retSignals = append(retSignals, sig)
	}

	val, err = PowertrainCombustionEngineEngineOilLevelFromV1Data(jsonData)
	if err != nil {
		if !errors.Is(err, errNotFound) {
			errs = append(errs, fmt.Errorf("failed to get 'PowertrainCombustionEngineEngineOilLevel': %w", err))
		}
	} else {
		sig := vss.Signal{
			Name:      "powertrainCombustionEngineEngineOilLevel",
			TokenID:   baseSignal.TokenID,
			Timestamp: baseSignal.Timestamp,
			Source:    baseSignal.Source,
		}
		sig.SetValue(val)
		retSignals = append(retSignals, sig)
	}

	val, err = PowertrainCombustionEngineEngineOilRelativeLevelFromV1Data(jsonData)
	if err != nil {
		if !errors.Is(err, errNotFound) {
			errs = append(errs, fmt.Errorf("failed to get 'PowertrainCombustionEngineEngineOilRelativeLevel': %w", err))
		}
	} else {
		sig := vss.Signal{
			Name:      "powertrainCombustionEngineEngineOilRelativeLevel",
			TokenID:   baseSignal.TokenID,
			Timestamp: baseSignal.Timestamp,
			Source:    baseSignal.Source,
		}
		sig.SetValue(val)
		retSignals = append(retSignals, sig)
	}

	val, err = PowertrainCombustionEngineSpeedFromV1Data(jsonData)
	if err != nil {
		if !errors.Is(err, errNotFound) {
			errs = append(errs, fmt.Errorf("failed to get 'PowertrainCombustionEngineSpeed': %w", err))
		}
	} else {
		sig := vss.Signal{
			Name:      "powertrainCombustionEngineSpeed",
			TokenID:   baseSignal.TokenID,
			Timestamp: baseSignal.Timestamp,
			Source:    baseSignal.Source,
		}
		sig.SetValue(val)
		retSignals = append(retSignals, sig)
	}

	val, err = PowertrainCombustionEngineTPSFromV1Data(jsonData)
	if err != nil {
		if !errors.Is(err, errNotFound) {
			errs = append(errs, fmt.Errorf("failed to get 'PowertrainCombustionEngineTPS': %w", err))
		}
	} else {
		sig := vss.Signal{
			Name:      "powertrainCombustionEngineTPS",
			TokenID:   baseSignal.TokenID,
			Timestamp: baseSignal.Timestamp,
			Source:    baseSignal.Source,
		}
		sig.SetValue(val)
		retSignals = append(retSignals, sig)
	}

	val, err = PowertrainFuelSystemAbsoluteLevelFromV1Data(jsonData)
	if err != nil {
		if !errors.Is(err, errNotFound) {
			errs = append(errs, fmt.Errorf("failed to get 'PowertrainFuelSystemAbsoluteLevel': %w", err))
		}
	} else {
		sig := vss.Signal{
			Name:      "powertrainFuelSystemAbsoluteLevel",
			TokenID:   baseSignal.TokenID,
			Timestamp: baseSignal.Timestamp,
			Source:    baseSignal.Source,
		}
		sig.SetValue(val)
		retSignals = append(retSignals, sig)
	}

	val, err = PowertrainFuelSystemRelativeLevelFromV1Data(jsonData)
	if err != nil {
		if !errors.Is(err, errNotFound) {
			errs = append(errs, fmt.Errorf("failed to get 'PowertrainFuelSystemRelativeLevel': %w", err))
		}
	} else {
		sig := vss.Signal{
			Name:      "powertrainFuelSystemRelativeLevel",
			TokenID:   baseSignal.TokenID,
			Timestamp: baseSignal.Timestamp,
			Source:    baseSignal.Source,
		}
		sig.SetValue(val)
		retSignals = append(retSignals, sig)
	}

	val, err = PowertrainRangeFromV1Data(jsonData)
	if err != nil {
		if !errors.Is(err, errNotFound) {
			errs = append(errs, fmt.Errorf("failed to get 'PowertrainRange': %w", err))
		}
	} else {
		sig := vss.Signal{
			Name:      "powertrainRange",
			TokenID:   baseSignal.TokenID,
			Timestamp: baseSignal.Timestamp,
			Source:    baseSignal.Source,
		}
		sig.SetValue(val)
		retSignals = append(retSignals, sig)
	}

	val, err = PowertrainTractionBatteryStateOfChargeCurrentFromV1Data(jsonData)
	if err != nil {
		if !errors.Is(err, errNotFound) {
			errs = append(errs, fmt.Errorf("failed to get 'PowertrainTractionBatteryStateOfChargeCurrent': %w", err))
		}
	} else {
		sig := vss.Signal{
			Name:      "powertrainTractionBatteryStateOfChargeCurrent",
			TokenID:   baseSignal.TokenID,
			Timestamp: baseSignal.Timestamp,
			Source:    baseSignal.Source,
		}
		sig.SetValue(val)
		retSignals = append(retSignals, sig)
	}

	val, err = PowertrainTransmissionTravelledDistanceFromV1Data(jsonData)
	if err != nil {
		if !errors.Is(err, errNotFound) {
			errs = append(errs, fmt.Errorf("failed to get 'PowertrainTransmissionTravelledDistance': %w", err))
		}
	} else {
		sig := vss.Signal{
			Name:      "powertrainTransmissionTravelledDistance",
			TokenID:   baseSignal.TokenID,
			Timestamp: baseSignal.Timestamp,
			Source:    baseSignal.Source,
		}
		sig.SetValue(val)
		retSignals = append(retSignals, sig)
	}

	val, err = PowertrainTypeFromV1Data(jsonData)
	if err != nil {
		if !errors.Is(err, errNotFound) {
			errs = append(errs, fmt.Errorf("failed to get 'PowertrainType': %w", err))
		}
	} else {
		sig := vss.Signal{
			Name:      "powertrainType",
			TokenID:   baseSignal.TokenID,
			Timestamp: baseSignal.Timestamp,
			Source:    baseSignal.Source,
		}
		sig.SetValue(val)
		retSignals = append(retSignals, sig)
	}

	val, err = SpeedFromV1Data(jsonData)
	if err != nil {
		if !errors.Is(err, errNotFound) {
			errs = append(errs, fmt.Errorf("failed to get 'Speed': %w", err))
		}
	} else {
		sig := vss.Signal{
			Name:      "speed",
			TokenID:   baseSignal.TokenID,
			Timestamp: baseSignal.Timestamp,
			Source:    baseSignal.Source,
		}
		sig.SetValue(val)
		retSignals = append(retSignals, sig)
	}
	return retSignals, errs
}

// ChassisAxleRow1WheelLeftTirePressureFromV1Data converts the given JSON data to a float64.
func ChassisAxleRow1WheelLeftTirePressureFromV1Data(jsonData []byte) (ret float64, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.signals.960")
	if result.Exists() && result.Value() != nil {
		val, ok := result.Value().(string)
		if ok {
			retVal, err := ToChassisAxleRow1WheelLeftTirePressure0(jsonData, val)
			if err == nil {
				return retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.signals.960': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.signals.960' is not of type 'string' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
		}
	}

	if errs == nil {
		return ret, fmt.Errorf("%w 'ChassisAxleRow1WheelLeftTirePressure'", errNotFound)
	}

	return ret, errs
}

// ChassisAxleRow1WheelRightTirePressureFromV1Data converts the given JSON data to a float64.
func ChassisAxleRow1WheelRightTirePressureFromV1Data(jsonData []byte) (ret float64, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.signals.961")
	if result.Exists() && result.Value() != nil {
		val, ok := result.Value().(string)
		if ok {
			retVal, err := ToChassisAxleRow1WheelRightTirePressure0(jsonData, val)
			if err == nil {
				return retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.signals.961': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.signals.961' is not of type 'string' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
		}
	}

	if errs == nil {
		return ret, fmt.Errorf("%w 'ChassisAxleRow1WheelRightTirePressure'", errNotFound)
	}

	return ret, errs
}

// ChassisAxleRow2WheelLeftTirePressureFromV1Data converts the given JSON data to a float64.
func ChassisAxleRow2WheelLeftTirePressureFromV1Data(jsonData []byte) (ret float64, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.signals.962")
	if result.Exists() && result.Value() != nil {
		val, ok := result.Value().(string)
		if ok {
			retVal, err := ToChassisAxleRow2WheelLeftTirePressure0(jsonData, val)
			if err == nil {
				return retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.signals.962': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.signals.962' is not of type 'string' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
		}
	}

	if errs == nil {
		return ret, fmt.Errorf("%w 'ChassisAxleRow2WheelLeftTirePressure'", errNotFound)
	}

	return ret, errs
}

// ChassisAxleRow2WheelRightTirePressureFromV1Data converts the given JSON data to a float64.
func ChassisAxleRow2WheelRightTirePressureFromV1Data(jsonData []byte) (ret float64, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.signals.963")
	if result.Exists() && result.Value() != nil {
		val, ok := result.Value().(string)
		if ok {
			retVal, err := ToChassisAxleRow2WheelRightTirePressure0(jsonData, val)
			if err == nil {
				return retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.signals.963': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.signals.963' is not of type 'string' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
		}
	}

	if errs == nil {
		return ret, fmt.Errorf("%w 'ChassisAxleRow2WheelRightTirePressure'", errNotFound)
	}

	return ret, errs
}

// CurrentLocationAltitudeFromV1Data converts the given JSON data to a float64.
func CurrentLocationAltitudeFromV1Data(jsonData []byte) (ret float64, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.pos.alt")
	if result.Exists() && result.Value() != nil {
		val, ok := result.Value().(float64)
		if ok {
			retVal, err := ToCurrentLocationAltitude0(jsonData, val)
			if err == nil {
				return retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.pos.alt': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.pos.alt' is not of type 'float64' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
		}
	}

	if errs == nil {
		return ret, fmt.Errorf("%w 'CurrentLocationAltitude'", errNotFound)
	}

	return ret, errs
}

// CurrentLocationLatitudeFromV1Data converts the given JSON data to a float64.
func CurrentLocationLatitudeFromV1Data(jsonData []byte) (ret float64, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.pos.lat")
	if result.Exists() && result.Value() != nil {
		val, ok := result.Value().(float64)
		if ok {
			retVal, err := ToCurrentLocationLatitude0(jsonData, val)
			if err == nil {
				return retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.pos.lat': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.pos.lat' is not of type 'float64' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
		}
	}

	if errs == nil {
		return ret, fmt.Errorf("%w 'CurrentLocationLatitude'", errNotFound)
	}

	return ret, errs
}

// CurrentLocationLongitudeFromV1Data converts the given JSON data to a float64.
func CurrentLocationLongitudeFromV1Data(jsonData []byte) (ret float64, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.pos.lon")
	if result.Exists() && result.Value() != nil {
		val, ok := result.Value().(float64)
		if ok {
			retVal, err := ToCurrentLocationLongitude0(jsonData, val)
			if err == nil {
				return retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.pos.lon': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.pos.lon' is not of type 'float64' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
		}
	}

	if errs == nil {
		return ret, fmt.Errorf("%w 'CurrentLocationLongitude'", errNotFound)
	}

	return ret, errs
}

// DIMOAftermarketHDOPFromV1Data converts the given JSON data to a float64.
func DIMOAftermarketHDOPFromV1Data(jsonData []byte) (ret float64, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.pos.hdop")
	if result.Exists() && result.Value() != nil {
		val, ok := result.Value().(float64)
		if ok {
			retVal, err := ToDIMOAftermarketHDOP0(jsonData, val)
			if err == nil {
				return retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.pos.hdop': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.pos.hdop' is not of type 'float64' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
		}
	}

	if errs == nil {
		return ret, fmt.Errorf("%w 'DIMOAftermarketHDOP'", errNotFound)
	}

	return ret, errs
}

// DIMOAftermarketNSATFromV1Data converts the given JSON data to a float64.
func DIMOAftermarketNSATFromV1Data(jsonData []byte) (ret float64, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.pos.sat")
	if result.Exists() && result.Value() != nil {
		val, ok := result.Value().(float64)
		if ok {
			retVal, err := ToDIMOAftermarketNSAT0(jsonData, val)
			if err == nil {
				return retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.pos.sat': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.pos.sat' is not of type 'float64' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
		}
	}

	if errs == nil {
		return ret, fmt.Errorf("%w 'DIMOAftermarketNSAT'", errNotFound)
	}

	return ret, errs
}

// ExteriorAirTemperatureFromV1Data converts the given JSON data to a float64.
func ExteriorAirTemperatureFromV1Data(jsonData []byte) (ret float64, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.signals.97")
	if result.Exists() && result.Value() != nil {
		val, ok := result.Value().(string)
		if ok {
			retVal, err := ToExteriorAirTemperature0(jsonData, val)
			if err == nil {
				return retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.signals.97': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.signals.97' is not of type 'string' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
		}
	}

	if errs == nil {
		return ret, fmt.Errorf("%w 'ExteriorAirTemperature'", errNotFound)
	}

	return ret, errs
}

// LowVoltageBatteryCurrentVoltageFromV1Data converts the given JSON data to a float64.
func LowVoltageBatteryCurrentVoltageFromV1Data(jsonData []byte) (ret float64, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.signals.29")
	if result.Exists() && result.Value() != nil {
		val, ok := result.Value().(string)
		if ok {
			retVal, err := ToLowVoltageBatteryCurrentVoltage0(jsonData, val)
			if err == nil {
				return retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.signals.29': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.signals.29' is not of type 'string' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
		}
	}

	if errs == nil {
		return ret, fmt.Errorf("%w 'LowVoltageBatteryCurrentVoltage'", errNotFound)
	}

	return ret, errs
}

// OBDDistanceWithMILFromV1Data converts the given JSON data to a float64.
func OBDDistanceWithMILFromV1Data(jsonData []byte) (ret float64, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.signals.102")
	if result.Exists() && result.Value() != nil {
		val, ok := result.Value().(string)
		if ok {
			retVal, err := ToOBDDistanceWithMIL0(jsonData, val)
			if err == nil {
				return retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.signals.102': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.signals.102' is not of type 'string' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
		}
	}

	if errs == nil {
		return ret, fmt.Errorf("%w 'OBDDistanceWithMIL'", errNotFound)
	}

	return ret, errs
}

// OBDRunTimeFromV1Data converts the given JSON data to a float64.
func OBDRunTimeFromV1Data(jsonData []byte) (ret float64, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.signals.107")
	if result.Exists() && result.Value() != nil {
		val, ok := result.Value().(string)
		if ok {
			retVal, err := ToOBDRunTime0(jsonData, val)
			if err == nil {
				return retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.signals.107': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.signals.107' is not of type 'string' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
		}
	}

	if errs == nil {
		return ret, fmt.Errorf("%w 'OBDRunTime'", errNotFound)
	}

	return ret, errs
}

// PowertrainCombustionEngineECTFromV1Data converts the given JSON data to a float64.
func PowertrainCombustionEngineECTFromV1Data(jsonData []byte) (ret float64, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.signals.96")
	if result.Exists() && result.Value() != nil {
		val, ok := result.Value().(string)
		if ok {
			retVal, err := ToPowertrainCombustionEngineECT0(jsonData, val)
			if err == nil {
				return retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.signals.96': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.signals.96' is not of type 'string' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
		}
	}

	if errs == nil {
		return ret, fmt.Errorf("%w 'PowertrainCombustionEngineECT'", errNotFound)
	}

	return ret, errs
}

// PowertrainCombustionEngineEngineOilLevelFromV1Data converts the given JSON data to a string.
func PowertrainCombustionEngineEngineOilLevelFromV1Data(jsonData []byte) (ret string, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.signals.964")
	if result.Exists() && result.Value() != nil {
		val, ok := result.Value().(string)
		if ok {
			retVal, err := ToPowertrainCombustionEngineEngineOilLevel0(jsonData, val)
			if err == nil {
				return retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.signals.964': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.signals.964' is not of type 'string' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
		}
	}

	if errs == nil {
		return ret, fmt.Errorf("%w 'PowertrainCombustionEngineEngineOilLevel'", errNotFound)
	}

	return ret, errs
}

// PowertrainCombustionEngineEngineOilRelativeLevelFromV1Data converts the given JSON data to a float64.
func PowertrainCombustionEngineEngineOilRelativeLevelFromV1Data(jsonData []byte) (ret float64, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.signals.964")
	if result.Exists() && result.Value() != nil {
		val, ok := result.Value().(string)
		if ok {
			retVal, err := ToPowertrainCombustionEngineEngineOilRelativeLevel0(jsonData, val)
			if err == nil {
				return retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.signals.964': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.signals.964' is not of type 'string' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
		}
	}

	if errs == nil {
		return ret, fmt.Errorf("%w 'PowertrainCombustionEngineEngineOilRelativeLevel'", errNotFound)
	}

	return ret, errs
}

// PowertrainCombustionEngineSpeedFromV1Data converts the given JSON data to a float64.
func PowertrainCombustionEngineSpeedFromV1Data(jsonData []byte) (ret float64, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.signals.94")
	if result.Exists() && result.Value() != nil {
		val, ok := result.Value().(string)
		if ok {
			retVal, err := ToPowertrainCombustionEngineSpeed0(jsonData, val)
			if err == nil {
				return retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.signals.94': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.signals.94' is not of type 'string' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
		}
	}

	if errs == nil {
		return ret, fmt.Errorf("%w 'PowertrainCombustionEngineSpeed'", errNotFound)
	}

	return ret, errs
}

// PowertrainCombustionEngineTPSFromV1Data converts the given JSON data to a float64.
func PowertrainCombustionEngineTPSFromV1Data(jsonData []byte) (ret float64, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.signals.103")
	if result.Exists() && result.Value() != nil {
		val, ok := result.Value().(string)
		if ok {
			retVal, err := ToPowertrainCombustionEngineTPS0(jsonData, val)
			if err == nil {
				return retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.signals.103': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.signals.103' is not of type 'string' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
		}
	}

	if errs == nil {
		return ret, fmt.Errorf("%w 'PowertrainCombustionEngineTPS'", errNotFound)
	}

	return ret, errs
}

// PowertrainFuelSystemAbsoluteLevelFromV1Data converts the given JSON data to a float64.
func PowertrainFuelSystemAbsoluteLevelFromV1Data(jsonData []byte) (ret float64, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.signals.642")
	if result.Exists() && result.Value() != nil {
		val, ok := result.Value().(string)
		if ok {
			retVal, err := ToPowertrainFuelSystemAbsoluteLevel0(jsonData, val)
			if err == nil {
				return retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.signals.642': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.signals.642' is not of type 'string' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
		}
	}
	result = gjson.GetBytes(jsonData, "data.signals.205")
	if result.Exists() && result.Value() != nil {
		val, ok := result.Value().(string)
		if ok {
			retVal, err := ToPowertrainFuelSystemAbsoluteLevel1(jsonData, val)
			if err == nil {
				return retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.signals.205': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.signals.205' is not of type 'string' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
		}
	}

	if errs == nil {
		return ret, fmt.Errorf("%w 'PowertrainFuelSystemAbsoluteLevel'", errNotFound)
	}

	return ret, errs
}

// PowertrainFuelSystemRelativeLevelFromV1Data converts the given JSON data to a float64.
func PowertrainFuelSystemRelativeLevelFromV1Data(jsonData []byte) (ret float64, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.signals.98")
	if result.Exists() && result.Value() != nil {
		val, ok := result.Value().(string)
		if ok {
			retVal, err := ToPowertrainFuelSystemRelativeLevel0(jsonData, val)
			if err == nil {
				return retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.signals.98': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.signals.98' is not of type 'string' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
		}
	}
	result = gjson.GetBytes(jsonData, "data.signals.207")
	if result.Exists() && result.Value() != nil {
		val, ok := result.Value().(string)
		if ok {
			retVal, err := ToPowertrainFuelSystemRelativeLevel1(jsonData, val)
			if err == nil {
				return retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.signals.207': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.signals.207' is not of type 'string' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
		}
	}

	if errs == nil {
		return ret, fmt.Errorf("%w 'PowertrainFuelSystemRelativeLevel'", errNotFound)
	}

	return ret, errs
}

// PowertrainRangeFromV1Data converts the given JSON data to a float64.
func PowertrainRangeFromV1Data(jsonData []byte) (ret float64, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.signals.723")
	if result.Exists() && result.Value() != nil {
		val, ok := result.Value().(string)
		if ok {
			retVal, err := ToPowertrainRange0(jsonData, val)
			if err == nil {
				return retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.signals.723': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.signals.723' is not of type 'string' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
		}
	}

	if errs == nil {
		return ret, fmt.Errorf("%w 'PowertrainRange'", errNotFound)
	}

	return ret, errs
}

// PowertrainTractionBatteryStateOfChargeCurrentFromV1Data converts the given JSON data to a float64.
func PowertrainTractionBatteryStateOfChargeCurrentFromV1Data(jsonData []byte) (ret float64, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.signals.722")
	if result.Exists() && result.Value() != nil {
		val, ok := result.Value().(string)
		if ok {
			retVal, err := ToPowertrainTractionBatteryStateOfChargeCurrent0(jsonData, val)
			if err == nil {
				return retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.signals.722': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.signals.722' is not of type 'string' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
		}
	}

	if errs == nil {
		return ret, fmt.Errorf("%w 'PowertrainTractionBatteryStateOfChargeCurrent'", errNotFound)
	}

	return ret, errs
}

// PowertrainTransmissionTravelledDistanceFromV1Data converts the given JSON data to a float64.
func PowertrainTransmissionTravelledDistanceFromV1Data(jsonData []byte) (ret float64, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.signals.645")
	if result.Exists() && result.Value() != nil {
		val, ok := result.Value().(string)
		if ok {
			retVal, err := ToPowertrainTransmissionTravelledDistance0(jsonData, val)
			if err == nil {
				return retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.signals.645': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.signals.645' is not of type 'string' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
		}
	}
	result = gjson.GetBytes(jsonData, "data.signals.114")
	if result.Exists() && result.Value() != nil {
		val, ok := result.Value().(string)
		if ok {
			retVal, err := ToPowertrainTransmissionTravelledDistance1(jsonData, val)
			if err == nil {
				return retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.signals.114': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.signals.114' is not of type 'string' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
		}
	}

	if errs == nil {
		return ret, fmt.Errorf("%w 'PowertrainTransmissionTravelledDistance'", errNotFound)
	}

	return ret, errs
}

// PowertrainTypeFromV1Data converts the given JSON data to a string.
func PowertrainTypeFromV1Data(jsonData []byte) (ret string, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.signals.99")
	if result.Exists() && result.Value() != nil {
		val, ok := result.Value().(string)
		if ok {
			retVal, err := ToPowertrainType0(jsonData, val)
			if err == nil {
				return retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.signals.99': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.signals.99' is not of type 'string' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
		}
	}
	result = gjson.GetBytes(jsonData, "data.signals.483")
	if result.Exists() && result.Value() != nil {
		val, ok := result.Value().(string)
		if ok {
			retVal, err := ToPowertrainType1(jsonData, val)
			if err == nil {
				return retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.signals.483': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.signals.483' is not of type 'string' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
		}
	}

	if errs == nil {
		return ret, fmt.Errorf("%w 'PowertrainType'", errNotFound)
	}

	return ret, errs
}

// SpeedFromV1Data converts the given JSON data to a float64.
func SpeedFromV1Data(jsonData []byte) (ret float64, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.signals.95")
	if result.Exists() && result.Value() != nil {
		val, ok := result.Value().(string)
		if ok {
			retVal, err := ToSpeed0(jsonData, val)
			if err == nil {
				return retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.signals.95': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.signals.95' is not of type 'string' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
		}
	}
	result = gjson.GetBytes(jsonData, "data.pos.spd")
	if result.Exists() && result.Value() != nil {
		val, ok := result.Value().(float64)
		if ok {
			retVal, err := ToSpeed1(jsonData, val)
			if err == nil {
				return retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.pos.spd': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.pos.spd' is not of type 'float64' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
		}
	}

	if errs == nil {
		return ret, fmt.Errorf("%w 'Speed'", errNotFound)
	}

	return ret, errs
}
