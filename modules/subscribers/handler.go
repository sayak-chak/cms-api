package subscribers

import (
	"cms-api/database/postgres"
	models "cms-api/models/requests"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service *service
}

func (h *handler) AddSubscriber(ctx *fiber.Ctx) error {
	var subscriberEmail models.SubscriberRequest

	err := json.Unmarshal(ctx.Body(), &subscriberEmail)
	if err != nil {
		return err
	}

	return h.service.AddSubscriber(subscriberEmail.Email)
}

func (h *handler) DeleteSubscriber(ctx *fiber.Ctx) error {
	var subscriberEmail models.SubscriberRequest

	err := json.Unmarshal(ctx.Body(), &subscriberEmail)
	if err != nil {
		return err
	}

	return h.service.DeleteSubscriber(subscriberEmail.Email)
}

func NewHandler() *handler {
	return &handler{
		service: &service{
			database: &postgres.Postgres{},
		},
	}
}
