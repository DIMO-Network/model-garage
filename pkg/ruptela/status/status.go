// Package status holds decoding functions for Ruptela status payloads.
package status

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/DIMO-Network/model-garage/pkg/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/ruptela"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/tidwall/gjson"
)

// DecodeStatusSignals decodes a status message into a slice of signals.
func DecodeStatusSignals(msgBytes []byte) ([]vss.Signal, error) {
	event := cloudevent.CloudEvent[struct{}]{}
	err := json.Unmarshal(msgBytes, &event)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal message: %w", err)
	}
	var signals []vss.Signal
	switch event.DataVersion {
	case ruptela.StatusEventDS:
		signals, err = SignalsFromV1Payload(msgBytes)
	case ruptela.LocationEventDS:
		signals, err = SignalsFromLocationPayload(msgBytes)
	case ruptela.DTCEventDS:
		signals, err = SignalsFromDTCPayload(msgBytes)
	default:
		return nil, fmt.Errorf("unknown data version: %s", event.DataVersion)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to decode signals: %w", err)
	}
	return signals, nil
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

// TokenIDFromData gets a tokenID from a V2 payload.
func TokenIDFromData(jsonData []byte) (uint32, error) {
	lookupKey := "subject"
	subject := gjson.GetBytes(jsonData, lookupKey)
	if !subject.Exists() {
		return 0, convert.FieldNotFoundError{Field: "tokenID", Lookup: lookupKey}
	}
	subjectStr, ok := subject.Value().(string)
	if !ok {
		return 0, fmt.Errorf("%s field is not a string", lookupKey)
	}
	subjectDID, err := cloudevent.DecodeNFTDID(subjectStr)
	if err != nil {
		return 0, fmt.Errorf("error decoding subject: %w", err)
	}
	return subjectDID.TokenID, nil
}

// SourceFromData gets a source from a V2 payload.
func SourceFromData(jsonData []byte) (string, error) {
	lookupKey := "source"
	source := gjson.GetBytes(jsonData, lookupKey)
	if !source.Exists() {
		return "", convert.FieldNotFoundError{Field: "source", Lookup: lookupKey}
	}
	src, ok := source.Value().(string)
	if !ok {
		return "", errors.New("source field is not a string")
	}
	return src, nil
}
