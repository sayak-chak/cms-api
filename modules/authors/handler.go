package authors

import (
	"cms-api/custom_errors"
	"cms-api/database/postgres"
	models "cms-api/models/requests"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service *service
}

func (h *handler) GetAuthors(ctx *fiber.Ctx) error {
	authorList := h.service.GetAuthors()
	authorListJson, err := json.Marshal(authorList)
	if err != nil {
		return err
	}
	fmt.Println(string(authorListJson))
	return nil
}

func (h *handler) RegisterNewAuthor(ctx *fiber.Ctx) error {
	var newAuthorAccCreationReq models.NewAuthorAccCreationRequest
	ctxBody := append([]byte(nil), ctx.Body()...)

	err := json.Unmarshal(ctxBody, &newAuthorAccCreationReq)
	if err != nil {
		return custom_errors.GenericError()
	}

	err = h.service.RegisterNewAuthor(&newAuthorAccCreationReq)
	if err != nil {
		return custom_errors.CantGetResource()
	}
	ctx.Response().SetStatusCode(201)
	return nil
}

func (h *handler) Login(ctx *fiber.Ctx) error {
	var loginReq models.LoginRequest
	ctxBody := append([]byte(nil), ctx.Body()...)

	err := json.Unmarshal(ctxBody, &loginReq)
	if err != nil {
		return custom_errors.GenericError()
	}

	loginResponse, err := h.service.Login(&loginReq)
	if err != nil {
		return custom_errors.CouldntCompleteYourOperation()
	}
	loginResponseJson, err := json.Marshal(&loginResponse)
	if err != nil {
		return custom_errors.CantGetResource()
	}

	return ctx.Status(200).Send([]byte(loginResponseJson))
}

func NewHandler() *handler {
	return &handler{
		service: &service{
			database: &postgres.Postgres{},
		},
	}
}
