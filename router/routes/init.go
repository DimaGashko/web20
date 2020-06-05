package router

import (
	"net/http"

	"web20.tk/router/handlers/common"
	"web20.tk/router/handlers/editor"
	"web20.tk/router/handlers/home"

	"github.com/gorilla/mux"
)

func Init(r *mux.Router) {
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("frontend/dist/static"))))
	r.NotFoundHandler = h(common.Send404Error)

	r.Handle("/", h(home.Home)).Methods("GET", "HEAD").Name("home")
	r.Handle("/add-article", h(editor.AddArticle)).Methods("GET", "HEAD").Name("add-article")
}

func h(handler common.RouteHandler) common.HttpHandler {
	return common.HttpHandler{HandlerFunc: handler}
}
