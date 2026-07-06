package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/heronhoga/sme-marketing-kit-be/utils"
)

func JWTMiddleware(c fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{
			"error":   true,
			"message": "Unauthorized",
		})
	}

	authHeader = authHeader[7:]
	secret := os.Getenv("JWT_SECRET")
	claims, err := utils.VerifyAccessToken(authHeader, secret)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error":   true,
			"message": "Unauthorized",
		})
	}
	c.Locals("userId", claims.UserID)

	return c.Next()
}
