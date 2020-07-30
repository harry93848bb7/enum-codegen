package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/harry93848bb7/enum-codegen"
	"gopkg.in/yaml.v2"
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
	// Write output files
	if outputFile == "" {
		fmt.Println(codegen)
		return
	}
	if !strings.HasSuffix(outputFile, ".gen.go") {
		outputFile = strings.TrimSuffix(outputFile, ".go") + ".gen.go"
	}
	if err := ioutil.WriteFile(outputFile, []byte(codegen), os.ModePerm); err != nil {
		fmt.Println("Failed to write enum codegen to the specified file:", err)
		os.Exit(1)
	}
	if generatedTests {
		codegentest, err := enum.GenerateTests(&config)
		if err != nil {
			fmt.Println("Failed to run enum test codegen from enum yaml spec file:", err)
			os.Exit(1)
		}
		if err := ioutil.WriteFile(strings.TrimSuffix(outputFile, ".gen.go")+"_test.go", []byte(codegentest), os.ModePerm); err != nil {
			fmt.Println("Failed to write enum codegen tests to the specified file:", err)
			os.Exit(1)
		}
	}
}
