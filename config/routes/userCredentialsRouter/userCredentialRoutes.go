package userCredentialsRouter

import (
	"api_go/internal/handlers/userCredentialHandler"
	"api_go/internal/handlers/userCredentialHandler/userCredVal"
	"github.com/gofiber/fiber/v2"
)

func SetUpUserCredentialsRoutes(router fiber.Router) {
	userCred := router.Group("/user-credentials")
	userCred.Post("/log-in", userCredentialHandler.Login)
	userCred.Post("sign-up", userCredVal.ValidateSignUo, userCredentialHandler.SignUp)
}
