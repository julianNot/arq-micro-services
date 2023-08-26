package routes

import (
	"github.com/gorilla/mux"
	"github.com/julianNot/golang-gorm-api/services"
)

func setRouterUsers(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/users", services.GetUsersHandler).Methods("GET")
	apiRouter.HandleFunc("/users/{id}", services.GetUserHandler).Methods("GET")
	apiRouter.HandleFunc("/users", services.PostUserHandler).Methods("POST")
	apiRouter.HandleFunc("/users/{id}", services.PutUserHandler).Methods("PUT")
	apiRouter.HandleFunc("/users/{id}", services.DeleteUserHandler).Methods("DELETE")
}
