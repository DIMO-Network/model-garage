# AutoPi Module Guide

## Overview

The AutoPi module handles data conversion for AutoPi devices, which are Raspberry Pi-based aftermarket telematics devices that communicate with vehicles via OBD-II.

## Payload Versions

AutoPi payloads today come in the **V2 format**. Old payloads used to come in the **V1 format**, which is why we need to look at both in case we need to decode one of the old payloads for some reason.

## Code Generation

This module uses the same code generation pattern as other modules like Ruptela.

Signal mappings are defined in `schema/autopi_definitions.yaml`. To regenerate the conversion code:

```bash
make generate-autopi
```
