package ruptela

import (
	"errors"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/tidwall/gjson"
)

// SignalsFromV1Payload gets a slice signals from a v1 payload.
//
// Position signals are skipped for records generated while the device was offline (see
// generatedOffline): their position predates their timestamp by minutes or hours.
func SignalsFromV1Payload(event cloudevent.RawEvent) ([]vss.Signal, error) {
	baseSignal := vss.Signal{
		CloudEventHeader: event.CloudEventHeader,
		Data: vss.SignalData{
			Timestamp: event.Time,
		},
	}
	hdr := event.CloudEventHeader
	hdr.Type = cloudevent.TypeSignal

	offline := generatedOffline(event.Data)
	sigDatas, errs := SignalsFromV1Data(baseSignal, event.Data)
	var sigs []vss.Signal
	for _, sd := range sigDatas {
		if offline && isPositionSignal(sd.Name) {
			continue
		}
		sigs = append(sigs, vss.Signal{CloudEventHeader: hdr, Data: sd})
	}
	if !offline {
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

// generatedOffline reports whether a status record was generated while the device had no data
// connection, from element 418 (GPRS status: 0 – disconnected, 1 – connected), which is captured
// with the record. The R1 buffers such records and sends them after it reconnects, stamped with the
// send time rather than the capture time, so their position is from minutes or hours earlier. The
// payload holds no absolute capture time to correct it: its only generation clock is element 525,
// seconds since the device booted.
func generatedOffline(jsonData []byte) bool {
	return gjson.GetBytes(jsonData, "signals.418").String() == "0"
}

// isPositionSignal reports whether a status signal is derived from the pos block.
func isPositionSignal(name string) bool {
	switch name {
	case vss.FieldCurrentLocationCoordinates, vss.FieldCurrentLocationAltitude, vss.FieldCurrentLocationHeading:
		return true
	}
	return false
}

// currentLocationCoordinatesFromV1Data extracts a vss.Location from v1 JSON data.
func currentLocationCoordinatesFromV1Data(jsonData []byte) (vss.Location, error) {
	pos := gjson.GetBytes(jsonData, "pos")
	if !pos.Exists() {
		return vss.Location{}, errNotFound
	}
	return posToLocation(pos, statusHDOPScale)
}
