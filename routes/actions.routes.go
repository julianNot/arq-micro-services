package routes

import (
	"github.com/gorilla/mux"
	"github.com/julianNot/golang-gorm-api/services"
)

func setRouterActions(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/actions", services.GetActionsHandler).Methods("GET")
	apiRouter.HandleFunc("/actions/{id}", services.GetActionHandler).Methods("GET")
	apiRouter.HandleFunc("/actions", services.PostActionHandler).Methods("POST")
}
