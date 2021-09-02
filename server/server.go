package server

import (
	"cms-api/database/postgres"
	"cms-api/modules/authors"
	"cms-api/modules/authors/contents"
	"cms-api/modules/subscribers"
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

	addAllRoutes(&database, app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	app.Listen(":" + string(port))
}

func addAllRoutes(database *postgres.Postgres, app *fiber.App) {
	authorsHandler := authors.NewHandler(database)
	subscribersHandler := subscribers.NewHandler(database)
	contentsHandler := contents.NewHandler(database)
	addRoutes(app, authorsHandler, subscribersHandler, contentsHandler)
}
