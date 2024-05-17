package schema

import (
	"regexp"
	"strings"
)

const (
	nameCol       = 0
	typeCol       = 1
	dataTypeCol   = 2
	deprecatedCol = 3
	unitCol       = 4
	minCol        = 5
	maxCol        = 6
	descCol       = 7
	colLen        = 8
)

var nonAlphaNum = regexp.MustCompile(`[^a-zA-Z0-9]+`)

// SignalInfo holds information about a signal that is accessed during template execution.
// This information comes from the combinations of the spec and definition files.
// The Types defined by this stuct are used to determine what strings to use in the template file.
type SignalInfo struct {
	// From spec CSV
	Name       string
	Type       string
	DataType   string
	Unit       string
	Min        string
	Max        string
	Desc       string
	Deprecated bool

	// Derived
	IsArray     bool
	GOName      string
	JSONName    string
	BaseGoType  string
	BaseGQLType string
	Conversions []*ConversionInfo
	Privileges  []string
}

// ConversionInfo contains the conversion information for a field.
type ConversionInfo struct {
	OriginalName string `json:"originalName" yaml:"originalName"`
	OriginalType string `json:"originalType" yaml:"originalType"`
	IsArray      bool   `json:"isArray"      yaml:"isArray"`
}

// DefinitionInfo contains the definition information for a field.
type DefinitionInfo struct {
	VspecName          string            `json:"vspecName"          yaml:"vspecName"`
	IsArray            *bool             `json:"isArray"            yaml:"isArray"`
	ClickHouseType     string            `json:"clickHouseType"     yaml:"clickHouseType"`
	GoType             string            `json:"goType"             yaml:"goType"`
	Conversions        []*ConversionInfo `json:"conversions"        yaml:"conversions"`
	RequiredPrivileges []string          `json:"requiredPrivileges" yaml:"requiredPrivileges"`
}

// Definitions is a map of definitions from clickhouse Name to definition info.
type Definitions struct {
	FromName map[string]*DefinitionInfo
	Signals  []*SignalInfo
}

// DefinedSignal returns a new slice of signals with the definition information applied.
// excluding signals that are not in the definition file.
func (m *Definitions) DefinedSignal(signal []*SignalInfo) []*SignalInfo {
	sigs := []*SignalInfo{}
	for _, sig := range signal {
		if definition, ok := m.FromName[sig.Name]; ok {
			newSignal := *sig
			newSignal.MergeWithDefinition(definition)
			sigs = append(sigs, &newSignal)
		}
	}
	return sigs
}

// NewSignalInfo creates a new SignalInfo from a record from the CSV file.
func NewSignalInfo(record []string) *SignalInfo {
	if len(record) < colLen {
		return nil
	}
	sig := &SignalInfo{
		Name:       record[nameCol],
		Type:       record[typeCol],
		DataType:   record[dataTypeCol],
		Deprecated: record[deprecatedCol] == "true",
		Unit:       record[unitCol],
		Min:        record[minCol],
		Max:        record[maxCol],
		Desc:       record[descCol],
	}
	// arrays are denoted by [] at the end of the type ex uint8[]
	sig.IsArray = strings.HasSuffix(sig.DataType, "[]")
	baseType := sig.DataType
	if sig.IsArray {
		// remove the [] from the type
		baseType = sig.DataType[:len(sig.DataType)-2]
	}
	if baseType != "" {
		//  if this is not a branch type, we can convert it to default golang and clickhouse types
		sig.BaseGoType = goTypeFromVSPEC(baseType)
		sig.BaseGQLType = gqlTypeFromVSPEC(baseType)
	}
	sig.GOName = goName(sig.Name)
	sig.JSONName = JSONName(sig.Name)

	return sig
}

// MergeWithDefinition merges the signal with the definition information.
func (s *SignalInfo) MergeWithDefinition(definition *DefinitionInfo) {
	if definition.GoType != "" {
		s.BaseGoType = definition.GoType
	}
	if definition.IsArray != nil {
		s.IsArray = *definition.IsArray
	}
	if len(definition.Conversions) != 0 {
		s.Conversions = definition.Conversions
		for _, conv := range s.Conversions {
			if conv.OriginalType == "" {
				conv.OriginalType = s.GOType()
			}
		}
	}
	s.Privileges = definition.RequiredPrivileges
}

// GOType returns the golang type of the signal.
func (s *SignalInfo) GOType() string {
	if s.IsArray {
		return "[]" + s.BaseGoType
	}
	return s.BaseGoType
}

// GQLType returns the graphql type of the signal.
func (s *SignalInfo) GQLType() string {
	return s.BaseGQLType
}

func goName(name string) string {
	return nonAlphaNum.ReplaceAllString(removePrefix(name), "")
}

// JSONName returns the json name of the signal.
func JSONName(name string) string {
	n := goName(name)
	// lowercase the first letter
	return strings.ToLower(n[:1]) + n[1:]
}

func removePrefix(name string) string {
	idx := strings.IndexByte(name, '.')
	if idx != -1 {
		return name[idx+1:]
	}
	return name
}

// goTypeFromVSPEC converts vspec type to golang types.
func goTypeFromVSPEC(dataType string) string {
	switch dataType {
	case "uint8", "int8", "uint16", "int16", "uint32", "int32", "uint64", "int64", "float", "double", "boolean":
		return "float64"
	default:
		return "string"
	}
}

// gqlTypeFromVSPEC converts vspec type to graphql types.
func gqlTypeFromVSPEC(baseType string) string {
	switch baseType {
	case "uint8", "int8", "uint16", "int16", "uint32", "int32", "float", "double", "uint64", "int64", "Boolean":
		return "Float"
	default:
		return "String"
	}
}
