package schema

import (
	"encoding/csv"
	"fmt"
	"io"
	"slices"
	"strings"

	"gopkg.in/yaml.v3"
)

// CoordinatesVSSDataType is a hardcoded reference to a VSS struct type that
// contains three properties: latitude, longitude, and HDOP.
//
// See the COVESA documentation for more on VSS structs:
// https://covesa.github.io/vehicle_signal_specification/rule_set/data_entry/data_types_struct/
//
// The type is in the VSS CSV file that we embed, but we are not yet willing
// to write a general mechanism for translating VSS structs into Go structs.
const CoordinatesVSSDataType = "Types.DIMO.Coordinates"

// LoadSignalsCSV loads the signals from a vss CSV file.
func LoadSignalsCSV(r io.Reader) ([]*SignalInfo, error) {
	reader := csv.NewReader(r)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read vspec: %w", err)
	}

	var signals []*SignalInfo
	for i := 1; i < len(records); i++ {
		sig := NewSignalInfo(records[i])
		if sig == nil {
			continue
		}
		signals = append(signals, sig)
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
		return nil, fmt.Errorf("failed to decode yaml: %w", err)
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

func LoadEventNames(r io.Reader) ([]*EventNameInfo, error) {
	decoder := yaml.NewDecoder(r)
	var eventNameInfos []*EventNameInfo
	err := decoder.Decode(&eventNameInfos)
	if err != nil {
		return nil, fmt.Errorf("failed to decode yaml: %w", err)
	}
	slices.SortStableFunc(eventNameInfos, func(a, b *EventNameInfo) int {
		return strings.Compare(a.Name, b.Name)
	})

	for _, eventNameInfo := range eventNameInfos {
		if eventNameInfo.GOName == "" {
			eventNameInfo.GOName = EventNameToGoName(eventNameInfo.Name)
		}
		if eventNameInfo.JSONName == "" {
			eventNameInfo.JSONName = eventNameInfo.Name
		}
		if err := ValidateEventName(eventNameInfo); err != nil {
			return nil, fmt.Errorf("error validating event name: %w", err)
		}
	}

	return eventNameInfos, nil
}

// GetDefaultSignals reads the default signals and definitions files and merges them.
func GetDefaultSignals() ([]*SignalInfo, error) {
	specReader := strings.NewReader(VssRel42DIMO())
	definitionReader := strings.NewReader(DefaultDefinitionsYAML())
	signalDefinitions, err := GetDefinedSignals(specReader, definitionReader)
	if err != nil {
		return nil, fmt.Errorf("error getting defined signals: %w", err)
	}
	return signalDefinitions.Signals, nil
}

// GetDefaultEventNames reads the default event names file and returns the event names.
func GetDefaultEventNames() ([]*EventNameInfo, error) {
	return LoadEventNames(strings.NewReader(DefaultEventNamesYAML()))
}

// GetDefinedSignals reads the signals and definitions files and merges them.
func GetDefinedSignals(specReader, definitionReader io.Reader) (SignalDefinitions, error) {
	signals, err := LoadSignalsCSV(specReader)
	if err != nil {
		return SignalDefinitions{}, fmt.Errorf("error reading signals: %w", err)
	}

	definitions, err := LoadDefinitionFile(definitionReader)
	if err != nil {
		return SignalDefinitions{}, fmt.Errorf("error reading definition file: %w", err)
	}
	signals = definitions.DefinedSignal(signals)

	signalDefs := SignalDefinitions{
		Signals:       signals,
		OriginalNames: createListOfOriginalNames(signals),
	}

	return signalDefs, nil
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
