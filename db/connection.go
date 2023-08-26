package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = ""
var DB *gorm.DB

func DBConnection(host, user, password, nameDb, port string) {
	DSN = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, nameDb, port)
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("DB Connected")
	}
}
