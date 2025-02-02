package main

import (
	"api_go/config"
	"api_go/config/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	config.LoadEnv()
	app := fiber.New()
	config.ConnectDB()
	fmt.Println("Hello, World!")
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("API GO running")
		return err
	})
	// registrar rutas
	routes.SetUpRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
