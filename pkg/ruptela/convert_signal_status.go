package ruptela

import (
	"fmt"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/vss"
)

// SignalsFromV1Payload gets a slice signals from a v1 payload.
func SignalsFromV1Payload(event cloudevent.RawEvent) ([]vss.Signal, error) {
	did, err := cloudevent.DecodeERC721DID(event.Subject)
	if err != nil {
		return nil, fmt.Errorf("failed to decode subject DID: %w", err)
	}

	baseSignal := vss.Signal{
		TokenID:   uint32(did.TokenID.Uint64()), //nolint:gosec // will not exceed uint32 max value
		Timestamp: event.Time,
		Source:    event.Source,
	}
	sigs, errs := SignalsFromV1Data(baseSignal, event.Data)
	if errs != nil {
		return nil, convert.ConversionError{
			Subject:        event.Subject,
			Source:         event.Source,
			DecodedSignals: sigs,
			Errors:         errs,
		}
	}
	AddCurrentLocationSignal(&sigs, baseSignal)
	return sigs, nil
}
