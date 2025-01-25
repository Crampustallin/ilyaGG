package models

type ScoreBoardResp struct {
	Users []User
	Games map[string]Game
}
