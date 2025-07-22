// Package nativestatus provides a function to generate conversion functions for a vehicle struct.
package nativestatus

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/tidwall/gjson"
)

// TokenIDGetter is an interface to get a tokenID from a subject.
type TokenIDGetter interface {
	TokenIDFromSubject(ctx context.Context, subject string) (uint32, error)
}

// SignalsFromV1Payload gets a slice signals from a v1 payload.
func SignalsFromV1Payload(ctx context.Context, tokenGetter TokenIDGetter, jsonData []byte) ([]vss.Signal, error) {
	ts, err := TimestampFromV1Data(jsonData)
	if err != nil {
		return nil, convert.ConversionError{
			Errors: []error{fmt.Errorf("error getting timestamp: %w", err)},
		}
	}
	tokenID, err := TokenIDFromV1Data(ctx, jsonData, tokenGetter)
	if err != nil {
		return nil, convert.ConversionError{
			Errors: []error{fmt.Errorf("error getting tokenId: %w", err)},
		}
	}

	source, err := SourceFromData(jsonData)
	if err != nil {
		return nil, convert.ConversionError{
			Subject: strconv.FormatUint(uint64(tokenID), 10),
			Errors:  []error{fmt.Errorf("error getting source: %w", err)},
		}
	}
	baseSignal := vss.Signal{
		TokenID:   tokenID,
		Timestamp: ts,
		Source:    source,
	}
	sigs, errs := SignalsFromV1Data(baseSignal, jsonData)
	if errs != nil {
		return nil, convert.ConversionError{
			Subject:        strconv.FormatUint(uint64(tokenID), 10),
			Source:         source,
			DecodedSignals: sigs,
			Errors:         errs,
		}
	}
	return sigs, nil
}

// SubjectFromV1Data gets a subject from a v1 payload.
func SubjectFromV1Data(jsonData []byte) (string, error) {
	result := gjson.GetBytes(jsonData, "subject")
	if !result.Exists() {
		return "", convert.FieldNotFoundError{Field: "subject", Lookup: "subject"}
	}
	sub, ok := result.Value().(string)
	if !ok {
		return "", errors.New("subject field is not a string")
	}
	return sub, nil
}

// TimestampFromV1Data gets a timestamp from a v1 payload.
func TimestampFromV1Data(jsonData []byte) (time.Time, error) {
	result := gjson.GetBytes(jsonData, "time")
	if !result.Exists() {
		return time.Time{}, convert.FieldNotFoundError{Field: "timestamp", Lookup: "time"}
	}

	timeStr, ok := result.Value().(string)
	if !ok {
		ms, ok := result.Value().(float64)
		if ok {
			return time.UnixMilli(int64(ms)), nil
		}
		return time.Time{}, errors.New("time field is not a string or float64")
	}
	ts, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("error parsing time: %w", err)
	}
	return ts, nil
}

// TokenIDFromV1Data gets a tokenID from a v1 payload.
// attempts to get the tokenID from the vehicleTokenId field.
// If the field is not found and tokenGetter is not nil it will attempt to get the tokenId using the subject field and the tokenGetter.
func TokenIDFromV1Data(ctx context.Context, jsonData []byte, tokenGetter TokenIDGetter) (uint32, error) {
	tokenID, err := TokenIDFromData(jsonData)
	if err == nil {
		return tokenID, nil
	}
	if !errors.As(err, &convert.FieldNotFoundError{}) {
		return 0, err
	}

	// if the tokenGetter is nil then we cannot get the tokenID from the subject
	if tokenGetter == nil {
		return 0, err
	}

	sub, err := SubjectFromV1Data(jsonData)
	if err != nil {
		return 0, err
	}
	tokenID, err = tokenGetter.TokenIDFromSubject(ctx, sub)
	if err != nil {
		return 0, fmt.Errorf("error getting tokenID: %w", err)
	}
	return tokenID, nil
}
