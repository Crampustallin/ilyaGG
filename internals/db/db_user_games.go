package db

import "github.com/Crampustallin/ilyaGG/internals/models"

// INFO: the file is for requests to user_games table

func (db *DataBase) GetUserGames() (userGames map[int]models.UserGame, err error) {
	userGames = make(map[int]models.UserGame)

	query := "SELECT user_id, game_title, game_status, add_date FROM user_games WHERE game_status = 0"
	rows, err := db.pg.Queryx(query)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
	}()

	for rows.Next() {
		var userGame models.UserGame
		if err := rows.StructScan(&userGame); err != nil {
			return nil, err
		}

		userGames[userGame.UserId] = userGame
	}

	return userGames, nil
}
