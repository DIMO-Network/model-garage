// Package cloudevent provides types for working with CloudEvents.
package cloudevent

import (
	"encoding/json"
	"reflect"
	"strings"
	"time"

	"github.com/tidwall/sjson"
)

const (
	// TypeStatus is the event type for status updates.
	TypeStatus = "dimo.status"

	// TypeFingerprint is the event type for fingerprint updates.
	TypeFingerprint = "dimo.fingerprint"

	// TypeVerifableCredential is the event type for verifiable credentials.
	TypeVerifableCredential = "dimo.verifiablecredential" //nolint:gosec // This is not a credential.

	// TypeUnknown is the event type for unknown events.
	TypeUnknown = "dimo.unknown"

	// SpecVersion is the version of the CloudEvents spec.
	SpecVersion = "1.0"
)

// RawEvent is a cloudevent with a json.RawMessage data field.
type RawEvent = CloudEvent[json.RawMessage]

var definedCloudeEventHdrFields = getJSONFieldNames(reflect.TypeOf(CloudEventHeader{}))

// CloudEvent represents an event according to the CloudEvents spec.
// To Add extra headers to the CloudEvent, add them to the Extras map.
// See https://github.com/cloudevents/spec/blob/v1.0.2/cloudevents/spec.md
type CloudEvent[A any] struct {
	CloudEventHeader
	// Data contains domain-specific information about the event.
	Data A `json:"data"`
}

// UnmarshalJSON implements custom JSON unmarshaling for CloudEvent.
func (c *CloudEvent[A]) UnmarshalJSON(data []byte) error {
	var err error
	c.CloudEventHeader, err = unmarshalCloudEvent(data, c.setDataField)
	return err
}

// MarshalJSON implements custom JSON marshaling for CloudEventHeader
func (c CloudEvent[A]) MarshalJSON() ([]byte, error) {
	// Marshal the base struct
	data, err := json.Marshal(c.CloudEventHeader)
	if err != nil {
		return nil, err
	}
	data, err = sjson.SetBytes(data, "data", c.Data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// CloudEventHeader contains the metadata for any CloudEvent.
// To add extra headers to the CloudEvent, add them to the Extras map.
type CloudEventHeader struct {
	// ID is an identifier for the event. The combination of ID and Source must
	// be unique.
	ID string `json:"id"`

	// Source is the context in which the event happened. In a distributed system it might consist of multiple Producers.
	Source string `json:"source"`

	// Producer is a specific instance, process or device that creates the data structure describing the CloudEvent.
	Producer string `json:"producer"`

	// SpecVersion is the version of CloudEvents specification used.
	// This is always hardcoded "1.0".
	SpecVersion string `json:"specversion"`

	// Subject is an optional field identifying the subject of the event within
	// the context of the event producer. In practice, we always set this.
	Subject string `json:"subject"`

	// Time is an optional field giving the time at which the event occurred. In
	// practice, we always set this.
	Time time.Time `json:"time"`

	// Type describes the type of event. It should generally be a reverse-DNS
	// name.
	Type string `json:"type"`

	// DataContentType is an optional MIME type for the data field. We almost
	// always serialize to JSON and in that case this field is implicitly
	// "application/json".
	DataContentType string `json:"datacontenttype,omitempty"`

	// DataSchema is an optional URI pointing to a schema for the data field.
	DataSchema string `json:"dataschema,omitempty"`

	// DataVersion is the version of the data type.
	DataVersion string `json:"dataversion,omitempty"`

	// Extras contains any additional fields that are not part of the CloudEvent excluding the data field.
	Extras map[string]any `json:"-"`
}

// Equals returns true if the two CloudEventHeaders share the same ID and Source.
func (c *CloudEventHeader) Equals(other CloudEventHeader) bool {
	return c.ID == other.ID && c.Source == other.Source
}

type cloudEventHeader CloudEventHeader

func ignoreDataField(json.RawMessage) error { return nil }

func (c *CloudEvent[A]) setDataField(data json.RawMessage) error {
	return json.Unmarshal(data, &c.Data)
}

// UnmarshalJSON implements custom JSON unmarshaling for CloudEventHeader.
func (c *CloudEventHeader) UnmarshalJSON(data []byte) error {
	var err error
	*c, err = unmarshalCloudEvent(data, ignoreDataField)
	return err
}

// MarshalJSON implements custom JSON marshaling for CloudEventHeader
func (c CloudEventHeader) MarshalJSON() ([]byte, error) {
	// Marshal the base struct
	aux := (cloudEventHeader)(c)
	aux.SpecVersion = SpecVersion
	data, err := json.Marshal(aux)
	if err != nil {
		return nil, err
	}
	// Add all extras using sjson]
	for k, v := range c.Extras {
		data, err = sjson.SetBytes(data, k, v)
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

func getJSONFieldNames(t reflect.Type) map[string]struct{} {
	fields := map[string]struct{}{}

	for i := range t.NumField() {
		field := t.Field(i)

		tag := field.Tag.Get("json")
		if tag == "" {
			continue
		}

		name := tag
		if comma := strings.Index(tag, ","); comma != -1 {
			name = tag[:comma]
		}

		if name == "-" {
			continue
		}

		fields[name] = struct{}{}
	}

	return fields
}

func unmarshalCloudEvent(data []byte, dataFunc func(json.RawMessage) error) (CloudEventHeader, error) {
	c := CloudEventHeader{}
	aux := cloudEventHeader{}
	// Unmarshal known fields directly into the struct
	if err := json.Unmarshal(data, &aux); err != nil {
		return c, err
	}
	aux.SpecVersion = SpecVersion
	c = (CloudEventHeader)(aux)
	// Create a map to hold all JSON fields
	rawFields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &rawFields); err != nil {
		return c, err
	}

	// Separate known and unknown fields
	for key, rawValue := range rawFields {
		if _, ok := definedCloudeEventHdrFields[key]; ok {
			// Skip defined fields
			continue
		}
		if key == "data" {
			if err := dataFunc(rawValue); err != nil {
				return c, err
			}
			continue
		}
		if c.Extras == nil {
			c.Extras = make(map[string]any)
		}
		var value any
		if err := json.Unmarshal(rawValue, &value); err != nil {
			return c, err
		}
		c.Extras[key] = value
	}
	return c, nil
}
