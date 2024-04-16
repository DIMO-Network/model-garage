// Code generated by "model-garage" DO NOT EDIT.
package vss

import (
	"errors"
	"fmt"
	"time"

	"github.com/tidwall/gjson"
)

// errInvalidType is returned when a field is not of the expected type or not found.
var errInvalidType = errors.New("invalid type")

// IsInvalidType returns true if the error is of type errInvalidType.
func IsInvalidType(err error) bool {
	return errors.Is(err, errInvalidType)
}

// FromData creates a new Dimo from JSON data. Using defined conversion functions.
// if a filed is not found it will not be set
func FromData(jsonData []byte) (*Dimo, error) {
	var dimo Dimo
	var errs error
	var err error

	dimo.VehicleChassisAxleRow1WheelLeftTirePressure, err = VehicleChassisAxleRow1WheelLeftTirePressureFromData(jsonData)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to get 'VehicleChassisAxleRow1WheelLeftTirePressure': %w", err))
	}

	dimo.VehicleChassisAxleRow1WheelRightTirePressure, err = VehicleChassisAxleRow1WheelRightTirePressureFromData(jsonData)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to get 'VehicleChassisAxleRow1WheelRightTirePressure': %w", err))
	}

	dimo.VehicleChassisAxleRow2WheelLeftTirePressure, err = VehicleChassisAxleRow2WheelLeftTirePressureFromData(jsonData)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to get 'VehicleChassisAxleRow2WheelLeftTirePressure': %w", err))
	}

	dimo.VehicleChassisAxleRow2WheelRightTirePressure, err = VehicleChassisAxleRow2WheelRightTirePressureFromData(jsonData)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to get 'VehicleChassisAxleRow2WheelRightTirePressure': %w", err))
	}

	dimo.VehicleCurrentLocationAltitude, err = VehicleCurrentLocationAltitudeFromData(jsonData)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to get 'VehicleCurrentLocationAltitude': %w", err))
	}

	dimo.VehicleCurrentLocationLatitude, err = VehicleCurrentLocationLatitudeFromData(jsonData)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to get 'VehicleCurrentLocationLatitude': %w", err))
	}

	dimo.VehicleCurrentLocationLongitude, err = VehicleCurrentLocationLongitudeFromData(jsonData)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to get 'VehicleCurrentLocationLongitude': %w", err))
	}

	dimo.VehicleCurrentLocationTimestamp, err = VehicleCurrentLocationTimestampFromData(jsonData)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to get 'VehicleCurrentLocationTimestamp': %w", err))
	}

	dimo.VehicleExteriorAirTemperature, err = VehicleExteriorAirTemperatureFromData(jsonData)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to get 'VehicleExteriorAirTemperature': %w", err))
	}

	dimo.VehicleLowVoltageBatteryCurrentVoltage, err = VehicleLowVoltageBatteryCurrentVoltageFromData(jsonData)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to get 'VehicleLowVoltageBatteryCurrentVoltage': %w", err))
	}

	dimo.VehicleOBDBarometricPressure, err = VehicleOBDBarometricPressureFromData(jsonData)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to get 'VehicleOBDBarometricPressure': %w", err))
	}

	dimo.VehicleOBDEngineLoad, err = VehicleOBDEngineLoadFromData(jsonData)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to get 'VehicleOBDEngineLoad': %w", err))
	}

	dimo.VehicleOBDIntakeTemp, err = VehicleOBDIntakeTempFromData(jsonData)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to get 'VehicleOBDIntakeTemp': %w", err))
	}

	dimo.VehicleOBDRunTime, err = VehicleOBDRunTimeFromData(jsonData)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to get 'VehicleOBDRunTime': %w", err))
	}

	dimo.VehiclePowertrainCombustionEngineECT, err = VehiclePowertrainCombustionEngineECTFromData(jsonData)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to get 'VehiclePowertrainCombustionEngineECT': %w", err))
	}

	dimo.VehiclePowertrainCombustionEngineEngineOilLevel, err = VehiclePowertrainCombustionEngineEngineOilLevelFromData(jsonData)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to get 'VehiclePowertrainCombustionEngineEngineOilLevel': %w", err))
	}

	dimo.VehiclePowertrainCombustionEngineSpeed, err = VehiclePowertrainCombustionEngineSpeedFromData(jsonData)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to get 'VehiclePowertrainCombustionEngineSpeed': %w", err))
	}

	dimo.VehiclePowertrainCombustionEngineTPS, err = VehiclePowertrainCombustionEngineTPSFromData(jsonData)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to get 'VehiclePowertrainCombustionEngineTPS': %w", err))
	}

	dimo.VehiclePowertrainFuelSystemAbsoluteLevel, err = VehiclePowertrainFuelSystemAbsoluteLevelFromData(jsonData)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to get 'VehiclePowertrainFuelSystemAbsoluteLevel': %w", err))
	}

	dimo.VehiclePowertrainFuelSystemSupportedFuelTypes, err = VehiclePowertrainFuelSystemSupportedFuelTypesFromData(jsonData)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to get 'VehiclePowertrainFuelSystemSupportedFuelTypes': %w", err))
	}

	dimo.VehiclePowertrainRange, err = VehiclePowertrainRangeFromData(jsonData)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to get 'VehiclePowertrainRange': %w", err))
	}

	dimo.VehiclePowertrainTractionBatteryChargingChargeLimit, err = VehiclePowertrainTractionBatteryChargingChargeLimitFromData(jsonData)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to get 'VehiclePowertrainTractionBatteryChargingChargeLimit': %w", err))
	}

	dimo.VehiclePowertrainTractionBatteryChargingIsCharging, err = VehiclePowertrainTractionBatteryChargingIsChargingFromData(jsonData)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to get 'VehiclePowertrainTractionBatteryChargingIsCharging': %w", err))
	}

	dimo.VehiclePowertrainTractionBatteryGrossCapacity, err = VehiclePowertrainTractionBatteryGrossCapacityFromData(jsonData)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to get 'VehiclePowertrainTractionBatteryGrossCapacity': %w", err))
	}

	dimo.VehiclePowertrainTractionBatteryStateOfChargeCurrent, err = VehiclePowertrainTractionBatteryStateOfChargeCurrentFromData(jsonData)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to get 'VehiclePowertrainTractionBatteryStateOfChargeCurrent': %w", err))
	}

	dimo.VehiclePowertrainTransmissionTravelledDistance, err = VehiclePowertrainTransmissionTravelledDistanceFromData(jsonData)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to get 'VehiclePowertrainTransmissionTravelledDistance': %w", err))
	}

	dimo.VehiclePowertrainType, err = VehiclePowertrainTypeFromData(jsonData)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to get 'VehiclePowertrainType': %w", err))
	}

	dimo.VehicleSpeed, err = VehicleSpeedFromData(jsonData)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to get 'VehicleSpeed': %w", err))
	}

	dimo.VehicleVehicleIdentificationBrand, err = VehicleVehicleIdentificationBrandFromData(jsonData)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to get 'VehicleVehicleIdentificationBrand': %w", err))
	}

	dimo.VehicleVehicleIdentificationModel, err = VehicleVehicleIdentificationModelFromData(jsonData)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to get 'VehicleVehicleIdentificationModel': %w", err))
	}

	dimo.VehicleVehicleIdentificationYear, err = VehicleVehicleIdentificationYearFromData(jsonData)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to get 'VehicleVehicleIdentificationYear': %w", err))
	}
	return &dimo, errs
}

// VehicleChassisAxleRow1WheelLeftTirePressureFromData converts the given JSON data to a *uint16.
func VehicleChassisAxleRow1WheelLeftTirePressureFromData(jsonData []byte) (ret *uint16, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.tires.frontLeft")
	if result.Exists() {
		val, ok := result.Value().(float64)
		if ok {
			retVal, err := ToVehicleChassisAxleRow1WheelLeftTirePressure0(val)
			if err == nil {
				return &retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.tires.frontLeft': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.tires.frontLeft' is not of type 'float64' got '%v' of type '%T'", errInvalidType, result.Value(), result.Value()))
		}
	}

	return ret, errs
}

// VehicleChassisAxleRow1WheelRightTirePressureFromData converts the given JSON data to a *uint16.
func VehicleChassisAxleRow1WheelRightTirePressureFromData(jsonData []byte) (ret *uint16, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.tires.frontRight")
	if result.Exists() {
		val, ok := result.Value().(float64)
		if ok {
			retVal, err := ToVehicleChassisAxleRow1WheelRightTirePressure0(val)
			if err == nil {
				return &retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.tires.frontRight': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.tires.frontRight' is not of type 'float64' got '%v' of type '%T'", errInvalidType, result.Value(), result.Value()))
		}
	}

	return ret, errs
}

// VehicleChassisAxleRow2WheelLeftTirePressureFromData converts the given JSON data to a *uint16.
func VehicleChassisAxleRow2WheelLeftTirePressureFromData(jsonData []byte) (ret *uint16, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.tires.backLeft")
	if result.Exists() {
		val, ok := result.Value().(float64)
		if ok {
			retVal, err := ToVehicleChassisAxleRow2WheelLeftTirePressure0(val)
			if err == nil {
				return &retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.tires.backLeft': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.tires.backLeft' is not of type 'float64' got '%v' of type '%T'", errInvalidType, result.Value(), result.Value()))
		}
	}

	return ret, errs
}

// VehicleChassisAxleRow2WheelRightTirePressureFromData converts the given JSON data to a *uint16.
func VehicleChassisAxleRow2WheelRightTirePressureFromData(jsonData []byte) (ret *uint16, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.tires.backRight")
	if result.Exists() {
		val, ok := result.Value().(float64)
		if ok {
			retVal, err := ToVehicleChassisAxleRow2WheelRightTirePressure0(val)
			if err == nil {
				return &retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.tires.backRight': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.tires.backRight' is not of type 'float64' got '%v' of type '%T'", errInvalidType, result.Value(), result.Value()))
		}
	}

	return ret, errs
}

// VehicleCurrentLocationAltitudeFromData converts the given JSON data to a *float64.
func VehicleCurrentLocationAltitudeFromData(jsonData []byte) (ret *float64, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.altitude")
	if result.Exists() {
		val, ok := result.Value().(float64)
		if ok {
			retVal, err := ToVehicleCurrentLocationAltitude0(val)
			if err == nil {
				return &retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.altitude': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.altitude' is not of type 'float64' got '%v' of type '%T'", errInvalidType, result.Value(), result.Value()))
		}
	}

	return ret, errs
}

// VehicleCurrentLocationLatitudeFromData converts the given JSON data to a *float64.
func VehicleCurrentLocationLatitudeFromData(jsonData []byte) (ret *float64, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.latitude")
	if result.Exists() {
		val, ok := result.Value().(float64)
		if ok {
			retVal, err := ToVehicleCurrentLocationLatitude0(val)
			if err == nil {
				return &retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.latitude': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.latitude' is not of type 'float64' got '%v' of type '%T'", errInvalidType, result.Value(), result.Value()))
		}
	}

	return ret, errs
}

// VehicleCurrentLocationLongitudeFromData converts the given JSON data to a *float64.
func VehicleCurrentLocationLongitudeFromData(jsonData []byte) (ret *float64, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.longitude")
	if result.Exists() {
		val, ok := result.Value().(float64)
		if ok {
			retVal, err := ToVehicleCurrentLocationLongitude0(val)
			if err == nil {
				return &retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.longitude': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.longitude' is not of type 'float64' got '%v' of type '%T'", errInvalidType, result.Value(), result.Value()))
		}
	}

	return ret, errs
}

// VehicleCurrentLocationTimestampFromData converts the given JSON data to a *time.Time.
func VehicleCurrentLocationTimestampFromData(jsonData []byte) (ret *time.Time, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.timestamp")
	if result.Exists() {
		val, ok := result.Value().(string)
		if ok {
			retVal, err := ToVehicleCurrentLocationTimestamp0(val)
			if err == nil {
				return &retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.timestamp': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.timestamp' is not of type 'string' got '%v' of type '%T'", errInvalidType, result.Value(), result.Value()))
		}
	}
	result = gjson.GetBytes(jsonData, "data.timestamp")
	if result.Exists() {
		val, ok := result.Value().(float64)
		if ok {
			retVal, err := ToVehicleCurrentLocationTimestamp1(val)
			if err == nil {
				return &retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.timestamp': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.timestamp' is not of type 'float64' got '%v' of type '%T'", errInvalidType, result.Value(), result.Value()))
		}
	}

	return ret, errs
}

// VehicleExteriorAirTemperatureFromData converts the given JSON data to a *float32.
func VehicleExteriorAirTemperatureFromData(jsonData []byte) (ret *float32, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.ambientTemp")
	if result.Exists() {
		val, ok := result.Value().(float64)
		if ok {
			retVal, err := ToVehicleExteriorAirTemperature0(val)
			if err == nil {
				return &retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.ambientTemp': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.ambientTemp' is not of type 'float64' got '%v' of type '%T'", errInvalidType, result.Value(), result.Value()))
		}
	}

	return ret, errs
}

// VehicleLowVoltageBatteryCurrentVoltageFromData converts the given JSON data to a *float32.
func VehicleLowVoltageBatteryCurrentVoltageFromData(jsonData []byte) (ret *float32, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.batteryVoltage")
	if result.Exists() {
		val, ok := result.Value().(float64)
		if ok {
			retVal, err := ToVehicleLowVoltageBatteryCurrentVoltage0(val)
			if err == nil {
				return &retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.batteryVoltage': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.batteryVoltage' is not of type 'float64' got '%v' of type '%T'", errInvalidType, result.Value(), result.Value()))
		}
	}

	return ret, errs
}

// VehicleOBDBarometricPressureFromData converts the given JSON data to a *float32.
func VehicleOBDBarometricPressureFromData(jsonData []byte) (ret *float32, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.barometricPressure")
	if result.Exists() {
		val, ok := result.Value().(float64)
		if ok {
			retVal, err := ToVehicleOBDBarometricPressure0(val)
			if err == nil {
				return &retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.barometricPressure': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.barometricPressure' is not of type 'float64' got '%v' of type '%T'", errInvalidType, result.Value(), result.Value()))
		}
	}

	return ret, errs
}

// VehicleOBDEngineLoadFromData converts the given JSON data to a *float32.
func VehicleOBDEngineLoadFromData(jsonData []byte) (ret *float32, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.engineLoad")
	if result.Exists() {
		val, ok := result.Value().(float64)
		if ok {
			retVal, err := ToVehicleOBDEngineLoad0(val)
			if err == nil {
				return &retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.engineLoad': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.engineLoad' is not of type 'float64' got '%v' of type '%T'", errInvalidType, result.Value(), result.Value()))
		}
	}

	return ret, errs
}

// VehicleOBDIntakeTempFromData converts the given JSON data to a *float32.
func VehicleOBDIntakeTempFromData(jsonData []byte) (ret *float32, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.intakeTemp")
	if result.Exists() {
		val, ok := result.Value().(float64)
		if ok {
			retVal, err := ToVehicleOBDIntakeTemp0(val)
			if err == nil {
				return &retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.intakeTemp': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.intakeTemp' is not of type 'float64' got '%v' of type '%T'", errInvalidType, result.Value(), result.Value()))
		}
	}

	return ret, errs
}

// VehicleOBDRunTimeFromData converts the given JSON data to a *float32.
func VehicleOBDRunTimeFromData(jsonData []byte) (ret *float32, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.runTime")
	if result.Exists() {
		val, ok := result.Value().(float64)
		if ok {
			retVal, err := ToVehicleOBDRunTime0(val)
			if err == nil {
				return &retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.runTime': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.runTime' is not of type 'float64' got '%v' of type '%T'", errInvalidType, result.Value(), result.Value()))
		}
	}

	return ret, errs
}

// VehiclePowertrainCombustionEngineECTFromData converts the given JSON data to a *int16.
func VehiclePowertrainCombustionEngineECTFromData(jsonData []byte) (ret *int16, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.coolantTemp")
	if result.Exists() {
		val, ok := result.Value().(float64)
		if ok {
			retVal, err := ToVehiclePowertrainCombustionEngineECT0(val)
			if err == nil {
				return &retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.coolantTemp': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.coolantTemp' is not of type 'float64' got '%v' of type '%T'", errInvalidType, result.Value(), result.Value()))
		}
	}

	return ret, errs
}

// VehiclePowertrainCombustionEngineEngineOilLevelFromData converts the given JSON data to a *string.
func VehiclePowertrainCombustionEngineEngineOilLevelFromData(jsonData []byte) (ret *string, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.oil")
	if result.Exists() {
		val, ok := result.Value().(float64)
		if ok {
			retVal, err := ToVehiclePowertrainCombustionEngineEngineOilLevel0(val)
			if err == nil {
				return &retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.oil': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.oil' is not of type 'float64' got '%v' of type '%T'", errInvalidType, result.Value(), result.Value()))
		}
	}

	return ret, errs
}

// VehiclePowertrainCombustionEngineSpeedFromData converts the given JSON data to a *uint16.
func VehiclePowertrainCombustionEngineSpeedFromData(jsonData []byte) (ret *uint16, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.engineSpeed")
	if result.Exists() {
		val, ok := result.Value().(float64)
		if ok {
			retVal, err := ToVehiclePowertrainCombustionEngineSpeed0(val)
			if err == nil {
				return &retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.engineSpeed': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.engineSpeed' is not of type 'float64' got '%v' of type '%T'", errInvalidType, result.Value(), result.Value()))
		}
	}

	return ret, errs
}

// VehiclePowertrainCombustionEngineTPSFromData converts the given JSON data to a *uint8.
func VehiclePowertrainCombustionEngineTPSFromData(jsonData []byte) (ret *uint8, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.throttlePosition")
	if result.Exists() {
		val, ok := result.Value().(float64)
		if ok {
			retVal, err := ToVehiclePowertrainCombustionEngineTPS0(val)
			if err == nil {
				return &retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.throttlePosition': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.throttlePosition' is not of type 'float64' got '%v' of type '%T'", errInvalidType, result.Value(), result.Value()))
		}
	}

	return ret, errs
}

// VehiclePowertrainFuelSystemAbsoluteLevelFromData converts the given JSON data to a *float32.
func VehiclePowertrainFuelSystemAbsoluteLevelFromData(jsonData []byte) (ret *float32, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.fuelPercentRemaining")
	if result.Exists() {
		val, ok := result.Value().(float64)
		if ok {
			retVal, err := ToVehiclePowertrainFuelSystemAbsoluteLevel0(val)
			if err == nil {
				return &retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.fuelPercentRemaining': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.fuelPercentRemaining' is not of type 'float64' got '%v' of type '%T'", errInvalidType, result.Value(), result.Value()))
		}
	}

	return ret, errs
}

// VehiclePowertrainFuelSystemSupportedFuelTypesFromData converts the given JSON data to a []string.
func VehiclePowertrainFuelSystemSupportedFuelTypesFromData(jsonData []byte) (ret []string, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.fuelType")
	if result.Exists() {
		val, ok := result.Value().(string)
		if ok {
			retVal, err := ToVehiclePowertrainFuelSystemSupportedFuelTypes0(val)
			if err == nil {
				return retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.fuelType': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.fuelType' is not of type 'string' got '%v' of type '%T'", errInvalidType, result.Value(), result.Value()))
		}
	}

	return ret, errs
}

// VehiclePowertrainRangeFromData converts the given JSON data to a *uint32.
func VehiclePowertrainRangeFromData(jsonData []byte) (ret *uint32, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.range")
	if result.Exists() {
		val, ok := result.Value().(float64)
		if ok {
			retVal, err := ToVehiclePowertrainRange0(val)
			if err == nil {
				return &retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.range': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.range' is not of type 'float64' got '%v' of type '%T'", errInvalidType, result.Value(), result.Value()))
		}
	}

	return ret, errs
}

// VehiclePowertrainTractionBatteryChargingChargeLimitFromData converts the given JSON data to a *uint8.
func VehiclePowertrainTractionBatteryChargingChargeLimitFromData(jsonData []byte) (ret *uint8, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.chargeLimit")
	if result.Exists() {
		val, ok := result.Value().(float64)
		if ok {
			retVal, err := ToVehiclePowertrainTractionBatteryChargingChargeLimit0(val)
			if err == nil {
				return &retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.chargeLimit': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.chargeLimit' is not of type 'float64' got '%v' of type '%T'", errInvalidType, result.Value(), result.Value()))
		}
	}

	return ret, errs
}

// VehiclePowertrainTractionBatteryChargingIsChargingFromData converts the given JSON data to a *bool.
func VehiclePowertrainTractionBatteryChargingIsChargingFromData(jsonData []byte) (ret *bool, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.charging")
	if result.Exists() {
		val, ok := result.Value().(bool)
		if ok {
			retVal, err := ToVehiclePowertrainTractionBatteryChargingIsCharging0(val)
			if err == nil {
				return &retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.charging': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.charging' is not of type 'bool' got '%v' of type '%T'", errInvalidType, result.Value(), result.Value()))
		}
	}

	return ret, errs
}

// VehiclePowertrainTractionBatteryGrossCapacityFromData converts the given JSON data to a *uint16.
func VehiclePowertrainTractionBatteryGrossCapacityFromData(jsonData []byte) (ret *uint16, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.batteryCapacity")
	if result.Exists() {
		val, ok := result.Value().(float64)
		if ok {
			retVal, err := ToVehiclePowertrainTractionBatteryGrossCapacity0(val)
			if err == nil {
				return &retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.batteryCapacity': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.batteryCapacity' is not of type 'float64' got '%v' of type '%T'", errInvalidType, result.Value(), result.Value()))
		}
	}

	return ret, errs
}

// VehiclePowertrainTractionBatteryStateOfChargeCurrentFromData converts the given JSON data to a *float32.
func VehiclePowertrainTractionBatteryStateOfChargeCurrentFromData(jsonData []byte) (ret *float32, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.soc")
	if result.Exists() {
		val, ok := result.Value().(float64)
		if ok {
			retVal, err := ToVehiclePowertrainTractionBatteryStateOfChargeCurrent0(val)
			if err == nil {
				return &retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.soc': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.soc' is not of type 'float64' got '%v' of type '%T'", errInvalidType, result.Value(), result.Value()))
		}
	}

	return ret, errs
}

// VehiclePowertrainTransmissionTravelledDistanceFromData converts the given JSON data to a *float32.
func VehiclePowertrainTransmissionTravelledDistanceFromData(jsonData []byte) (ret *float32, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.odometer")
	if result.Exists() {
		val, ok := result.Value().(float64)
		if ok {
			retVal, err := ToVehiclePowertrainTransmissionTravelledDistance0(val)
			if err == nil {
				return &retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.odometer': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.odometer' is not of type 'float64' got '%v' of type '%T'", errInvalidType, result.Value(), result.Value()))
		}
	}

	return ret, errs
}

// VehiclePowertrainTypeFromData converts the given JSON data to a *string.
func VehiclePowertrainTypeFromData(jsonData []byte) (ret *string, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.fuelType")
	if result.Exists() {
		val, ok := result.Value().(string)
		if ok {
			retVal, err := ToVehiclePowertrainType0(val)
			if err == nil {
				return &retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.fuelType': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.fuelType' is not of type 'string' got '%v' of type '%T'", errInvalidType, result.Value(), result.Value()))
		}
	}

	return ret, errs
}

// VehicleSpeedFromData converts the given JSON data to a *float32.
func VehicleSpeedFromData(jsonData []byte) (ret *float32, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.speed")
	if result.Exists() {
		val, ok := result.Value().(float64)
		if ok {
			retVal, err := ToVehicleSpeed0(val)
			if err == nil {
				return &retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.speed': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.speed' is not of type 'float64' got '%v' of type '%T'", errInvalidType, result.Value(), result.Value()))
		}
	}

	return ret, errs
}

// VehicleVehicleIdentificationBrandFromData converts the given JSON data to a *string.
func VehicleVehicleIdentificationBrandFromData(jsonData []byte) (ret *string, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.make")
	if result.Exists() {
		val, ok := result.Value().(string)
		if ok {
			retVal, err := ToVehicleVehicleIdentificationBrand0(val)
			if err == nil {
				return &retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.make': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.make' is not of type 'string' got '%v' of type '%T'", errInvalidType, result.Value(), result.Value()))
		}
	}

	return ret, errs
}

// VehicleVehicleIdentificationModelFromData converts the given JSON data to a *string.
func VehicleVehicleIdentificationModelFromData(jsonData []byte) (ret *string, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.model")
	if result.Exists() {
		val, ok := result.Value().(string)
		if ok {
			retVal, err := ToVehicleVehicleIdentificationModel0(val)
			if err == nil {
				return &retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.model': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.model' is not of type 'string' got '%v' of type '%T'", errInvalidType, result.Value(), result.Value()))
		}
	}

	return ret, errs
}

// VehicleVehicleIdentificationYearFromData converts the given JSON data to a *uint16.
func VehicleVehicleIdentificationYearFromData(jsonData []byte) (ret *uint16, err error) {
	var errs error
	var result gjson.Result
	result = gjson.GetBytes(jsonData, "data.year")
	if result.Exists() {
		val, ok := result.Value().(float64)
		if ok {
			retVal, err := ToVehicleVehicleIdentificationYear0(val)
			if err == nil {
				return &retVal, nil
			}
			errs = errors.Join(errs, fmt.Errorf("failed to convert 'data.year': %w", err))
		} else {
			errs = errors.Join(errs, fmt.Errorf("%w, field 'data.year' is not of type 'float64' got '%v' of type '%T'", errInvalidType, result.Value(), result.Value()))
		}
	}

	return ret, errs
}
