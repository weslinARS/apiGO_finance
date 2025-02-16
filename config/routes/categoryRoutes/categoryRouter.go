package categoryRoutes

import (
	"api_go/internal/handlers/categoryHandler"
	"api_go/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetUpCategoryRoutes(router fiber.Router) {
	category := router.Group("/categories")

	category.Get("/user/:idUser/categories/identification", middlewares.AuthMiddleware, categoryHandler.GetCategoriesIdentificationByUser)
	category.Get("/user/:idUser/categories", middlewares.AuthMiddleware, categoryHandler.GetCategoriesByUser)
}
