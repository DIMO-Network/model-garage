package status

import (
	"fmt"

	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/ruptela"
	"github.com/DIMO-Network/model-garage/pkg/vss"
)

// SignalsFromDTCPayload gets a slice signals from a dtc payload.
func SignalsFromDTCPayload(jsonData []byte) ([]vss.Signal, error) {
	ts, err := TimestampFromV1Data(jsonData)
	if err != nil {
		return nil, convert.ConversionError{
			Errors: []error{fmt.Errorf("error getting timestamp: %w", err)},
		}
	}
	tokenID, err := TokenIDFromData(jsonData)
	if err != nil {
		return nil, convert.ConversionError{
			Errors: []error{fmt.Errorf("error getting tokenId: %w", err)},
		}
	}
	source, err := SourceFromData(jsonData)
	if err != nil {
		return nil, convert.ConversionError{
			TokenID: tokenID,
			Errors:  []error{fmt.Errorf("error getting source: %w", err)},
		}
	}

	dtcValue, errs := ruptela.OBDDTCListFromV1Data(jsonData)

	dtcSignal := vss.Signal{
		TokenID:   tokenID,
		Timestamp: ts,
		Source:    source,
		Name:      vss.FieldOBDDTCList,
	}
	dtcSignal.SetValue(dtcValue)

	if errs != nil {
		return nil, convert.ConversionError{
			TokenID:        tokenID,
			Source:         source,
			DecodedSignals: []vss.Signal{dtcSignal},
			Errors:         []error{fmt.Errorf("error getting obdDTCList: %w", errs)},
		}
	}
	return []vss.Signal{dtcSignal}, nil
}
