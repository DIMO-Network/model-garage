package telemetry

import (
	"github.com/teslamotors/fleet-telemetry/protos"
)

// ConvertLocationToCurrentLocationLatitude converts a telemetry datum with key Location to the VSS signal CurrentLocationLatitude.
func ConvertLocationToCurrentLocationLatitude(val *protos.LocationValue) (float64, error) {
	return val.Latitude, nil
}

// ConvertLocationToCurrentLocationLongitude converts a telemetry datum with key Location to the VSS signal CurrentLocationLongitude.
func ConvertLocationToCurrentLocationLongitude(val *protos.LocationValue) (float64, error) {
	return val.Longitude, nil
}

// ConvertDetailedChargeStateToPowertrainTractionBatteryChargingIsCharging converts a telemetry datum with key DetailedChargeState to the VSS signal PowertrainTractionBatteryChargingIsCharging.
func ConvertDetailedChargeStateToPowertrainTractionBatteryChargingIsCharging(val protos.DetailedChargeStateValue) (float64, error) {
	switch val {
	case protos.DetailedChargeStateValue_DetailedChargeStateStarting, protos.DetailedChargeStateValue_DetailedChargeStateCharging:
		return 1, nil
	default:
		return 0, nil
	}
}

// ConvertDetailedChargeStateToPowertrainTractionBatteryChargingIsChargingCableConnected converts a telemetry datum with key DetailedChargeState to the VSS signal PowertrainTractionBatteryChargingIsChargingCableConnected.
func ConvertDetailedChargeStateToPowertrainTractionBatteryChargingIsChargingCableConnected(val protos.DetailedChargeStateValue) (float64, error) {
	switch val {
	case protos.DetailedChargeStateValue_DetailedChargeStateDisconnected:
		return 0, nil
	default:
		return 1, nil
	}
}

// ConvertACChargingPowerToPowertrainTractionBatteryCurrentPower converts a telemetry datum with key ACChargingPower to the VSS signal PowertrainTractionBatteryCurrentPower.
// The input value is expressed in W.
func ConvertACChargingPowerToPowertrainTractionBatteryCurrentPower(val float64) (float64, error) {
	return val, nil
}

// ConvertDCChargingPowerToPowertrainTractionBatteryCurrentPower converts a telemetry datum with key DCChargingPower to the VSS signal PowertrainTractionBatteryCurrentPower.
// The input value is expressed in W.
func ConvertDCChargingPowerToPowertrainTractionBatteryCurrentPower(val float64) (float64, error) {
	return val, nil
}

// ConvertDCChargingEnergyInToPowertrainTractionBatteryChargingAddedEnergy converts a telemetry datum with key DCChargingEnergyIn to the VSS signal PowertrainTractionBatteryChargingAddedEnergy.
// The input value is expressed in kWh.
func ConvertDCChargingEnergyInToPowertrainTractionBatteryChargingAddedEnergy(val float64) (float64, error) {
	return val, nil
}

// ConvertEnergyRemainingToPowertrainTractionBatteryStateOfChargeCurrentEnergy converts a telemetry datum with key EnergyRemaining to the VSS signal PowertrainTractionBatteryStateOfChargeCurrentEnergy.
// The input value is expressed in kWh.
func ConvertEnergyRemainingToPowertrainTractionBatteryStateOfChargeCurrentEnergy(val float64) (float64, error) {
	return val, nil
}

// ConvertSocToPowertrainTractionBatteryStateOfChargeCurrent converts a telemetry datum with key Soc to the VSS signal PowertrainTractionBatteryStateOfChargeCurrent.
// The input value is expressed in percent.
func ConvertSocToPowertrainTractionBatteryStateOfChargeCurrent(val float64) (float64, error) {
	return val, nil
}

// ConvertTpmsPressureFlToChassisAxleRow1WheelLeftTirePressure converts a telemetry datum with key TpmsPressureFl to the VSS signal ChassisAxleRow1WheelLeftTirePressure.
// The input value is expressed in kPa.
func ConvertTpmsPressureFlToChassisAxleRow1WheelLeftTirePressure(val float64) (float64, error) {
	return val, nil
}

// ConvertTpmsPressureFrToChassisAxleRow1WheelRightTirePressure converts a telemetry datum with key TpmsPressureFr to the VSS signal ChassisAxleRow1WheelRightTirePressure.
// The input value is expressed in kPa.
func ConvertTpmsPressureFrToChassisAxleRow1WheelRightTirePressure(val float64) (float64, error) {
	return val, nil
}

// ConvertTpmsPressureRlToChassisAxleRow2WheelLeftTirePressure converts a telemetry datum with key TpmsPressureRl to the VSS signal ChassisAxleRow2WheelLeftTirePressure.
// The input value is expressed in kPa.
func ConvertTpmsPressureRlToChassisAxleRow2WheelLeftTirePressure(val float64) (float64, error) {
	return val, nil
}

// ConvertTpmsPressureRrToChassisAxleRow2WheelRightTirePressure converts a telemetry datum with key TpmsPressureRr to the VSS signal ChassisAxleRow2WheelRightTirePressure.
// The input value is expressed in kPa.
func ConvertTpmsPressureRrToChassisAxleRow2WheelRightTirePressure(val float64) (float64, error) {
	return val, nil
}

// ConvertOutsideTempToExteriorAirTemperature converts a telemetry datum with key OutsideTemp to the VSS signal ExteriorAirTemperature.
// The input value is expressed in celsius.
func ConvertOutsideTempToExteriorAirTemperature(val float64) (float64, error) {
	return val, nil
}

// ConvertEstBatteryRangeToPowertrainRange converts a telemetry datum with key EstBatteryRange to the VSS signal PowertrainRange.
// The input value is expressed in km.
func ConvertEstBatteryRangeToPowertrainRange(val float64) (float64, error) {
	return val, nil
}

// ConvertChargeLimitSocToPowertrainTractionBatteryChargingChargeLimit converts a telemetry datum with key ChargeLimitSoc to the VSS signal PowertrainTractionBatteryChargingChargeLimit.
// The input value is expressed in percent.
func ConvertChargeLimitSocToPowertrainTractionBatteryChargingChargeLimit(val int32) (float64, error) {
	return float64(val), nil
}

// ConvertOdometerToPowertrainTransmissionTravelledDistance converts a telemetry datum with key Odometer to the VSS signal PowertrainTransmissionTravelledDistance.
// The input value is expressed in km.
func ConvertOdometerToPowertrainTransmissionTravelledDistance(val float64) (float64, error) {
	return val, nil
}

// ConvertVehicleSpeedToSpeed converts a telemetry datum with key VehicleSpeed to the VSS signal Speed.
// The input value is expressed in km/h.
func ConvertVehicleSpeedToSpeed(val float64) (float64, error) {
	return val, nil
}

// ConvertDoorStateToCabinDoorRow1DriverSideIsOpen converts a telemetry datum with key DoorState to the VSS signal CabinDoorRow1DriverSideIsOpen.
func ConvertDoorStateToCabinDoorRow1DriverSideIsOpen(val *protos.Doors) (float64, error) {
	return boolToFloat64(val.DriverFront), nil
}

// ConvertDoorStateToCabinDoorRow1PassengerSideIsOpen converts a telemetry datum with key DoorState to the VSS signal CabinDoorRow1PassengerSideIsOpen.
func ConvertDoorStateToCabinDoorRow1PassengerSideIsOpen(val *protos.Doors) (float64, error) {
	return boolToFloat64(val.PassengerFront), nil
}

// ConvertDoorStateToCabinDoorRow2DriverSideIsOpen converts a telemetry datum with key DoorState to the VSS signal CabinDoorRow2DriverSideIsOpen.
func ConvertDoorStateToCabinDoorRow2DriverSideIsOpen(val *protos.Doors) (float64, error) {
	return boolToFloat64(val.DriverRear), nil
}

// ConvertDoorStateToCabinDoorRow2PassengerSideIsOpen converts a telemetry datum with key DoorState to the VSS signal CabinDoorRow2PassengerSideIsOpen.
func ConvertDoorStateToCabinDoorRow2PassengerSideIsOpen(val *protos.Doors) (float64, error) {
	return boolToFloat64(val.PassengerRear), nil
}

// ConvertFdWindowToCabinDoorRow1DriverSideWindowIsOpen converts a telemetry datum with key FdWindow to the VSS signal CabinDoorRow1DriverSideWindowIsOpen.
func ConvertFdWindowToCabinDoorRow1DriverSideWindowIsOpen(val protos.WindowState) (float64, error) {
	return windowStateToIsOpen(val)
}

// ConvertFpWindowToCabinDoorRow1PassengerSideWindowIsOpen converts a telemetry datum with key FpWindow to the VSS signal CabinDoorRow1PassengerSideWindowIsOpen.
func ConvertFpWindowToCabinDoorRow1PassengerSideWindowIsOpen(val protos.WindowState) (float64, error) {
	return windowStateToIsOpen(val)
}

// ConvertRdWindowToCabinDoorRow2DriverSideWindowIsOpen converts a telemetry datum with key RdWindow to the VSS signal CabinDoorRow2DriverSideWindowIsOpen.
func ConvertRdWindowToCabinDoorRow2DriverSideWindowIsOpen(val protos.WindowState) (float64, error) {
	return windowStateToIsOpen(val)
}

// ConvertRpWindowToCabinDoorRow2PassengerSideWindowIsOpen converts a telemetry datum with key RpWindow to the VSS signal CabinDoorRow2PassengerSideWindowIsOpen.
func ConvertRpWindowToCabinDoorRow2PassengerSideWindowIsOpen(val protos.WindowState) (float64, error) {
	return windowStateToIsOpen(val)
}

// ConvertChargeAmpsToPowertrainTractionBatteryChargingChargeCurrentAC converts a telemetry datum with key ChargeAmps to the VSS signal PowertrainTractionBatteryChargingChargeCurrentAC.
// The input value is expressed in A.
func ConvertChargeAmpsToPowertrainTractionBatteryChargingChargeCurrentAC(val float64) (float64, error) {
	return val, nil
}

// ConvertChargerVoltageToPowertrainTractionBatteryChargingChargeVoltageUnknownType converts a telemetry datum with key ChargerVoltage to the VSS signal PowertrainTractionBatteryChargingChargeVoltageUnknownType.
// The input value is expressed in V.
func ConvertChargerVoltageToPowertrainTractionBatteryChargingChargeVoltageUnknownType(val float64) (float64, error) {
	return val, nil
}
