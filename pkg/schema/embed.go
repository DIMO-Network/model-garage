package schema

import (
	_ "embed"
)

//go:embed spec/vss_rel_4.2-DIMO-*.csv
var vssRel42DIMO string

//go:embed spec/default-definitions.yaml
var defaultDefinitionsYAML string

//go:embed spec/default-event-tags.yaml
var defaultEventTagsYAML string

// VssRel42DIMO is the embedded CSV file containing the VSS schema for DIMO.
func VssRel42DIMO() string {
	return vssRel42DIMO
}

// DefaultDefinitionsYAML is the embedded YAML file containing information about what signals will be displayed and used by the DIMO Node.
func DefaultDefinitionsYAML() string {
	return defaultDefinitionsYAML
}

// DefaultEventTagsYAML is the embedded YAML file containing information about event tags.
func DefaultEventTagsYAML() string {
	return defaultEventTagsYAML
}
