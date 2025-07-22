// Package defaultmodule provides a default implementation for decoding DIMO data.
package defaultmodule

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/schema"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/tidwall/gjson"
)

// SignalData is a struct for holding vss signal data.
type SignalData struct {
	Signals []*Signal `json:"signals"`
}

// Signal is a struct for holding vss signal data.
type Signal struct {
	// Timestamp is when this data was collected. (format: RFC3339)
	Timestamp time.Time `json:"timestamp"`
	// Name is the name of the signal collected.
	Name string `json:"name"`
	// Value is the value of the signal collected. If the signal base type is a number it will be converted to a float64.
	Value any `json:"value"`
}

// EventsData is a struct for holding a list of events.
type EventsData struct {
	Events []Event `json:"events"`
}

// Event is a struct for holding a single event.
type Event struct {
	// Name is the name of the event.
	Name string `json:"name"`
	// Timestamp is when this event occurred. (format: RFC3339)
	Timestamp time.Time `json:"timestamp"`
	// Duration is the duration of the event in nanoseconds.
	DurationNs uint64 `json:"durationNs,omitempty"`
	// Metadata is the metadata of the event.
	Metadata string `json:"metadata,omitempty"`
}

// Module holds dependencies for the default module. At present, there are none.
type Module struct {
	once      sync.Once
	signalMap map[string]*schema.SignalInfo
	loadErr   error
}

// LoadSignalMap loads the default signal map.
func LoadSignalMap() (map[string]*schema.SignalInfo, error) {
	defs, err := schema.LoadDefinitionFile(strings.NewReader(schema.DefaultDefinitionsYAML()))
	if err != nil {
		return nil, fmt.Errorf("failed to load default schema definitions: %w", err)
	}
	signalInfo, err := schema.LoadSignalsCSV(strings.NewReader(schema.VssRel42DIMO()))
	if err != nil {
		return nil, fmt.Errorf("failed to load default signal info: %w", err)
	}
	definedSignals := defs.DefinedSignal(signalInfo)
	signalMap := make(map[string]*schema.SignalInfo, len(definedSignals))
	for _, signal := range definedSignals {
		signalMap[signal.JSONName] = signal
	}

	return signalMap, nil
}

// SignalConvert converts a default CloudEvent to DIMO's vss signals.
func (m *Module) SignalConvert(_ context.Context, event cloudevent.RawEvent) ([]vss.Signal, error) {
	m.once.Do(func() {
		m.signalMap, m.loadErr = LoadSignalMap()
	})
	if m.loadErr != nil {
		return nil, fmt.Errorf("failed to load signal map: %w", m.loadErr)
	}

	return SignalConvert(event, m.signalMap)
}

// FingerprintConvert converts a default CloudEvent to a FingerprintEvent.
func (*Module) FingerprintConvert(_ context.Context, event cloudevent.RawEvent) (cloudevent.Fingerprint, error) {
	result := gjson.GetBytes(event.Data, "vin")
	if !result.Exists() {
		return cloudevent.Fingerprint{}, fmt.Errorf("vin not found in event data")
	}
	return cloudevent.Fingerprint{VIN: result.String()}, nil
}

// CloudEventConvert marshals the input message to Cloud Events and sets the type based on the message content.
func (*Module) CloudEventConvert(_ context.Context, msgData []byte) ([]cloudevent.CloudEventHeader, []byte, error) {
	var event cloudevent.CloudEvent[json.RawMessage]
	err := json.Unmarshal(msgData, &event)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal message: %w", err)
	}
	hdrs := []cloudevent.CloudEventHeader{}
	if gjson.GetBytes(event.Data, "vin").Exists() {
		fpHdr := event.CloudEventHeader
		fpHdr.Type = cloudevent.TypeFingerprint
		hdrs = append(hdrs, fpHdr)
	}
	signals := gjson.GetBytes(event.Data, "signals")
	if signals.Exists() && signals.IsArray() && len(signals.Array()) > 0 {
		statusHdr := event.CloudEventHeader
		statusHdr.Type = cloudevent.TypeStatus
		hdrs = append(hdrs, statusHdr)
	}
	events := gjson.GetBytes(event.Data, "events")
	if events.Exists() && events.IsArray() && len(events.Array()) > 0 {
		statusHdr := event.CloudEventHeader
		statusHdr.Type = cloudevent.TypeEvent
		hdrs = append(hdrs, statusHdr)
	}

	// if we can't infer the type, default to unknown so we don't drop the event.
	if len(hdrs) == 0 {
		unknownHdr := event.CloudEventHeader
		unknownHdr.Type = cloudevent.TypeUnknown
		hdrs = append(hdrs, unknownHdr)
	}

	return hdrs, event.Data, nil
}

// EventConvert converts a default CloudEvent to events.
func (*Module) EventConvert(_ context.Context, event cloudevent.RawEvent) ([]vss.Event, error) {
	// Parse the events array from the event data
	var eventsData EventsData
	err := json.Unmarshal(event.Data, &eventsData)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal events data: %w", err)
	}

	vssEvents := make([]vss.Event, 0, len(eventsData.Events))
	var decodeErrs []error
	for _, ev := range eventsData.Events {
		if ev.Name == "" {
			decodeErrs = append(decodeErrs, fmt.Errorf("event.name is empty"))
			continue
		}
		if ev.Timestamp.IsZero() {
			decodeErrs = append(decodeErrs, fmt.Errorf("event.timestamp is zero for event.name %s", ev.Name))
			continue
		}
		if len(ev.Metadata) > 0 && !json.Valid([]byte(ev.Metadata)) {
			// We do not expect to get this far if the metadata is not valid json. Since it would invalidate the entire cloudevent.
			decodeErrs = append(decodeErrs, fmt.Errorf("metadata for event.name %s, event.timestamp %s is not valid json", ev.Name, ev.Timestamp))
			continue
		}

		vssEvent := vss.Event{
			Subject:      event.Subject,
			Source:       event.Source,
			Producer:     event.Producer,
			CloudEventID: event.ID,
			Name:         ev.Name,
			Timestamp:    ev.Timestamp,
			DurationNs:   ev.DurationNs,
			Metadata:     ev.Metadata,
		}
		vssEvents = append(vssEvents, vssEvent)
	}

	if len(decodeErrs) > 0 {
		return nil, convert.ConversionError{
			DecodedEvents: vssEvents,
			Errors:        decodeErrs,
			Subject:       event.Subject,
			Source:        event.Source,
		}
	}

	return vssEvents, nil
}
