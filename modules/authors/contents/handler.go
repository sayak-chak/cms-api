package contents

import (
	"cms-api/custom_errors"
	"cms-api/database"
	models "cms-api/models/requests"
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	AddContent(ctx *fiber.Ctx) error
	TopContentsByTag(ctx *fiber.Ctx) error
	TopContents(ctx *fiber.Ctx) error
	GetContent(ctx *fiber.Ctx) error
	Upvote(ctx *fiber.Ctx) error
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

func (handler *handler) AddContent(ctx *fiber.Ctx) error {
	var requestBody models.AddContentRequest
	ctxBody := append([]byte(nil), ctx.Body()...)

	err := json.Unmarshal(ctxBody, &requestBody)
	if err != nil {
		return custom_errors.CouldntCompleteYourOperation()
	}

	err = handler.service.AddContent(&requestBody, ctx)
	return err
}

func (handler *handler) TopContentsByTag(ctx *fiber.Ctx) error {
	topContentsForThisTag, err := handler.service.TopContentsByTag(ctx.Params("tag"))
	if err != nil {
		return custom_errors.CantGetResource()
	}

	topContentsJsonResponse, err := json.Marshal(&topContentsForThisTag)
	if err != nil {
		return custom_errors.GenericError()
	}

	return ctx.Status(200).Send([]byte(topContentsJsonResponse))
}

func (handler *handler) TopContents(ctx *fiber.Ctx) error {
	topContents, err := handler.service.TopContents()
	if err != nil {
		return custom_errors.CantGetResource()
	}

	topContentsJsonResponse, err := json.Marshal(&topContents)
	if err != nil {
		return custom_errors.GenericError()
	}

	return ctx.Status(200).Send([]byte(topContentsJsonResponse))
}

func (handler *handler) GetContent(ctx *fiber.Ctx) error {
	contentId, err := strconv.Atoi(ctx.Params("contentId"))
	if err != nil {
		return custom_errors.GenericError()
	}

	content, err := handler.service.GetContent(contentId)
	if err != nil {
		return custom_errors.GenericError()
	}

	contentContentJsonResponse, err := json.Marshal(&content)
	if err != nil {
		return custom_errors.GenericError()
	}

	return ctx.Status(200).Send([]byte(contentContentJsonResponse))
}

func (handler *handler) Upvote(ctx *fiber.Ctx) error {
	var requestBody models.UpvoteRequest
	ctxBody := append([]byte(nil), ctx.Body()...)

	err := json.Unmarshal(ctxBody, &requestBody)
	if err != nil {
		return custom_errors.CouldntCompleteYourOperation()
	}

	err = handler.service.Upvote(requestBody.ContentId)
	if err != nil {
		return custom_errors.GenericError()
	}

	return nil
}
