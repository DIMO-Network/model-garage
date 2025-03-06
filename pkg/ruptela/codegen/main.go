//nolint:all // testFile
package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"math/big"
	"os"
	"strings"
	"text/template"

	"github.com/DIMO-Network/model-garage/pkg/codegen"
	rupschema "github.com/DIMO-Network/model-garage/pkg/ruptela/schema"
	"github.com/DIMO-Network/model-garage/pkg/schema"
)

//go:embed functions.tmpl
var functionsTemplate string

func main() {
	oidMap, err := loadCSVToMap(rupschema.OIDCSV())
	if err != nil {
		panic(err)
	}

	vssReader := strings.NewReader(schema.VssRel42DIMO())
	defReader := strings.NewReader(rupschema.RuptelaDefinitionsYAML())
	sigs, err := schema.GetDefinedSignals(vssReader, defReader)
	if err != nil {
		panic(err)
	}
	createRecords := make(map[string]Record)
	for _, sig := range sigs.Signals {
		for _, conv := range sig.Conversions {
			parts := strings.Split(conv.OriginalName, ".")
			if len(parts) < 1 || parts[0] != "signals" {
				continue
			}
			oid := parts[1]
			record, ok := oidMap[oid]
			if !ok {
				continue
			}
			offset, multiplier, err := getMultiplierAndOffset(record.MultiplierOffset)
			if err != nil {
				panic(err)
			}
			record.Multiplier = multiplier
			record.Offset = offset
			record.ErrorRange, record.ErrorSet, err = getErrorRange(record.ErrorValues)
			if err != nil {
				panic(err)
			}
			record.MinBig, err = getMinOrMax(record.MinValue)
			if err != nil {
				panic(err)
			}
			record.MaxBig, err = getMinOrMax(record.MaxValue)
			if err != nil {
				panic(err)
			}
			createRecords[oid] = record
		}
	}
	err = createFuncs(createRecords)
	if err != nil {
		panic(err)
	}
	return
}

func createFuncs(records map[string]Record) error {
	tmpl, err := template.New("ruptela-convert-functions").Funcs(
		template.FuncMap{
			"bigText": bigText,
		},
	).Parse(functionsTemplate)
	outputFile := "pkg/ruptela/multiplier_offset.go"
	customFileOutputFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("error creating Custom output file: %w", err)
	}
	defer func() {
		if cerr := customFileOutputFile.Close(); err == nil && cerr != nil {
			err = cerr
		}
	}()

	var outBuf bytes.Buffer
	data := struct {
		Records map[string]Record
	}{
		Records: records,
	}
	err = tmpl.Execute(&outBuf, &data)
	if err != nil {
		return fmt.Errorf("error executing Custom template: %w", err)
	}
	err = codegen.FormatAndWriteToFile(outBuf.Bytes(), outputFile)
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}
	return nil

}

func bigText(b *big.Int) string {
	return b.Text(10)
}
