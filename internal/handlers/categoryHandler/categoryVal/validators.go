package categoryVal

import (
	"api_go/internal/models"
	"api_go/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateCategoryValidator(ctx *fiber.Ctx) error {
	body := new(models.Category)
	return utils.CheckValidations(ctx, body)
}
