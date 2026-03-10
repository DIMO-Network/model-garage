package ruptela

import (
	"errors"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/tidwall/gjson"
)

// SignalsFromV1Payload gets a slice signals from a v1 payload.
func SignalsFromV1Payload(event cloudevent.RawEvent) ([]vss.Signal, error) {
	baseSignal := vss.Signal{
		CloudEventHeader: event.CloudEventHeader,
		Data: vss.SignalData{
			Timestamp: event.Time,
		},
	}
	hdr := event.CloudEventHeader
	hdr.Type = cloudevent.TypeSignal

	sigDatas, errs := SignalsFromV1Data(baseSignal, event.Data)
	var sigs []vss.Signal
	for _, sd := range sigDatas {
		sigs = append(sigs, vss.Signal{CloudEventHeader: hdr, Data: sd})
	}
	if coordLoc, err := currentLocationCoordinatesFromV1Data(event.Data); err == nil {
		sd := vss.SignalData{
			Name:         vss.FieldCurrentLocationCoordinates,
			Timestamp:    baseSignal.Data.Timestamp,
			CloudEventID: event.ID,
		}
		sd.SetValue(coordLoc)
		sigs = append(sigs, vss.Signal{CloudEventHeader: hdr, Data: sd})
	} else if !errors.Is(err, errNotFound) {
		errs = append(errs, err)
	}
	if errs != nil {
		return nil, convert.ConversionError{
			Subject:        event.Subject,
			Source:         event.Source,
			DecodedSignals: sigs,
			Errors:         errs,
		}
	}
	return sigs, nil
}

// currentLocationCoordinatesFromV1Data extracts a vss.Location from v1 JSON data.
func currentLocationCoordinatesFromV1Data(jsonData []byte) (vss.Location, error) {
	pos := gjson.GetBytes(jsonData, "pos")
	if !pos.Exists() {
		return vss.Location{}, errNotFound
	}
	return posToLocation(pos)
}
