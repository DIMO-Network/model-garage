// Package schema contains the embedded schema files for the ruptela devices
package schema

import _ "embed"

//go:embed oids.csv
var oidCSV string

//go:embed ruptela_definitions.yaml
var ruptelaDefinitions string

// OIDCSV is the embedded CSV file containing ruptela OID definitions.
func OIDCSV() string {
	return oidCSV
}

// RuptelaDefinitionsYAML is the embedded YAML file containing the ruptela-definitions.yaml for the VSS schema.
func RuptelaDefinitionsYAML() string {
	return ruptelaDefinitions
}
