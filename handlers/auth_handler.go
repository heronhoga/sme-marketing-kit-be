package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/heronhoga/sme-marketing-kit-be/models"
	"github.com/heronhoga/sme-marketing-kit-be/services"
	"github.com/heronhoga/sme-marketing-kit-be/utils"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Register(c fiber.Ctx) error {
	var registerRequest models.RegisterRequest
	err := c.Bind().Body(&registerRequest)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid Body Request",
		})
	}

	if registerRequest.Email == "" || registerRequest.Name == "" || registerRequest.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"error":   true,
			"message": "All Fields Are Required",
		})
	}

	// check if email pattern is right
	err = utils.CheckEmailPattern(registerRequest.Email)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	err = h.service.RegisterService(c, registerRequest)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"error":   false,
		"message": "Register Successful",
	})
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	var loginRequest models.LoginRequest
	err := c.Bind().Body(&loginRequest)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid Body Request",
		})
	}

	if loginRequest.Email == "" || loginRequest.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"error":   true,
			"message": "All Fields Are Required",
		})
	}

	err = utils.CheckEmailPattern(loginRequest.Email)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid Email Format",
		})
	}

	accessToken, refreshToken, err := h.service.LoginService(c, loginRequest)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"error":         false,
		"message":       "Login Successful",
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func (h *AuthHandler) GetProfile(c fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"error":   false,
		"message": "Profile Data Retrieved Successfully",
	})
}
