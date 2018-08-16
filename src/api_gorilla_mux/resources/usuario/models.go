package usuario

import "github.com/globalsign/mgo/bson"

type Usuario struct {
	Id    bson.ObjectId `json:"_id" bson:"_id"`
	Nome  string        `json:"nome" bson:"nome"`
	Email string        `json:"email" bson:"email"`
	Senha string        `json:"senha" bson:"senha"`
}

type Usuarios []Usuario
