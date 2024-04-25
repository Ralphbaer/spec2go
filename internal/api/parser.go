package api

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// ParseOpenAPISpec parses an OpenAPI specification from the provided file path.
//
// filePath: Path to the OpenAPI specification file.
// Returns *openapi3.T and error.
func ParseOpenAPISpec(filePath string) (*openapi3.T, error) {
	loader := openapi3.NewLoader()

	spec, err := loader.LoadFromFile(filePath)
	if err != nil {
		return nil, err
	}

	return spec, nil
}
