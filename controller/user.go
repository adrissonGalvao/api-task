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

var usersDAO = UserDAO{}

func init() {
	config.Read()

	usersDAO.Server = config.Server
	usersDAO.Database = config.Database
	usersDAO.Connect()
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user.ID = bson.NewObjectId()
	if err := usersDAO.Insert(user); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJson(w, http.StatusCreated, user)
}
func ListAllUser(w http.ResponseWriter, r *http.Request) {
	users, err := usersDAO.FindAll()
	if err != nil {
		util.RespondWithJson(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJson(w, http.StatusOK, users)
}
func ListOneUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetOneTask")
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateTask")
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteTask")
}
