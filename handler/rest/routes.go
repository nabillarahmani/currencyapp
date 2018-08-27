package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

//InitRoutes set all routes here
func InitRoutes(router *mux.Router) {
	//ping
	router.HandleFunc("/ping", ping).Methods("GET")

	// // modules
	// router.HandleFunc("/currency/add", ).Methods("POST")

	router.MethodNotAllowedHandler = http.HandlerFunc(notfound)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Currencyapp is up!!!"))
}

func notfound(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("404 Not Found"))
}
