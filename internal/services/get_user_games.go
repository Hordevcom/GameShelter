package services

import (
	"context"

	"github.com/Hordevcom/GameShelf/internal/models"
)

type UserGameDBGetter interface {
	GetUserGamesDB(ctx context.Context, username string) ([]models.UserGames, error)
}

type UserGamesGet struct {
	UserGameDBGetter
}

func (u *UserGamesGet) GetUserGames(ctx context.Context, username string) ([]models.UserGames, error) {
	return u.GetUserGamesDB(ctx, username)
}
