package common

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"web20.tk/core/db"
	"web20.tk/templates"
)

var AppConfig struct {
	Port int       `json:"port"`
	Db   db.Config `json:"db"`
}

type RouteHandler func(http.ResponseWriter, *http.Request, map[string]interface{}) (string, error)

type HttpHandler struct {
	HandlerFunc RouteHandler
}

func (h HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := make(map[string]interface{})
	err := initLayout(w, r, context)
	if err != nil {
		SendError(err, w, r, context)
		return
	}

	tmpl, err := h.handle(w, r, context)
	if err != nil {
		SendError(err, w, r, context)
		return
	}

	t, err := templates.ParseTemplate(tmpl)
	if err != nil {
		SendError(&AppError{err, http.StatusInternalServerError}, w, r, context)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	err = t.Execute(w, context)
	if err != nil {
		panic(err)
	}
}

func (h HttpHandler) handle(w http.ResponseWriter, r *http.Request, context map[string]interface{}) (string, error) {
	tmpl, err := h.HandlerFunc(w, r, context)
	if err != nil {
		return "", err
	}

	if tmpl == "" {
		panic("Empty template path")
	}

	return tmpl, nil
}

func SendError(err error, w http.ResponseWriter, r *http.Request, context map[string]interface{}) {
	code := GetErrorCode(err)
	w.WriteHeader(code)

	fmt.Printf("Couldn't load page: %s ({ msg: '%s', code: %d })", r.URL, err.Error(), code)

	var tmplPath string
	switch code {
	case http.StatusForbidden:
		tmplPath = templates.PAGES_PATH + "errors/err403/err403.tmpl"
	case http.StatusNotFound:
		tmplPath = templates.PAGES_PATH + "errors/err404/err404.tmpl"
	default:
		tmplPath = templates.PAGES_PATH + "errors/err500/err500.tmpl"
	}

	t, err := templates.ParseTemplate(tmplPath)
	if err != nil {
		panic(err)
	}

	err = t.Execute(w, context)
	if err != nil {
		panic(err)
	}
}

func Send404Error(res http.ResponseWriter, r *http.Request, context map[string]interface{}) (string, error) {
	return "", MakeError("", http.StatusNotFound)
}

func GetErrorCode(err error) int {
	appErr, ok := err.(*AppError)
	if !ok {
		return http.StatusInternalServerError
	}
	return appErr.Code
}

func initLayout(w http.ResponseWriter, r *http.Request, context map[string]interface{}) error {
	router := mux.CurrentRoute(r)
	routeName := "unnamed"
	if router != nil {
		routeName = router.GetName()
	}

	context["route"] = routeName
	context["year"] = time.Now().Year()
	return nil
}
