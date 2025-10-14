# Native Status Module Guide

## Overview

The Native Status module handles conversion of DIMO's native status payload format.

**Note**: This module isn't really used anymore.

## Historical Context

This module was previously used when AutoPi and Macaron (HashDog) were both transformed and used by `benthos-plugin`. Now that everything goes through DIS (DIMO Ingestion Service), AutoPi and Macaron have their own distinct paths and modules that they use.

## Payload Versions

The module supports two payload versions:

- **V1** (`v1.0.0`, `v1.1.0`) - Original format
- **V2** (`v2.0.0`) - Enhanced format with additional fields

## Code Generation

If needed, regenerate the native status module code:

```bash
make generate-nativestatus
```
