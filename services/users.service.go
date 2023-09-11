package services

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/julianNot/golang-gorm-api/db"
	"github.com/julianNot/golang-gorm-api/models"
	"golang.org/x/crypto/bcrypt"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.Directory
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.Directory
	params := mux.Vars(r)
	db.DB.Preload("Roles").First(&user, params["id"])
	if user.IdDirectory == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Not Found"))
		return
	}
	json.NewEncoder(w).Encode(&user)
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.Directory
	json.NewDecoder(r.Body).Decode(&user)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.HashPassword), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to hash password"))
		return
	}
	user.HashPassword = string(hashedPassword)
	createdUser := db.DB.Create(&user)
	err = createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.Directory
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])
	if user.IdDirectory == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Not Found"))
		return
	}
	db.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusOK)
}

func PutUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.Directory
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])
	if user.IdDirectory == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Not Found"))
		return
	}
	json.NewDecoder(r.Body).Decode(&user)
	db.DB.Save(&user)
	json.NewEncoder(w).Encode(&user)
}
