package contents

import (
	"cms-api/custom_errors"
	"cms-api/database"
	models "cms-api/models/requests"
	responseModels "cms-api/models/responses"
	"cms-api/utils"

	// "fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type Service interface {
	AddContent(image *[]byte, content *string) error
	TopContents() error
}

type service struct {
	database database.Database
}

func (s *service) TopContents() (*[]responseModels.TopContentsResponse, error) {
	return s.database.TopContents()
}

func (s *service) AddContent(addContentRequest *models.AddContentRequest, ctx *fiber.Ctx) error {

	// If needed, cleanup token when logged out/logged in from multiple devices

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authorId := claims["authorId"].(float64)

	if authorId != float64(addContentRequest.AuthorId) {
		return custom_errors.CouldntCompleteYourOperation()
	}

	//TODO: implement length checks and other stuff
	contentId, err := s.database.AddContent(addContentRequest)
	if err != nil {
		return err
	}

	for _, genre := range addContentRequest.Tags {
		if genreTable, isValid := utils.GetTagDBNameIfValid(genre); !isValid {
			return custom_errors.NoSuchTag()
		} else {
			err = s.database.AddContentToTag(utils.GetTagTableModelFor(genreTable, contentId))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *service) TopContentsByTag(genre string) (*[]responseModels.TopContentsResponse, error) {
	if genreTable, isValid := utils.GetTagDBNameIfValid(genre); !isValid {
		return nil, custom_errors.NoSuchTag()
	} else {
		return s.database.TopContentsByTag(genreTable)
	}

}

func (s *service) GetContent(contentId int) (*responseModels.ReadContentResponse, error) {
	return s.database.GetContent(contentId)

}

func (s *service) Upvote(aricleId int) error {
	return s.database.Upvote(aricleId)
}
