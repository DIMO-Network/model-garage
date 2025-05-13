package autopi

import (
	"errors"
	"fmt"
	"time"

	"github.com/tidwall/gjson"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/vss"
)

// SignalsFromV2Payload extracts signals from a V2 payload.
func SignalsFromV2Payload(event cloudevent.RawEvent) ([]vss.Signal, error) {
	did, err := cloudevent.DecodeERC721DID(event.Subject)
	if err != nil {
		return nil, fmt.Errorf("failed to decode DID: %w", err)
	}
	signals := gjson.GetBytes(event.Data, "vehicle.signals")
	if !signals.Exists() {
		return nil, convert.ConversionError{
			TokenID: uint32(did.TokenID.Uint64()), //nolint:gosec // will not exceed uint32 max value
			Source:  event.Source,
			Errors:  []error{convert.FieldNotFoundError{Field: "signals", Lookup: "data.vehicle.signals"}},
		}
	}
	if !signals.IsArray() {
		if signals.Value() == nil {
			// If the signals array is NULL treat it like an empty array.
			return []vss.Signal{}, nil
		}
		return nil, convert.ConversionError{
			TokenID: uint32(did.TokenID.Uint64()), //nolint:gosec // will not exceed uint32 max value
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
		TokenID: uint32(did.TokenID.Uint64()), //nolint:gosec // will not exceed uint32 max value
		Source:  event.Source,
	}
	for _, sigData := range signals.Array() {
		originalName, err := NameFromV2Signal(sigData)
		if err != nil {
			conversionErrors.Errors = append(conversionErrors.Errors, err)
			continue
		}
		ts, err := TimestampFromV2Signal(sigData)
		if err != nil {
			err = fmt.Errorf("error for '%s': %w", originalName, err)
			conversionErrors.Errors = append(conversionErrors.Errors, err)
			continue
		}
		signalMeta.Timestamp = ts
		sigs, err := SignalsFromV2Data(event.Data, signalMeta, originalName, sigData)
		if err != nil {
			conversionErrors.Errors = append(conversionErrors.Errors, err)
			continue
		}
		retSignals = append(retSignals, sigs...)
	}

	if len(conversionErrors.Errors) > 0 {
		conversionErrors.DecodedSignals = retSignals
		return nil, conversionErrors
	}
	return retSignals, nil
}

// TimestampFromV2Signal gets a timestamp from a V2 signal.
func TimestampFromV2Signal(sigResult gjson.Result) (time.Time, error) {
	lookupKey := "timestamp"
	timestamp := sigResult.Get(lookupKey)
	if !timestamp.Exists() {
		return time.Time{}, convert.FieldNotFoundError{Field: "timestamp", Lookup: lookupKey}
	}
	return time.UnixMilli(timestamp.Int()).UTC(), nil
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
