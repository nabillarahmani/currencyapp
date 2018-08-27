package main

import (
	// standard library from golang

	"github.com/gorilla/mux"

	// internal modules
	"github.com/nabillarahmani/currencyapp/handler/rest"
	"github.com/nabillarahmani/currencyapp/internal/common/configs"
	"github.com/nabillarahmani/currencyapp/internal/common/log"

	// library used from outside
	"gopkg.in/paytm/grace.v1"
)

var stdoutLog string

func main() {
	// init log
	stdoutLog = "../files/log/currencyapp.log"
	log.Init(stdoutLog)

	// init config variables
	configs.InitConfig()

	// init common modules like database or anything

	// init http router
	router := mux.NewRouter()
	rest.InitRoutes(router)

	// init modules

	// serve it
	err := grace.Serve(":7777", router)
	if err != nil {
		log.Error(err, "There's an error during starting the server!")
		return
	}
}
