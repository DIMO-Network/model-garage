// Code generated by github.com/DIMO-Network/model-garage.
package nativestatus

import "math"

// This file is automatically populated with conversion functions for each field of the model struct.
// any conversion functions already defined in this package will be coppied through.
// note: DO NOT mutate the orginalDoc parameter which is shared between all conversion functions.

// ToAngularVelocityYaw0 converts data from field 'yawRate' of type float64 to 'Vehicle.AngularVelocity.Yaw' of type float64.
// Vehicle.AngularVelocity.Yaw: Vehicle rotation rate along Z (vertical).
// Unit: 'degrees/s'
func ToAngularVelocityYaw0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToChassisAxleRow1WheelLeftSpeed0 converts data from field 'frontlLeftWheelSpeed' of type float64 to 'Vehicle.Chassis.Axle.Row1.Wheel.Left.Speed' of type float64.
// Vehicle.Chassis.Axle.Row1.Wheel.Left.Speed: Rotational speed of a vehicle's wheel.
// Unit: 'km/h'
func ToChassisAxleRow1WheelLeftSpeed0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToChassisAxleRow1WheelLeftTirePressure0 converts data from field 'tires.frontLeft' of type float64 to 'Vehicle.Chassis.Axle.Row1.Wheel.Left.Tire.Pressure' of type float64.
// Vehicle.Chassis.Axle.Row1.Wheel.Left.Tire.Pressure: Tire pressure in kilo-Pascal.
// Unit: 'kPa'
func ToChassisAxleRow1WheelLeftTirePressure0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToChassisAxleRow1WheelLeftTirePressure1 converts data from field 'tiresFrontLeft' of type float64 to 'Vehicle.Chassis.Axle.Row1.Wheel.Left.Tire.Pressure' of type float64.
// Vehicle.Chassis.Axle.Row1.Wheel.Left.Tire.Pressure: Tire pressure in kilo-Pascal.
// Unit: 'kPa'
func ToChassisAxleRow1WheelLeftTirePressure1(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToChassisAxleRow1WheelRightSpeed0 converts data from field 'frontRightWheelSpeed' of type float64 to 'Vehicle.Chassis.Axle.Row1.Wheel.Right.Speed' of type float64.
// Vehicle.Chassis.Axle.Row1.Wheel.Right.Speed: Rotational speed of a vehicle's wheel.
// Unit: 'km/h'
func ToChassisAxleRow1WheelRightSpeed0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToChassisAxleRow1WheelRightTirePressure0 converts data from field 'tires.frontRight' of type float64 to 'Vehicle.Chassis.Axle.Row1.Wheel.Right.Tire.Pressure' of type float64.
// Vehicle.Chassis.Axle.Row1.Wheel.Right.Tire.Pressure: Tire pressure in kilo-Pascal.
// Unit: 'kPa'
func ToChassisAxleRow1WheelRightTirePressure0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToChassisAxleRow1WheelRightTirePressure1 converts data from field 'tiresFrontRight' of type float64 to 'Vehicle.Chassis.Axle.Row1.Wheel.Right.Tire.Pressure' of type float64.
// Vehicle.Chassis.Axle.Row1.Wheel.Right.Tire.Pressure: Tire pressure in kilo-Pascal.
// Unit: 'kPa'
func ToChassisAxleRow1WheelRightTirePressure1(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToChassisAxleRow2WheelLeftTirePressure0 converts data from field 'tires.backLeft' of type float64 to 'Vehicle.Chassis.Axle.Row2.Wheel.Left.Tire.Pressure' of type float64.
// Vehicle.Chassis.Axle.Row2.Wheel.Left.Tire.Pressure: Tire pressure in kilo-Pascal.
// Unit: 'kPa'
func ToChassisAxleRow2WheelLeftTirePressure0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToChassisAxleRow2WheelLeftTirePressure1 converts data from field 'tiresBackLeft' of type float64 to 'Vehicle.Chassis.Axle.Row2.Wheel.Left.Tire.Pressure' of type float64.
// Vehicle.Chassis.Axle.Row2.Wheel.Left.Tire.Pressure: Tire pressure in kilo-Pascal.
// Unit: 'kPa'
func ToChassisAxleRow2WheelLeftTirePressure1(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToChassisAxleRow2WheelRightTirePressure0 converts data from field 'tires.backRight' of type float64 to 'Vehicle.Chassis.Axle.Row2.Wheel.Right.Tire.Pressure' of type float64.
// Vehicle.Chassis.Axle.Row2.Wheel.Right.Tire.Pressure: Tire pressure in kilo-Pascal.
// Unit: 'kPa'
func ToChassisAxleRow2WheelRightTirePressure0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToChassisAxleRow2WheelRightTirePressure1 converts data from field 'tiresBackRight' of type float64 to 'Vehicle.Chassis.Axle.Row2.Wheel.Right.Tire.Pressure' of type float64.
// Vehicle.Chassis.Axle.Row2.Wheel.Right.Tire.Pressure: Tire pressure in kilo-Pascal.
// Unit: 'kPa'
func ToChassisAxleRow2WheelRightTirePressure1(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToCurrentLocationAltitude0 converts data from field 'altitude' of type float64 to 'Vehicle.CurrentLocation.Altitude' of type float64.
// Vehicle.CurrentLocation.Altitude: Current altitude relative to WGS 84 reference ellipsoid, as measured at the position of GNSS receiver antenna.
// Unit: 'm'
func ToCurrentLocationAltitude0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToCurrentLocationIsRedacted0 converts data from field 'isRedacted' of type bool to 'Vehicle.CurrentLocation.IsRedacted' of type float64.
// Vehicle.CurrentLocation.IsRedacted: Indicates if the latitude and longitude signals at the current timestamp have been redacted using a privacy zone.
func ToCurrentLocationIsRedacted0(originalDoc []byte, val bool) (float64, error) {
	if val {
		return 1, nil
	}
	return 0, nil
}

// ToCurrentLocationLatitude0 converts data from field 'latitude' of type float64 to 'Vehicle.CurrentLocation.Latitude' of type float64.
// Vehicle.CurrentLocation.Latitude: Current latitude of vehicle in WGS 84 geodetic coordinates, as measured at the position of GNSS receiver antenna.
// Unit: 'degrees' Min: '-90' Max: '90'
func ToCurrentLocationLatitude0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToCurrentLocationLongitude0 converts data from field 'longitude' of type float64 to 'Vehicle.CurrentLocation.Longitude' of type float64.
// Vehicle.CurrentLocation.Longitude: Current longitude of vehicle in WGS 84 geodetic coordinates, as measured at the position of GNSS receiver antenna.
// Unit: 'degrees' Min: '-180' Max: '180'
func ToCurrentLocationLongitude0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToDIMOAftermarketHDOP0 converts data from field 'hdop' of type float64 to 'Vehicle.DIMO.Aftermarket.HDOP' of type float64.
// Vehicle.DIMO.Aftermarket.HDOP: Horizontal dilution of precision of GPS
func ToDIMOAftermarketHDOP0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToDIMOAftermarketNSAT0 converts data from field 'nsat' of type float64 to 'Vehicle.DIMO.Aftermarket.NSAT' of type float64.
// Vehicle.DIMO.Aftermarket.NSAT: Number of sync satellites for GPS
func ToDIMOAftermarketNSAT0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToDIMOAftermarketSSID0 converts data from field 'ssid' of type string to 'Vehicle.DIMO.Aftermarket.SSID' of type string.
// Vehicle.DIMO.Aftermarket.SSID: Service Set Identifier for the wifi.
func ToDIMOAftermarketSSID0(originalDoc []byte, val string) (string, error) {
	return val, nil
}

// ToDIMOAftermarketSSID1 converts data from field 'wifi.ssid' of type string to 'Vehicle.DIMO.Aftermarket.SSID' of type string.
// Vehicle.DIMO.Aftermarket.SSID: Service Set Identifier for the wifi.
func ToDIMOAftermarketSSID1(originalDoc []byte, val string) (string, error) {
	return val, nil
}

// ToDIMOAftermarketWPAState0 converts data from field 'wpa_state' of type string to 'Vehicle.DIMO.Aftermarket.WPAState' of type string.
// Vehicle.DIMO.Aftermarket.WPAState: Indicate the current WPA state for the device's wifi, e.g. "CONNECTED", "SCANNING", "DISCONNECTED"
func ToDIMOAftermarketWPAState0(originalDoc []byte, val string) (string, error) {
	return val, nil
}

// ToDIMOAftermarketWPAState1 converts data from field 'wifi.wpaState' of type string to 'Vehicle.DIMO.Aftermarket.WPAState' of type string.
// Vehicle.DIMO.Aftermarket.WPAState: Indicate the current WPA state for the device's wifi, e.g. "CONNECTED", "SCANNING", "DISCONNECTED"
func ToDIMOAftermarketWPAState1(originalDoc []byte, val string) (string, error) {
	return val, nil
}

// ToExteriorAirTemperature0 converts data from field 'ambientAirTemp' of type float64 to 'Vehicle.Exterior.AirTemperature' of type float64.
// Vehicle.Exterior.AirTemperature: Air temperature outside the vehicle.
// Unit: 'celsius'
func ToExteriorAirTemperature0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToExteriorAirTemperature1 converts data from field 'ambientTemp' of type float64 to 'Vehicle.Exterior.AirTemperature' of type float64.
// Vehicle.Exterior.AirTemperature: Air temperature outside the vehicle.
// Unit: 'celsius'
func ToExteriorAirTemperature1(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToLowVoltageBatteryCurrentVoltage0 converts data from field 'batteryVoltage' of type float64 to 'Vehicle.LowVoltageBattery.CurrentVoltage' of type float64.
// Vehicle.LowVoltageBattery.CurrentVoltage: Current Voltage of the low voltage battery.
// Unit: 'V'
func ToLowVoltageBatteryCurrentVoltage0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToOBDBarometricPressure0 converts data from field 'barometricPressure' of type float64 to 'Vehicle.OBD.BarometricPressure' of type float64.
// Vehicle.OBD.BarometricPressure: PID 33 - Barometric pressure
// Unit: 'kPa'
func ToOBDBarometricPressure0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToOBDCommandedEGR0 converts data from field 'commandedEgr' of type float64 to 'Vehicle.OBD.CommandedEGR' of type float64.
// Vehicle.OBD.CommandedEGR: PID 2C - Commanded exhaust gas recirculation (EGR)
// Unit: 'percent'
func ToOBDCommandedEGR0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToOBDCommandedEVAP0 converts data from field 'evap' of type float64 to 'Vehicle.OBD.CommandedEVAP' of type float64.
// Vehicle.OBD.CommandedEVAP: PID 2E - Commanded evaporative purge (EVAP) valve
// Unit: 'percent'
func ToOBDCommandedEVAP0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToOBDDistanceSinceDTCClear0 converts data from field 'distanceSinceDtcClear' of type float64 to 'Vehicle.OBD.DistanceSinceDTCClear' of type float64.
// Vehicle.OBD.DistanceSinceDTCClear: PID 31 - Distance traveled since codes cleared
// Unit: 'km'
func ToOBDDistanceSinceDTCClear0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToOBDDistanceWithMIL0 converts data from field 'distanceWMil' of type float64 to 'Vehicle.OBD.DistanceWithMIL' of type float64.
// Vehicle.OBD.DistanceWithMIL: PID 21 - Distance traveled with MIL on
// Unit: 'km'
func ToOBDDistanceWithMIL0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToOBDEngineLoad0 converts data from field 'engineLoad' of type float64 to 'Vehicle.OBD.EngineLoad' of type float64.
// Vehicle.OBD.EngineLoad: PID 04 - Engine load in percent - 0 = no load, 100 = full load
// Unit: 'percent'
func ToOBDEngineLoad0(originalDoc []byte, val float64) (float64, error) {
	schemaVersion := GetSchemaVersion(originalDoc)
	if hasV1Schema(schemaVersion) {
		return val * 100, nil
	}
	return val, nil
}

// ToOBDFuelPressure0 converts data from field 'fuelTankPressure' of type float64 to 'Vehicle.OBD.FuelPressure' of type float64.
// Vehicle.OBD.FuelPressure: PID 0A - Fuel pressure
// Unit: 'kPa'
func ToOBDFuelPressure0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToOBDIntakeTemp0 converts data from field 'intakeTemp' of type float64 to 'Vehicle.OBD.IntakeTemp' of type float64.
// Vehicle.OBD.IntakeTemp: PID 0F - Intake temperature
// Unit: 'celsius'
func ToOBDIntakeTemp0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToOBDLongTermFuelTrim10 converts data from field 'longTermFuelTrim1' of type float64 to 'Vehicle.OBD.LongTermFuelTrim1' of type float64.
// Vehicle.OBD.LongTermFuelTrim1: PID 07 - Long Term (learned) Fuel Trim - Bank 1 - negative percent leaner, positive percent richer
// Unit: 'percent'
func ToOBDLongTermFuelTrim10(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToOBDMAP0 converts data from field 'intakePressure' of type float64 to 'Vehicle.OBD.MAP' of type float64.
// Vehicle.OBD.MAP: PID 0B - Intake manifold pressure
// Unit: 'kPa'
func ToOBDMAP0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToOBDO2WRSensor1Voltage0 converts data from field 'oxygenSensor1' of type float64 to 'Vehicle.OBD.O2WR.Sensor1.Voltage' of type float64.
// Vehicle.OBD.O2WR.Sensor1.Voltage: PID 2x (byte CD) - Voltage for wide range/band oxygen sensor
// Unit: 'V'
func ToOBDO2WRSensor1Voltage0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToOBDO2WRSensor2Voltage0 converts data from field 'oxygenSensor2' of type float64 to 'Vehicle.OBD.O2WR.Sensor2.Voltage' of type float64.
// Vehicle.OBD.O2WR.Sensor2.Voltage: PID 2x (byte CD) - Voltage for wide range/band oxygen sensor
// Unit: 'V'
func ToOBDO2WRSensor2Voltage0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToOBDRunTime0 converts data from field 'runTime' of type float64 to 'Vehicle.OBD.RunTime' of type float64.
// Vehicle.OBD.RunTime: PID 1F - Engine run time
// Unit: 's'
func ToOBDRunTime0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToOBDShortTermFuelTrim10 converts data from field 'shortTermFuelTrim1' of type float64 to 'Vehicle.OBD.ShortTermFuelTrim1' of type float64.
// Vehicle.OBD.ShortTermFuelTrim1: PID 06 - Short Term (immediate) Fuel Trim - Bank 1 - negative percent leaner, positive percent richer
// Unit: 'percent'
func ToOBDShortTermFuelTrim10(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToOBDWarmupsSinceDTCClear0 converts data from field 'warmupsSinceDtcClear' of type float64 to 'Vehicle.OBD.WarmupsSinceDTCClear' of type float64.
// Vehicle.OBD.WarmupsSinceDTCClear: PID 30 - Number of warm-ups since codes cleared
func ToOBDWarmupsSinceDTCClear0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToPowertrainCombustionEngineECT0 converts data from field 'coolantTemp' of type float64 to 'Vehicle.Powertrain.CombustionEngine.ECT' of type float64.
// Vehicle.Powertrain.CombustionEngine.ECT: Engine coolant temperature.
// Unit: 'celsius'
func ToPowertrainCombustionEngineECT0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToPowertrainCombustionEngineEngineOilLevel0 converts data from field 'oil' of type float64 to 'Vehicle.Powertrain.CombustionEngine.EngineOilLevel' of type string.
// Vehicle.Powertrain.CombustionEngine.EngineOilLevel: Engine oil level.
func ToPowertrainCombustionEngineEngineOilLevel0(originalDoc []byte, val float64) (string, error) {
	switch {
	case val < 0.25:
		return "CRITICALLY_LOW", nil
	case val < 0.5:
		return "LOW", nil
	case val < 0.75:
		return "NORMAL", nil
	case val < .99:
		return "HIGH", nil
	default:
		return "CRITICALLY_HIGH", nil
	}
}

// ToPowertrainCombustionEngineEngineOilLevel1 converts data from field 'oilLife' of type float64 to 'Vehicle.Powertrain.CombustionEngine.EngineOilLevel' of type string.
// Vehicle.Powertrain.CombustionEngine.EngineOilLevel: Engine oil level.
func ToPowertrainCombustionEngineEngineOilLevel1(originalDoc []byte, val float64) (string, error) {
	panic("not implemented")
}

// ToPowertrainCombustionEngineEngineOilRelativeLevel0 converts data from field 'oil' of type float64 to 'Vehicle.Powertrain.CombustionEngine.EngineOilRelativeLevel' of type float64.
// Vehicle.Powertrain.CombustionEngine.EngineOilRelativeLevel: Engine oil level as a percentage.
// Unit: 'percent' Min: '0' Max: '100'
func ToPowertrainCombustionEngineEngineOilRelativeLevel0(originalDoc []byte, val float64) (float64, error) {
	// oil comes in as a value between 0 and 1, convert to percentage.
	return val * 100, nil
}

// ToPowertrainCombustionEngineMAF0 converts data from field 'maf' of type float64 to 'Vehicle.Powertrain.CombustionEngine.MAF' of type float64.
// Vehicle.Powertrain.CombustionEngine.MAF: Grams of air drawn into engine per second.
// Unit: 'g/s'
func ToPowertrainCombustionEngineMAF0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToPowertrainCombustionEngineSpeed0 converts data from field 'rpm' of type float64 to 'Vehicle.Powertrain.CombustionEngine.Speed' of type float64.
// Vehicle.Powertrain.CombustionEngine.Speed: Engine speed measured as rotations per minute.
// Unit: 'rpm'
func ToPowertrainCombustionEngineSpeed0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToPowertrainCombustionEngineSpeed1 converts data from field 'engineSpeed' of type float64 to 'Vehicle.Powertrain.CombustionEngine.Speed' of type float64.
// Vehicle.Powertrain.CombustionEngine.Speed: Engine speed measured as rotations per minute.
// Unit: 'rpm'
func ToPowertrainCombustionEngineSpeed1(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToPowertrainCombustionEngineTPS0 converts data from field 'throttlePosition' of type float64 to 'Vehicle.Powertrain.CombustionEngine.TPS' of type float64.
// Vehicle.Powertrain.CombustionEngine.TPS: Current throttle position.
// Unit: 'percent'  Max: '100'
func ToPowertrainCombustionEngineTPS0(originalDoc []byte, val float64) (float64, error) {
	schemaVersion := GetSchemaVersion(originalDoc)
	if hasV1Schema(schemaVersion) {
		return val * 100, nil
	}
	return val, nil
}

// ToPowertrainCombustionEngineTorque0 converts data from field 'engineTorque' of type float64 to 'Vehicle.Powertrain.CombustionEngine.Torque' of type float64.
// Vehicle.Powertrain.CombustionEngine.Torque: Current engine torque. Shall be reported as 0 during engine breaking.
// Unit: 'Nm'
func ToPowertrainCombustionEngineTorque0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToPowertrainFuelSystemAbsoluteLevel0 converts data from field 'fuelLevelLiters' of type float64 to 'Vehicle.Powertrain.FuelSystem.AbsoluteLevel' of type float64.
// Vehicle.Powertrain.FuelSystem.AbsoluteLevel: Current available fuel in the fuel tank expressed in liters.
// Unit: 'l'
func ToPowertrainFuelSystemAbsoluteLevel0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToPowertrainFuelSystemRelativeLevel0 converts data from field 'fuelLevel' of type float64 to 'Vehicle.Powertrain.FuelSystem.RelativeLevel' of type float64.
// Vehicle.Powertrain.FuelSystem.RelativeLevel: Level in fuel tank as percent of capacity. 0 = empty. 100 = full.
// Unit: 'percent' Min: '0' Max: '100'
func ToPowertrainFuelSystemRelativeLevel0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToPowertrainFuelSystemRelativeLevel1 converts data from field 'fuelPercentRemaining' of type float64 to 'Vehicle.Powertrain.FuelSystem.RelativeLevel' of type float64.
// Vehicle.Powertrain.FuelSystem.RelativeLevel: Level in fuel tank as percent of capacity. 0 = empty. 100 = full.
// Unit: 'percent' Min: '0' Max: '100'
func ToPowertrainFuelSystemRelativeLevel1(originalDoc []byte, val float64) (float64, error) {
	// fuelPercentRemaining comes in as a value between 0 and 1, convert to percentage.
	return val * 100, nil
}

// ToPowertrainFuelSystemSupportedFuelTypes0 converts data from field 'fuelType' of type string to 'Vehicle.Powertrain.FuelSystem.SupportedFuelTypes' of type string.
// Vehicle.Powertrain.FuelSystem.SupportedFuelTypes: High level information of fuel types supported
func ToPowertrainFuelSystemSupportedFuelTypes0(originalDoc []byte, val string) (string, error) {
	switch val {
	case "Gasoline":
		return "GASOLINE", nil
	case "Ethanol":
		return "E85", nil
	case "Diesel":
		return "DIESEL", nil
	case "LPG":
		return "LPG", nil
	default:
		return "OTHER", nil
	}
}

// ToPowertrainRange0 converts data from field 'range' of type float64 to 'Vehicle.Powertrain.Range' of type float64.
// Vehicle.Powertrain.Range: Remaining range in kilometers using all energy sources available in the vehicle.
// Unit: 'km'
func ToPowertrainRange0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToPowertrainTractionBatteryChargingChargeLimit0 converts data from field 'chargeLimit' of type float64 to 'Vehicle.Powertrain.TractionBattery.Charging.ChargeLimit' of type float64.
// Vehicle.Powertrain.TractionBattery.Charging.ChargeLimit: Target charge limit (state of charge) for battery.
// Unit: 'percent' Min: '0' Max: '100'
func ToPowertrainTractionBatteryChargingChargeLimit0(originalDoc []byte, val float64) (float64, error) {
	// chargeLimit comes in as a value between 0 and 1, convert to percentage.
	return val * 100, nil
}

// ToPowertrainTractionBatteryChargingIsCharging0 converts data from field 'charging' of type bool to 'Vehicle.Powertrain.TractionBattery.Charging.IsCharging' of type float64.
// Vehicle.Powertrain.TractionBattery.Charging.IsCharging: True if charging is ongoing. Charging is considered to be ongoing if energy is flowing from charger to vehicle.
func ToPowertrainTractionBatteryChargingIsCharging0(originalDoc []byte, val bool) (float64, error) {
	if val {
		return 1, nil
	}
	return 0, nil
}

// ToPowertrainTractionBatteryCurrentPower0 converts data from field 'charger.power' of type float64 to 'Vehicle.Powertrain.TractionBattery.CurrentPower' of type float64.
// Vehicle.Powertrain.TractionBattery.CurrentPower: Current electrical energy flowing in/out of battery. Positive = Energy flowing in to battery, e.g. during charging. Negative = Energy flowing out of battery, e.g. during driving.
// Unit: 'W'
func ToPowertrainTractionBatteryCurrentPower0(originalDoc []byte, val float64) (float64, error) {
	// V1 field is in kilowatts (kW), VSS field is in watts (W).
	return 1000 * val, nil
}

// ToPowertrainTractionBatteryCurrentVoltage0 converts data from field 'hvBatteryVoltage' of type float64 to 'Vehicle.Powertrain.TractionBattery.CurrentVoltage' of type float64.
// Vehicle.Powertrain.TractionBattery.CurrentVoltage: Current Voltage of the battery.
// Unit: 'V'
func ToPowertrainTractionBatteryCurrentVoltage0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToPowertrainTractionBatteryGrossCapacity0 converts data from field 'batteryCapacity' of type float64 to 'Vehicle.Powertrain.TractionBattery.GrossCapacity' of type float64.
// Vehicle.Powertrain.TractionBattery.GrossCapacity: Gross capacity of the battery.
// Unit: 'kWh'
func ToPowertrainTractionBatteryGrossCapacity0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToPowertrainTractionBatteryStateOfChargeCurrent0 converts data from field 'soc' of type float64 to 'Vehicle.Powertrain.TractionBattery.StateOfCharge.Current' of type float64.
// Vehicle.Powertrain.TractionBattery.StateOfCharge.Current: Physical state of charge of the high voltage battery, relative to net capacity. This is not necessarily the state of charge being displayed to the customer.
// Unit: 'percent' Min: '0' Max: '100.0'
func ToPowertrainTractionBatteryStateOfChargeCurrent0(originalDoc []byte, val float64) (float64, error) {
	schemaVersion := GetSchemaVersion(originalDoc)
	if hasV1Schema(schemaVersion) {
		// soc comes in as a value between 0 and 1, convert to percentage.
		return val * 100, nil
	}
	return val, nil
}

// ToPowertrainTractionBatteryTemperatureAverage0 converts data from field 'hvBatteryCoolantTemperature' of type float64 to 'Vehicle.Powertrain.TractionBattery.Temperature.Average' of type float64.
// Vehicle.Powertrain.TractionBattery.Temperature.Average: Current average temperature of the battery cells.
// Unit: 'celsius'
func ToPowertrainTractionBatteryTemperatureAverage0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToPowertrainTransmissionCurrentGear0 converts data from field 'gearSelection' of type float64 to 'Vehicle.Powertrain.Transmission.CurrentGear' of type float64.
// Vehicle.Powertrain.Transmission.CurrentGear: The current gear. 0=Neutral, 1/2/..=Forward, -1/-2/..=Reverse.
func ToPowertrainTransmissionCurrentGear0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToPowertrainTransmissionTemperature0 converts data from field 'atfTemperature' of type float64 to 'Vehicle.Powertrain.Transmission.Temperature' of type float64.
// Vehicle.Powertrain.Transmission.Temperature: The current gearbox temperature.
// Unit: 'celsius'
func ToPowertrainTransmissionTemperature0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToPowertrainTransmissionTravelledDistance0 converts data from field 'odometer' of type float64 to 'Vehicle.Powertrain.Transmission.TravelledDistance' of type float64.
// Vehicle.Powertrain.Transmission.TravelledDistance: Odometer reading, total distance travelled during the lifetime of the transmission.
// Unit: 'km'
func ToPowertrainTransmissionTravelledDistance0(originalDoc []byte, val float64) (float64, error) {
	if val > 999999 {
		// if the value is absurdly high, it is likely in meters, convert to kilometers
		// TODO: find a reliable way to determine if the value is in meters
		const metersToKilometers = 1000
		return math.Round(val / metersToKilometers), nil
	}
	return val, nil
}

// ToPowertrainType0 converts data from field 'fuelType' of type string to 'Vehicle.Powertrain.Type' of type string.
// Vehicle.Powertrain.Type: Defines the powertrain type of the vehicle.
func ToPowertrainType0(originalDoc []byte, val string) (string, error) {
	// possible arguments Gasoline, Ethanol, Diesel, Not available, Electric, LPG
	// deault to combustion
	if val == "Electric" {
		return "ELECTRIC", nil
	}
	return "COMBUSTION", nil
}

// ToServiceDistanceToService0 converts data from field 'serviceInterval' of type float64 to 'Vehicle.Service.DistanceToService' of type float64.
// Vehicle.Service.DistanceToService: Remaining distance to service (of any kind). Negative values indicate service overdue.
// Unit: 'km'
func ToServiceDistanceToService0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToSpeed0 converts data from field 'vehicleSpeed' of type float64 to 'Vehicle.Speed' of type float64.
// Vehicle.Speed: Vehicle speed.
// Unit: 'km/h'
func ToSpeed0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToSpeed1 converts data from field 'speed' of type float64 to 'Vehicle.Speed' of type float64.
// Vehicle.Speed: Vehicle speed.
// Unit: 'km/h'
func ToSpeed1(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}
