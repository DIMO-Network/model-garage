package compass

import (
	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/vss"
)

// DecodeSignals decodes a compass message into signals.
func DecodeSignals(ce cloudevent.RawEvent) ([]vss.Signal, error) {
	baseSignal := vss.Signal{
		Subject:   ce.Subject,
		Source:    ce.Source,
		Timestamp: ce.Time,
	}

	sigs, errs := SignalsFromCompass(baseSignal, ce.Data)
	if len(errs) != 0 {
		return nil, convert.ConversionError{
			Subject:        ce.Subject,
			Source:         ce.Source,
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
