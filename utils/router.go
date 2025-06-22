package utils

import (
	"fmt"
	"net/http"
	"strings"
)

type Router struct {
	router *http.ServeMux
}

func (r *Router) Handle(path string, handler http.Handler) {
	subpath := strings.TrimSuffix(path, "/")
	r.router.Handle(path, http.StripPrefix(subpath, handler))
	fmt.Printf("Registered path: %s (will strip %s)\n", path, subpath)
}

func (r *Router) HandleFunc(path string, handler http.HandlerFunc) {
	r.Handle(path, handler)
}

func (r *Router) GetServeMux() *http.ServeMux {
	return r.router
}

func NewRouter() *Router {
	router := &Router{
		router: http.NewServeMux(),
	}
	return router
}
