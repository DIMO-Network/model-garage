// Package main contains the code generation command for transforming a
// definitions file into code that executes a conversion from Tesla
// Fleet Telemetry protobufs into VSS signals.
package main

import "log"

const (
	defaultPackage        = "telemetry"
	defaultOuterFuncsPath = "pkg/tesla/telemetry/outer_convert_funcs_gen.go"
	defaultInnerFuncsPath = "pkg/tesla/telemetry/inner_convert_funcs_gen.go"
)

func main() {
	err := Generate(defaultPackage, defaultOuterFuncsPath, defaultInnerFuncsPath)
	if err != nil {
		log.Fatalf("Failure: %v", err)
	}
}
