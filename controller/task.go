package controller

import (
	. "api-task/dao"
	. "api-task/models"
	"api-task/util"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

var taskDAO = TaskDAO{}

func init() {
	config.Read()

	taskDAO.Server = config.Server
	taskDAO.Database = config.Database
	taskDAO.Connect()
}
func CreateTask(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	task.ID = bson.NewObjectId()
	if err := taskDAO.Insert(task); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJson(w, http.StatusCreated, task)
}
func ListAllTask(w http.ResponseWriter, r *http.Request) {
	tasks, err := taskDAO.FindAll()
	if err != nil {
		util.RespondWithJson(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJson(w, http.StatusOK, tasks)
}
func ListOneTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetOneTask")
}
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateTask")
}
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteTask")
}
