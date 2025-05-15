package compass

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/segmentio/ksuid"
)

type compassEvent struct {
	Time           string          `json:"time"`
	Data           json.RawMessage `json:"data"`
	VehicleTokenID *uint32         `json:"vehicleTokenId"`
	DeviceTokenID  *uint32         `json:"deviceTokenId"`
}

// Module is a module that converts compass messages to signals.
type Module struct {
	ChainID             uint64         `json:"chain_id"`
	SynthContractAddr   common.Address `json:"synth_contract_addr"`
	VehicleContractAddr common.Address `json:"vehicle_contract_addr"`
}

// FingerprintConvert converts a message to a fingerprint.
func (*Module) FingerprintConvert(_ context.Context, event cloudevent.RawEvent) (cloudevent.Fingerprint, error) {
	return DecodeFingerprint(event)
}

// SignalConvert converts a message to signals.
func (*Module) SignalConvert(_ context.Context, event cloudevent.RawEvent) ([]vss.Signal, error) {
	if event.Type != cloudevent.TypeStatus || event.Producer == event.Subject {
		return nil, nil
	}
	signals, err := DecodeSignals(event)
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
	var event compassEvent
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
		ContractAddress: m.SynthContractAddr,
		TokenID:         big.NewInt(int64(*event.DeviceTokenID)),
	}.String()

	// Construct the subject
	var subject string
	if event.VehicleTokenID != nil {
		subject = cloudevent.ERC721DID{
			ChainID:         m.ChainID,
			ContractAddress: m.VehicleContractAddr,
			TokenID:         big.NewInt(int64(*event.VehicleTokenID)),
		}.String()
	}

	statusHdr, err := createCloudEventHdr(&event, producer, subject, cloudevent.TypeStatus)
	if err != nil {
		return nil, nil, err
	}
	eventHeaders := []cloudevent.CloudEventHeader{statusHdr}

	_, err = DecodeFingerprint(cloudevent.RawEvent{Data: event.Data})
	isVinPresent := err == nil

	if isVinPresent {
		fpHdr, errCE := createCloudEventHdr(&event, producer, subject, cloudevent.TypeFingerprint)
		if errCE != nil {
			return nil, nil, errCE
		}
		eventHeaders = append(eventHeaders, fpHdr)
	}

	return eventHeaders, event.Data, nil
}

// createCloudEvent creates a cloud event from a compass event.
func createCloudEventHdr(event *compassEvent, producer, subject, eventType string) (cloudevent.CloudEventHeader, error) {
	timeValue, err := time.Parse(time.RFC3339, event.Time)
	if err != nil {
		return cloudevent.CloudEventHeader{}, fmt.Errorf("failed to parse time: %w", err)
	}
	return cloudevent.CloudEventHeader{
		DataContentType: "application/json",
		ID:              ksuid.New().String(),
		Subject:         subject,
		Source:          "dimo/integration/compass",
		SpecVersion:     "1.0",
		Time:            timeValue,
		Type:            eventType,
		DataVersion:     "1.0",
		Producer:        producer,
	}, nil
}
