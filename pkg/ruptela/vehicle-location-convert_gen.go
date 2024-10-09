// Code generated by github.com/DIMO-Network/model-garage DO NOT EDIT.
package ruptela

import (
	"errors"
	"fmt"

	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/tidwall/gjson"
)

// SignalsFromLocationData converts the given JSON data to a slice of signals.
func SignalsFromLocationData(originalDoc []byte, baseSignal vss.Signal, signalName string, valResult gjson.Result) ([]vss.Signal, error) {
	ret := make([]vss.Signal, 0)
	var retErrs error

	switch signalName {

	case "alt":
		val0, err := CurrentLocationAltitudeFromLocationData(originalDoc, valResult)
		if err != nil {
			retErrs = errors.Join(retErrs, fmt.Errorf("failed to convert 'pos.alt': %w", err))
		} else {
			sig := vss.Signal{
				TokenID:   baseSignal.TokenID,
				Timestamp: baseSignal.Timestamp,
				Source:    baseSignal.Source,
				Name:      "currentLocationAltitude",
			}
			sig.SetValue(val0)
			ret = append(ret, sig)
		}

	case "hdop":
		val0, err := DIMOAftermarketHDOPFromLocationData(originalDoc, valResult)
		if err != nil {
			retErrs = errors.Join(retErrs, fmt.Errorf("failed to convert 'pos.hdop': %w", err))
		} else {
			sig := vss.Signal{
				TokenID:   baseSignal.TokenID,
				Timestamp: baseSignal.Timestamp,
				Source:    baseSignal.Source,
				Name:      "dimoAftermarketHDOP",
			}
			sig.SetValue(val0)
			ret = append(ret, sig)
		}

	case "lat":
		val0, err := CurrentLocationLatitudeFromLocationData(originalDoc, valResult)
		if err != nil {
			retErrs = errors.Join(retErrs, fmt.Errorf("failed to convert 'pos.lat': %w", err))
		} else {
			sig := vss.Signal{
				TokenID:   baseSignal.TokenID,
				Timestamp: baseSignal.Timestamp,
				Source:    baseSignal.Source,
				Name:      "currentLocationLatitude",
			}
			sig.SetValue(val0)
			ret = append(ret, sig)
		}

	case "lon":
		val0, err := CurrentLocationLongitudeFromLocationData(originalDoc, valResult)
		if err != nil {
			retErrs = errors.Join(retErrs, fmt.Errorf("failed to convert 'pos.lon': %w", err))
		} else {
			sig := vss.Signal{
				TokenID:   baseSignal.TokenID,
				Timestamp: baseSignal.Timestamp,
				Source:    baseSignal.Source,
				Name:      "currentLocationLongitude",
			}
			sig.SetValue(val0)
			ret = append(ret, sig)
		}

	case "sat":
		val0, err := DIMOAftermarketNSATFromLocationData(originalDoc, valResult)
		if err != nil {
			retErrs = errors.Join(retErrs, fmt.Errorf("failed to convert 'pos.sat': %w", err))
		} else {
			sig := vss.Signal{
				TokenID:   baseSignal.TokenID,
				Timestamp: baseSignal.Timestamp,
				Source:    baseSignal.Source,
				Name:      "dimoAftermarketNSAT",
			}
			sig.SetValue(val0)
			ret = append(ret, sig)
		}

	case "spd":
		val0, err := SpeedFromLocationData(originalDoc, valResult)
		if err != nil {
			retErrs = errors.Join(retErrs, fmt.Errorf("failed to convert 'pos.spd': %w", err))
		} else {
			sig := vss.Signal{
				TokenID:   baseSignal.TokenID,
				Timestamp: baseSignal.Timestamp,
				Source:    baseSignal.Source,
				Name:      "speed",
			}
			sig.SetValue(val0)
			ret = append(ret, sig)
		}

	default:
		// do nothing
	}
	return ret, retErrs
}

// CurrentLocationAltitudeFromLocationData converts the given JSON data to a float64.
func CurrentLocationAltitudeFromLocationData(originalDoc []byte, result gjson.Result) (ret float64, err error) {
	var errs error
	val0, ok := result.Value().(float64)
	if ok {
		ret, err = ToCurrentLocationAltitude0(originalDoc, val0)
		if err == nil {
			return ret, nil
		}
		errs = errors.Join(errs, fmt.Errorf("failed to convert 'pos.alt': %w", err))
	} else {
		errs = errors.Join(errs, fmt.Errorf("%w, field 'pos.alt' is not of type 'float64' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
	}

	return ret, errs
}

// CurrentLocationLatitudeFromLocationData converts the given JSON data to a float64.
func CurrentLocationLatitudeFromLocationData(originalDoc []byte, result gjson.Result) (ret float64, err error) {
	var errs error
	val0, ok := result.Value().(float64)
	if ok {
		ret, err = ToCurrentLocationLatitude0(originalDoc, val0)
		if err == nil {
			return ret, nil
		}
		errs = errors.Join(errs, fmt.Errorf("failed to convert 'pos.lat': %w", err))
	} else {
		errs = errors.Join(errs, fmt.Errorf("%w, field 'pos.lat' is not of type 'float64' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
	}

	return ret, errs
}

// CurrentLocationLongitudeFromLocationData converts the given JSON data to a float64.
func CurrentLocationLongitudeFromLocationData(originalDoc []byte, result gjson.Result) (ret float64, err error) {
	var errs error
	val0, ok := result.Value().(float64)
	if ok {
		ret, err = ToCurrentLocationLongitude0(originalDoc, val0)
		if err == nil {
			return ret, nil
		}
		errs = errors.Join(errs, fmt.Errorf("failed to convert 'pos.lon': %w", err))
	} else {
		errs = errors.Join(errs, fmt.Errorf("%w, field 'pos.lon' is not of type 'float64' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
	}

	return ret, errs
}

// DIMOAftermarketHDOPFromLocationData converts the given JSON data to a float64.
func DIMOAftermarketHDOPFromLocationData(originalDoc []byte, result gjson.Result) (ret float64, err error) {
	var errs error
	val0, ok := result.Value().(float64)
	if ok {
		ret, err = ToDIMOAftermarketHDOP0(originalDoc, val0)
		if err == nil {
			return ret, nil
		}
		errs = errors.Join(errs, fmt.Errorf("failed to convert 'pos.hdop': %w", err))
	} else {
		errs = errors.Join(errs, fmt.Errorf("%w, field 'pos.hdop' is not of type 'float64' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
	}

	return ret, errs
}

// DIMOAftermarketNSATFromLocationData converts the given JSON data to a float64.
func DIMOAftermarketNSATFromLocationData(originalDoc []byte, result gjson.Result) (ret float64, err error) {
	var errs error
	val0, ok := result.Value().(float64)
	if ok {
		ret, err = ToDIMOAftermarketNSAT0(originalDoc, val0)
		if err == nil {
			return ret, nil
		}
		errs = errors.Join(errs, fmt.Errorf("failed to convert 'pos.sat': %w", err))
	} else {
		errs = errors.Join(errs, fmt.Errorf("%w, field 'pos.sat' is not of type 'float64' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
	}

	return ret, errs
}

// SpeedFromLocationData converts the given JSON data to a float64.
func SpeedFromLocationData(originalDoc []byte, result gjson.Result) (ret float64, err error) {
	var errs error
	val0, ok := result.Value().(string)
	if ok {
		ret, err = ToSpeed0(originalDoc, val0)
		if err == nil {
			return ret, nil
		}
		errs = errors.Join(errs, fmt.Errorf("failed to convert 'signals.95': %w", err))
	} else {
		errs = errors.Join(errs, fmt.Errorf("%w, field 'signals.95' is not of type 'string' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
	}
	val1, ok := result.Value().(float64)
	if ok {
		ret, err = ToSpeed1(originalDoc, val1)
		if err == nil {
			return ret, nil
		}
		errs = errors.Join(errs, fmt.Errorf("failed to convert 'pos.spd': %w", err))
	} else {
		errs = errors.Join(errs, fmt.Errorf("%w, field 'pos.spd' is not of type 'float64' got '%v' of type '%T'", convert.InvalidTypeError(), result.Value(), result.Value()))
	}

	return ret, errs
}
