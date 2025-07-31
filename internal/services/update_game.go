package services

import (
	"context"

	"github.com/Hordevcom/GameShelf/internal/middleware/auth"
	"github.com/Hordevcom/GameShelf/internal/models"
)

type UserGameDBUpdater interface {
	UpdateUserGame(ctx context.Context, userGameUpd models.UserGameUpdate, username string) error
}

type UserGameDbUpd struct {
	UserGameDBUpdater
}

func (u *UserGameDbUpd) UpdateGame(ctx context.Context, gameUpd models.UserGameUpdate, token string) error {
	username := auth.GetUsername(token)
	return u.UpdateUserGame(ctx, gameUpd, username)
}
