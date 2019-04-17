package middleware

import "net/http"

/*
Middleware interface
is used to provide additional functionality
to the view functions
*/
type Middleware func(http.HandlerFunc) http.HandlerFunc

/*
WithMiddleware function
is responsible for running all middleware
functions before running the actual handler h
*/
func WithMiddleware(h http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}
