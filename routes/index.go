package routes

import (
	"github.com/gorilla/mux"
)

func SetRouter(router *mux.Router) {
	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	setRouterUsers(apiRouter)
	setRouterTask(apiRouter)
}
