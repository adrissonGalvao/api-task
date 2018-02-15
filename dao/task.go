package dao

import (
	"log"

	"api-task/lib"
	. "api-task/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TaskDAO struct {
	lib.Data
}

func (m *TaskDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *TaskDAO) FindAll() ([]Task, error) {
	var tasks []Task
	err := db.C("task").Find(bson.M{}).All(&tasks)
	return tasks, err
}
func (m *TaskDAO) Insert(task Task) error {
	err := db.C("task").Insert(&task)
	return err
}
