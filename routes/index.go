package routes

import (
	"github.com/gorilla/mux"
	"github.com/julianNot/golang-gorm-api/auth"
)

func SetRouter(router *mux.Router) {
	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	apiRouter.HandleFunc("/auth/users", auth.AuthenticatePartnerHandler).Methods("POST")

	setRouterUsers(apiRouter)
	setRouterRoles(apiRouter)
	setRouterTenants(apiRouter)
}
