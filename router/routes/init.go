package router

import (
	"net/http"

	"web20.tk/router/common"
	"web20.tk/router/handlers/home"
	"web20.tk/router/handlers/user"

	"github.com/gorilla/mux"
)

func Init(router *mux.Router) {
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("frontend/dist/static"))))

	router.NotFoundHandler = common.HttpHandler{common.Send404Error}

	router.Handle("/", common.HttpHandler{home.Home}).Methods(http.MethodGet, http.MethodHead)
	router.Handle("/user", common.HttpHandler{user.User}).Methods(http.MethodGet, http.MethodHead)
}
