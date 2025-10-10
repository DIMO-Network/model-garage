# Ruptela Module Guide

## Overview

The Ruptela module handles data conversion for Ruptela hardware devices. These are aftermarket telematics devices that communicate vehicle data via cellular networks.

## Module Structure

```
pkg/ruptela/
├── module.go                      # Main module implementation
├── module_test.go                 # Module tests
├── convert_signal.go              # Signal conversion entry point
├── convert_signal_status.go       # Status signal conversions
├── convert_signal_status_gen.go   # Generated status conversion dispatcher
├── convert_signal_location.go     # Location signal conversions
├── convert_signal_location_gen.go # Generated location conversion dispatcher
├── convert_signal_funcs_gen.go    # Generated conversion function stubs
├── convert_signal_dtc.go          # DTC (diagnostic trouble codes) conversions
├── convert_event.go               # Event conversions
├── convert_fingerprint.go         # VIN/fingerprint extraction
├── convert_cloudevent.go          # CloudEvent generation (not currently used)
├── models.go                      # Data models
├── ruptela.go                     # Ruptela protocol definitions
├── multiplier_offset.go           # Ruptela-specific data transformations
├── schema/
│   ├── ruptela_definitions.yaml   # Signal mappings
│   ├── oids.csv                   # OBD parameter IDs
│   └── schema.go                  # Schema loading
└── codegen/
    ├── main.go                    # Custom codegen entry point
    ├── ruptela.go                 # Codegen utilities
    └── *.tmpl                     # Code generation templates
```

## Key Concepts

### Data Schemas (DS Types)

Ruptela devices send data in different schemas identified by DS (Data Schema) types:

- **StatusEventDS** - Regular status updates with sensor data
- **LocationEventDS** - GPS location data
- **DTCEventDS** - Diagnostic trouble codes
- **DevStatusDS** - Device health status

### Signal Types

Ruptela provides data through multiple channels:

1. **OBD Signals** - Data from the vehicle's OBD-II port (prefixed with `signals.`)
2. **CAN Signals** - Direct CAN bus data (also prefixed with `signals.`)
3. **Position Data** - GPS data (prefixed with `pos.`)
4. **DTC Codes** - Diagnostic trouble codes (`dtc_codes`)

### Multiplier and Offset

Many Ruptela signals require multiplier and offset transformations defined in the OID CSV file. The `multiplier_offset.go` file handles these transformations.

> **⚠️ Important: Firmware Updates**
>
> Ruptela sometimes updates their firmware, which may include an updated FMIO list. When the FMIO list gets updated, the offset multipliers might also change.
>
> If Ruptela releases a firmware update with a new FMIO list:
>
> 1. Add the new FMIO list to the repo (update `schema/oids.csv`)
> 2. Rerun code generation: `make generate-ruptela`
> 3. This ensures all signals get decoded correctly with the updated multipliers and offsets

## Signal Mapping

To add a new signal mapping:

1. Edit `schema/ruptela_definitions.yaml`
2. Add the VSS signal name and the corresponding Ruptela signal ID
3. Run `make generate-ruptela`

Example:

```yaml
- vspecName: Vehicle.Speed
  conversions:
    - originalName: "signals.95" # OBD vehicle speed
      originalType: string
    - originalName: "pos.spd" # GPS speed
      originalType: float64
```

## Code Generation

The Ruptela module uses both standard codegen and custom codegen:

```bash
# Standard conversion function generation
go run ./cmd/codegen -convert.package=ruptela -generators=convert \
  -convert.output-file=./pkg/ruptela/convert_signal_funcs_gen.go \
  -definitions=./pkg/ruptela/schema/ruptela_definitions.yaml

# Custom status conversion dispatcher
go run ./cmd/codegen -generators=custom \
  -custom.output-file=./pkg/ruptela/convert_signal_status_gen.go \
  -custom.template-file=./pkg/ruptela/codegen/convert_signal_status.tmpl \
  -custom.format=true \
  -definitions=./pkg/ruptela/schema/ruptela_definitions.yaml

# Custom location conversion dispatcher
go run ./cmd/codegen -generators=custom \
  -custom.output-file=./pkg/ruptela/convert_signal_location_gen.go \
  -custom.template-file=./pkg/ruptela/codegen/convert_signal_location.tmpl \
  -custom.format=true \
  -definitions=./pkg/ruptela/schema/ruptela_definitions.yaml

# Custom Ruptela-specific generation
go run ./pkg/ruptela/codegen
```

Or simply run:

```bash
make generate-ruptela
```

## Interface Implementations

The Ruptela module implements all four module interfaces:

- **SignalModule** - Converts status, location, and DTC data to VSS signals
- **CloudEventModule** - Generates CloudEvents with proper subject/producer metadata
- **FingerprintModule** - Extracts VIN from OBD signals (104, 105, 106)
- **EventModule** - Converts vehicle events (harsh braking, etc.)

## Testing

Run tests with:

```bash
go test ./pkg/ruptela/...
```

Key test files:

- `module_test.go` - Integration tests
- `convert_signal_status_test.go` - Status signal conversion tests
- `convert_signal_location_test.go` - Location signal conversion tests
- `convert_signal_dtc_test.go` - DTC conversion tests
- `convert_event_test.go` - Event conversion tests
- `convert_fingerprint_test.go` - Fingerprint extraction tests
