package categoryHandler

import (
	"api_go/config"
	"api_go/internal/models"
	"api_go/internal/utils/types"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/jsonapi"
	"gorm.io/gorm"
	"log"
)

// GetCategories handles the request to retrieve categories.
//
// Parameters:
// - ctx: *fiber.Ctx - The Fiber context which provides methods to handle the request and response.
//
// Returns:
// - error: Always returns nil in the current implementation.

type CategoryResponse struct {
	ID   string `json:"id" jsonapi:"primary,category"`
	Name string `json:"name" jsonapi:"attr,name"`
}
type CategoryList struct {
	Categories []CategoryResponse `json:"categories" jsonapi:"attr,categories"`
}

func GetCategoriesIdentificationByUser(ctx *fiber.Ctx) error {
	userId := ctx.Params("idUser", "")

	userCategories := &models.User{
		ID: userId,
	}
	categorySelection := func(tx *gorm.DB) *gorm.DB {
		return tx.Select("id", "name")
	}
	if err := config.DB.Preload("Categories", categorySelection).First(userCategories).Error; err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Error getting user: " + err.Error()})
	}

	ctx.Set(fiber.HeaderContentType, jsonapi.MediaType)
	err := jsonapi.MarshalPayload(ctx.Response().BodyWriter(), userCategories)
	if err != nil {
		log.Println("Error marshalling categories: " + err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.NewErrorResponse(fiber.StatusInternalServerError, "Error getting categories", ""))
	}
	return nil
}

func GetCategoriesByUser(ctx *fiber.Ctx) error {
	userId := ctx.Params("idUser", "")

	var categories []models.Category
	if err := config.DB.Model(&models.User{ID: userId}).Association("Categories").Find(&categories); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Error getting categories")
	}
	return ctx.Status(fiber.StatusOK).JSON(categories)
}

type userCategory struct {
	UserId     string
	CategoryId string
}

func LinkDefaultCategories(userId string) []error {
	// retrieve default categories Id
	var errorList []error
	var categories []models.Category
	config.DB.Select("id").Where("is_default = ?", true).Find(&categories)
	log.Println("Default categories: ", categories)
	for _, category := range categories {
		user := models.User{ID: userId}
		if err := config.DB.Model(&user).Association("Categories").Append(&category); err != nil {
			errorList = append(errorList, fmt.Errorf("error link user with default Category: %v", err))
		}
	}

	return errorList
}
