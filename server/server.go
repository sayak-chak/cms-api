package server

import (
	"cms-api/database/postgres"
	"os"

	"github.com/gofiber/fiber/v2/middleware/cors"

	"log"

	"github.com/gofiber/fiber/v2"
)

func StartNewServer() {
	database := postgres.Postgres{}
	err := database.SeedWithMockData()
	if err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	app.Use(cors.New()) //TODO : be more restrictive in prod

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	app.Listen(":" + string(port))
}
