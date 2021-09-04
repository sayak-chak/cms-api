package server

import (
	"cms-api/config"
	"cms-api/modules/authors"
	"cms-api/modules/authors/contents"
	"cms-api/modules/subscribers"

	jwtware "github.com/gofiber/jwt/v2"

	"github.com/gofiber/fiber/v2"
)

func addRoutes(app *fiber.App, authorsHandler authors.Handler, subscribersHandler subscribers.Handler, contentsHandler contents.Handler) {
	addPublicRoutes(app, authorsHandler, subscribersHandler, contentsHandler)
	app.Use(getJwtMiddleWare())
	addPrivateRoutes(app, contentsHandler)

}

func addPublicRoutes(app *fiber.App, authorsHandler authors.Handler, subscribersHandler subscribers.Handler, contentsHandler contents.Handler) {
	app.Get("/authors", authorsHandler.GetAuthors)
	app.Post("/subscribe", subscribersHandler.AddSubscriber)
	app.Patch("/unsubscribe", subscribersHandler.DeleteSubscriber) //TODO: patch or put?
	app.Get("content/:contentId", contentsHandler.GetContent)
	app.Get("/top-contents", contentsHandler.TopContents)
	app.Get("/top-contents/:tag", contentsHandler.TopContentsByTag)
	app.Post("/register", authorsHandler.RegisterNewAuthor) //TODO: something very similar for pass reset
	app.Post("/login", authorsHandler.Login)
	app.Post("/upvote", contentsHandler.Upvote)
}

func addPrivateRoutes(app *fiber.App, contentsHandler contents.Handler) {
	app.Post("/add-content", contentsHandler.AddContent) //authors only
}

func getJwtMiddleWare() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(config.Salt),
	})
}
