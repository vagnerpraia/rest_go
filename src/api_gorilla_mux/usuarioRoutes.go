package main

var getUsuarioRoute = Route{"GET", "/usuario/{id}", "getUsuario", getUsuario}
var getUsuariosRoute = Route{"GET", "/usuarios", "getUsuarios", getUsuarios}
var postUsuarioRoute = Route{"POST", "/usuario", "postUsuario", postUsuario}
var loginRoute = Route{"POST", "/login", "login", login}
