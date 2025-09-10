// Package vss holds the data structures and functions for working with events from DIMOs VSS schema.
package vss

import (
	"time"
)

const (
	// EventTableName is the name of the distributed table in Clickhouse.
	EventTableName = "event"
	// EventSubjectCol is the name of the subject column in Clickhouse.
	EventSubjectCol = "subject"
	// EventSourceCol is the name of the source column in Clickhouse.
	EventSourceCol = "source"
	// EventProducerCol is the name of the producer column in Clickhouse.
	EventProducerCol = "producer"
	// EventCloudEventIDCol is the name of the cloud_event_id column in Clickhouse.
	EventCloudEventIDCol = "cloud_event_id"
	// EventNameCol is the name of the name column in Clickhouse.
	EventNameCol = "name"
	// EventTimestampCol is the name of the timestamp column in Clickhouse.
	EventTimestampCol = "timestamp"
	// EventDurationNsCol is the name of the duration_ns column in Clickhouse.
	EventDurationNsCol = "duration_ns"
	// EventMetadataCol is the name of the metadata column in Clickhouse.
	EventMetadataCol = "metadata"
	// EventTagsCol is the name of the tags column in Clickhouse.
	EventTagsCol = "tags"
)

// Event represents a single event collected from a device.
// This is the data format that is stored in the database.
type Event struct {
	// Subject identifies the entity the event pertains to.
	Subject string `ch:"subject" json:"subject"`

	// Source is the entity that identified and submitted the event (oracle).
	Source string `ch:"source" json:"source"`

	// Producer is the specific origin of the data used to determine the event (device).
	Producer string `ch:"producer" json:"producer"`

	// CloudEventID is the identifier for the cloudevent.
	CloudEventID string `ch:"cloud_event_id" json:"cloudEventId"`

	// Name is the name of the event indicated by the oracle transmitting it.
	Name string `ch:"name" json:"name"`

	// Timestamp is the time at which the event described occurred, transmitted by oracle.
	Timestamp time.Time `ch:"timestamp" json:"timestamp"`

	// DurationNs is the duration in nanoseconds of the event.
	DurationNs uint64 `ch:"duration_ns" json:"durationNs"`

	// Metadata is arbitrary JSON metadata provided by the user, containing additional event-related information.
	Metadata string `ch:"metadata" json:"metadata"`

	// Tags is a list of tags associated with the event.
	Tags []string `ch:"tags" json:"tags"`
}

// EventToSlice converts an Event to an array of any for Clickhouse insertion.
// The order of the elements in the array is guaranteed to match the order of elements in the `EventColNames`.
func EventToSlice(obj Event) []any {
	return []any{
		obj.Subject,
		obj.Source,
		obj.Producer,
		obj.CloudEventID,
		obj.Name,
		obj.Timestamp,
		obj.DurationNs,
		obj.Metadata,
		obj.Tags,
	}
}

// EventColNames returns the column names of the Event struct.
func EventColNames() []string {
	return []string{
		EventSubjectCol,
		EventSourceCol,
		EventProducerCol,
		EventCloudEventIDCol,
		EventNameCol,
		EventTimestampCol,
		EventDurationNsCol,
		EventMetadataCol,
		EventTagsCol,
	}
}
