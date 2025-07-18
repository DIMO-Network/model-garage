package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"strings"
	"text/template"

	"github.com/DIMO-Network/model-garage/pkg/codegen"
	"github.com/DIMO-Network/model-garage/pkg/schema"
	tschema "github.com/DIMO-Network/model-garage/pkg/tesla/telemetry/schema"
	"github.com/teslamotors/fleet-telemetry/protos"
	"google.golang.org/protobuf/reflect/protoreflect"
	"gopkg.in/yaml.v3"
)

type TargetVSSSignal struct {
	GoVSSName    string
	JSONName     string
	GoOutputType string
	ConvertFunc  string
	GoInputUnit  string
	Body         string
}

type Rule struct {
	// TeslaField is the name of the enum value in the key field on the
	// Tesla Datum message.
	TeslaField string `yaml:"teslaField"`
	// TeslaType is the protobuf type of the value held within a Datum.
	TeslaType string `yaml:"teslaType"`
	// TeslaUnit is the unit of measure for a numeric Tesla value. It
	// may be empty.
	//
	// Tesla tends to use miles in all regions.
	TeslaUnit string `yaml:"teslaUnit"`
	// VSSSignals is a list of VSS paths that will be populated
	// from data with the given field.
	//
	// For example, this could contain a single element,
	// "Vehicle.Cabin.Door.Row1.Left.IsLocked".
	VSSSignals []string `yaml:"vssSignals"`
}

//go:embed inner.tmpl
var innerTmpl string

//go:embed outer.tmpl
var outerTmpl string

// protoTypeToGoType maps scalar protobuf types to Go types. The only
// point of disagreement here is for floating point numbers.
var protoTypeToGoType = map[string]string{
	"string": "string",
	"int32":  "int32",
	"int64":  "int64",
	"float":  "float32",
	"double": "float64",
	"bool":   "bool",
}

func snakeToPascal(s string) string {
	// This is not robust.
	words := strings.Split(s, "_")
	for i, w := range words {
		if len(w) != 0 {
			words[i] = strings.ToUpper(w[:1]) + w[1:]
		}
	}
	return strings.Join(words, "")
}

func createSignalLookup() (map[string]*schema.SignalInfo, error) {
	signalInfos, err := schema.LoadSignalsCSV(strings.NewReader(schema.VssRel42DIMO()))
	if err != nil {
		return nil, err
	}

	signalInfoBySignal := make(map[string]*schema.SignalInfo, len(signalInfos))
	for _, s := range signalInfos {
		signalInfoBySignal[s.Name] = s
	}

	return signalInfoBySignal, nil
}

func Generate(packageName, outerOutputPath, innerOutputPath string) error {
	signalInfoBySignal, err := createSignalLookup()
	if err != nil {
		return fmt.Errorf("failed to load VSS schema: %w", err)
	}

	rules, err := loadRules()
	if err != nil {
		return fmt.Errorf("failed to load rules: %w", err)
	}

	tmplInput := &TemplateInput{Package: packageName}

	teslaTypeToAttributes, err := assembleTeslaTypeInformation()
	if err != nil {
		return err
	}

	for _, r := range rules {
		_, ok := protos.Field_value[r.TeslaField]
		if !ok {
			return fmt.Errorf("unrecognized Tesla field %q", r.TeslaField)
		}

		teslaType, ok := teslaTypeToAttributes[r.TeslaType]
		if !ok {
			return fmt.Errorf("unsupported Tesla type %q", r.TeslaType)
		}

		if len(r.VSSSignals) == 0 {
			// It's fine to not specify any targets, but don't generate
			// code in this case.
			continue
		}

		var targets []*TargetVSSSignal
		for _, st := range r.VSSSignals {
			info, ok := signalInfoBySignal[st]
			if !ok {
				return fmt.Errorf("unrecognized VSS signal %q", st)
			}

			convertFunc := ""
			if r.TeslaUnit != "" && info.Unit != "" && r.TeslaUnit != info.Unit && r.TeslaType == "double" && info.GOType() == "float64" {
				if convertFrom, ok := conversions[r.TeslaUnit]; ok {
					// More to check here.
					convertFunc, ok = convertFrom[info.Unit]
					if !ok {
						return fmt.Errorf("no converters from unit %s to unit %s", r.TeslaUnit, info.Unit)
					}
				} else {
					return fmt.Errorf("no converters from unit %s", r.TeslaUnit)
				}
			}

			goInputUnit := ""
			if r.TeslaUnit != "" {
				if info.Unit == "" {
					goInputUnit = r.TeslaUnit
				} else {
					goInputUnit = info.Unit
				}
			}

			targets = append(targets, &TargetVSSSignal{
				GoVSSName:    info.GOName,
				JSONName:     info.JSONName,
				GoOutputType: info.GOType(),
				ConvertFunc:  convertFunc,
				GoInputUnit:  goInputUnit,
			})
		}

		tmplInput.Conversions = append(tmplInput.Conversions, &Conversion{
			TeslaField:       r.TeslaField,
			WrapperName:      teslaType.WrapperType,
			WrapperFieldName: teslaType.Field,
			Parser:           parsers[r.TeslaType],
			GoInputType:      teslaType.FieldType,
			VSSSignals:       targets,
		})
	}

	err = writeOuter(tmplInput, outerOutputPath)
	if err != nil {
		return err
	}

	err = writeInner(tmplInput, innerOutputPath)
	if err != nil {
		return err
	}

	return nil
}

func assembleTeslaTypeInformation() (map[string]TeslaGoValueType, error) {
	out := make(map[string]TeslaGoValueType)

	// This is protobuf reflection. The protos.Value type in Go does
	// not have all of these fields. Maybe we should reflect in Go,
	// instead?
	desc := (&protos.Value{}).ProtoReflect().Descriptor()

	for i := range desc.Fields().Len() {
		field := desc.Fields().Get(i)

		goFieldName := snakeToPascal(string(field.Name()))
		wrapperType := "Value_" + goFieldName // This is only a concept in Go.

		var protoType, goFieldType string

		switch kind := field.Kind(); kind {
		case protoreflect.MessageKind:
			protoType = string(field.Message().Name())
			goFieldType = "*protos." + protoType
		case protoreflect.EnumKind:
			// TODO(elffjs): Should we try to check if the number we
			// get for such fields is in-bounds?
			protoType = string(field.Enum().Name())
			goFieldType = "protos." + protoType
		default:
			// Primitive types, we hope.
			protoType = kind.String()
			var ok bool
			goFieldType, ok = protoTypeToGoType[protoType]
			if !ok {
				return nil, fmt.Errorf("no Go mapping for protobuf type %s", protoType)
			}
		}

		out[protoType] = TeslaGoValueType{
			WrapperType: wrapperType,
			Field:       goFieldName,
			FieldType:   goFieldType,
		}
	}

	return out, nil
}

func writeOuter(tmplInput *TemplateInput, outerPath string) error {
	t, err := template.New("outer").Parse(outerTmpl)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, tmplInput)
	if err != nil {
		return err
	}

	return codegen.FormatAndWriteToFile(buf.Bytes(), outerPath)
}

func writeInner(tmplInput *TemplateInput, innerPath string) error {
	existingBodies := make(map[string]string)

	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, innerPath, nil, parser.ParseComments)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else {
		for _, decl := range astFile.Decls {
			if fn, ok := decl.(*ast.FuncDecl); ok {
				var buf bytes.Buffer
				err := format.Node(&buf, fset, &printer.CommentedNode{
					Node:     fn.Body,
					Comments: astFile.Comments,
				})
				if err != nil {
					return err
				}
				existingBodies[fn.Name.Name] = buf.String()
			}
		}
	}

	for _, conv := range tmplInput.Conversions {
		for _, vs := range conv.VSSSignals {
			name := fmt.Sprintf("Convert%sTo%s", conv.TeslaField, vs.GoVSSName)
			if body, ok := existingBodies[name]; ok {
				vs.Body = body
			}
		}
	}

	t, err := template.New("inner").Parse(innerTmpl)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, tmplInput)
	if err != nil {
		return err
	}

	return codegen.FormatAndWriteToFile(buf.Bytes(), innerPath)
}

func loadRules() ([]Rule, error) {
	defs := tschema.TelemetryDefinitionsYAML()

	var rules []Rule

	err := yaml.Unmarshal([]byte(defs), &rules)
	if err != nil {
		return nil, fmt.Errorf("failed to parse rules YAML: %w", err)
	}

	return rules, nil
}

type Conversion struct {
	TeslaField       string
	WrapperName      string
	WrapperFieldName string
	Parser           string
	GoInputType      string

	VSSSignals []*TargetVSSSignal
}

type TemplateInput struct {
	Package     string
	Conversions []*Conversion
}

var conversions = map[string]map[string]string{
	"kW": {
		"W": "KilowattsToWatts",
	},
	"bar": {
		"kPa": "BarsToKilopascals",
	},
	"atm": {
		"kPa": "AtmospheresToKilopascals",
	},
	"mi": {
		"km": "MilesToKilometers",
	},
	"mph": {
		"km/h": "MilesPerHourToKilometersPerHour",
	},
}

var parsers = map[string]string{
	"double":      "Double",
	"int32":       "Int32",
	"WindowState": "WindowState",
	"Doors":       "Doors",
}

type TeslaGoValueType struct {
	// WrapperType is the name of a Go type T such that *T is
	// assignable to Value's Value field. The most common one is
	// protos.Value_StringValue.
	WrapperType string
	// Field is the name of the only field on the type named by
	// TeslaWrapperType. This field is what holds the value of
	// interest.
	Field string
	// FieldType is the type of the field named by Field.
	FieldType string
}
