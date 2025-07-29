package services

import (
	"context"

	"github.com/Hordevcom/GameShelf/internal/middleware/auth"
	"github.com/Hordevcom/GameShelf/internal/models"
)

func (s *Service) AddGameToUser(ctx context.Context, usergame models.UserGameJSON, token string) error {
	username := auth.GetUsername(token)

	err := s.db.AddGameToUser(ctx, models.UserGame{
		Username:   username,
		GameTitle:  usergame.GameTitle,
		GameStatus: usergame.GameStatus,
		GameStore:  usergame.GameStore,
	})

	return err
}
