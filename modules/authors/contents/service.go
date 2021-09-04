package contents

import (
	"cms-api/config"
	"cms-api/custom_errors"
	"cms-api/database"
	models "cms-api/models/requests"
	responseModels "cms-api/models/responses"
	"cms-api/modules/authors/contents/cache_layer"
	"cms-api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type Service interface {
	AddContent(image *[]byte, content *string) error
	TopContents() error
}

type service struct {
	database database.Database
}

func (s *service) TopContents() (*[]responseModels.TopContentsResponse, error) {
	cachedRes := cache_layer.GetCachedResponseIfPossible(config.CommonCacheTag)
	if cachedRes != nil {
		return cachedRes, nil
	}
	topContents, err := s.database.TopContents()
	if err != nil {
		return nil, err
	}
	cache_layer.CacheThis(topContents, config.CommonCacheTag)
	return topContents, nil
}

func (s *service) AddContent(addContentRequest *models.AddContentRequest, ctx *fiber.Ctx) error {

	//TODO: If needed, cleanup token when logged out/logged in from multiple devices

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

	for _, tag := range addContentRequest.Tags {
		if tagTable, isValid := utils.GetTagDBNameIfValid(tag); !isValid {
			return custom_errors.NoSuchTag()
		} else {
			err = s.database.AddContentToTag(utils.GetTagTableModelFor(tagTable, contentId))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *service) TopContentsByTag(inputTag string) (*[]responseModels.TopContentsResponse, error) {
	tag, isValid := utils.GetTagDBNameIfValid(inputTag)
	if !isValid {
		return nil, custom_errors.NoSuchTag()
	}

	cachedRes := cache_layer.GetCachedResponseIfPossible(tag)
	if cachedRes != nil {
		return cachedRes, nil
	}
	topContentsForThisTag, err := s.database.TopContentsByTag(tag)
	if err != nil {
		return nil, err
	}
	cache_layer.CacheThis(topContentsForThisTag, tag)
	return topContentsForThisTag, nil
}

func (s *service) GetContent(contentId int) (*responseModels.ReadContentResponse, error) {
	return s.database.GetContent(contentId)

}

func (s *service) Upvote(aricleId int) error {
	return s.database.Upvote(aricleId)
}
