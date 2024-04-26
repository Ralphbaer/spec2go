package api

import (
	"reflect"
	"strings"

	"github.com/Ralphbaer/spec2go/pkg/model"
	"github.com/getkin/kin-openapi/openapi3"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// PrepareData generates template data from the provided OpenAPI specification.
//
// Parameter:
// - spec: The OpenAPI specification to be processed.
// Return type:
// - *model.TemplateData: The generated template data.
func PrepareData(spec *openapi3.T) *model.TemplateData {
	apis := make([]model.API, 0)

	// Use reflection to access the private map 'm' from spec.Paths
	pathsValue := reflect.ValueOf(spec.Paths)
	mValue := reflect.Indirect(pathsValue).FieldByName("m")

	c := cases.Title(language.English)
	if mValue.IsValid() && mValue.CanInterface() && mValue.Interface() != nil {
		for iter := mValue.MapRange(); iter.Next(); {
			path := iter.Key().String()
			pathItem := iter.Value().Interface().(*openapi3.PathItem)
			if pathItem == nil {
				continue
			}
			api := model.API{
				ClassName: cleanClassName(c, path),
				HasMore:   false,
			}
			apis = append(apis, api)
		}

		for i := range apis {
			if i < len(apis)-1 {
				apis[i].HasMore = true
			}
		}
	}

	return &model.TemplateData{
		Apis:       apis,
		ServerPort: "8080",
	}
}

func cleanClassName(c cases.Caser, path string) string {
	var className string
	parts := strings.Split(path, "/")
	for _, part := range parts {
		if part != "" {
			className += c.String(strings.ToLower(part))
		}
	}
	return className
}
