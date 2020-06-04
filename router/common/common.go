package common

import (
	"errors"
	"html/template"
	"net/http"
)

type HttpHandler struct {
	HandlerFunc func(http.ResponseWriter, *http.Request, *map[string]interface{}) (string, *AppError)
}

func (h HttpHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var context map[string]interface{}

	tmplPath, appErr := h.HandlerFunc(w, req, &context)
	if appErr != nil {
		SendError(appErr, w)
		return
	}

	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		SendError(&AppError{err, http.StatusInternalServerError}, w)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	err = t.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func SendError(appErr *AppError, w http.ResponseWriter) {
	w.WriteHeader(appErr.Code)

	var tmplPath string
	switch appErr.Code {
	case http.StatusForbidden:
		tmplPath = "frontend/dist/templates/err403.tmpl"
	case http.StatusNotFound:
		tmplPath = "frontend/dist/templates/err404.tmpl"
	default:
		tmplPath = "frontend/dist/templates/err500.tmpl"
	}

	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		panic(err)
	}

	err = t.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func Send403Error(res http.ResponseWriter, req *http.Request, context *map[string]interface{}) (string, *AppError) {
	return "", &AppError{errors.New("Forbidden!"), http.StatusForbidden}
}

func Send404Error(res http.ResponseWriter, req *http.Request, context *map[string]interface{}) (string, *AppError) {
	return "", &AppError{errors.New("Page not found"), http.StatusNotFound}
}

func Send500Error(res http.ResponseWriter, req *http.Request, context *map[string]interface{}) (string, *AppError) {
	return "", &AppError{errors.New("Internal server error"), http.StatusInternalServerError}
}
