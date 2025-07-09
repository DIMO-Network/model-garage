// Package occurrences holds the data structures and functions for working with event data submitted by DIMO oracles.
package occurrences

import (
	"time"
)

const (
	// TableName is the name of the distributed table in Clickhouse.
	TableName = "event"
	// EventIDCol is the name of the unique event id column in Clickhouse.
	EventIDCol = "id"
	// CloudEventIDCol is the name of the cloud_event_id column in Clickhouse.
	CloudEventIDCol = "cloud_event_id"
	// Subject is the name of the timestamp column in Clickhouse.
	Subject = "subject"
	// SourceCol is the name of the source column in Clickhouse.
	SourceCol = "source"
	// ProducerCol is the name of the producer column in Clickhouse.
	ProducerCol = "producer"
	// EventNameCol is the name of the event_name column in Clickhouse.
	EventNameCol = "event_name"
	// EventTimeCol is the name of the event_time column in Clickhouse.
	EventTimeCol = "event_time"
	// EventDurationCol is the name of the event_duration column in Clickhouse.
	EventDurationCol = "event_duration"
	// EventMetaDataCol is the name of the event_metadata column in Clickhouse.
	EventMetaDataCol = "event_metadata"
)

// Event represents a single event submitted by an oracle with device data.
// This is the data format that is stored in the database.
type Event struct {
	EventID       string        `ch:"id" json:"id"`
	CloudEventID  string        `ch:"cloud_event_id" json:"cloud_event_id"`
	Subject       string        `ch:"subject" json:"subject"`
	Source        string        `ch:"source" json:"source"`
	Producer      string        `ch:"producer" json:"producer"`
	EventName     string        `ch:"event_name" json:"event_name"`
	EventTime     time.Time     `ch:"event_time" json:"event_time"`
	EventDuration time.Duration `ch:"event_duration" json:"event_duration"`
	EventMetaData string        `ch:"event_metadata" json:"event_metadata"`
}

// EventToSlice converts an Event to an array of any for Clickhouse insertion.
// The order of the elements in the array is guaranteed to match the order of elements in the `EventColNames`.
func EventToSlice(obj Event) []any {
	return []any{
		obj.EventID,
		obj.CloudEventID,
		obj.Subject,
		obj.Source,
		obj.Producer,
		obj.EventName,
		obj.EventTime,
		obj.EventDuration,
		obj.EventMetaData,
	}
}

// EventColNames returns the column names of the Event struct.
func EventColNames() []string {
	return []string{
		EventIDCol,
		CloudEventIDCol,
		Subject,
		SourceCol,
		ProducerCol,
		EventNameCol,
		EventTimeCol,
		EventDurationCol,
		EventMetaDataCol,
	}
}
