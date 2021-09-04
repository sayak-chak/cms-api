package subscribers

import (
	"cms-api/database"
	models "cms-api/models/requests"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	AddSubscriber(ctx *fiber.Ctx) error
	DeleteSubscriber(ctx *fiber.Ctx) error
}

func NewHandler(database database.Database) Handler {
	return &handler{
		service: &service{
			database: database,
		},
	}
}

type handler struct {
	service *service
}

func (h *handler) AddSubscriber(ctx *fiber.Ctx) error {
	var subscriberEmail models.SubscriberRequest

	err := json.Unmarshal(ctx.Body(), &subscriberEmail)
	if err != nil {
		return err
	}

	err = h.service.AddSubscriber(subscriberEmail.Email)
	if err != nil {
		return err
	}
	return ctx.SendStatus(204)
}

func (h *handler) DeleteSubscriber(ctx *fiber.Ctx) error {
	var subscriberEmail models.SubscriberRequest

	err := json.Unmarshal(ctx.Body(), &subscriberEmail)
	if err != nil {
		return err
	}

	err = h.service.DeleteSubscriber(subscriberEmail.Email)
	if err != nil {
		return err
	}
	return ctx.SendStatus(204)
}
