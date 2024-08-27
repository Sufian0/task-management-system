package api

import (
	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", GetTasks).Methods("GET")
	r.HandleFunc("/tasks", CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", GetTask).Methods("GET")
	r.HandleFunc("/tasks/{id}", UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", DeleteTask).Methods("DELETE")
}