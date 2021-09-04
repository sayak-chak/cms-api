package authors

import (
	"cms-api/config"
	"cms-api/database"
	models "cms-api/models/requests"
	responses "cms-api/models/responses"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

const (
	otpExpiryPeriod = 21600 // 6 hours
	bcryptCost      = 12
)

type Service interface {
	GetAuthors() []string
}

type service struct {
	database database.Database
}

type Author struct {
	Name string `json:"name"`
}

func (s *service) GetAuthors() []Author {
	var authorList []Author
	authorNames := s.database.GetAuthors()
	for _, authorName := range authorNames {
		authorList = append(authorList, Author{
			Name: authorName,
		})
	}

	return authorList
}

func (s *service) RegisterNewAuthor(newAuthorAccCreationReq *models.NewAuthorAccCreationRequest) error {
	//TODO: implement checks

	err := s.database.CleanUpOldOtps(time.Now().Unix(), otpExpiryPeriod)
	if err != nil {
		return err
	}

	phone, err := s.database.GetPhoneFor(newAuthorAccCreationReq.Otp)
	if err != nil {
		return err
	}

	password := []byte(newAuthorAccCreationReq.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcryptCost)
	if err != nil {
		return err
	}

	authorCredId, err := s.database.AddAuthorCreds(newAuthorAccCreationReq, phone, string(hashedPassword))
	if err != nil {
		return err
	}

	err = s.database.AddAuthorDetails(newAuthorAccCreationReq, authorCredId)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Login(loginReq *models.LoginRequest) (*responses.LoginResponse, error) {
	//TODO: implement checks

	mobileNumber := loginReq.Mobile
	hashedPassword, authCredsId, err := s.database.GetPasswordAndIdFor(mobileNumber)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(loginReq.Password))
	if err != nil {
		return nil, err
	}

	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["authorId"] = authCredsId
	claims["exp"] = time.Now().Add(config.JwtExpirationPeriod).Unix()

	// Generate encoded token and send it as response.

	encodedToken, err := token.SignedString([]byte(config.Salt))

	// To retrieve tokens,
	// user := c.Locals("user").(*jwt.Token)
	// claims := user.Claims.(jwt.MapClaims)
	// authorId := claims["authorId"].(string)

	if err != nil {
		return nil, err
	}

	return &responses.LoginResponse{
		Token:    encodedToken,
		AuthorId: authCredsId,
	}, nil
}
