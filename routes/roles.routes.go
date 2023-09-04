package routes

import (
	"github.com/gorilla/mux"
	"github.com/julianNot/golang-gorm-api/services"
)

func setRouterRoles(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/roles", services.GetRolesHandler).Methods("GET")
	apiRouter.HandleFunc("/roles/{id}", services.GetRoleHandler).Methods("GET")
	apiRouter.HandleFunc("/roles", services.PostRoleHandler).Methods("POST")
	apiRouter.HandleFunc("/roles/{id}", services.PutRoleHandler).Methods("PUT")
	apiRouter.HandleFunc("/roles/{id}", services.DeleteRoleHandler).Methods("DELETE")
}
