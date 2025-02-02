package config

import (
	"api_go/internal/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Error connecting to database", err)
	}
	// Migrate the schema
	errorMig := db.AutoMigrate(&models.User{}, &models.UserCredential{}, &models.Account{})
	if errorMig != nil {
		log.Panicln("Error migrating schema", errorMig)
		return
	}
	DB = db
}
