package router

import (
	"net/http"

	"web20.tk/router/handlers/common"
	"web20.tk/router/handlers/editor"
	"web20.tk/router/handlers/home"
	"web20.tk/router/handlers/info"
	"web20.tk/router/handlers/posts"

	"web20.tk/router/api-handlers/md"
	"web20.tk/router/api-handlers/postsApi"

	"github.com/gorilla/mux"
)

func Init(r *mux.Router) {
	r.StrictSlash(true)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("frontend/dist/static"))))
	r.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "frontend/dist/static/robots.txt")
	})

	r.NotFoundHandler = ssr(common.Send404Error)

	r.Handle("/", ssr(home.Home)).Methods("GET", "HEAD").Name("home")
	r.Handle("/posts", ssr(posts.List)).Methods("GET", "HEAD").Name("posts-list")
	r.Handle("/posts/{slug}", ssr(posts.Post)).Methods("GET", "HEAD").Name("post")
	r.Handle("/editor/new", ssr(editor.New)).Methods("GET", "HEAD").Name("editor")
	r.Handle("/editor/edit/{slug}", ssr(editor.Edit)).Methods("GET", "HEAD").Name("editor")
	r.Handle("/contact-us", ssr(info.ContactUs)).Methods("GET", "HEAD").Name("contact-us")
	r.Handle("/privacy", ssr(info.Privacy)).Methods("GET", "HEAD").Name("privacy")

	r.Handle("/api/posts", api(postsApi.GetAll)).Methods("GET")
	r.Handle("/api/posts", api(postsApi.Create)).Methods("POST")
	r.Handle("/api/posts/{slug}", api(postsApi.Get)).Methods("GET")
	r.Handle("/api/posts/{slug}", api(postsApi.Update)).Methods("PUT")
	r.Handle("/api/posts/{slug}", api(postsApi.Delete)).Methods("DELETE")

	r.Handle("/api/search", api(postsApi.Delete)).Methods("GET")
	r.Handle("/api/md2html", api(md.Md2Html)).Methods("POST")
}

func ssr(handler common.RouteHandler) common.HttpHandler {
	return common.HttpHandler{HandlerFunc: handler, Type: common.SSR_HANDLER}
}

func api(handler common.RouteApiHandler) common.HttpHandler {
	return common.HttpHandler{ApiHandlerFunc: handler, Type: common.API_HANDLER}
}
