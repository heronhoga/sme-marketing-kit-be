package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/heronhoga/sme-marketing-kit-be/models"
	"github.com/heronhoga/sme-marketing-kit-be/services"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	var loginRequest models.LoginRequest
	err := c.Bind().Body(&loginRequest)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": true,
			"message": "Invalid Body Request",
		})
	}

	if loginRequest.Email == "" || loginRequest.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": true,
			"message": "All Fields Are Required",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"error": false,
		"message": "Login Successful",
	})
}