package services

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/julianNot/golang-gorm-api/db"
	"github.com/julianNot/golang-gorm-api/models"
)

func GetActionsHandler(w http.ResponseWriter, r *http.Request) {
	var action []models.Action
	db.DB.Find(&action)
	json.NewEncoder(w).Encode(action)
}

func GetActionHandler(w http.ResponseWriter, r *http.Request) {
	var action models.Action
	params := mux.Vars(r)
	db.DB.First(&action, params["id"])
	if action.ActionID == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Role Not Found"))
		return
	}
	json.NewEncoder(w).Encode(&action)
}

func PostActionHandler(w http.ResponseWriter, r *http.Request) {
	var action models.Action
	json.NewDecoder(r.Body).Decode(&action)
	createdTask := db.DB.Create(&action)
	err := createdTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&action)
}