package ruptela

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/segmentio/ksuid"
)

const (
	EventNameHarshBraking   = vss.EventBehaviorHarshBrakingName
	EventNameExtremeBraking = vss.EventBehaviorExtremeBrakingName
	EventNameAcceleration   = vss.EventBehaviorHarshAccelerationName
	EventNameCornering      = vss.EventBehaviorHarshCorneringName
	EventNameEngineBlock    = vss.EventSecurityEngineBlockName
	EventNameEngineUnblock  = vss.EventSecurityEngineUnblockName
	zeroValue               = "0"
)

type eventSignals struct {
	Signals eventSignal `json:"signals"`
}
type eventSignal struct {
	Braking      string  `json:"135"`
	Acceleration string  `json:"136"`
	Cornering    string  `json:"143"`
	EngineBlock  *string `json:"405,omitempty"`
}

// CounterMetadata is the metadata for events with a counter value.
type CounterMetadata struct {
	CounterValue uint `json:"counterValue"`
}

// DecodeEvent converts a ruptela vehicle event to a vss event, in terms of a vehicle event or command that we can trigger on
func DecodeEvent(cEvent cloudevent.RawEvent) ([]vss.Event, error) {
	var signals eventSignals
	if err := json.Unmarshal(cEvent.Data, &signals); err != nil {
		return nil, fmt.Errorf("failed to unmarshal event data: %w", err)
	}

	events := make([]vss.Event, 0, 4)
	var errs []error
	if signals.Signals.Braking != "" && signals.Signals.Braking != zeroValue {
		brakingData, err := ToBrakingEventData(signals.Signals.Braking)
		if err == nil {
			for _, data := range brakingData {
				events = append(events, wrapEventData(cEvent, data))
			}
		} else {
			errs = append(errs, err)
		}
	}
	if signals.Signals.Acceleration != "" && signals.Signals.Acceleration != zeroValue {
		data, err := ToAccelerationEventData(signals.Signals.Acceleration)
		if err == nil {
			events = append(events, wrapEventData(cEvent, data))
		} else if !errors.Is(err, errNotFound) {
			errs = append(errs, err)
		}
	}
	if signals.Signals.Cornering != "" && signals.Signals.Cornering != zeroValue {
		data, err := ToCorneringEventData(signals.Signals.Cornering)
		if err == nil {
			events = append(events, wrapEventData(cEvent, data))
		} else if !errors.Is(err, errNotFound) {
			errs = append(errs, err)
		}
	}
	if signals.Signals.EngineBlock != nil && *signals.Signals.EngineBlock != "" {
		// this handles both engine block and unblock
		data, err := ToEngineSecurityEvent(*signals.Signals.EngineBlock)
		if err == nil {
			events = append(events, wrapEventData(cEvent, data))
		} else if !errors.Is(err, errNotFound) {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return nil, convert.ConversionError{
			DecodedEvents: events,
			Errors:        errs,
			Subject:       cEvent.Subject,
			Source:        cEvent.Source,
		}
	}

	return events, nil
}

// wrapEventData wraps EventData in a full vss.Event with CloudEvent header fields from the source event.
func wrapEventData(cEvent cloudevent.RawEvent, data vss.EventData) vss.Event {
	data.Timestamp = cEvent.Time
	data.CloudEventID = cEvent.ID
	if data.Tags == nil {
		data.Tags = []string{}
	}
	return vss.Event{
		CloudEventHeader: cloudevent.CloudEventHeader{
			Subject:     cEvent.Subject,
			Source:      cEvent.Source,
			Producer:    cEvent.Producer,
			ID:          ksuid.New().String(),
			SpecVersion: "1.0",
			Time:        cEvent.Time,
			Type:        cloudevent.TypeEvent,
			DataVersion: cEvent.DataVersion,
		},
		Data: data,
	}
}

// ToBrakingEventData parses a hex braking value into EventData entries.
func ToBrakingEventData(rawValue string) ([]vss.EventData, error) {
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return nil, fmt.Errorf("could not parse uint: %w", err)
	}

	// Ensure we're only working with 8 bits
	if rawInt > 0xFF {
		return nil, fmt.Errorf("value is greater than 8 bits: %d", rawInt)
	}
	value := uint8(rawInt)

	var data []vss.EventData

	// Check 4 LSB (bits 0-3)
	lsb := value & 0x0F
	if lsb != 0 {
		metaCounterJSON, err := json.Marshal(CounterMetadata{CounterValue: uint(lsb)})
		if err != nil {
			return nil, fmt.Errorf("failed to marshal metadata: %w", err)
		}
		data = append(data, vss.EventData{
			Name:     EventNameHarshBraking,
			Metadata: string(metaCounterJSON),
		})
	}

	// Check 4 MSB (bits 4-7)
	msb := (value >> 4) & 0x0F
	if msb != 0 {
		metaCounterJSON, err := json.Marshal(CounterMetadata{CounterValue: uint(msb)})
		if err != nil {
			return nil, fmt.Errorf("failed to marshal metadata: %w", err)
		}
		data = append(data, vss.EventData{
			Name:     EventNameExtremeBraking,
			Metadata: string(metaCounterJSON),
		})
	}

	return data, nil
}

// ToAccelerationEventData parses a hex acceleration value into EventData.
func ToAccelerationEventData(rawValue string) (vss.EventData, error) {
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return vss.EventData{}, fmt.Errorf("could not parse uint: %w", err)
	}
	if rawInt == 0 {
		return vss.EventData{}, errNotFound
	}

	metaCounterJSON, err := json.Marshal(CounterMetadata{CounterValue: uint(rawInt)})
	if err != nil {
		return vss.EventData{}, fmt.Errorf("failed to marshal metadata: %w", err)
	}
	return vss.EventData{
		Name:     EventNameAcceleration,
		Metadata: string(metaCounterJSON),
	}, nil
}

func ToEngineSecurityEvent(rawValue string) (vss.EventData, error) {
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return vss.EventData{}, fmt.Errorf("could not parse uint: %w", err)
	}
	if rawInt != 0 {
		return vss.EventData{Name: EventNameEngineBlock, Tags: []string{vss.EventSecurityEngineBlockName}}, nil
	}
	return vss.EventData{Name: EventNameEngineUnblock, Tags: []string{vss.EventSecurityEngineUnblockName}}, nil
}

// ToCorneringEventData parses a hex cornering value into EventData.
func ToCorneringEventData(rawValue string) (vss.EventData, error) {
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return vss.EventData{}, fmt.Errorf("could not parse uint: %w", err)
	}
	if rawInt == 0 {
		return vss.EventData{}, errNotFound
	}
	metaCounterJSON, err := json.Marshal(CounterMetadata{CounterValue: uint(rawInt)})
	if err != nil {
		return vss.EventData{}, fmt.Errorf("failed to marshal metadata: %w", err)
	}
	return vss.EventData{
		Name:     EventNameCornering,
		Metadata: string(metaCounterJSON),
	}, nil
}
