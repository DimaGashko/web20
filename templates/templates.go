package templates

import (
	"html/template"
	"os"
	"path/filepath"
)

const (
	PAGES_PATH  = "frontend/dist/templates/pages/"
	LAYOUT_PATH = "frontend/dist/templates/layout/"
)

var layoutTemplates []string

func init() {
	loadLayoutTemplates()
}

func ParseTemplate(path string) (*template.Template, error) {
	t, err := template.New(filepath.Base(layoutTemplates[0])).Funcs(getFuncMap()).ParseFiles(append(layoutTemplates, path)...)
	if err != nil {
		return nil, err
	}

	return t, err
}

func loadLayoutTemplates() {
	err := filepath.Walk(LAYOUT_PATH, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() {
			layoutTemplates = append(layoutTemplates, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
