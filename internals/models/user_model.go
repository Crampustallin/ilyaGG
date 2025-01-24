package models

type User struct {
	UserId    int    `json:"user_id"`
	UserName  string `json:"user_name"`
	UserScore int    `json:"user_score"`
}
