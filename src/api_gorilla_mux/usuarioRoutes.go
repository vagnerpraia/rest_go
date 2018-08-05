package main

var getUsuarioRoute = Route{"GET", "/usuario/{id}", "getUsuario", getUsuario}
var getUsuariosRoute = Route{"GET", "/usuarios", "getUsuarios", getUsuarios}
var loginRoute = Route{"POST", "/login", "login", login}
