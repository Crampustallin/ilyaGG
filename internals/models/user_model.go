package models

type User struct {
	UserId    int      `json:"user_id" db:"user_id"`
	UserName  string   `json:"user_name" db:"user_name"`
	UserScore int      `json:"user_score" db:"user_score"`
	UserGame  UserGame `json:"user_game"`
}

type UserGame struct {
	GameTitle  string `json:"game_title" db:"game_title"`
	GameStatus int    `json:"game_status" db:"game_status"`
	AddDate    string `json:"add_dt" db:"add_dt"`
}
