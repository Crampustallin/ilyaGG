package db

import "github.com/Crampustallin/ilyaGG/internals/models"

// INFO: the file is for requests to users table

func (db *DataBase) GetUsers() (users []models.User, err error) {
	query := "SELECT user_id, user_name, user_score FROM users ORDER BY user_score"
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

func (db *DataBase) GetUser(userName, passHash string) (user *models.User, err error) {
	query := "SELECT user_id, user_name, user_score FROM users WHERE user_name = $1 AND pass_hash = $2"
	rows, err := db.pg.Queryx(query, userName, passHash)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
	}()

	for rows.Next() {
		if err := rows.StructScan(&user); err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (db *DataBase) InsertUser(userName, passHash string) (err error) {
	query := "INSERT INTO users (user_name, pass_hash, user_score) VALUES ($1, $2, 0)"
	_, err = db.pg.Exec(query, userName, passHash)
	if err != nil {
		panic(err)
	}

	return nil
}
