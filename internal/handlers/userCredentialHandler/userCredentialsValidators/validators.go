package userCredentialsValidators

import (
	"api_go/internal/models"
	"api_go/internal/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type userCred struct {
	Email    string ` validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type signUpRequest struct {
	UserCred userCred    `validate:"required"`
	UserInfo models.User `validate:"required"`
}

func ValidateSignUo(c *fiber.Ctx) error {
	body := new(signUpRequest)
	fmt.Println("Validating sign up request")
	return utils.CheckValidations(c, body)
}
