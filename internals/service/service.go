package service

import (
	"os"

	"github.com/Crampustallin/ilyaGG/internals/db"
	"github.com/Crampustallin/ilyaGG/internals/models"
	"github.com/Crampustallin/ilyaGG/internals/service/user_service"
)

type Service struct {
	userService UserService
}

const CONNECTION_STRING_ENV_NAME = "DB_CONNECTION_STRING"

func New() *Service {
	dataBase := db.New("postgres", os.Getenv(CONNECTION_STRING_ENV_NAME))
	return &Service{
		userService: user_service.New(dataBase),
	}
}

type UserService interface {
	GetUsers() (users []models.User, err error)
	SetUser(userName, pass string) (err error)
}
