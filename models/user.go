package model

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID       bson.ObjectId `bson "_id" json:"id"`
	Name     string        `bson "name" json:"name"`
	Login    string        `bson "login" json:"login"`
	Password string        `bson "password" json:"password"`
}
