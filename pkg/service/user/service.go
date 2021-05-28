package user

import (
	"time"

	"github.com/ceng316/dentist-backend/pkg/model"
	"github.com/ceng316/dentist-backend/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

type Service struct {
	repository repository.Repository
}

func NewService(repo repository.Repository) (*Service, error) {
	return &Service{
		repository: repo,
	}, nil
}

func (s *Service) Login(user model.User) (*model.User, error) {

	u, err := s.repository.GetUserRepository().GetUser(user)
	if err != nil {
		return nil, err
	}
	atClaims := jwt.MapClaims{}
	atClaims["email"] = u.Userdata.Email
	atClaims["fullName"] = u.Userdata.Fullname
	atClaims["exp"] = time.Now().Add(time.Hour * 24 * 12).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte("your-super-secret"))
	u.Accesstoken = token
	u.Userdata.Password = ""
	if err != nil {
		return nil, err
	}

	return u, err
}
