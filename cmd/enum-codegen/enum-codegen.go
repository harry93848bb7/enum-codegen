package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/harry93848bb7/enum-codegen"
	"gopkg.in/yaml.v2"
	"mvdan.cc/gofumpt/format"
)

func main() {
	var (
		inputFile      string
		outputFile     string
		packageName    string
		generatedTests bool
	)
	flag.StringVar(&inputFile, "in", "", `Enum yaml spec filepath input`)
	flag.StringVar(&outputFile, "out", "", `Where to output generated code, stdout is default`)
	flag.StringVar(&packageName, "package", "main", `The package name for generated code`)
	flag.BoolVar(&generatedTests, "tests", true, `Generate the type tests alongside the types`)
	flag.Parse()

	// Load and parse the enum specification file
	if inputFile == "" {
		fmt.Println("Please specify a path to a enum yaml spec file")
		os.Exit(1)
	}
	b, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Failed to load enum yaml spec file:", err)
		os.Exit(1)
	}
	var config enum.Options
	if err := yaml.Unmarshal(b, &config); err != nil {
		fmt.Println("Failed to parse enum yaml spec file:", err)
		os.Exit(1)
	}
	config.PackageName = packageName

	// Run the codegen with the parsed config
	codegen, err := enum.Generate(&config)
	if err != nil {
		fmt.Println("Failed to run enum codegen from enum yaml spec file:", err)
		os.Exit(1)
	}
	formatted, err := format.Source(codegen, format.Options{
		LangVersion: "1.16.2",
	})
	if err != nil {
		fmt.Println("Failed to format generated code:", err)
		os.Exit(1)
	}
	// Write output files
	if outputFile == "" {
		fmt.Println(string(formatted))
		return
	}
	if !strings.HasSuffix(outputFile, ".gen.go") {
		outputFile = strings.TrimSuffix(outputFile, ".go") + ".gen.go"
	}
	if err := ioutil.WriteFile(outputFile, formatted, os.ModePerm); err != nil {
		fmt.Println("Failed to write enum codegen to the specified file:", err)
		os.Exit(1)
	}
	if generatedTests {
		codegentest, err := enum.GenerateTests(&config)
		if err != nil {
			fmt.Println("Failed to run enum test codegen from enum yaml spec file:", err)
			os.Exit(1)
		}
		formatted, err := format.Source(codegentest, format.Options{
			LangVersion: "1.16.2",
		})
		if err != nil {
			fmt.Println("Failed to format generated code tests:", err)
			os.Exit(1)
		}
		if err := ioutil.WriteFile(strings.TrimSuffix(outputFile, ".gen.go")+"_test.go", formatted, os.ModePerm); err != nil {
			fmt.Println("Failed to write enum codegen tests to the specified file:", err)
			os.Exit(1)
		}
	}
}
