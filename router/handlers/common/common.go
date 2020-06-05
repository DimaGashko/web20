package common

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"web20.tk/core/db"
)

const (
	PAGES_PATH  = "frontend/dist/templates/pages/"
	LAYOUT_PATH = "frontend/dist/templates/layout/"
)

var AppConfig struct {
	Port int       `json:"port"`
	Db   db.Config `json:"db"`
}

var layoutTemplates []string

type RouteHandler func(http.ResponseWriter, *http.Request, map[string]interface{}) (string, error)

type HttpHandler struct {
	HandlerFunc RouteHandler
}

func init() {
	loadLayoutTemplates()
}

func (h HttpHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	context := make(map[string]interface{})

	tmpl, err := h.handle(w, req, context)
	if err != nil {
		SendError(err, w, req, context)
		return
	}

	t, err := parseTemplate(tmpl)
	if err != nil {
		SendError(&AppError{err, http.StatusInternalServerError}, w, req, context)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	err = t.Execute(w, context)
	if err != nil {
		panic(err)
	}
}

func (h HttpHandler) handle(w http.ResponseWriter, req *http.Request, context map[string]interface{}) (string, error) {
	tmpl, err := h.HandlerFunc(w, req, context)
	if err != nil {
		return "", err
	}

	if tmpl == "" {
		panic("Empty template path")
	}

	return tmpl, nil
}

func parseTemplate(path string) (*template.Template, error) {
	t, err := template.ParseFiles(append(layoutTemplates, path)...)
	if err != nil {
		return nil, err
	}

	return t, err
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

	t, err := parseTemplate(tmplPath)
	if err != nil {
		panic(err)
	}

	err = t.Execute(w, context)
	if err != nil {
		panic(err)
	}
}

func Send404Error(res http.ResponseWriter, req *http.Request, context map[string]interface{}) (string, error) {
	return "", MakeError("", http.StatusNotFound)
}

func GetErrorCode(err error) int {
	appErr, ok := err.(*AppError)
	if !ok {
		return http.StatusInternalServerError
	}
	return appErr.Code
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
