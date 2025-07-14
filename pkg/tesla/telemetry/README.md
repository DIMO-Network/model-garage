The generator is a program in `codegen`. Running `make generate` should run this, along with the all of the other generators in this repository.

The core configuration is in [`schema/telemetry_definitions.yaml`](schema/telemetry_definitions.yaml). The top level is a sequence of mappings, each with the following fields:

| Name | Type | Required | Description | Example |
|-|-|-|-|-|
| `teslaField` | string | yes | Value of the Tesla [`Field` enum](https://github.com/teslamotors/fleet-telemetry/blob/7c3187a18777c24d096059e94ab91245da92cb64/protos/vehicle_data.proto#L11). | `DriveRail` |
| `teslaType` | string | yes | Standard [protobuf type](https://github.com/teslamotors/fleet-telemetry/blob/7c3187a18777c24d096059e94ab91245da92cb64/protos/vehicle_data.proto#L695) for `value` in `Datum`s with the given field as its key. | `string` |
| `teslaUnit` | string | no | VSS unit for the Tesla value, which should be numeric | `mph` |
| `vssSignals` | sequence of strings | no | Full paths for the target VSS signals. | `[Vehicle.Powertrain.TractionBattery.Charging.IsCharging]` |

The code generator tries to be helpful in two ways:

1. If a unit is specified on both Tesla and VSS sides, then we'll attempt to automatically convert the units. So, for example, if `teslaUnit` is `mi` and the VSS unit is `km` then `unit.MilesToKilometers` will be applied.
2. If we know how to parse `teslaType` from a string, then we'll accept both for the field. For example, if `VehicleSpeed` has Tesla type `double` and a value of `"2.5"` comes in then we'll convert this to the floating point number `2.5`.

Each combination of Tesla field and VSS signal gets an entry in [`inner_convert_funcs_gen.go`](inner_convert_funcs_gen.go), and edits to these functions will be preserved across invocations of `make generate`.
