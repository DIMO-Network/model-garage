package defaultmodule

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/schema"
	"github.com/DIMO-Network/model-garage/pkg/vss"
)

// SignalConvert converts a default CloudEvent to DIMO's vss signals.
func SignalConvert(event cloudevent.RawEvent, signalMap map[string]*schema.SignalInfo) ([]vss.Signal, error) {
	var sigData SignalData
	err := json.Unmarshal(event.Data, &sigData)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal signal data: %w", err)
	}
	var decodeErrs error
	vssSignals := make([]vss.Signal, 0)
	for _, signal := range sigData.Signals {
		vssSig, err := defaultSignalToVSS(signal, signalMap)
		if err != nil {
			// we want to return decoded signals even if some fail
			decodeErrs = errors.Join(decodeErrs, err)
			continue
		}
		vssSignals = append(vssSignals, vssSig)
	}
	return vssSignals, decodeErrs
}

func defaultSignalToVSS(signal *Signal, signalMap map[string]*schema.SignalInfo) (vss.Signal, error) {
	signalInfo, ok := signalMap[signal.Name]
	if !ok {
		return vss.Signal{}, fmt.Errorf("signal %s is not a defined signal name", signal.Name)
	}
	if signal.Value == nil {
		return vss.Signal{}, fmt.Errorf("signal %s is missing a value", signal.Name)
	}
	vssSig := vss.Signal{
		Timestamp: signal.Timestamp,
		Name:      signal.Name,
	}
	switch signalInfo.BaseGoType {
	case "float64":
		num, ok := signal.Value.(float64)
		if ok {
			vssSig.ValueNumber = num
		} else if str, ok := signal.Value.(string); ok {
			v, err := strconv.ParseFloat(str, 64)
			if err != nil {
				return vss.Signal{}, fmt.Errorf("signal %s can not be converted to a float64: %w", signal.Name, err)
			}
			vssSig.ValueNumber = v
		} else {
			return vss.Signal{}, fmt.Errorf("signal %s is not a float64", signal.Name)
		}
	case "string":
		str, ok := signal.Value.(string)
		if !ok {
			return vss.Signal{}, fmt.Errorf("signal %s is not a string", signal.Name)
		}
		vssSig.ValueString = str
	case "vss.Location":
		// This gets unmarshaled into type any.
		// This is ugly. Should we do more sanity checks here? Maybe we don't want (lat, 0, 0).
		m, ok := signal.Value.(map[string]any)
		if !ok {
			return vss.Signal{}, fmt.Errorf("signal %s is not a location object", signal.Name)
		}
		var loc vss.Location
		if v, exists := m["latitude"]; exists {
			loc.Latitude, ok = v.(float64)
			if !ok {
				return vss.Signal{}, fmt.Errorf("signal %s has a non-float64 latitude", signal.Name)
			}
		}
		if v, exists := m["longitude"]; exists {
			loc.Longitude, ok = v.(float64)
			if !ok {
				return vss.Signal{}, fmt.Errorf("signal %s has a non-float64 longitude", signal.Name)
			}
		}
		if v, exists := m["hdop"]; exists {
			loc.HDOP, ok = v.(float64)
			if !ok {
				return vss.Signal{}, fmt.Errorf("signal %s has a non-float64 hdop", signal.Name)
			}
		}
		vssSig.ValueLocation = loc
	default:
		return vss.Signal{}, fmt.Errorf("signal %s has an unsupported base type %s", signal.Name, signalInfo.BaseGoType)
	}

	return vssSig, nil
}
