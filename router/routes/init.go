package router

import (
	"net/http"

	"web20.tk/router/handlers/home"
	"web20.tk/router/handlers/user"

	"github.com/gorilla/mux"
)

func Init(router *mux.Router) {
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("frontend/dist/static"))))

	router.HandleFunc("/", home.Init).Methods(http.MethodGet, http.MethodHead)
	router.HandleFunc("/user", user.Init).Methods(http.MethodGet, http.MethodHead)
}
