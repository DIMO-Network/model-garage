// Package parse contains functions for parsing typed Tesla values out
// of strings. This is necessary to parse telemetry signals from
// vehicles with outdated firmware.
package parse

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/teslamotors/fleet-telemetry/protos"
)

func Int32(s string) (int32, error) {
	d, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("couldn't parse string into int32: %w", err)
	}
	return int32(d), nil
}

func Double(s string) (float64, error) {
	d, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("couldn't parse string into double: %w", err)
	}
	return d, nil
}

func Doors(s string) (*protos.Doors, error) {
	var out protos.Doors
	if s == "" {
		return &out, nil
	}
	// TODO(elffjs): Some kind of warning if it's none of these?
	for element := range strings.SplitSeq(s, "|") {
		switch element {
		case "DriverFront":
			out.DriverFront = true
		case "DriverRear":
			out.DriverRear = true
		case "PassengerFront":
			out.PassengerFront = true
		case "PassengerRear":
			out.PassengerRear = true
		case "TrunkFront":
			out.TrunkFront = true
		case "TrunkRear":
			out.TrunkRear = true
		}
	}
	return &out, nil
}

func WindowState(s string) (protos.WindowState, error) {
	// TODO(elffjs): Some kind of warning if it's none of these?
	// Have never seen Unknown in practice.
	switch s {
	case "Closed":
		return protos.WindowState_WindowStateClosed, nil
	case "PartiallyOpen":
		return protos.WindowState_WindowStatePartiallyOpen, nil
	case "Opened":
		return protos.WindowState_WindowStateOpened, nil
	default:
		return protos.WindowState_WindowStateUnknown, nil
	}
}
