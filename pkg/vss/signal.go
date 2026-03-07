// Package vss holds the data structures and functions for working with signals from DIMOs VSS schema.
package vss

import (
	"fmt"
	"time"
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

// Signal represents a single signal collected from a device.
// This is the data format that is stored in the database.
type Signal struct {
	// Subject is the subject of the signal. Typically a W3C DID.
	Subject string `ch:"subject" json:"subject"`

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

	// Source is the source of the signal collected.
	Source string `ch:"source" json:"source"`

	// Producer is the producer of the signal collected.
	Producer string `ch:"producer" json:"producer"`

	// CloudEventID is the ID of the CloudEvent that this signal was extracted from.
	CloudEventID string `ch:"cloud_event_id" json:"cloudEventId"`
}

// Location represents a point on the earth in WSG-84 coordinates,
// optionally with a Horizontal Dilution of Position (HDOP) value or a
// heading value.
type Location struct {
	Latitude  float64 `ch:"latitude"`
	Longitude float64 `ch:"longitude"`
	HDOP      float64 `ch:"hdop"`
	Heading   float64 `ch:"heading"`
}

// SetValue dynamically set the appropriate value field based on the type of the value.
func (s *Signal) SetValue(val any) {
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

// SignalToSlice converts a Signal to an array of any for Clickhouse insertion.
// The order of the elements in the array is guaranteed to match the order of elements in the `SignalColNames`.
func SignalToSlice(obj Signal) []any {
	return []any{
		obj.Subject,
		obj.Timestamp,
		obj.Name,
		obj.Source,
		obj.Producer,
		obj.CloudEventID,
		obj.ValueNumber,
		obj.ValueString,
		obj.ValueLocation,
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
