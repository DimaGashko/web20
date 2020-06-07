package info

import (
	"net/http"

	"web20.tk/templates"
)

func About(w http.ResponseWriter, r *http.Request, context map[string]interface{}) (string, error) {
	return templates.PATH + "about.tmpl", nil
}

func ContactUs(w http.ResponseWriter, r *http.Request, context map[string]interface{}) (string, error) {
	return templates.PATH + "contact-us.tmpl", nil
}
