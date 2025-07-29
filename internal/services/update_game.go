package services

import (
	"context"

	"github.com/Hordevcom/GameShelf/internal/middleware/auth"
	"github.com/Hordevcom/GameShelf/internal/models"
)

func (s *Service) UpdateGame(ctx context.Context, gameUpd models.UserGameUpdate, token string) error {
	username := auth.GetUsername(token)
	return s.db.UodateUserGame(ctx, gameUpd, username)
}
