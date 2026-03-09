// Package vss holds the data structures and functions for working with events from DIMOs VSS schema.
package vss

import (
	"time"

	"github.com/DIMO-Network/cloudevent"
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
	// EventTypeCol is the name of the type column in Clickhouse.
	EventTypeCol = "type"
	// EventDataVersionCol is the name of the data_version column in Clickhouse.
	EventDataVersionCol = "data_version"
)

// EventData holds the domain-specific payload for an event.
type EventData struct {
	// Name is the name of the event indicated by the oracle transmitting it.
	Name string `ch:"name" json:"name"`
	// Timestamp is the time at which the event described occurred.
	Timestamp time.Time `ch:"timestamp" json:"timestamp"`
	// DurationNs is the duration in nanoseconds of the event.
	DurationNs uint64 `ch:"duration_ns" json:"durationNs"`
	// Metadata is arbitrary JSON metadata.
	Metadata string `ch:"metadata" json:"metadata"`
	// CloudEventID is the ID of the source CloudEvent that this event was extracted from.
	// This is persisted to the database and is distinct from Event.ID, which is the
	// event's own CloudEvent identity (not persisted).
	CloudEventID string `ch:"cloud_event_id" json:"cloudEventId"`
	// Tags are filterable labels for the event, preserved through pack/unpack.
	// An empty slice is used instead of nil to ensure consistent JSON serialization.
	Tags []string `ch:"tags" json:"tags"`
}

// Event represents a single event collected from a device.
// This is a CloudEvent with EventData as the payload.
type Event = cloudevent.CloudEvent[EventData]

// EventToSlice converts an Event to an array of any for Clickhouse insertion.
// The order of the elements in the array is guaranteed to match the order of elements in the `EventColNames`.
func EventToSlice(obj Event) []any {
	return []any{
		obj.Subject,
		obj.Source,
		obj.Producer,
		obj.Data.CloudEventID,
		obj.Type,
		obj.DataVersion,
		obj.Data.Name,
		obj.Data.Timestamp,
		obj.Data.DurationNs,
		obj.Data.Metadata,
		obj.Data.Tags,
	}
}

// EventColNames returns the column names of the Event struct.
func EventColNames() []string {
	return []string{
		EventSubjectCol,
		EventSourceCol,
		EventProducerCol,
		EventCloudEventIDCol,
		EventTypeCol,
		EventDataVersionCol,
		EventNameCol,
		EventTimestampCol,
		EventDurationNsCol,
		EventMetadataCol,
		EventTagsCol,
	}
}

// EventsPayload is the data payload for an events CloudEvent on the wire.
type EventsPayload struct {
	Events []EventData `json:"events"`
}

// EventCloudEvent is a CloudEvent carrying multiple events on the wire.
type EventCloudEvent = cloudevent.CloudEvent[EventsPayload]

// PackEvents wraps extracted events into a single CloudEvent for wire transport.
// Only EventData fields are preserved; per-event header fields are discarded
// in favor of the shared header.
func PackEvents(header cloudevent.CloudEventHeader, events []Event) EventCloudEvent {
	payload := EventsPayload{Events: make([]EventData, 0, len(events))}
	for _, e := range events {
		payload.Events = append(payload.Events, e.Data)
	}
	return EventCloudEvent{CloudEventHeader: header, Data: payload}
}

// UnpackEvents extracts individual events from a wire CloudEvent.
// Each unpacked event inherits the envelope's header fields.
func UnpackEvents(msg EventCloudEvent) []Event {
	events := make([]Event, len(msg.Data.Events))
	for i, ed := range msg.Data.Events {
		hdr := msg.CloudEventHeader
		hdr.Type = cloudevent.TypeEvent
		events[i] = Event{CloudEventHeader: hdr, Data: ed}
	}
	return events
}
