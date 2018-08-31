package rest

import (
	"net/http"

	"github.com/nabillarahmani/currencyapp/handler/rest/currency"

	"github.com/gorilla/mux"
)

//InitRoutes set all routes here
func InitRoutes(router *mux.Router) {
	//ping
	router.HandleFunc("/ping", ping).Methods("GET")

	// modules routes
	// add new currency
	router.HandleFunc("/v1/currency/addremove", currency.AddRemoveCurrency).Methods("POST")
	// get all currency
	router.HandleFunc("/v1/currency/get", currency.GetCurrency).Methods("GET")

	// go templating routes

	// not found
	router.MethodNotAllowedHandler = http.HandlerFunc(notfound)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Currencyapp is up!!!"))
}

func notfound(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("404 Not Found"))
}
