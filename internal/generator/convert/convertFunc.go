package convert

// This package provides functions for converting data between different formats.

import (
	"bytes"
	"cmp"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"text/template"

	"github.com/DIMO-Network/model-garage/pkg/codegen"
	"github.com/DIMO-Network/model-garage/pkg/schema"
)

type FunctionInfo struct {
	Comments string
	Body     []byte
}

// getConversionFunctions returns the signals that need conversion functions.
func getConversionFunctions(signals []*schema.SignalInfo) []funcTmplData {
	var convertFunc []funcTmplData
	for _, signal := range signals {
		for i := range signal.Conversions {
			funcName := "To" + signal.GOName + strconv.Itoa(i)
			convData := funcTmplData{
				Signal:     signal,
				Conversion: signal.Conversions[i],
				FuncName:   funcName,
			}

			convertFunc = append(convertFunc, convData)
		}
	}
	return convertFunc
}

// createConvertFuncTemplate creates a template for generating conversion functions.
func createConvertFuncTemplate() (*template.Template, error) {
	tmpl, err := template.New("convertFuncTemplate").Parse(convertFuncTemplateStr)
	if err != nil {
		return nil, fmt.Errorf("error parsing go struct template: %w", err)
	}
	return tmpl, nil
}

// GetDeclaredFunctions returns a map of function names to their corresponding function information for a given directory.
// The function information includes comments and body.
func GetDeclaredFunctions(outputPath string) (map[string]FunctionInfo, error) {
	fset := token.NewFileSet()
	declaredFunctions := make(map[string]FunctionInfo)

	list, err := os.ReadDir(outputPath)
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %w", err)
	}

	for _, d := range list {
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".go") {
			continue
		}
		filename := filepath.Join(outputPath, d.Name())
		fileDeclaredFunctions, err := getDeclaredFunctionsForFile(fset, filename)
		if err != nil {
			return nil, err
		}
		for k, v := range fileDeclaredFunctions {
			declaredFunctions[k] = v
		}
	}

	return declaredFunctions, nil
}

func getDeclaredFunctionsForFile(fset *token.FileSet, filePath string) (map[string]FunctionInfo, error) {
	declaredFunctions := make(map[string]FunctionInfo)
	src, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("error parsing file: %w", err)
	}
	for _, decl := range src.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Recv != nil {
			continue
		}

		var docComments []string
		if fn.Doc != nil {
			for _, comment := range fn.Doc.List {
				docComments = append(docComments, comment.Text)
			}
		}
		// set comments to nil to avoid printing them in the body
		fn.Doc = nil

		// Capture function body including comments
		var buf bytes.Buffer
		err = format.Node(&buf, fset, &printer.CommentedNode{
			Node:     fn.Body,
			Comments: src.Comments,
		})
		if err != nil {
			return nil, fmt.Errorf("error formating function: %w", err)
		}

		declaredFunctions[fn.Name.Name] = FunctionInfo{
			Comments: strings.Join(docComments, "\n"),
			Body:     buf.Bytes(),
		}
	}
	return declaredFunctions, nil
}

// writeConvertFuncs writes the generated conversion functions to a file.
func writeConvertFuncs(convertFunc []funcTmplData, existingFuncs map[string]FunctionInfo, tmpl *template.Template, outputPath string, packageName string, copyComments bool) error {
	var convertBuff bytes.Buffer
	fmt.Fprintf(&convertBuff, header, packageName)
	slices.SortStableFunc(convertFunc, func(a, b funcTmplData) int {
		// split funcName to get digits at the end and compare the name then by the digit value
		// get the function name without the digits at the end
		aFuncName := a.FuncName
		aDigits := 0
		bFuncName := b.FuncName
		bDigits := 0
		for i := len(aFuncName) - 1; i >= 0; i-- {
			if aFuncName[i] < '0' || aFuncName[i] > '9' {
				aFuncName = aFuncName[:i+1]
				aDigits, _ = strconv.Atoi(a.FuncName[i+1:])
				break
			}
		}
		for i := len(bFuncName) - 1; i >= 0; i-- {
			if bFuncName[i] < '0' || bFuncName[i] > '9' {
				bFuncName = bFuncName[:i+1]
				bDigits, _ = strconv.Atoi(b.FuncName[i+1:])
				break
			}
		}
		val := cmp.Compare(aFuncName, bFuncName)
		if val != 0 {
			return val
		}
		return cmp.Compare(aDigits, bDigits)
	})

	// Add or update existing functions
	for _, convData := range convertFunc {
		funcName := convData.FuncName
		if fnInfo, exists := existingFuncs[funcName]; exists {
			convData.Body = string(fnInfo.Body)
			if copyComments {
				convData.DocComment = fnInfo.Comments
			}
		}

		err := tmpl.Execute(&convertBuff, convData)
		if err != nil {
			return fmt.Errorf("error executing template for function %s: %w", funcName, err)
		}
	}

	err := codegen.FormatAndWriteToFile(convertBuff.Bytes(), outputPath)
	if err != nil {
		return fmt.Errorf("error formatting and writing to file: %w", err)
	}
	return nil
}
