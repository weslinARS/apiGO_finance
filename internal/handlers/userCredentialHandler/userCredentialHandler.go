package userCredentialHandler

import (
	"api_go/config"
	"api_go/internal/models"
	"api_go/internal/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/jsonapi"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {

	return c.SendString("Login")
}

type userCred struct {
	Email    string ` validate:"required,email"`
	Password string `json:"password" validate:"required, len=50"`
}

type signUpRequest struct {
	UserCred userCred    `validate:"required"`
	UserInfo models.User `validate:"required"`
}

func SignUp(ctx *fiber.Ctx) error {
	var body signUpRequest
	// parse body to signUpRequest type
	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(400).SendString("Error parsing body")
	}
	//verify if there is a user with the same email
	var user models.User
	config.DB.Where("email = ?", body.UserInfo.Email).Select("id").First(&user)
	if id, err := uuid.Parse(user.ID); id != uuid.Nil && err == nil {
		fmt.Println(user.ID)
		return ctx.Status(fiber.StatusBadRequest).SendString("User already exists")
	}
	//verify if its a strong password
	if !utils.IsStrongPassword(body.UserCred.Password) {
		return ctx.Status(fiber.StatusBadRequest).SendString("Password is not strong enough")
	}
	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.UserCred.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password", err)
		return ctx.Status(500).SendString("Error hashing password")
	}
	// create user in DB
	newUser := models.User{
		Email:       body.UserInfo.Email,
		Name:        body.UserInfo.Name,
		Gender:      body.UserInfo.Gender,
		DateOfBirth: body.UserInfo.DateOfBirth,
		LastName:    body.UserInfo.LastName,
		ID:          body.UserInfo.ID,
		Model:       body.UserInfo.Model,
	}
	if err := config.DB.Create(&newUser).Error; err != nil {
		fmt.Println("Error creating user", err)
		return ctx.Status(500).SendString("Error creating user")
	} else {
		// create use credential in DB \
		userCred := models.UserCredential{
			UserId:    newUser.ID,
			UserEmail: body.UserCred.Email,
			Password:  string(hashedPassword),
		}
		if err := config.DB.Create(&userCred).Error; err != nil {
			return ctx.Status(500).SendString("Error creating user credential")
		} else {
			// enviar informacion de usuario

			//serialized in JSON:API
			err := jsonapi.MarshalOnePayloadEmbedded(ctx.Response().BodyWriter(), &newUser)
			if err != nil {
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			}
			return nil
		}

	}
}
