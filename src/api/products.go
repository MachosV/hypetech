package api

import (
	"fmt"
	"middleware"
	"mux"
	"net/http"
)

var sortMap map[int]string

func init() {
	mux := mux.GetMux()
	mux.HandleFunc("/products", middleware.WithMiddleware(productHandler,
		middleware.Time(),
	))
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case "GET":
		getProducts(w, r)
	default:
		notHandled(w, r)
	}
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	sorting := r.FormValue("sort")
}

func notHandled(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Error, method not handled")
}
