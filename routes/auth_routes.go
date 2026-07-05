package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/heronhoga/sme-marketing-kit-be/handlers"
)

func AuthRoutes(r fiber.Router, authHandler *handlers.AuthHandler) {
	r.Post("/login", authHandler.Login)
	r.Post("/register", authHandler.Register)
}
