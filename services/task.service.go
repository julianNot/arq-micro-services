package services

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/julianNot/golang-gorm-api/db"
	"github.com/julianNot/golang-gorm-api/models"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var task []models.Task
	db.DB.Find(&task)
	json.NewEncoder(w).Encode(task)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)
	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task Not Found"))
		return
	}
	json.NewEncoder(w).Encode(&task)
}

func PostTasksHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	createdTask := db.DB.Create(&task)
	err := createdTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&task)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)
	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task Not Found"))
		return
	}
	db.DB.Unscoped().Delete(&task)
	w.WriteHeader(http.StatusNoContent) //204
}

func PutTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)
	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)

		w.Write([]byte("Task Not Found"))
		return
	}
	json.NewDecoder(r.Body).Decode(&task)
	db.DB.Save(&task)
	json.NewEncoder(w).Encode(&task)
}
