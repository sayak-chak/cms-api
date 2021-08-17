package database

import (
	requestModels "cms-api/models/requests"
	responseModels "cms-api/models/responses"
)

type repository interface {
	GetAuthors() []string
	AddSubscriber(subscriberEmail string) error
	DeleteSubscriber(subscriberEmail string) error
	AddContent(addContentRequest *requestModels.AddContentRequest) (int, error)
	AddContentToTag(tagTableModel interface{}) error
	TopContents() (*[]responseModels.TopContentsResponse, error)
	TopContentsByTag(tagTable string) (*[]responseModels.TopContentsResponse, error)
	GetPhoneFor(otp string) (int, error)
	CleanUpOldOtps(currTime int64, otpExpiryPeriod int) error
	AddAuthorCreds(newAuthorAccCreationReq *requestModels.NewAuthorAccCreationRequest, phone int, hashedPassword string) (int, error)
	AddAuthorDetails(newAuthorAccCreationReq *requestModels.NewAuthorAccCreationRequest, authorCredId int) error
	GetPasswordAndIdFor(mobileNumber int) (string, int, error)
	GetContent(contentId int) (*responseModels.ReadContentResponse, error)
	Upvote(aricleId int) error
}
