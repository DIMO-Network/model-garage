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
		Subject: event.Subject,
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
			retSignals = append(retSignals, sigs...)
			return true
		})
		// Special case location.
		if coordLoc, err := posToLocation(sigData); err == nil {
			sig := vss.Signal{
				Name:      vss.FieldCurrentLocationCoordinates,
				Subject:   signalMeta.Subject,
				Timestamp: ts,
				Source:    signalMeta.Source,
			}
			sig.SetValue(coordLoc)
			retSignals = append(retSignals, sig)
		} else if !errors.Is(err, errNotFound) {
			conversionErrors.Errors = append(conversionErrors.Errors, err)
		}
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

// posToLocation converts a gjson.Result representing an Ruptela location object with lat, lon,
// and optional hdop fields into a vss.Location.
//
// If lat or lon are absent or equal to the sentinel value -0x80000000 then this function returns
// errNotFound. If hdop is absent or equal to the sentinel 0xff then HDOP is set to zero.
// Appropriate scaling factors are applied.
//
// This function is outside of the code generation regime because it combines two or three fields
// of a Ruptela location object into one, and these fields can appear in at least two kinds of
// locations in JSON, depending on the dataversion.
func posToLocation(loc gjson.Result) (vss.Location, error) {
	latResult := loc.Get("lat")
	lonResult := loc.Get("lon")
	if !latResult.Exists() || latResult.Value() == nil || !lonResult.Exists() || lonResult.Value() == nil {
		return vss.Location{}, errNotFound
	}
	latVal, ok := latResult.Value().(float64)
	if !ok {
		return vss.Location{}, fmt.Errorf("%w, field 'lat' is not of type 'float64' got '%v' of type '%T'", convert.InvalidTypeError(), latResult.Value(), latResult.Value())
	}
	lonVal, ok := lonResult.Value().(float64)
	if !ok {
		return vss.Location{}, fmt.Errorf("%w, field 'lon' is not of type 'float64' got '%v' of type '%T'", convert.InvalidTypeError(), lonResult.Value(), lonResult.Value())
	}
	if latVal == -0x80000000 || lonVal == -0x80000000 {
		return vss.Location{}, errNotFound
	}
	var hdop float64
	hdopResult := loc.Get("hdop")
	if hdopResult.Exists() && hdopResult.Value() != nil {
		hdopVal, ok := hdopResult.Value().(float64)
		if !ok {
			return vss.Location{}, fmt.Errorf("%w, field 'hdop' is not of type 'float64' got '%v' of type '%T'", convert.InvalidTypeError(), hdopResult.Value(), hdopResult.Value())
		}
		if hdopVal != 0xff {
			hdop = hdopVal / 10
		}
	}
	return vss.Location{
		Latitude:  latVal / 10_000_000,
		Longitude: lonVal / 10_000_000,
		HDOP:      hdop,
	}, nil
}
