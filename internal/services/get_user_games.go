package services

import (
	"context"

	"github.com/Hordevcom/GameShelf/internal/models"
)

func (s *Service) GetUserGames(ctx context.Context, username string) ([]models.UserGames, error) {
	return s.db.GetUserGames(ctx, username)
}
