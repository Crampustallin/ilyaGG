package user_service

import (
	"encoding/base64"

	"github.com/Crampustallin/ilyaGG/internals/interfaces"
	"github.com/Crampustallin/ilyaGG/internals/models"
	"golang.org/x/crypto/argon2"
)

type UserService struct {
	dataBase interfaces.DataBase
}

func New(dataBase interfaces.DataBase) *UserService {
	return &UserService{
		dataBase: dataBase,
	}
}

func (s *UserService) GetUsers() (users []models.User, err error) {
	return s.dataBase.GetUsers()
}

func (s *UserService) SetUser(userName, pass string) (err error) {
	hashBytes := argon2.IDKey([]byte(pass), []byte{}, 1, 64*1024, 4, 32)
	passHash := base64.RawStdEncoding.EncodeToString(hashBytes)

	err = s.dataBase.InsertUser(userName, passHash)
	if err != nil {
		return err
	}

	return nil
}
