// Package runner is a package that provides a programmatic interface to the code generation tool.
package runner

import (
	"fmt"
	"io"
	"slices"

	"github.com/DIMO-Network/model-garage/internal/generator/convert"
	"github.com/DIMO-Network/model-garage/internal/generator/custom"
	"github.com/DIMO-Network/model-garage/internal/generator/model"
	"github.com/DIMO-Network/model-garage/pkg/codegen"
	"github.com/DIMO-Network/model-garage/pkg/schema"
)

const (
	// AllGenerator is a constant to run all generators.
	AllGenerator = "all"
	// ModelGenerator is a constant to run the model generator.
	ModelGenerator = "model"
	// ConvertGenerator is a constant to run the convert generator.
	ConvertGenerator = "convert"
	// CustomGenerator is a constant to run the custom generator.
	CustomGenerator = "custom"
)

// Config is the configuration for the code generation tool.
type Config struct {
	PackageName string
	OutputDir   string
	Custom      custom.Config
	Convert     convert.Config
}

// Execute runs the code generation tool.
func Execute(vspecReader, definitionsReader io.Reader, generators []string, cfg Config) error {
	// TODO move params to a config struct.

	if len(generators) == 0 {
		generators = []string{AllGenerator}
	}
	// if none of the generators are selected, return an error.
	switch {
	case slices.Contains(generators, AllGenerator):
	case slices.Contains(generators, ModelGenerator):
	case slices.Contains(generators, ConvertGenerator):
	case slices.Contains(generators, CustomGenerator):
	default:
		return fmt.Errorf("no generator selected")
	}

	err := codegen.EnsureDir(cfg.OutputDir)
	if err != nil {
		return fmt.Errorf("failed to ensure output directory: %w", err)
	}

	tmplData, err := schema.GetDefinedSignals(vspecReader, definitionsReader)
	if err != nil {
		return fmt.Errorf("failed to get defined signals: %w", err)
	}

	tmplData.PackageName = cfg.PackageName

	if slices.Contains(generators, AllGenerator) || slices.Contains(generators, ModelGenerator) {
		err = model.Generate(tmplData, cfg.OutputDir)
		if err != nil {
			return fmt.Errorf("failed to generate model file: %w", err)
		}
	}

	if slices.Contains(generators, AllGenerator) || slices.Contains(generators, ConvertGenerator) {
		err = convert.Generate(tmplData, cfg.OutputDir, cfg.Convert)
		if err != nil {
			return fmt.Errorf("failed to generate convert file: %w", err)
		}
	}

	if slices.Contains(generators, AllGenerator) || slices.Contains(generators, CustomGenerator) {
		err = custom.Generate(tmplData, cfg.OutputDir, cfg.Custom)
		if err != nil {
			return fmt.Errorf("failed to generate custom file: %w", err)
		}
	}

	return nil
}
