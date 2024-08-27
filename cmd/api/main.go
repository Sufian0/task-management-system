package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/Sufian0/task-management-system/internal/api"
	"github.com/Sufian0/task-management-system/internal/database"
)

func main() {
	database.InitDB()
	defer database.DB.Close()

	r := mux.NewRouter()
	api.SetupRoutes(r)

	log.Println("Server starting on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}