package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Method  string
	Path    string
	Name    string
	Handler http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"GET", "/", "getIndex", getIndex},
	getUsuarioRoute,
	getUsuariosRoute,
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(route.Handler)
	}

	return router
}
