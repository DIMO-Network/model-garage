package ruptela

import (
	"fmt"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/vss"
)

// DecodeStatusSignals decodes a status message into a slice of signals.
func DecodeStatusSignals(event cloudevent.RawEvent) ([]vss.Signal, error) {
	var err error
	var signals []vss.Signal
	switch event.DataVersion {
	case StatusEventDS:
		signals, err = SignalsFromV1Payload(event)
	case LocationEventDS:
		signals, err = SignalsFromLocationPayload(event)
	case DTCEventDS:
		signals, err = SignalsFromDTCPayload(event)
	default:
		return nil, fmt.Errorf("unknown data version: %s", event.DataVersion)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to decode signals: %w", err)
	}
	return signals, nil
}
