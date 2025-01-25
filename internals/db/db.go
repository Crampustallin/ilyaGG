package db

import (
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

func (db *DataBase) Close() (err error) {
	return db.Close()
}
