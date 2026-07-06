package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/heronhoga/sme-marketing-kit-be/handlers"
	"github.com/heronhoga/sme-marketing-kit-be/middlewares"
)

func AuthRoutes(r fiber.Router, authHandler *handlers.AuthHandler) {
	r.Post("/login", authHandler.Login)
	r.Post("/register", authHandler.Register)
	r.Get("/profile", middlewares.JWTMiddleware, authHandler.GetProfile)
}
