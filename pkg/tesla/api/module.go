package api

import (
	"errors"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/tidwall/gjson"
)

const DataVersion = "fleet_api/v1.0.0"

// SignalConvert converts a Tesla Fleet API response CloudEvent to DIMO's VSS rows.
func SignalConvert(event cloudevent.RawEvent) ([]vss.Signal, error) {
	source := event.Source

	baseSignal := vss.Signal{
		Subject: event.Subject,
		Source:  source,
	}

	sigs, errs := SignalsFromTesla(baseSignal, event.Data)
	if len(errs) != 0 {
		return nil, convert.ConversionError{
			Subject:        event.Subject,
			Source:         source,
			DecodedSignals: sigs,
			Errors:         errs,
		}
	}

	return sigs, nil
}

// IsFingerprint returns whether the Fleet API response contains an extractable
// VIN. This should always return true.
func IsFingerprint(event cloudevent.RawEvent) bool {
	result := gjson.GetBytes(event.Data, "vin")

	// This check should always pass.
	return result.Exists() && result.Type == gjson.String
}

// FingerprintConvert extracts a fingerprint from a Fleet API vehicle_data response.
// We expect this to always succeed.
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
