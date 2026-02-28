package telemetry

import (
	"testing"
	"time"

	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/stretchr/testify/assert"
	"github.com/teslamotors/fleet-telemetry/protos"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestConvertUntyped(t *testing.T) {
	teslaConnection := "0xc4035Fecb1cc906130423EF05f9C20977F643722" // This is the real value in dev and prod.
	subject := "did:erc721:137:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:7"

	ts, err := time.Parse(time.RFC3339, "2025-01-01T09:00:00Z")
	if err != nil {
		t.Fatal("Failed to create test timestamp.")
	}

	vin := "5YJYGDEF2LFR00942"

	pl := &protos.Payload{
		Data: []*protos.Datum{
			{Key: protos.Field_Location, Value: &protos.Value{Value: &protos.Value_LocationValue{LocationValue: &protos.LocationValue{Latitude: 30.267222, Longitude: -97.743056}}}},
			{Key: protos.Field_ACChargingPower, Value: &protos.Value{Value: &protos.Value_StringValue{StringValue: "5.700000084936619"}}},
			{Key: protos.Field_DCChargingEnergyIn, Value: &protos.Value{Value: &protos.Value_StringValue{StringValue: "2.380388924359452"}}},
			{Key: protos.Field_DetailedChargeState, Value: &protos.Value{Value: &protos.Value_DetailedChargeStateValue{DetailedChargeStateValue: protos.DetailedChargeStateValue_DetailedChargeStateCharging}}},
			{Key: protos.Field_Soc, Value: &protos.Value{Value: &protos.Value_StringValue{StringValue: "18.155283129013426"}}},
			{Key: protos.Field_TpmsPressureFl, Value: &protos.Value{Value: &protos.Value_StringValue{StringValue: "2.9250000435858965"}}},
			{Key: protos.Field_TpmsPressureFr, Value: &protos.Value{Value: &protos.Value_StringValue{StringValue: "2.425000036135316"}}},
			{Key: protos.Field_TpmsPressureRl, Value: &protos.Value{Value: &protos.Value_StringValue{StringValue: "2.8000000417232513"}}},
			{Key: protos.Field_TpmsPressureRr, Value: &protos.Value{Value: &protos.Value_StringValue{StringValue: "2.8000000417232513"}}},
			{Key: protos.Field_OutsideTemp, Value: &protos.Value{Value: &protos.Value_StringValue{StringValue: "2.5"}}},
			{Key: protos.Field_EstBatteryRange, Value: &protos.Value{Value: &protos.Value_StringValue{StringValue: "19.80471193262205"}}},
			{Key: protos.Field_ChargeLimitSoc, Value: &protos.Value{Value: &protos.Value_StringValue{StringValue: "80"}}},
			{Key: protos.Field_Odometer, Value: &protos.Value{Value: &protos.Value_StringValue{StringValue: "61026.774055062444"}}},
			{Key: protos.Field_VehicleSpeed, Value: &protos.Value{Value: &protos.Value_StringValue{StringValue: "21"}}},
			{Key: protos.Field_EnergyRemaining, Value: &protos.Value{Value: &protos.Value_StringValue{StringValue: "39.61999911442399"}}},
			{Key: protos.Field_DoorState, Value: &protos.Value{Value: &protos.Value_StringValue{StringValue: "DriverFront|PassengerFront"}}},
			{Key: protos.Field_FdWindow, Value: &protos.Value{Value: &protos.Value_StringValue{StringValue: "Opened"}}},
			{Key: protos.Field_FpWindow, Value: &protos.Value{Value: &protos.Value_StringValue{StringValue: "Opened"}}},
			{Key: protos.Field_RdWindow, Value: &protos.Value{Value: &protos.Value_StringValue{StringValue: "Closed"}}},
			{Key: protos.Field_RpWindow, Value: &protos.Value{Value: &protos.Value_StringValue{StringValue: "PartiallyOpen"}}},
			{Key: protos.Field_ChargerVoltage, Value: &protos.Value{Value: &protos.Value_DoubleValue{DoubleValue: 114.774}}}, // Newer fields are always typed.
			{Key: protos.Field_ChargeAmps, Value: &protos.Value{Value: &protos.Value_StringValue{StringValue: "12.0"}}},
		},
		CreatedAt: timestamppb.New(ts),
		Vin:       vin,
	}

	signals, errors := ProcessPayload(pl, subject, teslaConnection)
	if len(errors) != 0 {
		t.Fatalf("Unexpected errors from conversion: %v", errors)
	}

	expectedSignals := []vss.Signal{
		{Subject: subject, Timestamp: ts, Name: "currentLocationCoordinates", ValueLocation: vss.Location{Latitude: 30.267222, Longitude: -97.743056}, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "powertrainTractionBatteryCurrentPower", ValueNumber: 5700.000084936619, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "powertrainTractionBatteryChargingIsCharging", ValueNumber: 1, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "powertrainTractionBatteryChargingIsChargingCableConnected", ValueNumber: 1, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "powertrainTractionBatteryChargingAddedEnergy", ValueNumber: 2.380388924359452, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "powertrainTractionBatteryStateOfChargeCurrent", ValueNumber: 18.155283129013426, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "chassisAxleRow1WheelLeftTirePressure", ValueNumber: 296.375629416341, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "chassisAxleRow1WheelRightTirePressure", ValueNumber: 245.71312866141088, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "chassisAxleRow2WheelLeftTirePressure", ValueNumber: 283.71000422760847, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "chassisAxleRow2WheelRightTirePressure", ValueNumber: 283.71000422760847, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "exteriorAirTemperature", ValueNumber: 2.5, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "powertrainRange", ValueNumber: 31.872594320493704, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "powertrainTractionBatteryChargingChargeLimit", ValueNumber: 80, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "speed", ValueNumber: 33.796224, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "powertrainTransmissionTravelledDistance", ValueNumber: 98213.07266487041, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "powertrainTractionBatteryStateOfChargeCurrentEnergy", ValueNumber: 39.61999911442399, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "cabinDoorRow1DriverSideIsOpen", ValueNumber: 1, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "cabinDoorRow1PassengerSideIsOpen", ValueNumber: 1, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "cabinDoorRow2DriverSideIsOpen", ValueNumber: 0, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "cabinDoorRow2PassengerSideIsOpen", ValueNumber: 0, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "cabinDoorRow1DriverSideWindowIsOpen", ValueNumber: 1, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "cabinDoorRow1PassengerSideWindowIsOpen", ValueNumber: 1, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "cabinDoorRow2DriverSideWindowIsOpen", ValueNumber: 0, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "cabinDoorRow2PassengerSideWindowIsOpen", ValueNumber: 1, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "powertrainTractionBatteryChargingChargeVoltageUnknownType", ValueNumber: 114.774, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "powertrainTractionBatteryChargingChargeCurrentAC", ValueNumber: 12.0, Source: teslaConnection},
	}

	assert.ElementsMatch(t, expectedSignals, signals)
}

func TestConvertTyped(t *testing.T) {
	teslaConnection := "0xc4035Fecb1cc906130423EF05f9C20977F643722" // This is the real value in dev and prod.
	subject := "did:erc721:137:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:7"

	ts, err := time.Parse(time.RFC3339, "2025-01-01T09:00:00Z")
	if err != nil {
		t.Fatal("Failed to create test timestamp.")
	}

	vin := "5YJYGDEF2LFR00942"

	pl := &protos.Payload{
		Data: []*protos.Datum{
			{Key: protos.Field_Location, Value: &protos.Value{Value: &protos.Value_LocationValue{LocationValue: &protos.LocationValue{Latitude: 30.267222, Longitude: -97.743056}}}},
			{Key: protos.Field_ACChargingPower, Value: &protos.Value{Value: &protos.Value_DoubleValue{DoubleValue: 5.700000084936619}}},
			{Key: protos.Field_DCChargingEnergyIn, Value: &protos.Value{Value: &protos.Value_DoubleValue{DoubleValue: 2.380388924359452}}},
			{Key: protos.Field_DetailedChargeState, Value: &protos.Value{Value: &protos.Value_DetailedChargeStateValue{DetailedChargeStateValue: protos.DetailedChargeStateValue_DetailedChargeStateCharging}}},
			{Key: protos.Field_Soc, Value: &protos.Value{Value: &protos.Value_DoubleValue{DoubleValue: 18.155283129013426}}},
			{Key: protos.Field_TpmsPressureFl, Value: &protos.Value{Value: &protos.Value_DoubleValue{DoubleValue: 2.9250000435858965}}},
			{Key: protos.Field_TpmsPressureFr, Value: &protos.Value{Value: &protos.Value_DoubleValue{DoubleValue: 2.425000036135316}}},
			{Key: protos.Field_TpmsPressureRl, Value: &protos.Value{Value: &protos.Value_DoubleValue{DoubleValue: 2.8000000417232513}}},
			{Key: protos.Field_TpmsPressureRr, Value: &protos.Value{Value: &protos.Value_DoubleValue{DoubleValue: 2.8000000417232513}}},
			{Key: protos.Field_OutsideTemp, Value: &protos.Value{Value: &protos.Value_DoubleValue{DoubleValue: 2.5}}},
			{Key: protos.Field_EstBatteryRange, Value: &protos.Value{Value: &protos.Value_DoubleValue{DoubleValue: 19.80471193262205}}},
			{Key: protos.Field_ChargeLimitSoc, Value: &protos.Value{Value: &protos.Value_IntValue{IntValue: 80}}},
			{Key: protos.Field_Odometer, Value: &protos.Value{Value: &protos.Value_DoubleValue{DoubleValue: 61026.774055062444}}},
			{Key: protos.Field_VehicleSpeed, Value: &protos.Value{Value: &protos.Value_DoubleValue{DoubleValue: 21}}},
			{Key: protos.Field_EnergyRemaining, Value: &protos.Value{Value: &protos.Value_DoubleValue{DoubleValue: 39.61999911442399}}},
			{Key: protos.Field_DoorState, Value: &protos.Value{Value: &protos.Value_DoorValue{DoorValue: &protos.Doors{DriverFront: true, PassengerFront: true}}}},
			{Key: protos.Field_FdWindow, Value: &protos.Value{Value: &protos.Value_WindowStateValue{WindowStateValue: protos.WindowState_WindowStateOpened}}},
			{Key: protos.Field_FpWindow, Value: &protos.Value{Value: &protos.Value_WindowStateValue{WindowStateValue: protos.WindowState_WindowStateOpened}}},
			{Key: protos.Field_RdWindow, Value: &protos.Value{Value: &protos.Value_WindowStateValue{WindowStateValue: protos.WindowState_WindowStateClosed}}},
			{Key: protos.Field_RpWindow, Value: &protos.Value{Value: &protos.Value_WindowStateValue{WindowStateValue: protos.WindowState_WindowStatePartiallyOpen}}},
			{Key: protos.Field_ChargerVoltage, Value: &protos.Value{Value: &protos.Value_DoubleValue{DoubleValue: 114.774}}},
			{Key: protos.Field_ChargeAmps, Value: &protos.Value{Value: &protos.Value_DoubleValue{DoubleValue: 12.0}}},
		},
		CreatedAt: timestamppb.New(ts),
		Vin:       vin,
	}

	signals, errors := ProcessPayload(pl, subject, teslaConnection)
	if len(errors) != 0 {
		t.Fatalf("Unexpected errors from conversion: %v", errors)
	}

	expectedSignals := []vss.Signal{
		{Subject: subject, Timestamp: ts, Name: "currentLocationCoordinates", ValueLocation: vss.Location{Latitude: 30.267222, Longitude: -97.743056}, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "powertrainTractionBatteryCurrentPower", ValueNumber: 5700.000084936619, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "powertrainTractionBatteryChargingIsCharging", ValueNumber: 1, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "powertrainTractionBatteryChargingIsChargingCableConnected", ValueNumber: 1, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "powertrainTractionBatteryChargingAddedEnergy", ValueNumber: 2.380388924359452, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "powertrainTractionBatteryStateOfChargeCurrent", ValueNumber: 18.155283129013426, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "chassisAxleRow1WheelLeftTirePressure", ValueNumber: 296.375629416341, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "chassisAxleRow1WheelRightTirePressure", ValueNumber: 245.71312866141088, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "chassisAxleRow2WheelLeftTirePressure", ValueNumber: 283.71000422760847, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "chassisAxleRow2WheelRightTirePressure", ValueNumber: 283.71000422760847, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "exteriorAirTemperature", ValueNumber: 2.5, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "powertrainRange", ValueNumber: 31.872594320493704, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "powertrainTractionBatteryChargingChargeLimit", ValueNumber: 80, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "speed", ValueNumber: 33.796224, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "powertrainTransmissionTravelledDistance", ValueNumber: 98213.07266487041, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "powertrainTractionBatteryStateOfChargeCurrentEnergy", ValueNumber: 39.61999911442399, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "cabinDoorRow1DriverSideIsOpen", ValueNumber: 1, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "cabinDoorRow1PassengerSideIsOpen", ValueNumber: 1, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "cabinDoorRow2DriverSideIsOpen", ValueNumber: 0, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "cabinDoorRow2PassengerSideIsOpen", ValueNumber: 0, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "cabinDoorRow1DriverSideWindowIsOpen", ValueNumber: 1, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "cabinDoorRow1PassengerSideWindowIsOpen", ValueNumber: 1, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "cabinDoorRow2DriverSideWindowIsOpen", ValueNumber: 0, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "cabinDoorRow2PassengerSideWindowIsOpen", ValueNumber: 1, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "powertrainTractionBatteryChargingChargeVoltageUnknownType", ValueNumber: 114.774, Source: teslaConnection},
		{Subject: subject, Timestamp: ts, Name: "powertrainTractionBatteryChargingChargeCurrentAC", ValueNumber: 12.0, Source: teslaConnection},
	}

	assert.ElementsMatch(t, expectedSignals, signals)
}

func TestIgnoreInvalids(t *testing.T) {
	teslaConnection := "0xc4035Fecb1cc906130423EF05f9C20977F643722" // This is the real value in dev and prod.
	subject := "did:erc721:137:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:7"

	ts, err := time.Parse(time.RFC3339, "2025-01-01T09:00:00Z")
	if err != nil {
		t.Fatal("Failed to create test timestamp.")
	}

	vin := "5YJYGDEF2LFR00942"

	pl := &protos.Payload{
		Data: []*protos.Datum{
			{Key: protos.Field_Odometer, Value: &protos.Value{Value: &protos.Value_DoubleValue{DoubleValue: 61026.774055062444}}},
			{Key: protos.Field_VehicleSpeed, Value: &protos.Value{Value: &protos.Value_Invalid{Invalid: true}}},
		},
		CreatedAt: timestamppb.New(ts),
		Vin:       vin,
	}

	signals, errors := ProcessPayload(pl, subject, teslaConnection)
	if len(errors) != 0 {
		t.Fatalf("Unexpected errors from conversion: %v", err)
	}

	expectedSignals := []vss.Signal{
		{Subject: subject, Timestamp: ts, Name: "powertrainTransmissionTravelledDistance", ValueNumber: 98213.07266487041, Source: teslaConnection},
	}

	assert.ElementsMatch(t, expectedSignals, signals)
}

func TestConversionErrors(t *testing.T) {
	teslaConnection := "0xc4035Fecb1cc906130423EF05f9C20977F643722" // This is the real value in dev and prod.
	subject := "did:erc721:137:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:7"

	ts, err := time.Parse(time.RFC3339, "2025-01-01T09:00:00Z")
	if err != nil {
		t.Fatal("Failed to create test timestamp.")
	}

	vin := "5YJYGDEF2LFR00942"

	pl := &protos.Payload{
		Data: []*protos.Datum{
			{Key: protos.Field_Odometer, Value: &protos.Value{Value: &protos.Value_BooleanValue{BooleanValue: true}}},
		},
		CreatedAt: timestamppb.New(ts),
		Vin:       vin,
	}

	signals, errors := ProcessPayload(pl, subject, teslaConnection)
	if len(signals) != 0 {
		t.Errorf("Should not have gotten any signals.")
	}
	if len(errors) != 1 {
		t.Fatalf("Should have gotten exactly one error.")
	}
}
