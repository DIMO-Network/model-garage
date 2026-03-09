// Package vss holds the data structures and functions for working with signals from DIMOs VSS schema.
package vss

import (
	"fmt"
	"time"

	"github.com/DIMO-Network/cloudevent"
)

const (
	// TableName is the name of the distributed table in Clickhouse.
	TableName = "signal"
	// SubjectCol is the name of the subject column in ClickHouse.
	SubjectCol = "subject"
	// TimestampCol is the name of the timestamp column in Clickhouse.
	TimestampCol = "timestamp"
	// SourceCol is the name of the source column in Clickhouse.
	SourceCol = "source"
	// NameCol is the name of the name column in Clickhouse.
	NameCol = "name"
	// ProducerCol is the name of the producer column in Clickhouse.
	ProducerCol = "producer"
	// CloudEventIDCol is the name of the cloud_event_id column in Clickhouse.
	CloudEventIDCol = "cloud_event_id"
	// ValueNumberCol is the name of the value_number column in Clickhouse.
	ValueNumberCol = "value_number"
	// ValueStringCol is the name of the value_string column in Clickhouse.
	ValueStringCol = "value_string"
	// ValueLocationCol is the name of the value_location column in Clickhouse.
	ValueLocationCol = "value_location"
)

// SignalData holds the per-signal data fields.
type SignalData struct {
	// Timestamp is when this data was collected.
	Timestamp time.Time `ch:"timestamp" json:"timestamp"`
	// Name is the name of the signal collected.
	Name string `ch:"name" json:"name"`
	// ValueNumber is the value of the signal collected.
	ValueNumber float64 `ch:"value_number" json:"valueNumber"`
	// ValueString is the value of the signal collected.
	ValueString string `ch:"value_string" json:"valueString"`
	// ValueLocation is the value of the signal collected.
	ValueLocation Location `ch:"value_location" json:"valueLocation"`
	// CloudEventID is the ID of the source CloudEvent that this signal was extracted from.
	// This is persisted to the database and is distinct from Signal.ID, which is the
	// signal's own CloudEvent identity (not persisted).
	CloudEventID string `ch:"cloud_event_id" json:"cloudEventId"`
}

// SetValue dynamically set the appropriate value field based on the type of the value.
func (s *SignalData) SetValue(val any) {
	switch typedVal := val.(type) {
	case float64:
		s.ValueNumber = typedVal
	case string:
		s.ValueString = typedVal
	case Location:
		s.ValueLocation = typedVal
	default:
		s.ValueString = fmt.Sprintf("%v", val)
	}
}

// Signal represents a single signal collected from a device.
// This is a CloudEvent with SignalData as the payload.
type Signal = cloudevent.CloudEvent[SignalData]

// Location represents a point on the earth in WSG-84 coordinates,
// optionally with a Horizontal Dilution of Position (HDOP) value or a
// heading value.
type Location struct {
	Latitude  float64 `ch:"latitude" json:"latitude"`
	Longitude float64 `ch:"longitude" json:"longitude"`
	HDOP      float64 `ch:"hdop" json:"hdop"`
	Heading   float64 `ch:"heading" json:"heading"`
}

// SignalToSlice converts a Signal to an array of any for Clickhouse insertion.
// The order of the elements in the array is guaranteed to match the order of elements in the `SignalColNames`.
func SignalToSlice(obj Signal) []any {
	return []any{
		obj.Subject,
		obj.Data.Timestamp,
		obj.Data.Name,
		obj.Source,
		obj.Producer,
		obj.Data.CloudEventID,
		obj.Data.ValueNumber,
		obj.Data.ValueString,
		obj.Data.ValueLocation,
	}
}

// SignalColNames returns the column names of the Signal struct.
func SignalColNames() []string {
	return []string{
		SubjectCol,
		TimestampCol,
		NameCol,
		SourceCol,
		ProducerCol,
		CloudEventIDCol,
		ValueNumberCol,
		ValueStringCol,
		ValueLocationCol,
	}
}

// SignalsPayload is the data payload for a signals CloudEvent on the wire.
type SignalsPayload struct {
	Signals []SignalData `json:"signals"`
}

// SignalCloudEvent is a CloudEvent carrying multiple signals on the wire.
type SignalCloudEvent = cloudevent.CloudEvent[SignalsPayload]

// PackSignals wraps extracted signals into a single CloudEvent for wire transport.
// Only data fields (Timestamp, Name, Value*) are preserved from each Signal.
// Header fields on individual Signal structs are ignored — use the header param.
func PackSignals(header cloudevent.CloudEventHeader, signals []Signal) SignalCloudEvent {
	payload := SignalsPayload{Signals: make([]SignalData, 0, len(signals))}
	for _, s := range signals {
		payload.Signals = append(payload.Signals, s.Data)
	}
	return SignalCloudEvent{CloudEventHeader: header, Data: payload}
}

// UnpackSignals extracts individual signals from a wire CloudEvent.
// Each unpacked Signal gets header fields from the envelope.
func UnpackSignals(msg SignalCloudEvent) []Signal {
	signals := make([]Signal, len(msg.Data.Signals))
	for i, sd := range msg.Data.Signals {
		hdr := msg.CloudEventHeader
		hdr.Type = cloudevent.TypeSignal
		signals[i] = Signal{CloudEventHeader: hdr, Data: sd}
	}
	return signals
}
