package usuario

import (
	"github.com/treinamento_go/src/api_gorilla_mux/model"
)

var GetUsuarioRoute = model.Route{"GET", "/usuario/{id}", "getUsuario", getUsuarioService}
var GetUsuariosRoute = model.Route{"GET", "/usuarios", "getUsuarios", getUsuariosService}
var PostUsuarioRoute = model.Route{"POST", "/usuario", "postUsuario", postUsuarioService}
var LoginRoute = model.Route{"POST", "/login", "login", loginService}
