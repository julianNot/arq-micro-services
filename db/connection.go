package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DSN = ""
var DB *gorm.DB

func DBConnection(host, user, password, nameDb, port string) {
	DSN = fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", user, password, host, port, nameDb)
	var err error
	DB, err = gorm.Open(sqlserver.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("DB Connected")
	}
}
