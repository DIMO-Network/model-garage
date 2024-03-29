-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS vehicle (
	Chassis_Axle_Row1_Wheel_Left_Tire_Pressure UInt16 COMMENT 'Tire pressure in kilo-Pascal.',
	Chassis_Axle_Row1_Wheel_Right_Tire_Pressure UInt16 COMMENT 'Tire pressure in kilo-Pascal.',
	Chassis_Axle_Row2_Wheel_Left_Tire_Pressure UInt16 COMMENT 'Tire pressure in kilo-Pascal.',
	Chassis_Axle_Row2_Wheel_Right_Tire_Pressure UInt16 COMMENT 'Tire pressure in kilo-Pascal.',
	CurrentLocation_Altitude Float64 COMMENT 'Current altitude relative to WGS 84 reference ellipsoid, as measured at the position of GNSS receiver antenna.',
	CurrentLocation_Latitude Float64 COMMENT 'Current latitude of vehicle in WGS 84 geodetic coordinates, as measured at the position of GNSS receiver antenna.',
	CurrentLocation_Longitude Float64 COMMENT 'Current longitude of vehicle in WGS 84 geodetic coordinates, as measured at the position of GNSS receiver antenna.',
	CurrentLocation_Timestamp DateTime COMMENT 'Timestamp from GNSS system for current location, formatted according to ISO 8601 with UTC time zone.',
	DIMO_DefinitionID String COMMENT 'ID for the vehicles definition',
	DIMO_Source String COMMENT 'where the data was sourced from',
	DIMO_Subject String COMMENT 'subjet of this vehicle data',
	DIMO_Timestamp DateTime COMMENT 'timestamp of when this data was colllected',
	DIMO_Type String COMMENT 'type of data collected',
	DIMO_VehicleID String COMMENT 'unque DIMO ID for the vehicle',
	Exterior_AirTemperature Float32 COMMENT 'Air temperature outside the vehicle.',
	LowVoltageBattery_CurrentVoltage Float32 COMMENT 'Current Voltage of the low voltage battery.',
	OBD_BarometricPressure Float32 COMMENT 'PID 33 - Barometric pressure',
	OBD_EngineLoad Float32 COMMENT 'PID 04 - Engine load in percent - 0 = no load, 100 = full load',
	OBD_IntakeTemp Float32 COMMENT 'PID 0F - Intake temperature',
	OBD_RunTime Float32 COMMENT 'PID 1F - Engine run time',
	Powertrain_CombustionEngine_ECT Int16 COMMENT 'Engine coolant temperature.',
	Powertrain_CombustionEngine_EngineOilLevel String COMMENT 'Engine oil level.',
	Powertrain_CombustionEngine_Speed UInt16 COMMENT 'Engine speed measured as rotations per minute.',
	Powertrain_CombustionEngine_TPS UInt8 COMMENT 'Current throttle position.',
	Powertrain_FuelSystem_AbsoluteLevel Float32 COMMENT 'Current available fuel in the fuel tank expressed in liters.',
	Powertrain_FuelSystem_SupportedFuelTypes Array(String) COMMENT 'High level information of fuel types supported',
	Powertrain_Range UInt32 COMMENT 'Remaining range in meters using all energy sources available in the vehicle.',
	Powertrain_TractionBattery_Charging_ChargeLimit UInt8 COMMENT 'Target charge limit (state of charge) for battery.',
	Powertrain_TractionBattery_Charging_IsCharging Bool COMMENT 'True if charging is ongoing. Charging is considered to be ongoing if energy is flowing from charger to vehicle.',
	Powertrain_TractionBattery_GrossCapacity UInt16 COMMENT 'Gross capacity of the battery.',
	Powertrain_TractionBattery_StateOfCharge_Current Float32 COMMENT 'Physical state of charge of the high voltage battery, relative to net capacity. This is not necessarily the state of charge being displayed to the customer.',
	Powertrain_Transmission_TravelledDistance Float32 COMMENT 'Odometer reading, total distance travelled during the lifetime of the transmission.',
	Powertrain_Type String COMMENT 'Defines the powertrain type of the vehicle.',
	Speed Float32 COMMENT 'Vehicle speed.',
	VehicleIdentification_Brand String COMMENT 'Vehicle brand or manufacturer.',
	VehicleIdentification_Model String COMMENT 'Vehicle model.',
	VehicleIdentification_VIN String COMMENT '17-character Vehicle Identification Number (VIN) as defined by ISO 3779.',
	VehicleIdentification_Year UInt16 COMMENT 'Model year of the vehicle.',
)
ENGINE = MergeTree()
ORDER BY (Vehicle_DIMO_Subject, Vehicle_DIMO_Timestamp)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
