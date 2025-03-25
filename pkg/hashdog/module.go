package hashdog

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/ethereum/go-ethereum/common"
)

// Module is a module that converts Macaron messages to signals.
type Module struct {
	ChainID                 uint64         `json:"chain_id"`
	AftermarketContractAddr common.Address `json:"aftermarket_contract_addr"`
	VehicleContractAddr     common.Address `json:"vehicle_contract_addr"`
}
type HashdogPayload struct {
	DecodedPayload cloudevent.Fingerprint `json:"decodedPayload"`
}

// FingerprintConvert converts a message to a fingerprint.
func (*Module) FingerprintConvert(_ context.Context, event cloudevent.RawEvent) (cloudevent.Fingerprint, error) {
	var fpData HashdogPayload
	err := json.Unmarshal(event.Data, &fpData)
	if err != nil {
		return cloudevent.Fingerprint{}, fmt.Errorf("failed to unmarshal hashdog fingerprint data: %w", err)
	}
	return fpData.DecodedPayload, nil
}

// SignalConvert converts a message to signals.
func (*Module) SignalConvert(_ context.Context, event cloudevent.RawEvent) ([]vss.Signal, error) {
	if event.Producer == event.Subject {
		return nil, nil
	}
	signals, err := SignalsFromV2Payload(event)
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
	return ConvertToCloudEvents(msgData, m.ChainID, m.AftermarketContractAddr, m.VehicleContractAddr)
}
