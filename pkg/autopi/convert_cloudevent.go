// Package autopi holds decoding functions for Ruptela status payloads.
package autopi

import (
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/DIMO-Network/cloudevent"
	"github.com/ethereum/go-ethereum/common"
	"github.com/segmentio/ksuid"
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

// ConvertToCloudEvents converts a message data payload into a slice of CloudEvents.
// It handles both status and fingerprint events, creating separate CloudEvents for each.
func ConvertToCloudEvents(msgData []byte, chainID uint64, aftermarketContractAddr, vehicleContractAddr common.Address) ([]cloudevent.CloudEventHeader, []byte, error) {
	var result []cloudevent.CloudEventHeader

	var event AutopiEvent
	err := json.Unmarshal(msgData, &event)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal record data: %w", err)
	}
	if event.DeviceTokenID == nil {
		return nil, nil, fmt.Errorf("device token id is missing")
	}

	// handle both status and fingerprint events
	var eventType string
	switch event.Type {
	case StatusEventType:
		eventType = cloudevent.TypeStatus
	case FingerprintEventType:
		eventType = cloudevent.TypeFingerprint
	default:
		return nil, nil, fmt.Errorf("unknown event type: %s", event.Type)
	}

	// Construct the producer DID
	producer := cloudevent.ERC721DID{
		ChainID:         chainID,
		ContractAddress: aftermarketContractAddr,
		TokenID:         big.NewInt(int64(*event.DeviceTokenID)),
	}.String()

	// Construct the subject
	var subject string
	if event.VehicleTokenID != nil {
		subject = cloudevent.ERC721DID{
			ChainID:         chainID,
			ContractAddress: vehicleContractAddr,
			TokenID:         big.NewInt(int64(*event.VehicleTokenID)),
		}.String()
	}

	cloudEvent, err := createCloudEventHeader(event, producer, subject, eventType)
	if err != nil {
		return nil, nil, err
	}
	// Append the status event to the result
	result = append(result, cloudEvent)

	// Each AP payload has device information, so we need to create separate status event where subject == producer
	cloudEventDevice, err := createCloudEventHeader(event, producer, producer, cloudevent.TypeStatus)
	if err != nil {
		return nil, nil, err
	}

	// Append the status event to the result
	result = append(result, cloudEventDevice)

	return result, event.Data, nil
}

// createCloudEvent creates a cloud event from autopi event.
func createCloudEventHeader(event AutopiEvent, producer, subject, eventType string) (cloudevent.CloudEventHeader, error) {
	timeValue, err := time.Parse(time.RFC3339, event.Time)
	if err != nil {
		return cloudevent.CloudEventHeader{}, fmt.Errorf("failed to parse time: %v", err)
	}
	return cloudevent.CloudEventHeader{
		DataContentType: "application/json",
		ID:              ksuid.New().String(),
		Subject:         subject,
		SpecVersion:     "1.0",
		Time:            timeValue,
		Type:            eventType,
		DataVersion:     DataVersion,
		Producer:        producer,
		Signature:       event.Signature,
	}, nil
}
