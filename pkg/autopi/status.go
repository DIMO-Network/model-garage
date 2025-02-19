// Package autopi holds decoding functions for Ruptela status payloads.
package autopi

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/DIMO-Network/model-garage/pkg/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/segmentio/ksuid"
	"github.com/tidwall/gjson"
	"golang.org/x/mod/semver"
)

const (
	StatusEventType      = "com.dimo.device.status.v2"
	FingerprintEventType = "zone.dimo.aftermarket.device.fingerprint"
	DataVersion          = "v2"
)

type AutopiEvent struct {
	Data           json.RawMessage `json:"data"`
	VehicleTokenID *uint32         `json:"vehicleTokenId"`
	DeviceTokenID  *uint32         `json:"deviceTokenId"`
	Signature      string          `json:"signature"`
	Time           string          `json:"time"`
	Type           string          `json:"type"`
}

const (
	// StatusV1 is the version string for payloads with the version 1.0 schema.
	StatusV1 = "v1.0.0"
	// StatusV1Converted is the version string for payloads that have been converted to the 1.0 schema.
	StatusV1Converted = "v1.1.0"
)

// GetDataVersion returns the version string used in the payload.
func GetDataVersion(jsonData []byte) string {
	dataVersion := gjson.GetBytes(jsonData, "dataversion")
	if !dataVersion.Exists() {
		return ""
	}
	return dataVersion.String()
}

// HasV1Data checks if the payload has the same version as v1.0.0.
func HasV1Data(version string) bool {
	return version == "" || semver.Compare(StatusV1, version) == 0 || semver.Compare(StatusV1Converted, version) == 0
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

// ConvertToCloudEvents converts a message data payload into a slice of CloudEvents.
// It handles both status and fingerprint events, creating separate CloudEvents for each.
func ConvertToCloudEvents(msgData []byte, chainID uint64, aftermarketContractAddr, vehicleContractAddr string) ([][]byte, error) {
	var result [][]byte

	var event AutopiEvent
	err := json.Unmarshal(msgData, &event)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal record data: %w", err)
	}
	if event.DeviceTokenID == nil {
		return nil, fmt.Errorf("device token id is missing")
	}

	// handle both status and fingerprint events
	var eventType string
	switch event.Type {
	case StatusEventType:
		eventType = cloudevent.TypeStatus
	case FingerprintEventType:
		eventType = cloudevent.TypeFingerprint
	default:
		return nil, fmt.Errorf("unknown event type: %s", event.Type)
	}

	// Construct the producer DID
	producer := cloudevent.NFTDID{
		ChainID:         chainID,
		ContractAddress: common.HexToAddress(aftermarketContractAddr),
		TokenID:         *event.DeviceTokenID,
	}.String()

	// Construct the subject
	var subject string
	if event.VehicleTokenID != nil {
		subject = cloudevent.NFTDID{
			ChainID:         chainID,
			ContractAddress: common.HexToAddress(vehicleContractAddr),
			TokenID:         *event.VehicleTokenID,
		}.String()
	}

	cloudEvent, err := convertToCloudEvent(event, producer, subject, eventType)
	if err != nil {
		return nil, err
	}
	// Append the status event to the result
	result = append(result, cloudEvent)

	// Each AP payload has device information, so we need to create separate status event where subject == producer
	cloudEventDevice, err := convertToCloudEvent(event, producer, producer, cloudevent.TypeStatus)
	if err != nil {
		return nil, err
	}

	// Append the status event to the result
	result = append(result, cloudEventDevice)

	return result, nil
}

// convertToCloudEvent wraps an AutopiEvent into a CloudEvent.
// Returns:
//   - A byte slice containing the JSON representation of the CloudEvent.
//   - An error if the CloudEvent creation or marshaling fails.
func convertToCloudEvent(event AutopiEvent, producer, subject, eventType string) ([]byte, error) {
	cloudEvent, err := createCloudEvent(event, producer, subject, eventType)
	if err != nil {
		return nil, err
	}

	cloudEventBytes, err := json.Marshal(cloudEvent)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal cloudEvent: %w", err)
	}
	return cloudEventBytes, nil
}

// createCloudEvent creates a cloud event from autopi event.
func createCloudEvent(event AutopiEvent, producer, subject, eventType string) (cloudevent.CloudEvent[json.RawMessage], error) {
	timeValue, err := time.Parse(time.RFC3339, event.Time)
	if err != nil {
		return cloudevent.CloudEvent[json.RawMessage]{}, fmt.Errorf("failed to parse time: %v", err)
	}
	return cloudevent.CloudEvent[json.RawMessage]{
		CloudEventHeader: cloudevent.CloudEventHeader{
			DataContentType: "application/json",
			ID:              ksuid.New().String(),
			Subject:         subject,
			SpecVersion:     "1.0",
			Time:            timeValue,
			Type:            eventType,
			DataVersion:     DataVersion,
			Producer:        producer,
			Extras: map[string]any{
				"signature": event.Signature,
			},
		},
		Data: event.Data,
	}, nil
}
