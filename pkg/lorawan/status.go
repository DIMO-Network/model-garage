// Package lorawan holds decoding functions for lorawan (Macaron) status payloads.
package lorawan

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
)

const (
	FingerprintFrame = 0x01
	SleepFrame       = 0x05
	DataVersion      = "v2"
)

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

	var event LorawanEvent
	err := json.Unmarshal(msgData, &event)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal record data: %w", err)
	}
	if event.DeviceTokenID == nil {
		return nil, fmt.Errorf("device token id is missing")
	}

	if event.Data.Header <= 0 {
		return nil, fmt.Errorf("unknown event frame header: %d", event.Data.Header)
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

	if event.Data.Header == FingerprintFrame {
		cloudEvent, err := convertToCloudEvent(event, producer, subject, cloudevent.TypeFingerprint)
		if err != nil {
			return nil, err
		}
		// Append the status event to the result
		result = append(result, cloudEvent)
	}

	cloudEvent, err := convertToCloudEvent(event, producer, subject, cloudevent.TypeStatus)
	if err != nil {
		return nil, err
	}
	// Append the status event to the result
	result = append(result, cloudEvent)

	// Each Lorawan payload has device information, so we need to create separate status event where subject == producer
	cloudEventDevice, err := convertToCloudEvent(event, producer, producer, cloudevent.TypeStatus)
	if err != nil {
		return nil, err
	}
	// Append the status event to the result
	result = append(result, cloudEventDevice)

	return result, nil
}

// convertToCloudEvent wraps a LorawanEvent into a CloudEvent.
// Returns:
//   - A byte slice containing the JSON representation of the CloudEvent.
//   - An error if the CloudEvent creation or marshaling fails.
func convertToCloudEvent(event LorawanEvent, producer, subject, eventType string) ([]byte, error) {
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

// createCloudEvent creates a CloudEvent from LorawanEvent.
func createCloudEvent(event LorawanEvent, producer, subject, eventType string) (cloudevent.CloudEvent[json.RawMessage], error) {
	timeValue, err := time.Parse(time.RFC3339, event.Time)
	if err != nil {
		return cloudevent.CloudEvent[json.RawMessage]{}, fmt.Errorf("failed to parse time: %v", err)
	}
	rawData, err := json.Marshal(event.Data)
	if err != nil {
		return cloudevent.CloudEvent[json.RawMessage]{}, fmt.Errorf("failed to marshall data: %v", err)
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
		Data: rawData,
	}, nil
}
