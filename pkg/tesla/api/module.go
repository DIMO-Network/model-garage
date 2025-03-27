package api

import (
	"errors"
	"fmt"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/tidwall/gjson"
)

const DataVersion = "fleet_api/v1.0.0"

// SignalConvert converts a Tesla CloudEvent to DIMO's VSS rows.
func SignalConvert(event cloudevent.RawEvent) ([]vss.Signal, error) {
	did, err := cloudevent.DecodeNFTDID(event.Subject)
	if err != nil {
		return nil, fmt.Errorf("failed to decode subject DID: %w", err)
	}

	tokenID := did.TokenID
	source := event.Source

	baseSignal := vss.Signal{
		TokenID: tokenID,
		Source:  source,
	}

	sigs, errs := SignalsFromTesla(baseSignal, event.Data)
	if len(errs) != 0 {
		return nil, convert.ConversionError{
			TokenID:        tokenID,
			Source:         source,
			DecodedSignals: sigs,
			Errors:         errs,
		}
	}

	return sigs, nil
}

func IsFingerprint(event cloudevent.RawEvent) bool {
	result := gjson.GetBytes(event.Data, "vin")

	// This check should always pass.
	return result.Exists() && result.Type == gjson.String
}

func FingerprintConvert(event cloudevent.RawEvent) (cloudevent.Fingerprint, error) {
	var fp cloudevent.Fingerprint

	result := gjson.GetBytes(event.Data, "vin")

	if !result.Exists() {
		return fp, errors.New("data object has no VIN field")
	}
	if result.Type != gjson.String {
		return fp, errors.New("data object has a VIN field, but it is not a string")
	}

	fp.VIN = result.String()
	return fp, nil
}
