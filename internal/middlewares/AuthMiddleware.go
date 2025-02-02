package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
)

func AuthMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).SendString("Missing or malformed JWT")
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
		}
		return []byte(os.Getenv("JWTSECRET")), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid or expired JWT")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Locals("userId", claims["userId"])
		return c.Next()
	} else {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid JWT claims")
	}
}
