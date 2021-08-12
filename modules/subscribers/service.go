package subscribers

import (
	"cms-api/database"
	"regexp"
)

type service struct {
	database database.Database
}

func (s *service) AddSubscriber(subscriberEmail string) error {
	if emailValid := isEmailValid(subscriberEmail); emailValid {
		return s.database.AddSubscriber(subscriberEmail)
	}
	return nil
}

func (s *service) DeleteSubscriber(subscriberEmail string) error {
	return s.database.DeleteSubscriber(subscriberEmail)
}

func isEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}
