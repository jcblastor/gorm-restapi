package main

import (
	"log"

	"github.com/jcblastor/gorm-restapi/db"
	"github.com/jcblastor/gorm-restapi/models"
	"github.com/joho/godotenv"
)

func main() {
	// read .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db.DBConection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})
}
