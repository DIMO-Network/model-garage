// Package status provides a functions for managing lorawan (Macaron) status payloads.
package status

import (
	"errors"
	"fmt"
	"time"

	"github.com/tidwall/gjson"

	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/lorawan"
	"github.com/DIMO-Network/model-garage/pkg/vss"
)

// SignalsFromV2Payload extracts signals from a V2 payload.
func SignalsFromV2Payload(jsonData []byte) ([]vss.Signal, error) {
	tokenID, err := lorawan.TokenIDFromData(jsonData)
	if err != nil {
		return nil, convert.ConversionError{
			Errors: []error{fmt.Errorf("error getting tokenId: %w", err)},
		}
	}
	source, err := lorawan.SourceFromData(jsonData)
	if err != nil {
		return nil, convert.ConversionError{
			TokenID: tokenID,
			Errors:  []error{fmt.Errorf("error getting source: %w", err)},
		}
	}
	signals := gjson.GetBytes(jsonData, "data.vehicle.signals")
	if !signals.Exists() {
		return nil, convert.ConversionError{
			TokenID: tokenID,
			Source:  source,
			Errors:  []error{convert.FieldNotFoundError{Field: "signals", Lookup: "data.vehicle.signals"}},
		}
	}
	if !signals.IsArray() {
		if signals.Value() == nil {
			// If the signals array is NULL treat it like an empty array.
			return []vss.Signal{}, nil
		}
		return nil, convert.ConversionError{
			TokenID: tokenID,
			Source:  source,
			Errors:  []error{errors.New("signals field is not an array")},
		}
	}
	retSignals := []vss.Signal{}
	signalMeta := vss.Signal{
		TokenID: tokenID,
		Source:  source,
	}

	conversionErrors := convert.ConversionError{
		TokenID: tokenID,
		Source:  source,
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
		sigs, err := lorawan.SignalsFromV2Data(jsonData, signalMeta, originalName, sigData)
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
