// Package main provides the entrypoint for the code generation tool.
package main

import (
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/DIMO-Network/model-garage/internal/generator/convert"
	"github.com/DIMO-Network/model-garage/internal/generator/custom"
	"github.com/DIMO-Network/model-garage/pkg/runner"
	"github.com/DIMO-Network/model-garage/pkg/schema"
	"github.com/DIMO-Network/model-garage/pkg/version"
)

func main() {
	// Command-line flags
	printVersion := flag.Bool("version", false, "Print the version of the codegen tool")
	outputDir := flag.String("output", ".", "Output directory for the generated Go file")
	vspecPath := flag.String("spec", "", "Path to the vspec CSV file if empty, the embedded vspec will be used")
	definitionPath := flag.String("definitions", "", "Path to the definitions file if empty, the definitions will be used")
	packageName := flag.String("package", "vspec", "Name of the package to generate")
	generators := flag.String("generators", "all", "Comma separated list of generators to run. Options: all, model, convert, custom.")
	// Convert flags
	copyComments := flag.Bool("convert.copy-comments", false, "Copy through comments on conversion functions. Default is false.")
	convertPackageName := flag.String("convert.package", "", "Name of the package to generate the conversion functions. If empty, the base package name is used.")
	convertOutputDir := flag.String("convert.output", "", "Output directory for the generated conversion files. If empty, the output directory is used.")
	// Custom flags
	customOutFile := flag.String("custom.output-file", "", "Path of the generate gql file that is appened to the outputDir.")
	customTemplateFile := flag.String("custom.template-file", "", "Path to the template file. Which is executed with codegen.TemplateData data.")
	customFormat := flag.Bool("custom.format", false, "Format the generated file with goimports.")

	flag.Parse()

	if *printVersion {
		log.Printf("codegen version: %s", version.GetVersion())
		return
	}

	var vspecReader io.Reader
	if *vspecPath != "" {
		f, err := os.Open(filepath.Clean(*vspecPath))
		if err != nil {
			log.Fatalf("failed to open file: %v", err)
		}
		vspecReader = f
		//nolint:errcheck // we don't care about the error since we are not writing to the file
		defer f.Close()
	} else {
		vspecReader = strings.NewReader(schema.VssRel42DIMO())
	}
	var definitionReader io.Reader
	if *definitionPath != "" {
		f, err := os.Open(filepath.Clean(*definitionPath))
		if err != nil {
			defer log.Fatalf("failed to open file: %v", err)
			return
		}
		definitionReader = f
		//nolint:errcheck // we don't care about the error since we are not writing to the file
		defer f.Close()
	} else {
		definitionReader = strings.NewReader(schema.DefinitionsYAML())
	}
	gens := strings.Split(*generators, ",")

	cfg := runner.Config{
		PackageName: *packageName,
		OutputDir:   *outputDir,
		Custom: custom.Config{
			OutputFile:   *customOutFile,
			TemplateFile: *customTemplateFile,
			Format:       *customFormat,
		},
		Convert: convert.Config{
			CopyComments: *copyComments,
			PackageName:  *convertPackageName,
			OutputDir:    *convertOutputDir,
		},
	}

	err := runner.Execute(vspecReader, definitionReader, gens, cfg)
	if err != nil {
		defer log.Fatal(err)
		return
	}
}
