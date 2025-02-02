package userCredentialHandler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {

	return c.SendString("Login")
}

func SignUp(ctx *fiber.Ctx) error {
	fmt.Print("entrooo")
	// get the body of the request
	body := new()
	return ctx.Status(200).SendString("UsuarioLoggeado")
}
