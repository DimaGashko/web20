package common

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
)

const BASE_LAYOUT_PATH = "frontend/dist/templates/base-layout.tmpl"

var AppConfig struct {
	Port int `json:"port"`
	Db   struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"pass"`
		Name     string `json:"name"`
	} `json:"db"`
}

type HttpHandler struct {
	HandlerFunc func(http.ResponseWriter, *http.Request, map[string]interface{}) (string, error)
}

func (h HttpHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	context := make(map[string]interface{})

	tmplPath, err := h.HandlerFunc(w, req, context)
	if err != nil {
		SendError(err, w, context)
		return
	}

	t, err := template.ParseFiles(BASE_LAYOUT_PATH, tmplPath)
	if err != nil {
		SendError(&AppError{err, http.StatusInternalServerError}, w, context)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	err = t.Execute(w, context)
	if err != nil {
		panic(err)
	}
}

func SendError(err error, w http.ResponseWriter, context map[string]interface{}) {
	fmt.Print(err.Error())
	code := GetErrorCode(err)
	w.WriteHeader(code)

	var tmplPath string
	switch code {
	case http.StatusForbidden:
		tmplPath = "frontend/dist/templates/err403.tmpl"
	case http.StatusNotFound:
		tmplPath = "frontend/dist/templates/err404.tmpl"
	default:
		tmplPath = "frontend/dist/templates/err500.tmpl"
	}

	t, err := template.ParseFiles(BASE_LAYOUT_PATH, tmplPath)
	if err != nil {
		panic(err)
	}

	err = t.Execute(w, context)
	if err != nil {
		panic(err)
	}
}

func Send403Error(res http.ResponseWriter, req *http.Request, context map[string]interface{}) (string, error) {
	return "", &AppError{errors.New("Forbidden!"), http.StatusForbidden}
}

func Send404Error(res http.ResponseWriter, req *http.Request, context map[string]interface{}) (string, error) {
	return "", &AppError{errors.New("Page not found"), http.StatusNotFound}
}

func Send500Error(res http.ResponseWriter, req *http.Request, context map[string]interface{}) (string, error) {
	return "", &AppError{errors.New("Internal server error"), http.StatusInternalServerError}
}

func GetErrorCode(err error) int {
	appErr, ok := err.(*AppError)
	if !ok {
		return http.StatusInternalServerError
	}
	return appErr.Code
}
