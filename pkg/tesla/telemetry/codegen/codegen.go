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
	"log"
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

const (
	ParseFloatFlag  = "PARSE_FLOAT"
	ConvertUnitFlag = "CONVERT_UNIT"
)

type TargetVSSSignal struct {
	GoVSSName    string
	JsonName     string
	GoOutputType string
	ConvertFunc  string
	Body         string
}

type Rule struct {
	TeslaField string `yaml:"teslaField"`
	// TeslaType is the protobuf type of the value field on the Datum. If not specified, it is
	// assumed to be string. This is the dominant case.
	TeslaType      string `yaml:"teslaType"`
	TeslaUnit      string `yaml:"teslaUnit"`
	DisableConvert bool   `yaml:"disableConvert"`
	// VSSSignal is the full path to the VSS signal that will be produced from this Datum. For
	// example, one might write "Vehicle.Cabin.Door.Row1.Left.IsLocked".
	VSSSignals []string `yaml:"vssSignals"`
}

//go:embed inner.tmpl
var innerTmpl string

//go:embed outer.tmpl
var outerTmpl string

// protoToGoTypes maps protobuf types to Go types. The only point of disagreement here
// is for floating point numbers.
var protoToGoTypes = map[string]string{
	"string": "string",
	"int32":  "int32",
	"int64":  "int64",
	"float":  "float32",
	"double": "float64",
	"bool":   "bool",
}

func snakeToPascal(s string) string {
	words := strings.Split(s, "_")
	for i, w := range words {
		if len(w) != 0 {
			words[i] = strings.ToUpper(w[:1]) + w[1:]
		}
	}
	return strings.Join(words, "")
}

func Generate(packageName, outerOutputPath, innerOutputPath string) error {
	signalInfos, err := schema.LoadSignalsCSV(strings.NewReader(schema.VssRel42DIMO()))
	if err != nil {
		log.Fatalf("Failed to load VSS schema: %v", err)
	}

	signalInfoBySignal := make(map[string]*schema.SignalInfo, len(signalInfos))
	for _, s := range signalInfos {
		signalInfoBySignal[s.Name] = s
	}

	rules, err := loadRules()
	if err != nil {
		return fmt.Errorf("failed to load rules: %w", err)
	}

	tmplInput := &TemplateInput{
		Package: packageName,
	}

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

		parseString := parsers[r.TeslaType]

		if r.TeslaUnit != "" {
			if r.TeslaType != "double" {
				return fmt.Errorf("unit specified for Tesla signal of non-double type %s", r.TeslaType)
			}
		}

		var targets []*TargetVSSSignal

		for _, st := range r.VSSSignals {
			info, ok := signalInfoBySignal[st]
			if !ok {
				return fmt.Errorf("unrecognized VSS signal %q", st)
			}

			var convertFunc string
			if !r.DisableConvert && r.TeslaUnit != "" && info.Unit != "" && r.TeslaUnit != info.Unit {
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

			targets = append(targets, &TargetVSSSignal{
				GoVSSName:    info.GOName,
				JsonName:     info.JSONName,
				GoOutputType: info.GOType(),
				ConvertFunc:  convertFunc,
			})
		}

		tmplInput.Conversions = append(tmplInput.Conversions, &Conversion{
			TeslaField:       r.TeslaField,
			WrapperName:      teslaType.TeslaWrapperType,
			WrapperFieldName: teslaType.TeslaWrapperFieldName,
			Parser:           parseString,
			GoInputType:      teslaType.ValueType,
			TeslaTypeName:    teslaType.NiceName,
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

func assembleTeslaTypeInformation() (map[string]TeslaTypeDescription, error) {
	out := make(map[string]TeslaTypeDescription)

	desc := (&protos.Value{}).ProtoReflect().Descriptor()
	for i := range desc.Fields().Len() {
		field := desc.Fields().Get(i)
		fieldName := field.Name()

		teslaWrapperFieldName := snakeToPascal(string(fieldName))
		teslaWrapperType := "Value_" + teslaWrapperFieldName
		var protoType, valueType string
		switch field.Kind() {
		case protoreflect.MessageKind:
			protoType = string(field.Message().Name())
			valueType = "*protos." + protoType
		case protoreflect.EnumKind:
			protoType = string(field.Enum().Name())
			valueType = "protos." + protoType
		default:
			// Primitive types.
			protoType = field.Kind().String()
			goType, ok := protoToGoTypes[protoType]
			if !ok {
				return nil, fmt.Errorf("no Go mapping for protobuf type %s", protoType)
			}
			valueType = goType
		}

		niceName := strings.ToUpper(protoType[:1]) + protoType[1:]

		out[protoType] = TeslaTypeDescription{
			TeslaWrapperType:      teslaWrapperType,
			TeslaWrapperFieldName: teslaWrapperFieldName,
			ValueType:             valueType,
			NiceName:              niceName,
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
	TeslaTypeName    string
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
	"WindowState": "WindowState",
	"Doors":       "Doors",
}

type TeslaTypeDescription struct {
	TeslaWrapperType      string
	TeslaWrapperFieldName string
	ValueType             string
	NiceName              string
}
