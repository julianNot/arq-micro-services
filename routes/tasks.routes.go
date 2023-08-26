package routes

import (
	"github.com/gorilla/mux"
	"github.com/julianNot/golang-gorm-api/services"
)

func setRouterTask(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/tasks", services.GetTasksHandler).Methods("GET")
	apiRouter.HandleFunc("/tasks/{id}", services.GetTaskHandler).Methods("GET")
	apiRouter.HandleFunc("/tasks", services.PostTasksHandler).Methods("POST")
	apiRouter.HandleFunc("/tasks/{id}", services.PutTaskHandler).Methods("PUT")
	apiRouter.HandleFunc("/tasks/{id}", services.DeleteTaskHandler).Methods("DELETE")
}
