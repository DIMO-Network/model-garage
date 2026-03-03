.PHONY: clean run build install dep test lint format docker gql tools tools-golangci-lint tools-gotestsum
# Set the bin path
PATHINSTBIN = $(abspath ./bin)
export PATH := $(PATHINSTBIN):$(PATH)

BIN_NAME					?= codegen
DEFAULT_INSTALL_DIR			:= $(go env GOPATH)/bin
DEFAULT_ARCH				:= $(shell go env GOARCH)
DEFAULT_GOOS				:= $(shell go env GOOS)
ARCH						?= $(DEFAULT_ARCH)
GOOS						?= $(DEFAULT_GOOS)
INSTALL_DIR					?= $(DEFAULT_INSTALL_DIR)
.DEFAULT_GOAL := run

# Get the app version from the git tag and commit
GIT_COMMIT := $(shell git rev-parse --short HEAD)
TAG := $(shell git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")
TAG_COMMIT := $(shell git rev-list -n 1 $(TAG))
VERSION := $(TAG)
ifneq ($(TAG_COMMIT), $(shell git rev-parse HEAD))
	VERSION := $(TAG)-$(GIT_COMMIT)
endif

# Dependency versions
GOLANGCI_VERSION := latest
help:
	@echo "Specify a subcommand:"
	@grep -hE '^[0-9a-zA-Z_-]+:.*?## .*$$' ${MAKEFILE_LIST} | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[0;36m%-20s\033[m %s\n", $$1, $$2}'
	@echo ""

clean:
	@rm -rf $(PATHINSTBIN)

dep:
	@go mod tidy

test: # Run all tests
	@go test ./... -race

lint: # Run all linters
	@golangci-lint version
	@golangci-lint run --timeout=5m


add-migration: # Add a new migration to the database
	go tool goose create ${name} sql -s --dir=pkg/migrations

tools-golangci-lint: # Install golangci-lint
	@mkdir -p $(PATHINSTBIN)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | BINARY=golangci-lint bash -s -- ${GOLANGCI_VERSION}

tools: tools-golangci-lint # Install all tools
	go install tool

clickhouse: # Run the clickhouse container
	go run ./cmd/clickhouse-container

generate: generate-ruptela generate-tesla generate-hashdog generate-autopi generate-vss# Generate all files for the repository

generate-ruptela: # Generate all files for ruptela
	go run ./cmd/codegen -convert.package=ruptela -generators=convert -convert.output-file=./pkg/ruptela/convert_signal_funcs_gen.go -definitions=./pkg/ruptela/schema/ruptela_definitions.yaml
	go run ./cmd/codegen -generators=custom -custom.output-file=./pkg/ruptela/convert_signal_status_gen.go -custom.template-file=./pkg/ruptela/codegen/convert_signal_status.tmpl -custom.format=true -definitions=./pkg/ruptela/schema/ruptela_definitions.yaml
	go run ./cmd/codegen -generators=custom -custom.output-file=./pkg/ruptela/convert_signal_location_gen.go -custom.template-file=./pkg/ruptela/codegen/convert_signal_location.tmpl -custom.format=true -definitions=./pkg/ruptela/schema/ruptela_definitions.yaml
	go run ./pkg/ruptela/codegen

generate-autopi: # Generate all files for autopi
	go run ./cmd/codegen -convert.package=autopi -generators=convert -convert.output-file=./pkg/autopi/convert_signal_funcs_gen.go -definitions=./pkg/autopi/schema/autopi_definitions.yaml
	go run ./cmd/codegen -generators=custom -custom.output-file=./pkg/autopi/convert_signal_status_gen.go -custom.template-file=./pkg/autopi/codegen/convert_signal_status.tmpl -custom.format=true -definitions=./pkg/autopi/schema/autopi_definitions.yaml

generate-hashdog: # Generate all files for hashdog (macaron)
	go run ./cmd/codegen -convert.package=hashdog -generators=convert -convert.output-file=./pkg/hashdog/convert_signal_funcs_gen.go -definitions=./pkg/hashdog/schema/hashdog_definitions.yaml
	go run ./cmd/codegen -generators=custom -custom.output-file=./pkg/hashdog/convert_signal_status_gen.go -custom.template-file=./pkg/hashdog/codegen/convert_signal_status.tmpl -custom.format=true -definitions=./pkg/hashdog/schema/hashdog_definitions.yaml


generate-tesla: # Generate all files for tesla
	go run ./cmd/codegen -convert.package=api -generators=convert -convert.output-file=./pkg/tesla/api/convert_signal_funcs_gen.go -definitions=./pkg/tesla/api/schema/tesla_definitions.yaml
	go run ./cmd/codegen -generators=custom -custom.output-file=./pkg/tesla/api/convert_signal_status_gen.go -custom.template-file=./pkg/tesla/api/codegen/convert_signal_status.tmpl -custom.format=true -definitions=./pkg/tesla/api/schema/tesla_definitions.yaml
	go run ./pkg/tesla/telemetry/codegen

generate-vss: # Generate all files for vss
	go run ./cmd/codegen -generators=custom -custom.output-file=./pkg/vss/vehicle-structs.go -custom.template-file=./internal/generator/vehicle.tmpl -custom.format=true
