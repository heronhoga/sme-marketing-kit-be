package handlers

import (
	"github.com/gofiber/fiber/v3"
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
	return c.Status(200).JSON(fiber.Map{
		"error": false,
		"message": "Login Successful",
	})
}