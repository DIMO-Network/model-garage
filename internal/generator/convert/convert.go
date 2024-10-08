// Package convert provides a function to generate conversion functions for a vehicle struct.
package convert

import (
	_ "embed"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/DIMO-Network/model-garage/pkg/schema"
)

const (
	// DefaultConversionFile is the default name of the conversion file.
	DefaultConversionFile = "convert-funcs_gen.go"
)

//go:embed convertFunc.tmpl
var convertFuncTemplateStr string

const header = `// Code generated by github.com/DIMO-Network/model-garage.
package %s

// This file is automatically populated with conversion functions for each field of the model struct.
// any conversion functions already defined in this package will be coppied through.
// note: DO NOT mutate the orginalDoc parameter which is shared between all conversion functions.
`

// Config is the configuration for the conversion generator.
type Config struct {
	// CopyComments determines if comments for the conversion functions should be copied through.
	CopyComments bool
	// OutputFile is the output directory for the generated conversion files.
	// if empty, the base output directory is used.
	OutputFile string
	// PackageName is the name of the package to generate the conversion functions.
	PackageName string
}

// funcTmplData contains the data to be used during template execution for writing a single conversion function.
type funcTmplData struct {
	// Signal is the signal that we are converting to.
	Signal *schema.SignalInfo
	// Conversion is the information about the signal we are converting from.
	Conversion *schema.ConversionInfo
	// FuncName is the name of the conversion function.
	FuncName string
	// DocComment is the original doc comment for the conversion function if it exists.
	DocComment string
	// Body of the original conversion function if it exists.
	Body string
}

// Generate creates a conversion functions for each field of a model struct.
// as well as the entire model struct.
func Generate(tmplData *schema.TemplateData, cfg Config) (err error) {
	cfg.OutputFile = filepath.Clean(cfg.OutputFile)
	if cfg.OutputFile == "" {
		cfg.OutputFile = DefaultConversionFile
	}
	if cfg.PackageName == "" {
		cfg.PackageName = strings.ToLower(tmplData.ModelName)
	}

	// Get the conversion functions that need to be generated.
	convertFunc := getConversionFunctions(tmplData.Signals)
	if len(convertFunc) == 0 {
		return nil
	}

	outputDir := filepath.Dir(cfg.OutputFile)
	// Get existing functions in the output directory.
	existingFuncs, err := GetDeclaredFunctions(outputDir)
	if err != nil {
		return fmt.Errorf("error getting declared functions: %w", err)
	}

	// Create the conversion functions.
	convertFuncTemplate, err := createConvertFuncTemplate()
	if err != nil {
		return err
	}

	err = writeConvertFuncs(convertFunc, existingFuncs, convertFuncTemplate, cfg.OutputFile, cfg.PackageName, cfg.CopyComments)
	if err != nil {
		return err
	}
	return nil
}
