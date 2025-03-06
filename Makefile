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

build:
	@CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(ARCH) \
		go build -ldflags "-X 'github.com/DIMO-Network/model-garage/pkg/version.version=$(VERSION)'" \
		-o bin/$(BIN_NAME) ./cmd/$(BIN_NAME)


run: build
	@./bin/$(BIN_NAME)

all: clean target

clean:
	@rm -rf $(PATHINSTBIN)

install: build
	@install -d $(INSTALL_DIR)
	@rm -f $(INSTALL_DIR)/$(BIN_NAME)
	@cp $(PATHINSTBIN)/* $(INSTALL_DIR)/

dep:
	@go mod tidy

test:
	@go test ./...

lint:
	@golangci-lint version
	@golangci-lint run --timeout=5m

format:
	@golangci-lint run --fix

migration:
	migration -output=./pkg/migrations -package=migrations -filename="${name}"

tools-golangci-lint:
	@mkdir -p $(PATHINSTBIN)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | BINARY=golangci-lint bash -s -- ${GOLANGCI_VERSION}

tools: tools-golangci-lint
	go install tool

clickhouse:
	go run ./cmd/clickhouse-container

generate: generate-nativestatus generate-ruptela generate-tesla generate-compass generate-lorawan# Generate all files for the repository
	go run ./cmd/codegen -generators=custom -custom.output-file=./pkg/vss/vehicle-structs.go -custom.template-file=./internal/generator/vehicle.tmpl -custom.format=true

generate-nativestatus: # Generate all files for nativestatus
	go run ./cmd/codegen -convert.package=nativestatus -generators=convert -convert.output-file=./pkg/nativestatus/vehicle-convert-funcs_gen.go -definitions=./pkg/nativestatus/schema/native-definitions.yaml
	go run ./cmd/codegen -generators=custom -custom.output-file=./pkg/nativestatus/vehicle-v1-convert_gen.go -custom.template-file=./pkg/nativestatus/convertv1.tmpl -custom.format=true -definitions=./pkg/nativestatus/schema/native-definitions.yaml
	go run ./cmd/codegen -generators=custom -custom.output-file=./pkg/nativestatus/vehicle-v2-convert_gen.go -custom.template-file=./pkg/nativestatus/convertv2.tmpl -custom.format=true -definitions=./pkg/nativestatus/schema/native-definitions.yaml

generate-ruptela: # Generate all files for ruptela
	go run ./cmd/codegen -convert.package=ruptela -generators=convert -convert.output-file=./pkg/ruptela/conver_signal_funcs_gen.go -definitions=./pkg/ruptela/schema/ruptela_definitions.yaml
	go run ./cmd/codegen -generators=custom -custom.output-file=./pkg/ruptela/convert_signal_status_gen.go -custom.template-file=./pkg/ruptela/codegen/convert_signal_status.tmpl -custom.format=true -definitions=./pkg/ruptela/schema/ruptela_definitions.yaml
	go run ./cmd/codegen -generators=custom -custom.output-file=./pkg/ruptela/convert_signal_location_gen.go -custom.template-file=./pkg/ruptela/codegen/convert_signal_location.tmpl -custom.format=true -definitions=./pkg/ruptela/schema/ruptela_definitions.yaml
	go run ./pkg/ruptela/codegen

generate-autopi: # Generate all files for autopi
	go run ./cmd/codegen -convert.package=autopi -generators=convert -convert.output-file=./pkg/autopi/convert_signal_funcs_gen.go -definitions=./pkg/autopi/schema/autopi_definitions.yaml
	go run ./cmd/codegen -generators=custom -custom.output-file=./pkg/autopi/convert_signal_status_gen.go -custom.template-file=./pkg/autopi/codegen/convert_signal_status.tmpl -custom.format=true -definitions=./pkg/autopi/schema/autopi_definitions.yaml

generate-lorawan: # Generate all files for lorawan (macaron)
	go run ./cmd/codegen -convert.package=lorawan -generators=convert -convert.output-file=./pkg/lorawan/vehicle-convert-funcs_gen.go -definitions=./pkg/lorawan/schema/lorawan-definitions.yaml
	go run ./cmd/codegen -generators=custom -custom.output-file=./pkg/lorawan/lorawan-convert_gen.go -custom.template-file=./pkg/lorawan/codegen/convert-status.tmpl -custom.format=true -definitions=./pkg/lorawan/schema/lorawan-definitions.yaml


generate-tesla: # Generate all files for tesla
	go run ./cmd/codegen -convert.package=tesla -generators=convert -convert.output-file=./pkg/tesla/convert_signal_funcs_gen.go -definitions=./pkg/tesla/schema/tesla_definitions.yaml
	go run ./cmd/codegen -generators=custom -custom.output-file=./pkg/tesla/convert_signal_status_gen.go -custom.template-file=./pkg/tesla/codegen/convert_signal_status.tmpl -custom.format=true -definitions=./pkg/tesla/schema/tesla_definitions.yaml

generate-compass: # Generate all files for compass
	go run ./cmd/codegen -convert.package=compass -generators=convert -convert.output-file=./pkg/compass/vehicle-convert-funcs_gen.go -definitions=./pkg/compass/schema/compass-definitions.yaml
	go run ./cmd/codegen -generators=custom -custom.output-file=./pkg/compass/compass-convert_gen.go -custom.template-file=./pkg/compass/codegen/convert-status.tmpl -custom.format=true -definitions=./pkg/compass/schema/compass-definitions.yaml
