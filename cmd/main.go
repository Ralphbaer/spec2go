package main

import (
	"fmt"
	"os"

	"github.com/Ralphbaer/spec2go/internal/api"
	"github.com/Ralphbaer/spec2go/internal/generator"
	"github.com/getkin/kin-openapi/openapi3"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: spec2go [path to OpenAPI spec file]")
		os.Exit(1)
	}

	specPath := os.Args[1]

	spec, err := openapi3.NewLoader().LoadFromFile(specPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing OpenAPI spec: %v\n", err)
		os.Exit(1)
	}

	data := api.PrepareData(spec)
	templateDir := "internal/template/templates"

	err = generator.GenerateCode(data, templateDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating code: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Code generation successful!")
}
