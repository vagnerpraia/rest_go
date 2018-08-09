package main

import (
	"github.com/globalsign/mgo/bson"
)

type Usuario struct {
	Id    bson.ObjectId `json:"_id" bson:"_id"`
	Nome  string        `json:"nome" bson:"nome"`
	Email string        `json:"email" bson:"email"`
	Senha string        `json:"senha" bson:"senha"`
}

type Usuarios []Usuario

var usuarios = Usuarios{
	Usuario{bson.NewObjectId(), "Maria da Silva", "maria@mailtest.com", "123456"},
	Usuario{bson.NewObjectId(), "João da Silva", "joao@mailtest.com", "456789"},
	Usuario{bson.NewObjectId(), "Paulo da Silva", "paulo@mailtest.com", "789123"},
}

var session = getSessionMongoDB()
var db = session.DB("godb")
var collection = db.C("usuarios")

func insertUsuario(usuario Usuario) *Usuario {
	usuarioDb := &Usuario{Id: bson.NewObjectId(), Nome: usuario.Nome, Email: usuario.Email, Senha: usuario.Senha}
	err := collection.Insert(usuarioDb)
	showError(err)

	return usuarioDb
}

func updateUsuario(usuario Usuario) *Usuario {
	usuarioDb := &Usuario{Id: usuario.Id, Nome: usuario.Nome, Email: usuario.Email, Senha: usuario.Senha}

	if len(usuario.Id) == 12 {
		err := collection.Update(usuarioDb.Id, usuarioDb)
		showError(err)

		return usuarioDb
	} else {
		return nil
	}
}
