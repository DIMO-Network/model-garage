package schema

import (
	"encoding/csv"
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/DIMO-Network/model-garage/pkg/vss"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

// coordinatesVSSDataType is a fake reference to a VSS struct type that
// includes the three fields we want. See
// https://covesa.github.io/vehicle_signal_specification/rule_set/data_entry/data_types_struct/index.html
const coordinatesVSSDataType = "Types.DIMO.Coordinates"

// LoadSignalsCSV loads the signals from a vss CSV file.
func LoadSignalsCSV(r io.Reader) ([]*SignalInfo, error) {
	reader := csv.NewReader(r)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read vspec: %w", err)
	}

	locRec := make([]string, colLen)
	locRec[nameCol] = vss.FieldCurrentLocationCoordinates
	locRec[typeCol] = "struct"
	locRec[dataTypeCol] = coordinatesVSSDataType
	locRec[deprecatedCol] = "false"
	locRec[descCol] = "Vehicle's current location in WGS 84 coordinates, possibly with a measure of HDOP."
	records = append(records, locRec)

	var signals []*SignalInfo
	for i := 1; i < len(records); i++ {
		record := records[i]
		signals = append(signals, NewSignalInfo(record))
	}

	// Sort the signals by name
	slices.SortStableFunc(signals, func(a, b *SignalInfo) int {
		return strings.Compare(a.Name, b.Name)
	})

	return signals, nil
}

// LoadDefinitionFile loads the definitions from a definitions.yaml file.
func LoadDefinitionFile(r io.Reader) (*Definitions, error) {
	decoder := yaml.NewDecoder(r)
	var defInfos []*DefinitionInfo
	err := decoder.Decode(&defInfos)
	if err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	}
	definitions := &Definitions{
		FromName: map[string]*DefinitionInfo{},
	}
	for _, info := range defInfos {
		if err := Validate(info); err != nil {
			return nil, fmt.Errorf("error validating definitions: %w", err)
		}
		definitions.FromName[info.VspecName] = info
	}

	return definitions, nil
}

// GetDefinedSignals reads the signals and definitions files and merges them.
func GetDefinedSignals(specReader, definitionReader io.Reader) (*TemplateData, error) {
	signals, err := LoadSignalsCSV(specReader)
	if err != nil {
		return nil, fmt.Errorf("error reading signals: %w", err)
	}

	definitions, err := LoadDefinitionFile(definitionReader)
	if err != nil {
		return nil, fmt.Errorf("error reading definition file: %w", err)
	}
	signals = definitions.DefinedSignal(signals)
	modelName := "Model"
	if len(signals) > 0 {
		idx := strings.IndexByte(signals[0].Name, '.')
		if idx > 0 {
			modelName = signals[0].Name[:idx]
			modelName = cases.Title(language.English).String(modelName)
		}
	}

	originalNameMap := map[string]map[string]*SignalInfo{}
	for _, signal := range signals {
		for _, conv := range signal.Conversions {
			signalsForName := originalNameMap[conv.OriginalName]
			if signalsForName == nil {
				signalsForName = map[string]*SignalInfo{}
			}
			signalsForName[signal.Name] = signal
			originalNameMap[conv.OriginalName] = signalsForName
		}
	}

	tmplData := &TemplateData{
		Signals:       signals,
		ModelName:     modelName,
		OriginalNames: createListOfOriginalNames(signals),
	}

	return tmplData, nil
}

// createListOfOriginalNames reverse the mapping of signalInfo => []conversions to conversions.OriginalName => []signalsInfo
func createListOfOriginalNames(signals []*SignalInfo) []*OriginalNameInfo {
	originalNameMap := map[string]map[string]*SignalInfo{}
	for _, signal := range signals {
		for _, conv := range signal.Conversions {
			signalsForName := originalNameMap[conv.OriginalName]
			if signalsForName == nil {
				signalsForName = map[string]*SignalInfo{}
			}
			signalsForName[signal.Name] = signal
			originalNameMap[conv.OriginalName] = signalsForName
		}
	}
	originalNameMapList := []*OriginalNameInfo{}
	for originalName, signalsForName := range originalNameMap {
		origNameInfo := &OriginalNameInfo{
			Name: originalName,
		}
		for _, signal := range signalsForName {
			origNameInfo.Signals = append(origNameInfo.Signals, signal)
		}
		slices.SortStableFunc(origNameInfo.Signals, func(a, b *SignalInfo) int {
			return strings.Compare(a.Name, b.Name)
		})
		originalNameMapList = append(originalNameMapList, origNameInfo)
	}
	slices.SortStableFunc(originalNameMapList, func(a, b *OriginalNameInfo) int {
		return strings.Compare(a.Name, b.Name)
	})
	return originalNameMapList
}
