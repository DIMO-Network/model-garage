package ruptela

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/vss"
)

const (
	EventNameHarshBraking   = "HarshBraking"
	EventNameExtremeBraking = "ExtremeBraking"
	EventNameAcceleration   = "HarshAcceleration"
	EventNameCornering      = "HarshCornering"
	zeroValue               = "0"
)

type eventSignals struct {
	Signals eventSignal `json:"signals"`
}
type eventSignal struct {
	Braking      string `json:"135"`
	Acceleration string `json:"136"`
	Cornering    string `json:"143"`
}

// CounterMetadata is the metadata for events with a counter value.
type CounterMetadata struct {
	CounterValue uint `json:"counterValue"`
}

// EventConvert converts a ruptela event to a vss event.
func DecodeEvent(cEvent cloudevent.RawEvent) ([]vss.Event, error) {
	var signals eventSignals
	if err := json.Unmarshal(cEvent.Data, &signals); err != nil {
		return nil, fmt.Errorf("failed to unmarshal event data: %w", err)
	}

	var events []vss.Event
	var errs []error
	if signals.Signals.Braking != zeroValue {
		brakingEvents, err := ToBrakingEvents(signals.Signals.Braking)
		if err != nil {
			errs = append(errs, err)
		}
		events = append(events, brakingEvents...)
	}
	if signals.Signals.Acceleration != zeroValue {
		accelerationEvent, err := ToAccelerationEvent(signals.Signals.Acceleration)
		if err == nil {
			events = append(events, accelerationEvent)
		} else if !errors.Is(err, errNotFound) {
			errs = append(errs, err)
		}
	}
	if signals.Signals.Cornering != zeroValue {
		corneringEvent, err := ToCorneringEvent(signals.Signals.Cornering)
		if err == nil {
			events = append(events, corneringEvent)
		} else if !errors.Is(err, errNotFound) {
			errs = append(errs, err)
		}
	}

	for i := range events {
		events[i].Subject = cEvent.Subject
		events[i].Source = cEvent.Source
		events[i].Producer = cEvent.Producer
		events[i].CloudEventID = cEvent.ID
		events[i].Timestamp = cEvent.Time
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

func ToBrakingEvents(rawValue string) ([]vss.Event, error) {
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return nil, fmt.Errorf("could not parse uint: %w", err)
	}

	// Ensure we're only working with 8 bits
	if rawInt > 0xFF {
		return nil, fmt.Errorf("value is greater than 8 bits: %d", rawInt)
	}
	value := uint8(rawInt)

	var events []vss.Event

	// Check 4 LSB (bits 0-3)
	lsb := value & 0x0F
	if lsb != 0 {
		metaCounter := CounterMetadata{
			CounterValue: uint(lsb),
		}
		metaCounterJSON, err := json.Marshal(metaCounter)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal metadata: %w", err)
		}
		events = append(events, vss.Event{
			Name:     EventNameHarshBraking,
			Metadata: string(metaCounterJSON),
		})
	}

	// Check 4 MSB (bits 4-7)
	msb := (value >> 4) & 0x0F
	if msb != 0 {
		metaCounter := CounterMetadata{
			CounterValue: uint(msb),
		}
		metaCounterJSON, err := json.Marshal(metaCounter)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal metadata: %w", err)
		}
		events = append(events, vss.Event{
			Name:     EventNameExtremeBraking,
			Metadata: string(metaCounterJSON),
		})
	}

	return events, nil
}

func ToAccelerationEvent(rawValue string) (vss.Event, error) {
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return vss.Event{}, fmt.Errorf("could not parse uint: %w", err)
	}
	if rawInt == 0 {
		return vss.Event{}, errNotFound
	}

	metaCounter := CounterMetadata{
		CounterValue: uint(rawInt),
	}
	metaCounterJSON, err := json.Marshal(metaCounter)
	if err != nil {
		return vss.Event{}, fmt.Errorf("failed to marshal metadata: %w", err)
	}
	return vss.Event{
		Name:     EventNameAcceleration,
		Metadata: string(metaCounterJSON),
	}, nil
}

func ToCorneringEvent(rawValue string) (vss.Event, error) {
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return vss.Event{}, fmt.Errorf("could not parse uint: %w", err)
	}
	if rawInt == 0 {
		return vss.Event{}, errNotFound
	}
	metaCounter := CounterMetadata{
		CounterValue: uint(rawInt),
	}
	metaCounterJSON, err := json.Marshal(metaCounter)
	if err != nil {
		return vss.Event{}, fmt.Errorf("failed to marshal metadata: %w", err)
	}
	return vss.Event{
		Name:     EventNameCornering,
		Metadata: string(metaCounterJSON),
	}, nil
}
