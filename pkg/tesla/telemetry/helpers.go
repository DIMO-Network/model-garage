package telemetry

// windowStateToIsOpen converts the Tesla WindowState enum, which we typically receive
// as a string, to 1 (open) or 0 (closed).
// See https://github.com/teslamotors/fleet-telemetry/blob/646fce2fb2ddd607ce4e76c865ce411e32ded81f/protos/vehicle_data.proto#L465
func windowStateToIsOpen(state string) float64 {
	if state == "PartiallyOpen" || state == "Opened" {
		return 1
	}
	return 0
}
