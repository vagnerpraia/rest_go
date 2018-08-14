package router

import (
	"github.com/treinamento_go/src/api_gorilla_mux/model"
	"github.com/treinamento_go/src/api_gorilla_mux/resources/sobre"
	"github.com/treinamento_go/src/api_gorilla_mux/resources/usuario"
)

var routes = model.Routes{
	sobre.GetIndexRoute,
	sobre.GetSobreRoute,
	usuario.GetUsuarioRoute,
	usuario.GetUsuariosRoute,
	usuario.PostUsuarioRoute,
	usuario.LoginRoute,
}
