package ruptela

import (
	"errors"
	"fmt"
	"time"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/tidwall/gjson"
)

// SignalsFromLocationPayload extracts signals from a V2 payload.
func SignalsFromLocationPayload(event cloudevent.RawEvent) ([]vss.Signal, error) {
	did, err := cloudevent.DecodeERC721DID(event.Subject)
	if err != nil {
		return nil, fmt.Errorf("failed to decode subject DID: %w", err)
	}
	signals := gjson.GetBytes(event.Data, "location")
	if !signals.Exists() {
		return nil, convert.ConversionError{
			Subject: event.Subject,
			Source:  event.Source,
			Errors:  []error{convert.FieldNotFoundError{Field: "signals", Lookup: "data.location"}},
		}
	}
	if !signals.IsArray() {
		if signals.Value() == nil {
			// If the signals array is NULL treat it like an empty array.
			return []vss.Signal{}, nil
		}
		return nil, convert.ConversionError{
			Subject: event.Subject,
			Source:  event.Source,
			Errors:  []error{errors.New("signals field is not an array")},
		}
	}
	retSignals := []vss.Signal{}
	signalMeta := vss.Signal{
		TokenID: uint32(did.TokenID.Uint64()), //nolint:gosec // will not exceed uint32 max value
		Source:  event.Source,
	}

	conversionErrors := convert.ConversionError{
		Subject: event.Subject,
		Source:  event.Source,
	}
	for _, sigData := range signals.Array() {
		if !sigData.IsObject() {
			err := errors.New("signal data is not an object")
			conversionErrors.Errors = append(conversionErrors.Errors, err)
			continue
		}
		// first get the timestamp field from the signal object
		ts, err := TimestampFromLocationSignal(sigData)
		if err != nil {
			err = fmt.Errorf("error getting location signal: %w", err)
			conversionErrors.Errors = append(conversionErrors.Errors, err)
			continue
		}
		// loop over other fields in the signal object and create a signal for each.
		sigData.ForEach(func(key, value gjson.Result) bool {
			if key.String() == "ts" {
				// skip the timestamp field
				return true
			}
			signalMeta.Timestamp = ts
			sigs, err := SignalsFromLocationData(event.Data, signalMeta, key.String(), value)
			if err != nil {
				conversionErrors.Errors = append(conversionErrors.Errors, err)
				return true
			}
			AddCurrentLocationSignal(&sigs, signalMeta)
			retSignals = append(retSignals, sigs...)
			return true
		})
		// Create the location tuple, if possible.
		locSig := vss.Signal{
			TokenID:   signalMeta.TokenID,
			Timestamp: ts,
			Name:      "currentLocation",
			Source:    signalMeta.Source,
		}
		locSig.SetValue(
			vss.Location{
				Latitude:  sigData.Get("lat").Num,
				Longitude: sigData.Get("lon").Num,
				HDOP:      sigData.Get("hdop").Num,
			},
		)

	}

	if len(conversionErrors.Errors) > 0 {
		conversionErrors.DecodedSignals = retSignals
		return nil, conversionErrors
	}
	return retSignals, nil
}

// TimestampFromLocationSignal gets a timestamp from a V2 signal.
func TimestampFromLocationSignal(sigResult gjson.Result) (time.Time, error) {
	lookupKey := "ts"
	timestamp := sigResult.Get(lookupKey)
	if !timestamp.Exists() {
		return time.Time{}, convert.FieldNotFoundError{Field: "timestamp", Lookup: lookupKey}
	}
	switch timestamp.Type {
	case gjson.Number:
		return time.Unix(timestamp.Int(), 0).UTC(), nil
	case gjson.String:
		ts, err := time.Parse(time.RFC3339, timestamp.String())
		if err != nil {
			return time.Time{}, fmt.Errorf("error parsing timestamp: %w", err)
		}
		return ts, nil
	}
	return time.Time{}, fmt.Errorf("timestamp field is not a number or string")
}

// NameFromV2Signal gets a name from a V2 signal.
func NameFromV2Signal(sigResult gjson.Result) (string, error) {
	lookupKey := "name"
	signalName := sigResult.Get(lookupKey)
	if !signalName.Exists() {
		return "", convert.FieldNotFoundError{Field: "name", Lookup: lookupKey}
	}
	return signalName.String(), nil
}
