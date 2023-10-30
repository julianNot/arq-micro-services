package routes

import (
	"github.com/gorilla/mux"
	"github.com/julianNot/golang-gorm-api/services"
)

func setRouterTenants(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/tenants", services.GetAllTenantsHandler).Methods("GET")
	apiRouter.HandleFunc("/tenants/{id}", services.GetAllTenantsHandler).Methods("GET")
	apiRouter.HandleFunc("/tenants/{id}/users", services.GetUsersByTenant).Methods("GET")
	apiRouter.HandleFunc("/tenants", services.PostTenantHandler).Methods("POST")
}
