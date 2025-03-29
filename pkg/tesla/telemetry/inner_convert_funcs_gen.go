package telemetry

import (
	"strings"

	"github.com/DIMO-Network/model-garage/pkg/tesla/telemetry/unit"
	"github.com/teslamotors/fleet-telemetry/protos"
)

// ConvertLocationLocationValueToCurrentLocationLatitude converts a telemetry datum with key Location to the VSS signal CurrentLocationLatitude.
func ConvertLocationLocationValueToCurrentLocationLatitude(val *protos.LocationValue) (float64, error) {
	return val.Latitude, nil
}

// ConvertLocationLocationValueToCurrentLocationLongitude converts a telemetry datum with key Location to the VSS signal CurrentLocationLongitude.
func ConvertLocationLocationValueToCurrentLocationLongitude(val *protos.LocationValue) (float64, error) {
	return val.Longitude, nil
}

// ConvertDetailedChargeStateDetailedChargeStateValueToPowertrainTractionBatteryChargingIsCharging converts a telemetry datum with key DetailedChargeState to the VSS signal PowertrainTractionBatteryChargingIsCharging.
func ConvertDetailedChargeStateDetailedChargeStateValueToPowertrainTractionBatteryChargingIsCharging(val protos.DetailedChargeStateValue) (float64, error) {
	switch val {
	case protos.DetailedChargeStateValue_DetailedChargeStateStarting, protos.DetailedChargeStateValue_DetailedChargeStateCharging:
		return 1, nil
	default:
		return 0, nil
	}
}

// ConvertACChargingPowerStringToPowertrainTractionBatteryCurrentPower converts a telemetry datum with key ACChargingPower to the VSS signal PowertrainTractionBatteryCurrentPower.
// Before being passed to this function, the string value is parsed into a float64.
// Before being passed to this function, the units are converted using KilowattsToWatts
func ConvertACChargingPowerStringToPowertrainTractionBatteryCurrentPower(val float64) (float64, error) {
	return val, nil
}

// ConvertDCChargingPowerStringToPowertrainTractionBatteryCurrentPower converts a telemetry datum with key DCChargingPower to the VSS signal PowertrainTractionBatteryCurrentPower.
// Before being passed to this function, the string value is parsed into a float64.
// Before being passed to this function, the units are converted using KilowattsToWatts
func ConvertDCChargingPowerStringToPowertrainTractionBatteryCurrentPower(val float64) (float64, error) {
	return val, nil
}

// ConvertDCChargingEnergyInStringToPowertrainTractionBatteryChargingAddedEnergy converts a telemetry datum with key DCChargingEnergyIn to the VSS signal PowertrainTractionBatteryChargingAddedEnergy.
// Before being passed to this function, the string value is parsed into a float64.
func ConvertDCChargingEnergyInStringToPowertrainTractionBatteryChargingAddedEnergy(val float64) (float64, error) {
	return val, nil
}

// ConvertEnergyRemainingStringToPowertrainTractionBatteryStateOfChargeCurrentEnergy converts a telemetry datum with key EnergyRemaining to the VSS signal PowertrainTractionBatteryStateOfChargeCurrentEnergy.
// Before being passed to this function, the string value is parsed into a float64.
func ConvertEnergyRemainingStringToPowertrainTractionBatteryStateOfChargeCurrentEnergy(val float64) (float64, error) {
	return val, nil
}

// ConvertSocStringToPowertrainTractionBatteryStateOfChargeCurrent converts a telemetry datum with key Soc to the VSS signal PowertrainTractionBatteryStateOfChargeCurrent.
// Before being passed to this function, the string value is parsed into a float64.
func ConvertSocStringToPowertrainTractionBatteryStateOfChargeCurrent(val float64) (float64, error) {
	return val, nil
}

// ConvertTpmsPressureFlStringToChassisAxleRow1WheelLeftTirePressure converts a telemetry datum with key TpmsPressureFl to the VSS signal ChassisAxleRow1WheelLeftTirePressure.
// Before being passed to this function, the string value is parsed into a float64.
// Before being passed to this function, the units are converted using AtmospheresToKilopascals
func ConvertTpmsPressureFlStringToChassisAxleRow1WheelLeftTirePressure(val float64) (float64, error) {
	return val, nil
}

// ConvertTpmsPressureFrStringToChassisAxleRow1WheelRightTirePressure converts a telemetry datum with key TpmsPressureFr to the VSS signal ChassisAxleRow1WheelRightTirePressure.
// Before being passed to this function, the string value is parsed into a float64.
// Before being passed to this function, the units are converted using AtmospheresToKilopascals
func ConvertTpmsPressureFrStringToChassisAxleRow1WheelRightTirePressure(val float64) (float64, error) {
	return val, nil
}

// ConvertTpmsPressureRlStringToChassisAxleRow2WheelLeftTirePressure converts a telemetry datum with key TpmsPressureRl to the VSS signal ChassisAxleRow2WheelLeftTirePressure.
// Before being passed to this function, the string value is parsed into a float64.
// Before being passed to this function, the units are converted using AtmospheresToKilopascals
func ConvertTpmsPressureRlStringToChassisAxleRow2WheelLeftTirePressure(val float64) (float64, error) {
	return val, nil
}

// ConvertTpmsPressureRrStringToChassisAxleRow2WheelRightTirePressure converts a telemetry datum with key TpmsPressureRr to the VSS signal ChassisAxleRow2WheelRightTirePressure.
// Before being passed to this function, the string value is parsed into a float64.
// Before being passed to this function, the units are converted using AtmospheresToKilopascals
func ConvertTpmsPressureRrStringToChassisAxleRow2WheelRightTirePressure(val float64) (float64, error) {
	return val, nil
}

// ConvertOutsideTempStringToExteriorAirTemperature converts a telemetry datum with key OutsideTemp to the VSS signal ExteriorAirTemperature.
// Before being passed to this function, the string value is parsed into a float64.
func ConvertOutsideTempStringToExteriorAirTemperature(val float64) (float64, error) {
	return val, nil
}

// ConvertEstBatteryRangeStringToPowertrainRange converts a telemetry datum with key EstBatteryRange to the VSS signal PowertrainRange.
// Before being passed to this function, the string value is parsed into a float64.
func ConvertEstBatteryRangeStringToPowertrainRange(val float64) (float64, error) {
	// We have to do this because we've violated the VSS spec and used km here instead of m.
	return unit.MilesToKilometers(val), nil
}

// ConvertChargeLimitSocStringToPowertrainTractionBatteryChargingChargeLimit converts a telemetry datum with key ChargeLimitSoc to the VSS signal PowertrainTractionBatteryChargingChargeLimit.
// Before being passed to this function, the string value is parsed into a float64.
func ConvertChargeLimitSocStringToPowertrainTractionBatteryChargingChargeLimit(val float64) (float64, error) {
	return val, nil
}

// ConvertOdometerStringToPowertrainTransmissionTravelledDistance converts a telemetry datum with key Odometer to the VSS signal PowertrainTransmissionTravelledDistance.
// Before being passed to this function, the string value is parsed into a float64.
// Before being passed to this function, the units are converted using MilesToKilometers
func ConvertOdometerStringToPowertrainTransmissionTravelledDistance(val float64) (float64, error) {
	return val, nil
}

// ConvertVehicleSpeedStringToSpeed converts a telemetry datum with key VehicleSpeed to the VSS signal Speed.
// Before being passed to this function, the string value is parsed into a float64.
// Before being passed to this function, the units are converted using MilesPerHourToKilometersPerHour
func ConvertVehicleSpeedStringToSpeed(val float64) (float64, error) {
	return val, nil
}

// ConvertDoorStateStringToCabinDoorRow1DriverSideIsOpen converts a telemetry datum with key DoorState to the VSS signal CabinDoorRow1DriverSideIsOpen.
func ConvertDoorStateStringToCabinDoorRow1DriverSideIsOpen(val string) (float64, error) {
	if strings.Contains(val, "DriverFront") {
		return 1, nil
	}
	return 0, nil
}

// ConvertDoorStateStringToCabinDoorRow1PassengerSideIsOpen converts a telemetry datum with key DoorState to the VSS signal CabinDoorRow1PassengerSideIsOpen.
func ConvertDoorStateStringToCabinDoorRow1PassengerSideIsOpen(val string) (float64, error) {
	if strings.Contains(val, "PassengerFront") {
		return 1, nil
	}
	return 0, nil
}

// ConvertDoorStateStringToCabinDoorRow2DriverSideIsOpen converts a telemetry datum with key DoorState to the VSS signal CabinDoorRow2DriverSideIsOpen.
func ConvertDoorStateStringToCabinDoorRow2DriverSideIsOpen(val string) (float64, error) {
	if strings.Contains(val, "DriverRear") {
		return 1, nil
	}
	return 0, nil
}

// ConvertDoorStateStringToCabinDoorRow2PassengerSideIsOpen converts a telemetry datum with key DoorState to the VSS signal CabinDoorRow2PassengerSideIsOpen.
func ConvertDoorStateStringToCabinDoorRow2PassengerSideIsOpen(val string) (float64, error) {
	if strings.Contains(val, "PassengerRear") {
		return 1, nil
	}
	return 0, nil
}

// ConvertFdWindowStringToCabinDoorRow1DriverSideWindowIsOpen converts a telemetry datum with key FdWindow to the VSS signal CabinDoorRow1DriverSideWindowIsOpen.
func ConvertFdWindowStringToCabinDoorRow1DriverSideWindowIsOpen(val string) (float64, error) {
	return windowStateToIsOpen(val), nil
}

// ConvertFpWindowStringToCabinDoorRow1PassengerSideWindowIsOpen converts a telemetry datum with key FpWindow to the VSS signal CabinDoorRow1PassengerSideWindowIsOpen.
func ConvertFpWindowStringToCabinDoorRow1PassengerSideWindowIsOpen(val string) (float64, error) {
	return windowStateToIsOpen(val), nil
}

// ConvertRdWindowStringToCabinDoorRow2DriverSideWindowIsOpen converts a telemetry datum with key RdWindow to the VSS signal CabinDoorRow2DriverSideWindowIsOpen.
func ConvertRdWindowStringToCabinDoorRow2DriverSideWindowIsOpen(val string) (float64, error) {
	return windowStateToIsOpen(val), nil
}

// ConvertRpWindowStringToCabinDoorRow2PassengerSideWindowIsOpen converts a telemetry datum with key RpWindow to the VSS signal CabinDoorRow2PassengerSideWindowIsOpen.
func ConvertRpWindowStringToCabinDoorRow2PassengerSideWindowIsOpen(val string) (float64, error) {
	return windowStateToIsOpen(val), nil
}

// ConvertChargeAmpsStringToPowertrainTractionBatteryChargingChargeCurrentAC converts a telemetry datum with key ChargeAmps to the VSS signal PowertrainTractionBatteryChargingChargeCurrentAC.
// Before being passed to this function, the string value is parsed into a float64.
func ConvertChargeAmpsStringToPowertrainTractionBatteryChargingChargeCurrentAC(val float64) (float64, error) {
	return val, nil
}

// ConvertChargerVoltageDoubleToPowertrainTractionBatteryChargingChargeVoltageUnknownType converts a telemetry datum with key ChargerVoltage to the VSS signal PowertrainTractionBatteryChargingChargeVoltageUnknownType.
func ConvertChargerVoltageDoubleToPowertrainTractionBatteryChargingChargeVoltageUnknownType(val float64) (float64, error) {
	return val, nil
}
