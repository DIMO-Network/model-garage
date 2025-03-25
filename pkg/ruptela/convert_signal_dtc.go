package ruptela

import (
	"fmt"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/vss"
)

// SignalsFromDTCPayload gets a slice signals from a dtc payload.
func SignalsFromDTCPayload(event cloudevent.RawEvent) ([]vss.Signal, error) {
	did, err := cloudevent.DecodeNFTDID(event.Subject)
	if err != nil {
		return nil, fmt.Errorf("failed to decode subject DID: %w", err)
	}
	dtcValue, errs := OBDDTCListFromV1Data(event.Data)

	dtcSignal := vss.Signal{
		TokenID:   did.TokenID,
		Timestamp: event.Time,
		Source:    event.Source,
		Name:      vss.FieldOBDDTCList,
	}
	dtcSignal.SetValue(dtcValue)

	if errs != nil {
		return nil, convert.ConversionError{
			TokenID:        did.TokenID,
			Source:         event.Source,
			DecodedSignals: []vss.Signal{dtcSignal},
			Errors:         []error{fmt.Errorf("error getting obdDTCList: %w", errs)},
		}
	}
	return []vss.Signal{dtcSignal}, nil
}
