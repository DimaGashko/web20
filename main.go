package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"web20.tk/core/db"
	"web20.tk/router/handlers/common"
	appRouter "web20.tk/router/routes"
)

var server_stop_rq = make(chan bool, 1)

func main() {
	configure()
	go waitStopSignal(stopServer)
	runServer()
}

func configure() {
	var configPath = flag.String("config", "config.json", "configuration file")
	var argPort = flag.Int("port", 9191, "servering port")
	flag.Parse()

	parseAppConfig(*configPath)

	if *argPort > 0 {
		common.AppConfig.Port = *argPort
	}

	err := db.Configure(common.AppConfig.Db)
	if err != nil {
		panic(err)
	}
}

func parseAppConfig(configPath string) {
	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Printf("Couldn't parse config file (%s)\n", configPath)
		panic(err)
	}
	err = json.Unmarshal(b, &common.AppConfig)
	if err != nil {
		fmt.Printf("Couldn't unmarshal config file (%s)\n", configPath)
		panic(err)
	}
}

func runServer() {
	router := mux.NewRouter()
	appRouter.Init(router)
	listen := fmt.Sprintf(":%d", common.AppConfig.Port)
	fmt.Printf("SSR server is running on port %d\n", common.AppConfig.Port)
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
