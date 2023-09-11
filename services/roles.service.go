package services

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/julianNot/golang-gorm-api/db"
	"github.com/julianNot/golang-gorm-api/models"
)

func GetRolesHandler(w http.ResponseWriter, r *http.Request) {
	var roles []models.Role
	db.DB.Find(&roles)
	json.NewEncoder(w).Encode(roles)
}

func GetRoleHandler(w http.ResponseWriter, r *http.Request) {
	var role models.Role
	params := mux.Vars(r)
	db.DB.First(&role, params["id"])
	if role.IdRole == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Role Not Found"))
		return
	}
	json.NewEncoder(w).Encode(&role)
}

func PostRoleHandler(w http.ResponseWriter, r *http.Request) {
	var role models.Role
	json.NewDecoder(r.Body).Decode(&role)
	createdTask := db.DB.Create(&role)
	err := createdTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&role)
}

func DeleteRoleHandler(w http.ResponseWriter, r *http.Request) {
	var role models.Role
	params := mux.Vars(r)
	db.DB.First(&role, params["id"])
	if role.IdRole == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Role Not Found"))
		return
	}
	db.DB.Unscoped().Delete(&role)
	w.WriteHeader(http.StatusNoContent) //204
}

func PutRoleHandler(w http.ResponseWriter, r *http.Request) {
	var role models.Role
	params := mux.Vars(r)
	db.DB.First(&role, params["id"])
	if role.IdRole == "" {
		w.WriteHeader(http.StatusNotFound)

		w.Write([]byte("Role Not Found"))
		return
	}
	json.NewDecoder(r.Body).Decode(&role)
	db.DB.Save(&role)
	json.NewEncoder(w).Encode(&role)
}
