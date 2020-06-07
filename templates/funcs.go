package templates

import (
	"html/template"
	"time"
)

func getFuncMap() template.FuncMap {
	return template.FuncMap{
		"makeSlice": makeSlice,
		"curYear":   curYear,
	}
}

func makeSlice(len int) []struct{} {
	return make([]struct{}, len)
}

func curYear() int {
	return time.Now().Year()
}
