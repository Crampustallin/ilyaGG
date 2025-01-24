package interfaces

import "github.com/Crampustallin/ilyaGG/internals/models"

type DataBase interface {
	Close() (err error)
	GetUsers() (users []models.User, err error)
	InsertUser(userName, passHash string) (err error)
}
