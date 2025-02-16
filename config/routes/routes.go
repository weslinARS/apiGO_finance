package routes

import (
	"api_go/config/routes/categoryRoutes"
	"api_go/config/routes/userCredentialsRouter"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(
	app *fiber.App) {
	userCredentialsRouter.SetUpUserCredentialsRoutes(app)
	categoryRoutes.SetUpCategoryRoutes(app)
}
