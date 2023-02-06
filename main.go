package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jcblastor/gorm-restapi/db"
	"github.com/jcblastor/gorm-restapi/routes"
	"github.com/joho/godotenv"
)

func main() {
	// read .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// connected to database
	db.DBConection()

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/users", routes.GetUsers).Methods("GET")
	r.HandleFunc("/users", routes.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", routes.GetUserById).Methods("GET")
	r.HandleFunc("/users/{id}", routes.DeleteUser).Methods("DELETE")

	r.HandleFunc("/tasks", routes.GetTasks).Methods("GET")
	r.HandleFunc("/tasks", routes.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.GetTaskById).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.DeleteTaskById).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
