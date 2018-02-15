package dao

import (
	"api-task/lib"
	. "api-task/models"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserDAO struct {
	lib.Data
}

func (m *UserDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *UserDAO) FindAll() ([]User, error) {
	var users []User
	err := db.C("user").Find(bson.M{}).All(&users)
	return users, err
}

func (m *UserDAO) Insert(user User) error {
	err := db.C("user").Insert(&user)
	return err
}
