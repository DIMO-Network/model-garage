// Code generated by "model-garage" DO NOT EDIT.
package vss

import "time"

const (
	// FieldDefinitionID ID for the vehicles definition
	FieldDefinitionID = "DefinitionID"
	// FieldSource where the data was sourced from
	FieldSource = "Source"
	// FieldSubject subjet of this vehicle data
	FieldSubject = "Subject"
	// FieldTimestamp timestamp of when this data was colllected
	FieldTimestamp = "Timestamp"
	// FieldType type of data collected
	FieldType = "Type"
	// FieldVehicleCurrentLocationAltitude Current altitude relative to WGS 84 reference ellipsoid, as measured at the position of GNSS receiver antenna.
	FieldVehicleCurrentLocationAltitude = "Vehicle_CurrentLocation_Altitude"
	// FieldVehicleCurrentLocationLatitude Current latitude of vehicle in WGS 84 geodetic coordinates, as measured at the position of GNSS receiver antenna.
	FieldVehicleCurrentLocationLatitude = "Vehicle_CurrentLocation_Latitude"
	// FieldVehicleCurrentLocationLongitude Current longitude of vehicle in WGS 84 geodetic coordinates, as measured at the position of GNSS receiver antenna.
	FieldVehicleCurrentLocationLongitude = "Vehicle_CurrentLocation_Longitude"
	// FieldVehicleCurrentLocationTimestamp Timestamp from GNSS system for current location, formatted according to ISO 8601 with UTC time zone.
	FieldVehicleCurrentLocationTimestamp = "Vehicle_CurrentLocation_Timestamp"
	// FieldVehicleLowVoltageBatteryCurrentVoltage Current Voltage of the low voltage battery.
	FieldVehicleLowVoltageBatteryCurrentVoltage = "Vehicle_LowVoltageBattery_CurrentVoltage"
	// FieldVehicleSpeed Vehicle speed.
	FieldVehicleSpeed = "Vehicle_Speed"
	// FieldVehicleVehicleIdentificationBrand Vehicle brand or manufacturer.
	FieldVehicleVehicleIdentificationBrand = "Vehicle_VehicleIdentification_Brand"
	// FieldVehicleVehicleIdentificationModel Vehicle model.
	FieldVehicleVehicleIdentificationModel = "Vehicle_VehicleIdentification_Model"
	// FieldVehicleVehicleIdentificationVIN 17-character Vehicle Identification Number (VIN) as defined by ISO 3779.
	FieldVehicleVehicleIdentificationVIN = "Vehicle_VehicleIdentification_VIN"
	// FieldVehicleVehicleIdentificationYear Model year of the vehicle.
	FieldVehicleVehicleIdentificationYear = "Vehicle_VehicleIdentification_Year"
	// FieldVehicleID unque DIMO ID for the vehicle
	FieldVehicleID = "VehicleID"
)

type Dimo struct {
	// DefinitionID ID for the vehicles definition
	DefinitionID string `ch:"DefinitionID" json:"DefinitionID,omitempty"`
	// Source where the data was sourced from
	Source string `ch:"Source" json:"Source,omitempty"`
	// Subject subjet of this vehicle data
	Subject string `ch:"Subject" json:"Subject,omitempty"`
	// Timestamp timestamp of when this data was colllected
	Timestamp time.Time `ch:"Timestamp" json:"Timestamp,omitempty"`
	// Type type of data collected
	Type string `ch:"Type" json:"Type,omitempty"`
	// VehicleCurrentLocationAltitude Current altitude relative to WGS 84 reference ellipsoid, as measured at the position of GNSS receiver antenna.
	VehicleCurrentLocationAltitude float64 `ch:"Vehicle_CurrentLocation_Altitude" json:"Vehicle_CurrentLocation_Altitude,omitempty"`
	// VehicleCurrentLocationLatitude Current latitude of vehicle in WGS 84 geodetic coordinates, as measured at the position of GNSS receiver antenna.
	VehicleCurrentLocationLatitude float64 `ch:"Vehicle_CurrentLocation_Latitude" json:"Vehicle_CurrentLocation_Latitude,omitempty"`
	// VehicleCurrentLocationLongitude Current longitude of vehicle in WGS 84 geodetic coordinates, as measured at the position of GNSS receiver antenna.
	VehicleCurrentLocationLongitude float64 `ch:"Vehicle_CurrentLocation_Longitude" json:"Vehicle_CurrentLocation_Longitude,omitempty"`
	// VehicleCurrentLocationTimestamp Timestamp from GNSS system for current location, formatted according to ISO 8601 with UTC time zone.
	VehicleCurrentLocationTimestamp time.Time `ch:"Vehicle_CurrentLocation_Timestamp" json:"Vehicle_CurrentLocation_Timestamp,omitempty"`
	// VehicleLowVoltageBatteryCurrentVoltage Current Voltage of the low voltage battery.
	VehicleLowVoltageBatteryCurrentVoltage float32 `ch:"Vehicle_LowVoltageBattery_CurrentVoltage" json:"Vehicle_LowVoltageBattery_CurrentVoltage,omitempty"`
	// VehicleSpeed Vehicle speed.
	VehicleSpeed float32 `ch:"Vehicle_Speed" json:"Vehicle_Speed,omitempty"`
	// VehicleVehicleIdentificationBrand Vehicle brand or manufacturer.
	VehicleVehicleIdentificationBrand string `ch:"Vehicle_VehicleIdentification_Brand" json:"Vehicle_VehicleIdentification_Brand,omitempty"`
	// VehicleVehicleIdentificationModel Vehicle model.
	VehicleVehicleIdentificationModel string `ch:"Vehicle_VehicleIdentification_Model" json:"Vehicle_VehicleIdentification_Model,omitempty"`
	// VehicleVehicleIdentificationVIN 17-character Vehicle Identification Number (VIN) as defined by ISO 3779.
	VehicleVehicleIdentificationVIN string `ch:"Vehicle_VehicleIdentification_VIN" json:"Vehicle_VehicleIdentification_VIN,omitempty"`
	// VehicleVehicleIdentificationYear Model year of the vehicle.
	VehicleVehicleIdentificationYear uint16 `ch:"Vehicle_VehicleIdentification_Year" json:"Vehicle_VehicleIdentification_Year,omitempty"`
	// VehicleID unque DIMO ID for the vehicle
	VehicleID string `ch:"VehicleID" json:"VehicleID,omitempty"`
}