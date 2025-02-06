// Package status converts Compass CloudEvents to ClickHouse-ready slices of signals.
package status

import (
	"encoding/json"
	"fmt"
	"github.com/DIMO-Network/model-garage/pkg/compass"

	"github.com/DIMO-Network/model-garage/pkg/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/vss"
)

func Decode(msgBytes []byte) ([]vss.Signal, error) {
	// Only interested in the top-level CloudEvent fields.
	var ce cloudevent.CloudEventHeader

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

	sigs, errs := compass.SignalsFromV1Data(baseSignal, msgBytes)
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
