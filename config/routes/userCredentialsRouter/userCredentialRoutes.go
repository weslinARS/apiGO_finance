package userCredentialsRouter

import (
	"api_go/internal/handlers/userCredentialHandler"
	"api_go/internal/handlers/userCredentialHandler/userCredentialsValidators"
	"github.com/gofiber/fiber/v2"
)

func SetUpUserCredentialsRoutes(router fiber.Router) {
	userCred := router.Group("/user-credentials")
	userCred.Post("/log-in", userCredentialHandler.Login)
	userCred.Post("sign-up", userCredentialsValidators.ValidateSignUo, userCredentialHandler.SignUp)
}
