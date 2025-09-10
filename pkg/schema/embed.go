package schema

import (
	_ "embed"

	"github.com/DIMO-Network/model-garage/pkg/nativestatus/schema"
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

// DefinitionsYAML is the embedded YAML file containing the definitions.yaml for the VSS schema.
//
// Deprecated: Use pkg/nativestatus/shema.DefinitionsYAML instead.
func DefinitionsYAML() string {
	return schema.DefinitionsYAML()
}

// DefaultDefinitionsYAML is the embedded YAML file containing information about what signals will be displayed and used by the DIMO Node.
func DefaultDefinitionsYAML() string {
	return defaultDefinitionsYAML
}

// DefaultEventTagsYAML is the embedded YAML file containing information about event tags.
func DefaultEventTagsYAML() string {
	return defaultEventTagsYAML
}
