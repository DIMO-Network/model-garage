package tesla

const (
	FleetTelemetryDataVersion = "fleet_telemetry/v1.0.0"
)

type TelemetryData struct {
	Payloads [][]byte `json:"payloads"`
}
