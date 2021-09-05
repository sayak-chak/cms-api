package server

import (
	"cms-api/config"
	"cms-api/database/postgres"
	"cms-api/modules/authors"
	"cms-api/modules/authors/contents"
	"cms-api/modules/subscribers"

	"github.com/gofiber/fiber/v2/middleware/cors"

	"log"

	"github.com/gofiber/fiber/v2"
)

func StartNewServer(env string) {
	database := postgres.Postgres{}
	err := database.InitializeDatabase()
	if err != nil {
		log.Fatal(err)
	}
	err = database.SeedWithMockData(env)
	if err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	app.Use(cors.New()) //TODO : be more restrictive in prod

	addAllRoutes(&database, app)

	port := config.Port

	app.Listen(":" + string(port))
}

func addAllRoutes(database *postgres.Postgres, app *fiber.App) {
	authorsHandler := authors.NewHandler(database)
	subscribersHandler := subscribers.NewHandler(database)
	contentsHandler := contents.NewHandler(database)
	addRoutes(app, authorsHandler, subscribersHandler, contentsHandler)
}
