package ruptela

import (
	"errors"
	"fmt"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/tidwall/gjson"
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
	if coordLoc, err := currentLocationCoordinatesFromV1Data(event.Data); err == nil {
		sig := vss.Signal{
			Name:      vss.FieldCurrentLocationCoordinates,
			TokenID:   baseSignal.TokenID,
			Timestamp: baseSignal.Timestamp,
			Source:    baseSignal.Source,
		}
		sig.SetValue(coordLoc)
		sigs = append(sigs, sig)
	} else if !errors.Is(err, errNotFound) {
		errs = append(errs, err)
	}
	if errs != nil {
		return nil, convert.ConversionError{
			Subject:        event.Subject,
			Source:         event.Source,
			DecodedSignals: sigs,
			Errors:         errs,
		}
	}
	return sigs, nil
}

// currentLocationCoordinatesFromV1Data extracts a vss.Location from v1 JSON data.
func currentLocationCoordinatesFromV1Data(jsonData []byte) (vss.Location, error) {
	pos := gjson.GetBytes(jsonData, "pos")
	if !pos.Exists() {
		return vss.Location{}, errNotFound
	}
	return posToLocation(pos)
}
