package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/julianNot/golang-gorm-api/db"
	"github.com/julianNot/golang-gorm-api/helpers"
	"github.com/julianNot/golang-gorm-api/routes"
)

func main() {
	db.SetupDatabase()

	router := mux.NewRouter()
	routes.SetRouter(router)

	http.ListenAndServe(":"+helpers.GetPortServer(), router)
}
