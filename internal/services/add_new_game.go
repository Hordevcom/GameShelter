package services

import (
	"context"

	"github.com/Hordevcom/GameShelf/internal/models"
)

func (s *Service) AddNewGame(ctx context.Context, game models.Game) error {
	return s.db.AddGameToServer(ctx, game)
}
