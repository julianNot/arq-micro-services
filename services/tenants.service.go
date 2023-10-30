package services

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/julianNot/golang-gorm-api/db"
	"github.com/julianNot/golang-gorm-api/models"
)

type Tenant struct {
	NameTenant   string `json:"name_tenant"`
	UrlLogo      string `json:"url_logo"`
	EmailContact string `json:"email_contact"`
}

func (Tenant) TableName() string {
	return "Tenant"
}

func GetAllTenantsHandler(w http.ResponseWriter, r *http.Request) {
	var tenants []models.Tenant
	db.DB.Find(&tenants)
	json.NewEncoder(w).Encode(&tenants)
}

func GetUsersByTenant(w http.ResponseWriter, r *http.Request) {
	var users []models.Directory
	params := mux.Vars(r)
	db.DB.Where("id_tenant = ?", params["id"]).Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func PostTenantHandler(w http.ResponseWriter, r *http.Request) {
	var tenant Tenant
	json.NewDecoder(r.Body).Decode(&tenant)
	createdUser := db.DB.Create(&tenant)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&tenant)
}
