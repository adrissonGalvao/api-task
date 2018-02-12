package controller

import (
	"fmt"
	"net/http"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateTask")
}
func ListAllTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ListAllTask")
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
