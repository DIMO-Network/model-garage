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
