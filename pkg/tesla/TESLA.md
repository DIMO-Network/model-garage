# Tesla Module Guide

## Overview

The Tesla module handles data conversion for Tesla vehicles. This module is unique because it handles two completely different payload formats with separate code paths.

## Two Payload Types

### 1. Tesla API (`api/`) - Older Models

Used when we poll older Tesla models via the Tesla API.

- **Format**: JSON
- **Payload Size**: Large payloads containing many signals in a single payload
- **Structure**: Nested JSON with multiple vehicle states bundled together
- **Code Path**: Uses standard code generation pattern (similar to Ruptela)

### 2. Fleet Telemetry (`telemetry/`) - Newer Streaming Service

Used with Tesla's newer Fleet Telemetry streaming service.

- **Format**: Protobuf
- **Payload Size**: Small chunks with 1-5 signals per payload
- **Structure**: High-frequency streaming updates
- **Code Path**: Uses a very different conversion approach with custom protobuf parsing

## Module Structure

```
pkg/tesla/
├── api/                           # Tesla API (JSON, older models)
│   ├── module.go
│   ├── convert_signal_funcs_gen.go
│   ├── convert_signal_status_gen.go
│   └── schema/
│       └── tesla_definitions.yaml
└── telemetry/                     # Fleet Telemetry (Protobuf, streaming)
    ├── module.go
    ├── inner_convert_funcs_gen.go
    ├── outer_convert_funcs_gen.go
    └── schema/
        └── tesla_telemetry_definitions.yaml
```

## Code Generation

Generate Tesla module code:

```bash
make generate-tesla
```

This runs both API codegen and Fleet Telemetry codegen.
