package common

import (
	"fmt"
	"html/template"
	"net/http"

	"web20.tk/core/db"
)

const (
	PAGES_PATH   = "frontend/dist/templates/pages/"
	LAYOUTS_PATH = "frontend/dist/templates/layouts/"
	DEF_LAYOUT   = LAYOUTS_PATH + "base-layout/base-layout.tmpl"
)

var AppConfig struct {
	Port int       `json:"port"`
	Db   db.Config `json:"db"`
}

type RouteHandler func(http.ResponseWriter, *http.Request, map[string]interface{}) (string, string, error)

type HttpHandler struct {
	HandlerFunc RouteHandler
}

func (h HttpHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	context := make(map[string]interface{})

	tmpl, layout, err := h.handle(w, req, context)
	if err != nil {
		SendError(err, w, req, context)
		return
	}

	t, err := template.ParseFiles(layout, tmpl)
	if err != nil {
		SendError(&AppError{err, http.StatusInternalServerError}, w, req, context)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	err = t.Execute(w, context)
	if err != nil {
		panic(err)
	}
}

func (h HttpHandler) handle(w http.ResponseWriter, req *http.Request, context map[string]interface{}) (string, string, error) {
	tmpl, layout, err := h.HandlerFunc(w, req, context)
	if err != nil {
		return "", "", err
	}

	if tmpl == "" {
		panic("Empty template path")
	}

	if layout == "" {
		layout = DEF_LAYOUT
	}

	return tmpl, layout, nil
}

func SendError(err error, w http.ResponseWriter, req *http.Request, context map[string]interface{}) {
	code := GetErrorCode(err)
	w.WriteHeader(code)

	fmt.Printf("Couldn't load page: %s ({ msg: '%s', code: %d })", req.URL, err.Error(), code)

	var tmplPath string
	switch code {
	case http.StatusForbidden:
		tmplPath = PAGES_PATH + "errors/err403/err403.tmpl"
	case http.StatusNotFound:
		tmplPath = PAGES_PATH + "errors/err404/err404.tmpl"
	default:
		tmplPath = PAGES_PATH + "errors/err500/err500.tmpl"
	}

	t, err := template.ParseFiles(DEF_LAYOUT, tmplPath)
	if err != nil {
		panic(err)
	}

	err = t.Execute(w, context)
	if err != nil {
		panic(err)
	}
}

func Send404Error(res http.ResponseWriter, req *http.Request, context map[string]interface{}) (string, string, error) {
	return "", "", MakeError("", http.StatusNotFound)
}

func GetErrorCode(err error) int {
	appErr, ok := err.(*AppError)
	if !ok {
		return http.StatusInternalServerError
	}
	return appErr.Code
}
