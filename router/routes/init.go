package router

import (
	"net/http"

	"web20.tk/router/handlers/common"
	"web20.tk/router/handlers/editor"
	"web20.tk/router/handlers/home"
	"web20.tk/router/handlers/info"
	"web20.tk/router/handlers/posts"

	"github.com/gorilla/mux"
)

func Init(r *mux.Router) {
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("frontend/dist/static"))))
	r.NotFoundHandler = ssr(common.Send404Error)

	r.Handle("/", ssr(home.Home)).Methods("GET", "HEAD").Name("home")
	r.Handle("/posts", ssr(posts.List)).Methods("GET", "HEAD").Name("posts-list")
	r.Handle("/posts/{slug}", ssr(posts.Post)).Methods("GET", "HEAD").Name("post")
	r.Handle("/editor/new", ssr(editor.Editor)).Methods("GET", "HEAD").Name("editor")
	r.Handle("/about", ssr(info.About)).Methods("GET", "HEAD").Name("about")
	r.Handle("/contact-us", ssr(info.ContactUs)).Methods("GET", "HEAD").Name("contact-us")
	r.Handle("/privacy", ssr(info.Privacy)).Methods("GET", "HEAD").Name("privacy")
}

func ssr(handler common.RouteHandler) common.HttpHandler {
	return common.HttpHandler{HandlerFunc: handler}
}

func api(handler common.RouteHandler) common.HttpHandler {
	return common.HttpHandler{HandlerFunc: handler}
}
