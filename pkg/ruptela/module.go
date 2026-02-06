package ruptela

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/segmentio/ksuid"
)

// Module is a module that converts ruptela messages to signals.
type Module struct {
	ChainID                 uint64         `json:"chain_id"`
	AftermarketContractAddr common.Address `json:"aftermarket_contract_addr"`
	VehicleContractAddr     common.Address `json:"vehicle_contract_addr"`
}

// FingerprintConvert converts a message to a fingerprint.
func (*Module) FingerprintConvert(_ context.Context, event cloudevent.RawEvent) (cloudevent.Fingerprint, error) {
	return DecodeFingerprint(event)
}

// SignalConvert converts a message to signals.
func (*Module) SignalConvert(_ context.Context, event cloudevent.RawEvent) ([]vss.Signal, error) {
	if event.DataVersion == DevStatusDS || event.DataVersion == BattDS || event.Type != cloudevent.TypeStatus {
		return nil, nil
	}
	signals, err := DecodeStatusSignals(event)
	if err == nil {
		return signals, nil
	}
	convertErr := convert.ConversionError{}
	if !errors.As(err, &convertErr) {
		// Add the error to the batch and continue to the next message.
		return nil, fmt.Errorf("failed to convert signals: %w", err)
	}

	return convertErr.DecodedSignals, convertErr
}

// CloudEventConvert converts a message to cloud events.
func (m Module) CloudEventConvert(_ context.Context, msgData []byte) ([]cloudevent.CloudEventHeader, []byte, error) {
	var event RuptelaEvent
	err := json.Unmarshal(msgData, &event)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal record data: %w", err)
	}
	if event.DeviceTokenID == nil {
		return nil, nil, fmt.Errorf("device token id is missing")
	}
	// Construct the producer DID
	producer := cloudevent.ERC721DID{
		ChainID:         m.ChainID,
		ContractAddress: m.AftermarketContractAddr,
		TokenID:         big.NewInt(int64(*event.DeviceTokenID)),
	}.String()
	subject, err := m.determineSubject(&event, producer)
	if err != nil {
		return nil, nil, err
	}
	cloudEventTypes, err := getCloudEventTypes(&event)
	if err != nil {
		return nil, nil, err
	}
	hdrs := make([]cloudevent.CloudEventHeader, 0, len(cloudEventTypes))
	for _, cloudEventType := range cloudEventTypes {
		cloudEventHdr := createCloudEventHdr(&event, producer, subject, cloudEventType)
		hdrs = append(hdrs, cloudEventHdr)
	}

	return hdrs, event.Data, nil
}

// EventConvert converts a message to events.
func (*Module) EventConvert(_ context.Context, event cloudevent.RawEvent) ([]vss.Event, error) {
	return DecodeEvent(event)
}

// determineSubject determines the subject of the cloud event based on the DS type.
func (m Module) determineSubject(event *RuptelaEvent, producer string) (string, error) {
	var subject string
	switch event.DS {
	case StatusEventDS, LocationEventDS, DTCEventDS:
		if event.VehicleTokenID != nil {
			subject = cloudevent.ERC721DID{
				ChainID:         m.ChainID,
				ContractAddress: m.VehicleContractAddr,
				TokenID:         big.NewInt(int64(*event.VehicleTokenID)),
			}.String()
		}
	case DevStatusDS, BattDS:
		subject = producer
	default:
		return "", fmt.Errorf("unknown DS type: %s", event.DS)
	}
	return subject, nil
}

// createCloudEvent creates a cloud event from a ruptela event.
func createCloudEventHdr(event *RuptelaEvent, producer, subject, eventType string) cloudevent.CloudEventHeader {
	return cloudevent.CloudEventHeader{
		DataContentType: "application/json",
		ID:              ksuid.New().String(),
		Subject:         subject,
		Source:          "dimo/integration/2lcaMFuCO0HJIUfdq8o780Kx5n3",
		SpecVersion:     "1.0",
		Time:            event.Time,
		Type:            eventType,
		DataVersion:     event.DS,
		Producer:        producer,
		Signature:       event.Signature,
	}
}

// getCloudEventTypes gets the cloud event types contained in the ruptela event.
func getCloudEventTypes(event *RuptelaEvent) ([]string, error) {
	// always include the status event
	cloudEventTypes := []string{cloudevent.TypeStatus}
	if event.DS != StatusEventDS {
		return cloudEventTypes, nil
	}
	var dataContent DataContent
	err := json.Unmarshal(event.Data, &dataContent)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}
	if checkVINPresenceInPayload(event, dataContent.Signals) {
		cloudEventTypes = append(cloudEventTypes, cloudevent.TypeFingerprint)
	}
	if checkEventPresenceInPayload(event, dataContent.Signals) {
		cloudEventTypes = append(cloudEventTypes, cloudevent.TypeEvent)
	}
	return cloudEventTypes, nil
}

// checkVINPresenceInPayload checks if the VIN is present in the payload.
func checkVINPresenceInPayload(event *RuptelaEvent, dataMap map[string]string) bool {
	if event.DS != StatusEventDS {
		return false
	}

	// VIN keys in the ruptela payload including CAN IO's as a backup(smart5 devices)
	vinKeys := []string{"104", "105", "106", "123", "124", "125"}

	for _, key := range vinKeys {
		value, ok := dataMap[key]
		// ruptela smart5 sends empty vin as long 0s
		if ok && (value != "0" && value != "0000000000000000") {
			// key has non-zero value
			return true
		}
	}
	return false
}

// checkEventPresenceInPayload checks if the event is present in the payload.
func checkEventPresenceInPayload(event *RuptelaEvent, dataMap map[string]string) bool {
	if event.DS != StatusEventDS {
		return false
	}
	eventKeys := []string{"135", "136", "143"}
	for _, key := range eventKeys {
		value, ok := dataMap[key]
		if ok && value != "0" {
			return true
		}
	}
	return false
}
