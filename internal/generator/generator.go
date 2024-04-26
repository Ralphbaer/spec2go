package generator

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/cbroglie/mustache"

	"github.com/Ralphbaer/spec2go/pkg/model"
)

// GenerateCode generates code from given template data using Mustache templates.
func GenerateCode(data *model.TemplateData, templateDir string) error {
	templatePath := filepath.Join(templateDir, "main.mustache")

	// Render the template with the provided data
	output, err := mustache.RenderFile(templatePath, data)
	if err != nil {
		return err
	}

	// Prepare the buffer to hold the generated code
	var buf bytes.Buffer
	buf.WriteString(output)

	outputDir := "./output"
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		return err
	}

	// Write the generated code to a file
	filePath := filepath.Join(outputDir, "main.go")
	return os.WriteFile(filePath, buf.Bytes(), 0o644)
}
