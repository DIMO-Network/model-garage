// Package status converts Tesla CloudEvents to ClickHouse-ready slices of signals.
package status

import (
	"encoding/json"
	"fmt"

	"github.com/DIMO-Network/model-garage/pkg/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/tesla"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	ftconv "github.com/DIMO-Network/tesla-vss/pkg/convert"
	"github.com/teslamotors/fleet-telemetry/protos"
	"google.golang.org/protobuf/proto"
)

func Decode(msgBytes []byte) ([]vss.Signal, error) {
	// Only interested in the top-level CloudEvent fields.
	var ce cloudevent.CloudEvent[json.RawMessage]

	if err := json.Unmarshal(msgBytes, &ce); err != nil {
		return nil, fmt.Errorf("failed to unmarshal payload: %w", err)
	}

	did, err := cloudevent.DecodeNFTDID(ce.Subject)
	if err != nil {
		return nil, fmt.Errorf("failed to decode subject DID: %w", err)
	}

	tokenID := did.TokenID
	source := ce.Source

	baseSignal := vss.Signal{
		TokenID: tokenID,
		Source:  source,
	}

	switch ce.DataVersion {
	case tesla.FleetTelemetryDataVersion:
		var td tesla.TelemetryData
		if err := json.Unmarshal(ce.Data, &td); err != nil {
			return nil, fmt.Errorf("failed to unmarshal telemetry wrapper: %w", err)
		}

		var batchedSigs []vss.Signal
		var batchedErrs []error
		for i, payload := range td.Payloads {
			var pl protos.Payload
			err := proto.Unmarshal(payload, &pl)
			if err != nil {
				batchedErrs = append(batchedErrs, fmt.Errorf("failed to unmarshal payload at index %d: %w", i, err))
				continue
			}
			sigs, errs := ftconv.ProcessPayload(&pl, tokenID, source)
			batchedSigs = append(batchedSigs, sigs...)
			batchedErrs = append(batchedErrs, errs...)
		}

		if len(batchedErrs) != 0 {
			return nil, convert.ConversionError{
				TokenID:        tokenID,
				Source:         source,
				DecodedSignals: batchedSigs,
				Errors:         batchedErrs,
			}
		}
		return batchedSigs, nil
	default:
		sigs, errs := tesla.SignalsFromTesla(baseSignal, msgBytes)
		if len(errs) != 0 {
			return nil, convert.ConversionError{
				TokenID:        tokenID,
				Source:         source,
				DecodedSignals: sigs,
				Errors:         errs,
			}
		}

		return sigs, nil
	}
}
