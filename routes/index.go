package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/heronhoga/sme-marketing-kit-be/handlers"
)

func RoutesInit(r *fiber.App, authHandler *handlers.AuthHandler) {
	api := r.Group("/api")
	AuthRoutes(api, authHandler)
	
}