package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	appRouter "web20.tk/router/routes"
)

const PORT = 9191

var server_stop_rq = make(chan bool, 1)

func main() {
	go waitStopSignal(stopServer)
	runServer()
}

func runServer() {
	router := mux.NewRouter()
	appRouter.Init(router)
	listen := fmt.Sprintf(":%d", PORT)
	fmt.Println("SSR server is active")
	go func() {
		err := http.ListenAndServe(listen, router)
		if err != nil {
			fmt.Println("ListenAndServe Failed")
			panic(err)
		}
	}()
	<-server_stop_rq
	fmt.Println("Api server has been stopped")
}

func waitStopSignal(stopHandler func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	s := <-c
	fmt.Printf("Got stop signal: %s\n", s.String())
	stopHandler()
}

func stopServer() {
	server_stop_rq <- true
}
