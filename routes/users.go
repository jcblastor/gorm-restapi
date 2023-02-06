package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jcblastor/gorm-restapi/db"
	"github.com/jcblastor/gorm-restapi/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := []models.User{}
	db.DB.Find(&users)

	json.NewEncoder(w).Encode(&users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	params := mux.Vars(r)

	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}

	// llenamos los campos relacionados
	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)

	json.NewEncoder(w).Encode(&user)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)

	newUser := db.DB.Create(&user)
	err := newUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	params := mux.Vars(r)

	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	// aca le decimos a gorm que marque al usuario como eliminado
	db.DB.Delete(&user)

	// aca le decimos que eliminme fisicamente al usuario
	// db.DB.Unscoped().Delete(&user)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}
