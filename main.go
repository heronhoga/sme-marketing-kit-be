package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/heronhoga/sme-marketing-kit-be/handlers"
	"github.com/heronhoga/sme-marketing-kit-be/repositories"
	"github.com/heronhoga/sme-marketing-kit-be/routes"
	"github.com/heronhoga/sme-marketing-kit-be/services"
)

func main() {
	app := fiber.New()

	authRepo := repositories.NewAuthRepository()
	authService := services.NewAuthService(authRepo)
	authHandler := handlers.NewAuthHandler(authService)

	routes.RoutesInit(app, authHandler)
	log.Fatal(app.Listen(":8000"))
}