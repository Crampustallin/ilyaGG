package models

type Game struct {
	GameTitle  string `json:"game_title" db:"game_title"`
	GameArtUrl string `json:"game_art_url" db:"game_art_url"`
}
