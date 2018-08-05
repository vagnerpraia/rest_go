package main

import (
	"net/http"
)

type Route struct {
	Method  string
	Path    string
	Name    string
	Handler http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	getIndexRoute,
	getSobreRoute,
	getUsuarioRoute,
	getUsuariosRoute,
	postUsuarioRoute,
	loginRoute,
}
