package compass

import (
	"fmt"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/vss"
)

// DecodeSignals decodes a compass message into signals.
func DecodeSignals(ce cloudevent.RawEvent) ([]vss.Signal, error) {
	did, err := cloudevent.DecodeERC721DID(ce.Subject)
	if err != nil {
		return nil, fmt.Errorf("failed to decode subject DID: %w", err)
	}

	tokenID := did.TokenID
	source := ce.Source

	baseSignal := vss.Signal{
		TokenID:   uint32(tokenID.Uint64()), //nolint:gosec // will not exceed uint32 max value
		Source:    source,
		Timestamp: ce.Time,
	}

	sigs, errs := SignalsFromCompass(baseSignal, ce.Data)
	if len(errs) != 0 {
		return nil, convert.ConversionError{
			TokenID:        uint32(tokenID.Uint64()), //nolint:gosec // will not exceed uint32 max value
			Source:         source,
			DecodedSignals: sigs,
			Errors:         errs,
		}
	}

	return sigs, nil
}

// ConvertPSIToKPa converts a pressure value from psi to kPa.
func ConvertPSIToKPa(psi float64) float64 {
	return psi * 6.89476
}

// ConvertBarToKPa converts a pressure value from bar to kPa.
func ConvertBarToKPa(bar float64) float64 {
	return bar * 100
}
