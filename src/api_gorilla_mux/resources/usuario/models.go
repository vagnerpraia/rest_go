package usuario

import (
	"github.com/globalsign/mgo/bson"
	"github.com/treinamento_go/src/api_gorilla_mux/database"
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

var session = database.GetSessionMongoDB()
var db = session.DB("godb")
var collection = db.C("usuarios")

func getUsuariosModel() (data []Usuario, err error) {
	var usuariosDb Usuarios
	err = collection.Find(nil).Sort("nome").All(&usuariosDb)

	if err == nil {
		data = usuariosDb
	}

	return data, err
}

func getUsuarioModel(id string) (data Usuario, err error) {
	if bson.IsObjectIdHex(id) {
		idBson := bson.ObjectIdHex(id)

		var usuarioDb Usuario
		err = collection.FindId(idBson).One(&usuarioDb)
		if err == nil {
			data = usuarioDb
		}
	}

	return data, err
}

func insertUsuario(usuario Usuario) (data Usuario, err error) {
	usuarioDb := Usuario{Id: bson.NewObjectId(), Nome: usuario.Nome, Email: usuario.Email, Senha: usuario.Senha}
	err = collection.Insert(&usuarioDb)
	if err == nil {
		data = usuarioDb
	}

	return data, err
}

func updateUsuario(usuario Usuario) (data Usuario, err error) {
	if len(usuario.Id) == 12 {
		usuarioDb := Usuario{Id: usuario.Id, Nome: usuario.Nome, Email: usuario.Email, Senha: usuario.Senha}
		err = collection.Update(&usuarioDb.Id, &usuarioDb)
		if err == nil {
			data = usuarioDb
		}
	}

	return data, err
}
