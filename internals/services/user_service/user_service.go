package user_service

import (
	"encoding/base64"
	"errors"
	"fmt"

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

func (s *UserService) GetUser(userName, pass string) (err error) {
	hashBytes := argon2.IDKey([]byte(pass), []byte{}, 1, 64*1024, 4, 32)
	passHash := base64.RawStdEncoding.EncodeToString(hashBytes)
	u, err := s.dataBase.GetUser(userName, passHash)
	if err != nil {
		return err
	}

	if u == nil {
		return errors.New("no data")
	}

	fmt.Println(u)

	return nil
}

func (s *UserService) GetScoreBoard() (err error) {
	// TODO: get users since it's the beta get all users anyway
	// in the ftuture maybe we should create something like groups and get score board informatioin only for users in this group

	// TODO: get user_games that is status is in progress
	// if some users don't have games like that then just NA the game. But send their data of the score and the name anyway

	return
}
