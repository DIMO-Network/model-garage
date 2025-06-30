// Package hashdog holds decoding functions for hashdog (Macaron) status payloads.
package hashdog

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
	FingerprintFrame = 0x01
	SleepFrame       = 0x05
	DataVersion      = "v2"
)

// ConvertToCloudEvents converts a message data payload into a slice of CloudEvents.
// It handles both status and fingerprint events, creating separate CloudEvents for each.
func ConvertToCloudEvents(msgData []byte, chainID uint64, aftermarketContractAddr, vehicleContractAddr common.Address) ([]cloudevent.CloudEventHeader, []byte, error) {
	var result []cloudevent.CloudEventHeader

	var event LorawanEvent
	err := json.Unmarshal(msgData, &event)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal record data: %w", err)
	}
	if event.DeviceTokenID == nil {
		return nil, nil, fmt.Errorf("device token id is missing")
	}

	var statusData Data
	err = json.Unmarshal(event.Data, &statusData)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal event data field: %w", err)
	}

	if statusData.Header <= 0 {
		return nil, nil, fmt.Errorf("unknown event frame header: %d", statusData.Header)
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

	if statusData.Header == FingerprintFrame {
		cloudEvent, err := createCloudEventHeader(event, producer, subject, cloudevent.TypeFingerprint)
		if err != nil {
			return nil, nil, err
		}
		// Append the status event to the result
		result = append(result, cloudEvent)
	}

	cloudEvent, err := createCloudEventHeader(event, producer, subject, cloudevent.TypeStatus)
	if err != nil {
		return nil, nil, err
	}
	// Append the status event to the result
	result = append(result, cloudEvent)

	// Each Lorawan payload has device information, so we need to create separate status event where subject == producer
	cloudEventDevice, err := createCloudEventHeader(event, producer, producer, cloudevent.TypeStatus)
	if err != nil {
		return nil, nil, err
	}
	// Append the status event to the result
	result = append(result, cloudEventDevice)

	return result, event.Data, nil
}

// createCloudEventHeader creates a CloudEvent header from LorawanEvent.
func createCloudEventHeader(event LorawanEvent, producer, subject, eventType string) (cloudevent.CloudEventHeader, error) {
	timeValue, err := time.Parse(time.RFC3339, event.Time)
	if err != nil {
		return cloudevent.CloudEventHeader{}, fmt.Errorf("failed to parse time: %w", err)
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
