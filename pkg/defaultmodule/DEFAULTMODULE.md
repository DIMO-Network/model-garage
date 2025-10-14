# Default Module Guide

## Overview

The Default Module is used when a data producer doesn't want to create a custom Go package for their decoding logic. This module is important because it allows any Oracle or data producer to send vehicle data to DIMO as long as they format their payloads according to the expected structure.

Unlike other modules (AutoPi, Ruptela, etc.) which require custom conversion logic and code generation, the Default Module expects data to already be in a standardized format.

## How It Works

The Default Module expects incoming CloudEvents to contain data in a specific JSON structure. Producers simply need to:

1. Use VSS signal names directly in their payload
2. Format their data according to the expected schema
3. Send the data - no custom conversion code needed

Things will naturally get converted as long as the payload matches the expected format.

## Expected Payload Format

### Signals

To send vehicle signals, use the `signals` array:

```json
{
  "signals": [
    {
      "name": "speed",
      "value": 65.5,
      "timestamp": "2024-12-01T15:31:12.378075897Z"
    },
    {
      "name": "powertrainType",
      "value": "HYBRID",
      "timestamp": "2024-12-01T15:31:12.378075897Z"
    }
  ]
}
```

- **name**: Must be a valid VSS signal name (e.g., `speed`, `powertrainType`, `currentLocationLatitude`)
- **value**: The signal value (number or string, depending on the signal type)
- **timestamp**: RFC3339 timestamp when the data was collected

### Events

To send vehicle events, use the `events` array:

```json
{
  "events": [
    {
      "name": "harsh_braking",
      "timestamp": "2024-12-01T15:31:12.378075897Z",
      "durationNs": 0,
      "metadata": "{\"side\":\"left\"}",
      "tags": ["behavior.harsh_acceleration"]
    }
  ]
}
```

- **name**: Event name
- **timestamp**: RFC3339 timestamp when the event occurred
- **durationNs**: Optional duration of the event in nanoseconds
- **metadata**: Optional JSON string with additional event data
- **tags**: Optional array of event tags

### Fingerprint (VIN)

To send vehicle identification, include a `vin` field:

```json
{
  "vin": "1HGBH41JXMN109186",
  "signals": [...]
}
```

### Complete Example

```json
{
  "id": "2pcYwspbaBFJ7NPGZ2kivkuJ12a",
  "source": "0x1234567890123456789012345678901234567890",
  "producer": "did:erc721:80003:0x78513c8CB4D6B6079f813850376bc9c7fc8aE67f:12",
  "specversion": "1.0",
  "subject": "did:erc721:80003:0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8:15",
  "time": "2024-12-01T15:31:12.378075897Z",
  "type": "dimo.status",
  "data": {
    "vin": "1HGBH41JXMN109186",
    "signals": [
      {
        "name": "speed",
        "value": 65.5,
        "timestamp": "2024-12-01T15:31:12.378075897Z"
      },
      {
        "name": "currentLocationLatitude",
        "value": 37.7749,
        "timestamp": "2024-12-01T15:31:12.378075897Z"
      }
    ],
    "events": [
      {
        "name": "harsh_braking",
        "timestamp": "2024-12-01T15:31:12.378075897Z",
        "durationNs": 0,
        "metadata": "{\"severity\":\"high\"}",
        "tags": ["behavior.harsh_acceleration"]
      }
    ]
  }
}
```

## Interface Implementations

The Default Module implements all four module interfaces:

- **SignalModule** - Converts signal data to VSS signals
- **CloudEventModule** - Automatically determines CloudEvent types based on payload content
- **FingerprintModule** - Extracts VIN from the data
- **EventModule** - Converts event data to VSS events

## Validation

The module validates:

- Signal names must be defined VSS signals
- Signal values must match the expected type (number or string)
- Event timestamps must not be zero
- Event metadata (if provided) must be valid JSON
- Event tags must be defined in the event tag schema

## No Code Generation Required

Unlike other modules, the Default Module does not require code generation. It dynamically validates signal names against the VSS schema at runtime.

## When to Use

Use the Default Module when:

- You're building a new data producer and want to get started quickly
- Your device can easily format data into the expected structure
- You don't need complex conversion logic or data transformations

Use a custom module (like Ruptela, AutoPi, etc.) when:

- You're working with existing hardware with a fixed data format
- You need complex data transformations or calculations
