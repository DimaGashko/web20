package home

import (
	"html/template"
	"net/http"
)

func Init(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("frontend/dist/templates/home.tmpl")
	if err != nil {
		panic(err)
	}

	err = t.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
