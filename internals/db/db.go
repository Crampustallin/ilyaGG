package db

import (
	"github.com/Crampustallin/ilyaGG/internals/models"
	"github.com/jmoiron/sqlx"
)

type DataBase struct {
	pg *sqlx.DB
}

func New(driver, connectionString string) *DataBase {
	pg, err := sqlx.Connect(driver, connectionString)
	if err != nil {
		panic(err)
	}

	return &DataBase{
		pg: pg,
	}
}

func (db *DataBase) GetUsers() (users []models.User, err error) {
	query := "SELECT user_id, user_name, user_score FROM users"
	rows, err := db.pg.Queryx(query)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
	}()

	for rows.Next() {
		var user models.User
		if err := rows.StructScan(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (db *DataBase) InsertUser(userName, passHash string) (err error) {
	query := "INSERT INTO users (user_name, pass_hash) VALUES ($1, $2)"
	_, err = db.pg.Exec(query, userName, passHash)
	if err != nil {
		return err
	}

	return nil
}
