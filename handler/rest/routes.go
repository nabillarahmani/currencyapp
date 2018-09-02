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
	// add new currency rates
	router.HandleFunc("/v1/currency_rates/add", currency.AddCurrencyRates).Methods("POST")
	// get specific currency rates
	router.HandleFunc("/v1/currency_rates/get/{date}", currency.GetCurrencyRates).Methods("GET")
	// get trend
	router.HandleFunc("/v1/currency_rates/get/trend", currency.GetCurrencyRatesTrend).Methods("POST")

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
