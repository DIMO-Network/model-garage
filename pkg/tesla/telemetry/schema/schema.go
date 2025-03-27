package schema

import (
	_ "embed"
)

//go:embed telemetry_definitions.yaml
var schema string

func TelemetryDefinitionsYAML() string {
	return schema
}
