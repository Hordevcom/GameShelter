package services

import (
	"context"

	"github.com/Hordevcom/GameShelf/internal/middleware/auth"
	"github.com/Hordevcom/GameShelf/internal/models"
)

type UserGameAdderDBAdder interface {
	AddGameToUserDB(ctx context.Context, usergame models.UserGame) error
	CheckGameInUserLibDB(ctx context.Context, gametitle, username string) (error, bool)
}

type UserAddGame struct {
	UserGameAdderDBAdder
}

func (u *UserAddGame) AddGameToUser(ctx context.Context, usergame models.UserGameJSON, token string) error {
	username := auth.GetUsername(token)

	err := u.AddGameToUserDB(ctx, models.UserGame{
		Username:   username,
		GameTitle:  usergame.GameTitle,
		GameStatus: usergame.GameStatus,
		GameStore:  usergame.GameStore,
	})

	return err
}

func (u *UserAddGame) CheckGameInUserLib(ctx context.Context, gametitle string, token string) (error, bool) {
	username := auth.GetUsername(token)
	return u.CheckGameInUserLibDB(ctx, gametitle, username)
}
