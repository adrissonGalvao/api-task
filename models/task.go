package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Task struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Description string        `bson:"description" json:"description"`
	user_task   User          `bson:"user" json:"user_task"`
}
