package database

import (
	"go-echo-crud/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("store.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	DB.AutoMigrate(&models.Category{}, &models.Product{}, &models.Cart{})
}
