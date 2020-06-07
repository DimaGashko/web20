package templates

import (
	"html/template"
)

const (
	TMPL_PATTERN = "frontend/dist/templates/layout/*.tmpl"
	PATH         = "frontend/dist/templates/pages/"
)

func ParseTemplate(path string) (*template.Template, error) {
	t, err := template.New("base.tmpl").Funcs(getFuncMap()).ParseGlob(TMPL_PATTERN)
	if err != nil {
		return nil, err
	}

	t, err = t.ParseFiles(path)
	if err != nil {
		return nil, err
	}

	return t, err
}
