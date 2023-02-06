package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jcblastor/gorm-restapi/db"
	"github.com/jcblastor/gorm-restapi/models"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := []models.Task{}
	db.DB.Find(&tasks)

	json.NewEncoder(w).Encode(&tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	task := models.Task{}
	json.NewDecoder(r.Body).Decode(&task)

	newTask := db.DB.Create(&task)
	err := newTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error al crear el usuario"))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	task := models.Task{}

	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("El usuario no existe"))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func DeleteTaskById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	task := models.Task{}

	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("task not found"))
		return
	}

	db.DB.Delete(&task)

	w.Write([]byte("Task deleted successfully"))
}
