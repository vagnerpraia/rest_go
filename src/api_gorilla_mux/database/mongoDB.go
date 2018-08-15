package database

import (
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/treinamento_go/src/api_gorilla_mux/util"
)

func GetSessionMongoDB() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	util.HandlerError(&err, http.StatusInternalServerError, "Ocorreu um erro.")

	return session
}

func closeSessionMongoDB() {

}
