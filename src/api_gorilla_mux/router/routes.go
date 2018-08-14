package router

import (
	"github.com/treinamento_go/src/api_gorilla_mux/model"
	"github.com/treinamento_go/src/api_gorilla_mux/usuario"
)

var routes = model.Routes{
	usuario.GetUsuarioRoute,
	usuario.GetUsuariosRoute,
	usuario.PostUsuarioRoute,
	usuario.LoginRoute,
}
