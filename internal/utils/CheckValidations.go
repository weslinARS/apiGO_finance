package utils

import (
	"api_go/internal/utils/types"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"time"
)

var Validator = validator.New()

func init() {
	fmt.Println("Initializing validators...")
	err := Validator.RegisterValidation("datetimeF", IsDate)
	if err != nil {
		fmt.Println("Error registering datetimeF validator")
	}
}
func CheckValidations(ctx *fiber.Ctx, s interface{}) error {
	var errors []*types.IError
	errParse := ctx.BodyParser(s)
	if errParse != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": "Error parsing body: " + errParse.Error(),
			},
		)
	}
	err := Validator.Struct(s)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el types.IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}
	return ctx.Next()
}

func IsDate(fl validator.FieldLevel) bool {
	if fl.Field().String() == "" {
		return false
	}
	dateFormats := []string{
		"2006-01-02",
		"2006-01-02T15:04:05Z07:00",
		"01-02-2006",
		"01/02/2006",
		"02-01-2006",
		"02/01/2006",
		"Mon Jan 2 15:04:05 MST 2006",
		"Mon, 02 Jan 2006 15:04:05 MST",
		"Monday, 02-Jan-06 15:04:05 MST",
		"Mon Jan _2 15:04:05 2006",
	}

	for _, format := range dateFormats {
		if _, err := time.Parse(format, fl.Field().String()); err == nil {
			return true
		}
	}
	return false

}
