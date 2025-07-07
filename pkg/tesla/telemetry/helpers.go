package telemetry

import "github.com/teslamotors/fleet-telemetry/protos"

// windowStateToIsOpen converts the Tesla WindowState enum, which we typically receive
// as a string, to 1 (open) or 0 (closed).
// See https://github.com/teslamotors/fleet-telemetry/blob/646fce2fb2ddd607ce4e76c865ce411e32ded81f/protos/vehicle_data.proto#L465
func windowStateToIsOpen(s protos.WindowState) float64 {
	switch s {
	case protos.WindowState_WindowStatePartiallyOpen, protos.WindowState_WindowStateOpened:
		return 1
	default:
		return 0
	}
}

func boolToFloat64(b bool) float64 {
	if b {
		return 1
	}
	return 0
}
