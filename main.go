package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/heronhoga/sme-marketing-kit-be/config"
	"github.com/heronhoga/sme-marketing-kit-be/handlers"
	"github.com/heronhoga/sme-marketing-kit-be/repositories"
	"github.com/heronhoga/sme-marketing-kit-be/routes"
	"github.com/heronhoga/sme-marketing-kit-be/services"
	"github.com/joho/godotenv"
)

func main() {
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Print("Using System Env")
	} else {
		log.Print("Using App Env")
	}

	// connect to database
	dbconn, err := config.DBInit()
	if err != nil {
		log.Fatal(err)
	}

	
	app := fiber.New()

	authRepo := repositories.NewAuthRepository(dbconn)
	authService := services.NewAuthService(authRepo)
	authHandler := handlers.NewAuthHandler(authService)

	routes.RoutesInit(app, authHandler)
	log.Fatal(app.Listen(":8000"))
}