package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/julianNot/golang-gorm-api/models"
)

func SetupDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	DBConnection(host, user, password, dbname, port)

	DB.AutoMigrate(
		// &models.Tenant{},
		&models.Action{},
		&models.Role{},
		&models.Directory{},
	// 	// &models.ActionRole{},
	)
}
