package mux

import (
	"net/http"
)

var mux *http.ServeMux

func init() {
	mux = http.NewServeMux()
	fs := http.FileServer(http.Dir("js"))
	mux.Handle("/js/", http.StripPrefix("/js/", fs))
}

/*
GetMux function returns
the mux object responsible to route all requests
*/
func GetMux() *http.ServeMux {
	return mux
}
