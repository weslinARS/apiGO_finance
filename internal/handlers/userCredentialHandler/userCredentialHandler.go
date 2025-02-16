package userCredentialHandler

import (
	"api_go/config"
	"api_go/internal/handlers/categoryHandler"
	"api_go/internal/models"
	"api_go/internal/utils"
	"api_go/internal/utils/types"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/jsonapi"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

func Login(c *fiber.Ctx) error {
	// parse body
	var body userCred
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(400).SendString("Error parsing body")
	}
	var userCredential models.UserCredential
	// get user credential from DB
	config.DB.Where("user_email =?", body.Email).First(&userCredential)
	// verify if user exists
	if userCredential.UserId == "" {
		return c.Status(fiber.StatusBadRequest).SendString("User not found")
	}
	// verify if password is correct
	err = bcrypt.CompareHashAndPassword([]byte(userCredential.Password), []byte(body.Password))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid password")
	}
	var user struct {
		ID          string    `json:"id"`
		Email       string    `json:"email"`
		Name        string    `json:"name"`
		LastName    string    `json:"last_name"`
		Gender      string    `json:"gender"`
		DateOfBirth time.Time `json:"date_of_birth"`
	}
	//get user information
	config.DB.Model(&models.User{}).Select(
		[]string{
			"id", "email", "name", "last_name", "gender", "date_of_birth",
		}).Where("id =?", userCredential.UserId).First(&user)
	if user.ID == "" {
		return c.Status(fiber.StatusBadRequest).SendString("User Information not found")
	}
	// return token
	claims := jwt.MapClaims{
		"userId": userCredential.UserId,
		"email":  userCredential.UserEmail,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWTSECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error generating token" + err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"userInfo": user,
		"token":    t,
	})
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
	newUser := models.User{}
	config.DB.Where("email = ?", body.UserInfo.Email).Select("id").First(&user)
	if id, err := uuid.Parse(user.ID); id != uuid.Nil && err == nil {
		fmt.Println(user.ID)
		return ctx.Status(fiber.StatusBadRequest).JSON(types.NewErrorResponse(fiber.StatusInternalServerError, "user already exists", ""))
	}
	//verify if its a strong password
	if !utils.IsStrongPassword(body.UserCred.Password) {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.NewErrorResponse(fiber.StatusBadRequest, "Password is not strong enough", ""))
	}
	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.UserCred.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password", err)
		return ctx.Status(500).JSON(types.NewErrorResponse(fiber.StatusBadRequest, "Error hashing password", ""))
	}
	err = config.DB.Transaction(func(tx *gorm.DB) error {
		// create user in DB
		newUser = models.User{
			Email:       body.UserInfo.Email,
			Name:        body.UserInfo.Name,
			Gender:      body.UserInfo.Gender,
			DateOfBirth: body.UserInfo.DateOfBirth,
			LastName:    body.UserInfo.LastName,
			ID:          body.UserInfo.ID,
			Model:       body.UserInfo.Model,
		}
		if err := tx.Create(&newUser).Error; err != nil {
			fmt.Println("Error creating user", err)
			return errors.New("Error creating user: " + err.Error())
			//return ctx.Status(500).SendString("Error creating user")
		} else {
			// create use credential in DB \
			userCred := models.UserCredential{
				UserId:    newUser.ID,
				UserEmail: body.UserCred.Email,
				Password:  string(hashedPassword),
			}
			if err := tx.Create(&userCred).Error; err != nil {
				//return ctx.Status(500).SendString("Error creating user credential")
				return errors.New("Error creating user credentials: " + err.Error())
			}
		}
		return nil
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.NewErrorResponse(fiber.StatusInternalServerError, "Error creating user: "+err.Error(), ""))
	}
	// link default categories
	// link categories to user
	errlist := categoryHandler.LinkDefaultCategories(newUser.ID)
	if len(errlist) > 0 {
		log.Print("Error linking default categories")
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.NewErrorResponse(fiber.StatusInternalServerError, errlist, "Error linking default categories"))
	}
	// enviar informacion de usuario

	//serialized in JSON:API
	fmt.Println(newUser)
	ctx.Set("Content-Type", jsonapi.MediaType)
	return ctx.Status(fiber.StatusOK).JSON(types.JSONAPIResponse{
		Data: types.JSONAPIResource{
			Type: "user",
			Id:   newUser.ID,
			Atributes: map[string]interface{}{
				"email":     newUser.Email,
				"name":      newUser.Name,
				"last_name": newUser.LastName,
			},
			Relationships: map[string]interface{}{},
		},
	})
}
