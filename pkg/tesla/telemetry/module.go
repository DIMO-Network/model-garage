// Package telemetry converts batches of Tesla protobuf Payloads into VSS signals.
package telemetry

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/teslamotors/fleet-telemetry/protos"
	"github.com/tidwall/gjson"
	"google.golang.org/protobuf/proto"
)

const DataVersion = "fleet_telemetry/v1.0.0"

type TelemetryData struct {
	Payloads [][]byte `json:"payloads"`
}

// SignalConvert converts a CloudEvent containing a batch of Fleet Telemetry
// protobuf Payloads into DIMO's VSS rows.
func SignalConvert(event cloudevent.RawEvent) ([]vss.Signal, error) {
	did, err := cloudevent.DecodeERC721DID(event.Subject)
	if err != nil {
		return nil, fmt.Errorf("failed to decode subject DID: %w", err)
	}

	tokenID := uint32(did.TokenID.Uint64()) //nolint:gosec // will not exceed uint32 max value
	source := event.Source

	var td TelemetryData
	if err := json.Unmarshal(event.Data, &td); err != nil {
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
		sigs, errs := ProcessPayload(&pl, tokenID, source)
		batchedSigs = append(batchedSigs, sigs...)
		batchedErrs = append(batchedErrs, errs...)
	}

	if len(batchedErrs) != 0 {
		return nil, convert.ConversionError{
			TokenID:        tokenID,
			DecodedSignals: batchedSigs,
			Errors:         batchedErrs,
		}
	}
	return batchedSigs, nil
}

// IsFingerprint returns whether the Fleet Telemetry batch contains an extractable
// VIN. This should always return true.
func IsFingerprint(event cloudevent.RawEvent) bool {
	result := gjson.GetBytes(event.Data, "payloads")

	// This check should always pass.
	return result.Exists() && result.IsArray() && len(result.Array()) != 0
}

// FingerprintConvert extracts a fingerprint from the first Fleet Telemetry protobuf
// Payload. We expect this to always succeed.
func FingerprintConvert(event cloudevent.RawEvent) (cloudevent.Fingerprint, error) {
	var fp cloudevent.Fingerprint

	var tlmData TelemetryData
	if err := json.Unmarshal(event.Data, &tlmData); err != nil {
		return fp, fmt.Errorf("failed to unmarshal Telemetry payload: %w", err)
	}

	if len(tlmData.Payloads) == 0 {
		return fp, errors.New("no payloads")
	}

	var pl protos.Payload
	if err := proto.Unmarshal(tlmData.Payloads[0], &pl); err != nil {
		return fp, fmt.Errorf("failed to unmarshal tesla payload at index 0: %w", err)
	}

	fp.VIN = pl.Vin

	return fp, nil
}
