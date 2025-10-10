# Compass Module Guide

## Overview

The Compass module handles data conversion for Compass IoT telematics devices.

**Note**: We don't currently use this module in production.

## Code Generation

This module uses the same code generation pattern as other modules like Ruptela.

Signal mappings are defined in `schema/compass_definitions.yaml`. To regenerate the conversion code:

```bash
make generate-compass
```
