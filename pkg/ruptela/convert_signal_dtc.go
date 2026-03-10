package ruptela

import (
	"fmt"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/vss"
)

// SignalsFromDTCPayload gets a slice signals from a dtc payload.
func SignalsFromDTCPayload(event cloudevent.RawEvent) ([]vss.Signal, error) {
	dtcValue, errs := OBDDTCListFromV1Data(event.Data)

	hdr := event.CloudEventHeader
	hdr.Type = cloudevent.TypeSignal
	dtcSignal := vss.Signal{
		CloudEventHeader: hdr,
		Data: vss.SignalData{
			Timestamp:    event.Time,
			Name:         vss.FieldOBDDTCList,
			CloudEventID: event.ID,
		},
	}
	dtcSignal.Data.SetValue(dtcValue)

	if errs != nil {
		return nil, convert.ConversionError{
			Subject:        event.Subject,
			Source:         event.Source,
			DecodedSignals: []vss.Signal{dtcSignal},
			Errors:         []error{fmt.Errorf("error getting obdDTCList: %w", errs)},
		}
	}
	return []vss.Signal{dtcSignal}, nil
}
