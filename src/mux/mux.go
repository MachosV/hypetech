package mux

import (
	"net/http"
)

var mux *http.ServeMux

func init() {
	mux = http.NewServeMux()
}

/*
GetMux function returns
the mux object responsible to route all requests
*/
func GetMux() *http.ServeMux {
	return mux
}
