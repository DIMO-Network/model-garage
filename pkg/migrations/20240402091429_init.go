// Package migrations provides the database migration for the DIMO clickhouse database.
package migrations

import (
	"context"
	"database/sql"
	"fmt"
	"runtime"

	"github.com/pressly/goose/v3"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	addFunc := func() { goose.AddNamedMigrationContext(filename, upInit, downInit) }
	addFuncs = append(addFuncs, addFunc)
	addFunc()
}

func upInit(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.ExecContext(ctx, createDIMOtableStatement)
	if err != nil {
		return fmt.Errorf("failed to create dimo table: %w", err)
	}
	return nil
}

func downInit(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.ExecContext(ctx, "DROP TABLE IF EXISTS dimo")
	if err != nil {
		return fmt.Errorf("failed to drop dimo table: %w", err)
	}
	return nil
}

const createDIMOtableStatement = `CREATE TABLE IF NOT EXISTS dimo (
	DefinitionID String COMMENT 'ID for the vehicles definition',
	Source String COMMENT 'where the data was sourced from',
	Subject String COMMENT 'subjet of this vehicle data',
	Timestamp DateTime('UTC') COMMENT 'timestamp of when this data was colllected',
	Type String COMMENT 'type of data collected',
	Vehicle_Chassis_Axle_Row1_Wheel_Left_Tire_Pressure UInt16 COMMENT 'Tire pressure in kilo-Pascal.',
	Vehicle_Chassis_Axle_Row1_Wheel_Right_Tire_Pressure UInt16 COMMENT 'Tire pressure in kilo-Pascal.',
	Vehicle_Chassis_Axle_Row2_Wheel_Left_Tire_Pressure UInt16 COMMENT 'Tire pressure in kilo-Pascal.',
	Vehicle_Chassis_Axle_Row2_Wheel_Right_Tire_Pressure UInt16 COMMENT 'Tire pressure in kilo-Pascal.',
	Vehicle_CurrentLocation_Altitude Float64 COMMENT 'Current altitude relative to WGS 84 reference ellipsoid, as measured at the position of GNSS receiver antenna.',
	Vehicle_CurrentLocation_Latitude Float64 COMMENT 'Current latitude of vehicle in WGS 84 geodetic coordinates, as measured at the position of GNSS receiver antenna.',
	Vehicle_CurrentLocation_Longitude Float64 COMMENT 'Current longitude of vehicle in WGS 84 geodetic coordinates, as measured at the position of GNSS receiver antenna.',
	Vehicle_CurrentLocation_Timestamp DateTime('UTC') COMMENT 'Timestamp from GNSS system for current location, formatted according to ISO 8601 with UTC time zone.',
	Vehicle_Exterior_AirTemperature Float32 COMMENT 'Air temperature outside the vehicle.',
	Vehicle_LowVoltageBattery_CurrentVoltage Float32 COMMENT 'Current Voltage of the low voltage battery.',
	Vehicle_OBD_BarometricPressure Float32 COMMENT 'PID 33 - Barometric pressure',
	Vehicle_OBD_EngineLoad Float32 COMMENT 'PID 04 - Engine load in percent - 0 = no load, 100 = full load',
	Vehicle_OBD_IntakeTemp Float32 COMMENT 'PID 0F - Intake temperature',
	Vehicle_OBD_RunTime Float32 COMMENT 'PID 1F - Engine run time',
	Vehicle_Powertrain_CombustionEngine_ECT Int16 COMMENT 'Engine coolant temperature.',
	Vehicle_Powertrain_CombustionEngine_EngineOilLevel String COMMENT 'Engine oil level.',
	Vehicle_Powertrain_CombustionEngine_Speed UInt16 COMMENT 'Engine speed measured as rotations per minute.',
	Vehicle_Powertrain_CombustionEngine_TPS UInt8 COMMENT 'Current throttle position.',
	Vehicle_Powertrain_FuelSystem_AbsoluteLevel Float32 COMMENT 'Current available fuel in the fuel tank expressed in liters.',
	Vehicle_Powertrain_FuelSystem_SupportedFuelTypes Array(String) COMMENT 'High level information of fuel types supported',
	Vehicle_Powertrain_Range UInt32 COMMENT 'Remaining range in meters using all energy sources available in the vehicle.',
	Vehicle_Powertrain_TractionBattery_Charging_ChargeLimit UInt8 COMMENT 'Target charge limit (state of charge) for battery.',
	Vehicle_Powertrain_TractionBattery_Charging_IsCharging Bool COMMENT 'True if charging is ongoing. Charging is considered to be ongoing if energy is flowing from charger to vehicle.',
	Vehicle_Powertrain_TractionBattery_GrossCapacity UInt16 COMMENT 'Gross capacity of the battery.',
	Vehicle_Powertrain_TractionBattery_StateOfCharge_Current Float32 COMMENT 'Physical state of charge of the high voltage battery, relative to net capacity. This is not necessarily the state of charge being displayed to the customer.',
	Vehicle_Powertrain_Transmission_TravelledDistance Float32 COMMENT 'Odometer reading, total distance travelled during the lifetime of the transmission.',
	Vehicle_Powertrain_Type String COMMENT 'Defines the powertrain type of the vehicle.',
	Vehicle_Speed Float32 COMMENT 'Vehicle speed.',
	Vehicle_VehicleIdentification_Brand String COMMENT 'Vehicle brand or manufacturer.',
	Vehicle_VehicleIdentification_Model String COMMENT 'Vehicle model.',
	Vehicle_VehicleIdentification_Year UInt16 COMMENT 'Model year of the vehicle.',
	VehicleID String COMMENT 'unque DIMO ID for the vehicle',
)
ENGINE = MergeTree()
ORDER BY (Subject, Timestamp)`