package models

import "time"

type UserAuth struct {
	Username string `json:"login"`
	Password string `json:"password,omitempty"`
}

type Game struct {
	Title string `json:"title,omitempty"`
	Genre string `json:"genre,omitempty"`
}

type UserGameJSON struct {
	GameTitle  string `json:"title,omitempty"`
	GameStatus string `json:"status,omitempty"`
	GameStore  string `json:"store,omitempty"`
}

type UserGame struct {
	Username   string
	GameTitle  string
	GameStatus string
	GameStore  string
}

type UserGames struct {
	GameTitle  string    `json:"title"`
	GameStatus string    `json:"status"`
	GameStore  string    `json:"store"`
	AddedAt    time.Time `json:"added_at"`
}

type UserGameUpdate struct {
	GameTitle  string `json:"title"`
	GameStatus string `json:"status"`
}
