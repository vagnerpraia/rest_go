package main

import (
	"github.com/treinamento_go/src/api_gorilla_mux/model"
)

var getIndexRoute = model.Route{"GET", "/", "getIndex", getSobre}
var getSobreRoute = model.Route{"GET", "/sobre", "getSobre", getSobre}
