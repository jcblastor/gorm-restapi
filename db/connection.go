package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConection() {
	var err error
	dsn := os.Getenv("URL_DB")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Error connected to database")
		log.Fatal(err)
	}

	log.Println("DB connected")
}
