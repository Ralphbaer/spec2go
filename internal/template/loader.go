package template

import (
	"os"
	"path/filepath"
	"text/template"
)

// LoadTemplates loads templates from the specified directory.
func LoadTemplates(dir string) (*template.Template, error) {
	tmpl := template.New("base")

	if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".mustache" {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			_, err = tmpl.New(filepath.Base(path)).Parse(string(content))
			if err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return tmpl, nil
}
