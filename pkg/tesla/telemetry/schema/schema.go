// Package schema merely embeds the definitions file for conversions
// from Tesla Fleet Telemetry to VSS signals.
package schema

import (
	_ "embed"
)

//go:embed telemetry_definitions.yaml
var schema string

func TelemetryDefinitionsYAML() string {
	return schema
}
