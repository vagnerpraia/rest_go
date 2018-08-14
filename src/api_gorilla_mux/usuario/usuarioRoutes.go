package usuario

import (
	"github.com/treinamento_go/src/api_gorilla_mux/model"
)

var GetUsuarioRoute = model.Route{"GET", "/usuario/{id}", "getUsuario", getUsuario}
var GetUsuariosRoute = model.Route{"GET", "/usuarios", "getUsuarios", getUsuarios}
var PostUsuarioRoute = model.Route{"POST", "/usuario", "postUsuario", postUsuario}
var LoginRoute = model.Route{"POST", "/login", "login", login}
