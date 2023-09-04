package services

import (
	"encoding/json"
	"net/http"

	"github.com/julianNot/golang-gorm-api/db"
	"github.com/julianNot/golang-gorm-api/models"
)

func PostTenantHandler(w http.ResponseWriter, r *http.Request) {
	var tenant models.Tenant
	json.NewDecoder(r.Body).Decode(&tenant)
	createdUser := db.DB.Create(&tenant)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&tenant)
}
