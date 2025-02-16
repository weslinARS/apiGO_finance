package config

import (
	"api_go/internal/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Panic("Error connecting to database", err)
	}
	// Migrate the schema
	errorMig := db.AutoMigrate(&models.User{}, &models.UserCredential{}, &models.Account{}, &models.Currency{}, &models.Category{})
	if errorMig != nil {
		log.Panicln("Error migrating schema", errorMig)
		return
	}

	DB = db
	// create default categories if they do not exist
	err = createDefaultCategories()
	if err != nil {
		log.Panicln("Error creating default categories", err)
		return
	}

}

func createDefaultCategories() error {
	// create default categories
	defaultCategories := []models.Category{
		{
			Name:        "Food",
			Description: "Food category",
			IsDefault:   true,
		},
		{
			Name:        "Medicine",
			Description: "Medicine category",
			IsDefault:   true,
		},
		{
			Name:        "Clothes",
			Description: "Clothes category",
			IsDefault:   true,
		},
		{
			Name:        "Transport",
			Description: "Transport category",
			IsDefault:   true,
		},
		{
			Name:        "Entertainment",
			Description: "Entertainment category",
			IsDefault:   true,
		},
		{
			Name:        "Groceries",
			Description: "Groceries category",
			IsDefault:   true,
		},
	}
	for _, category := range defaultCategories {
		if err := DB.Where(models.Category{
			Name: category.Name,
		}).Attrs().FirstOrCreate(&category).Error; err != nil {
			return err
		}
	}
	return nil
}
