package dao

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	. "api-task/models"
)

type TaskDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const COLLECTION = "task"

func (m *TaskDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *TaskDAO) FindAll() ([]Task, error) {
	var task []Task
	err := db.C(COLLECTION).Find(bson.M{}).All(&task)
	return task, err
}
func (m *TaskDAO) Insert(task Task) error {
	err := db.C(COLLECTION).Insert(&task)
	return err
}
