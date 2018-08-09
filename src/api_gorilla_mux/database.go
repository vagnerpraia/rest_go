package main

import (
	"github.com/globalsign/mgo"
)

func getSessionMongoDB() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	showError(err)

	return session
}

func closeSessionMongoDB() {

}
