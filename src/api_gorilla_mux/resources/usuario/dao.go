package usuario

import (
	"github.com/globalsign/mgo/bson"
	"github.com/vagnerpraia/treinamento_go/src/api_gorilla_mux/database"
)

var session = database.GetSessionMongoDB()
var db = session.DB("godb")
var collection = db.C("usuarios")

func getUsuariosDao() (data []Usuario, err error) {
	var usuariosDb Usuarios
	err = collection.Find(nil).Sort("nome").All(&usuariosDb)

	if err == nil {
		data = usuariosDb
	}

	return data, err
}

func getUsuarioDao(id string) (data Usuario, err error) {
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

func insertUsuarioDao(usuario Usuario) (data Usuario, err error) {
	usuarioDb := Usuario{Id: bson.NewObjectId(), Nome: usuario.Nome, Email: usuario.Email, Senha: usuario.Senha}
	err = collection.Insert(&usuarioDb)
	if err == nil {
		data = usuarioDb
	}

	return data, err
}

func updateUsuarioDao(usuario Usuario) (data Usuario, err error) {
	if len(usuario.Id) == 12 {
		usuarioDb := Usuario{Id: usuario.Id, Nome: usuario.Nome, Email: usuario.Email, Senha: usuario.Senha}
		err = collection.Update(&usuarioDb.Id, &usuarioDb)
		if err == nil {
			data = usuarioDb
		}
	}

	return data, err
}

func loginDao(usuario Usuario) (data Usuario, err error) {
	var usuarioDb Usuario

	err = collection.Find(bson.M{"email": usuario.Email, "senha": usuario.Senha}).One(&usuarioDb)
	if err == nil {
		data = usuarioDb
	}

	return data, err
}
