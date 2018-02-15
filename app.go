package main

import (
	"api-task/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/task", controller.CreateTask).Methods("POST")
	r.HandleFunc("/api/task", controller.ListAllTask).Methods("GET")
	r.HandleFunc("/api/task/{id}", controller.ListOneTask).Methods("GET")
	r.HandleFunc("/api/task/{id}", controller.UpdateTask).Methods("PUT")
	r.HandleFunc("/api/task/{id}", controller.DeleteTask).Methods("DELETE")
	r.HandleFunc("/api/user", controller.CreateUser).Methods("POST")
	r.HandleFunc("/api/user", controller.ListAllUser).Methods("GET")
	r.HandleFunc("/api/user/{id}", controller.ListOneUser).Methods("GET")
	r.HandleFunc("/api/user/{id}", controller.UpdateUser).Methods("PUT")
	r.HandleFunc("/api/user/{id}", controller.DeleteUser).Methods("DELETE")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}

}
