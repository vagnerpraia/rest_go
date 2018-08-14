package database

import (
	"github.com/globalsign/mgo"
	"github.com/treinamento_go/src/api_gorilla_mux/util"
)

func GetSessionMongoDB() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	util.ShowError(err)

	return session
}

func closeSessionMongoDB() {

}
