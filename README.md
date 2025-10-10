# Model Garage

![GitHub license](https://img.shields.io/badge/license-Apache%202.0-blue.svg)
[![GoDoc](https://godoc.org/github.com/DIMO-Network/model-garage?status.svg)](https://godoc.org/github.com/DIMO-Network/model-garage)
[![Go Report Card](https://goreportcard.com/badge/github.com/DIMO-Network/model-garage)](https://goreportcard.com/report/github.com/DIMO-Network/model-garage)

Welcome to the **Model Garage**, a Golang toolkit for managing and working with DIMO vehicle signal data.

## Features

1. **Signal Conversion**: Converting raw data from various hardware devices (AutoPi, Ruptela, Tesla, etc.) into standardized VSS (Vehicle Signal Specification) format
2. **Cloud Event Processing**: Transforming raw messages into CloudEvents for the DIMO data pipeline
3. **Modular Architecture**: A plugin-like system where each hardware provider has its own conversion module
4. **Code Generation**: Automatic generation of conversion functions from YAML definitions

## Signal Definitions

Signal definitions can be found in the [spec package](./pkg/schema/spec/spec.md).
This package is the source of truth for signals used for DIMO Data.

## Migrations

To create a new migration, run the following command:

```bash
Make migration name=<migration_name>
```

This will create a new migration file the given name in the `migrations` directory.
this creation should be used over the goose binary to ensure expected behavior of embedded migrations.

## Repo structure

### Codegen

The `codegen` directory contains the code generation tool for creating models from vspec CSV schemas. The tool is a standalone application that can be run from the command line.

Example usage:

```bash
go run github.com/DIMO-Network/model-garage/cmd/codegen -generators=custom -custom.output-file=./pkg/vss/vehicle-structs.go -custom.template-file=./internal/generator/vehicle.tmpl -custom.format=true
```

```
codegen is a tool to generate code for the model-garage project.
Available generators:
        - custom: Runs a given golang template with pkg/schema.TemplateData data.
        - convert: Generates conversion functions for converting between raw data into signals.Usage:
  -convert.copy-comments
        Copy through comments on conversion functions. Default is false.
  -convert.output-file string
        Output file for the conversion functions. (default "convert-funcs_gen.go")
  -convert.package string
        Name of the package to generate the conversion functions. If empty, the base model name is used.
  -custom.format
        Format the generated file with goimports.
  -custom.output-file string
        Path of the generate gql file (default "custom.txt")
  -custom.template-file string
        Path to the template file. Which is executed with codegen.TemplateData data.
  -definitions string
        Path to the definitions file if empty, the definitions will be used
  -generators string
        Comma separated list of generators to run. Options: convert, custom. (default "all")
  -spec string
        Path to the vspec CSV file if empty, the embedded vspec will be used
```

#### Generation Info

The codegen tool is typically used to create files based on arbitrary signal definitions. The tool reads the signal definitions and custom templates and executes the templates to create the output files.

#### Custom Generator

The custom generator takes in a custom template file and output file. The template file is a Go template that is executed with the signal definitions. The data struct passed into the template is defined by [pkg/schema/signal.go.(TemplateData)](pkg/schema/signal.go)
see [vehicle.tmpl](internal/generator/vehicle.tmpl) for an example template.

#### Convert Generator

The convert generator is a built-in generator that creates conversion functions for each signal. The conversion functions are created based on the signal definitions. The conversion functions are meant to be overridden with custom logic as needed. When generation is re-run, the conversion functions are not overwritten.

## Getting Started

For comprehensive documentation on how to work with Model Garage, see the [Developer Guide](./DEVELOPER_GUIDE.md).

Quick links:

- [Understanding Modules](./DEVELOPER_GUIDE.md#understanding-modules)
- [Adding New Signals](./DEVELOPER_GUIDE.md#adding-new-signals)
- [Code Generation](./DEVELOPER_GUIDE.md#code-generation)
- [Module-Specific Guides](./DEVELOPER_GUIDE.md#module-specific-guides)

## Common Tasks

### Adding a New Signal

1. Update `pkg/schema/spec/default-definitions.yaml` to add your signal
2. Update module-specific definitions (e.g., `pkg/ruptela/schema/ruptela_definitions.yaml`)
3. Run `make generate` to regenerate conversion code
4. Implement custom conversion logic if needed
5. Update tests

See the [Adding New Signals](./DEVELOPER_GUIDE.md#adding-new-signals) guide for detailed steps.

### Propagating Changes

After releasing changes to model-garage, update the following services:

1. **dis** - Signal conversion and storage
2. **telemetry-api** - API access to signals
3. **vehicle-triggers-api** - Webhook support for signals

### Adding Custom VSS Signals

When the COVESA standard doesn't have a signal you need:

1. Go to the [DIMO VSS repository](https://github.com/DIMO-Network/VSS)
2. Add your signal to `overlays/DIMO/dimo.vspec`
3. Follow the process documented in that repository
