package common

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"web20.tk/core/db"
	"web20.tk/templates"
)

var AppConfig struct {
	Port   int       `json:"port"`
	AppUrl string    `json:"app-url"`
	Db     db.Config `json:"db"`
}

type RouteHandler func(http.ResponseWriter, *http.Request, map[string]interface{}) (string, error)
type RouteApiHandler func(http.ResponseWriter, *http.Request) (interface{}, error)
type HandlerType string

type HttpHandler struct {
	HandlerFunc    RouteHandler
	ApiHandlerFunc RouteApiHandler
	Type           HandlerType
}

const (
	SSR_HANDLER HandlerType = "ssr"
	API_HANDLER HandlerType = "api"
)

func (h HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.Type == SSR_HANDLER {
		h.serveSsr(w, r)
	} else {
		h.serveApi(w, r)
	}
}

func (h HttpHandler) serveSsr(w http.ResponseWriter, r *http.Request) {
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
		SendError(err, w, r, context)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	err = t.ExecuteTemplate(w, "base.tmpl", context)
	if err != nil {
		panic(err)
	}
}

func (h HttpHandler) serveApi(w http.ResponseWriter, r *http.Request) {
	context, err := h.ApiHandlerFunc(w, r)
	if err != nil {
		SendApiError(err, w, r, context)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	resp, err := json.Marshal(context)
	if err != nil {
		SendApiError(err, w, r, context)
	}

	w.Write(resp)
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

	var tmpl string
	switch code {
	case http.StatusForbidden:
		tmpl = templates.PATH + "err403.tmpl"
	case http.StatusNotFound:
		tmpl = templates.PATH + "err404.tmpl"
	default:
		tmpl = templates.PATH + "err500.tmpl"
	}

	t, err := templates.ParseTemplate(tmpl)
	if err != nil {
		panic(err)
	}

	err = t.ExecuteTemplate(w, "base.tmpl", context)
	if err != nil {
		panic(err)
	}
}

func SendApiError(err error, w http.ResponseWriter, r *http.Request, context interface{}) {
	code := GetErrorCode(err)
	w.WriteHeader(code)

	fmt.Printf("Couldn't load page: %s ({ msg: '%s', code: %d })", r.URL, err.Error(), code)

	resp, err := json.Marshal(err)
	if err != nil {
		SendApiError(err, w, r, context)
	}

	w.Write(resp)
}

func Send404Error(res http.ResponseWriter, r *http.Request, context map[string]interface{}) (string, error) {
	return "", New404()
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
