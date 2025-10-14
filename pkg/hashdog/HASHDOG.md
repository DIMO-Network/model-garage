# HashDog Module Guide

## Overview

The HashDog module (also known as Macaron) handles data conversion for HashDog LoRaWAN-based telematics devices.

## Current Implementation

HashDog devices use LoRaWAN (Long Range Wide Area Network) for low-power, long-range communication. Currently, the `ingest-lorawan` service (a Benthos instance) performs the initial conversion of the raw encoded binary payloads, then sends the data to DIS in a format closer to AutoPi.

This module then handles the conversion of that pre-processed data.

## Code Generation

This module uses the same code generation pattern as other modules like Ruptela.

Signal mappings are defined in `schema/hashdog_definitions.yaml`. To regenerate the conversion code:

```bash
make generate-hashdog
```

## Future Improvement

In a perfect world, this module would handle the actual raw payloads that come from our Macaron devices (encoded binary) and convert those directly as needed, rather than relying on the `ingest-lorawan` service for the initial conversion.
