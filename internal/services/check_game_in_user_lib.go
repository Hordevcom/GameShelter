package services

import (
	"context"

	"github.com/Hordevcom/GameShelf/internal/middleware/auth"
)

func (s *Service) CheckGameInUserLib(ctx context.Context, gametitle string, token string) (error, bool) {
	username := auth.GetUsername(token)
	return s.db.CheckGameInUserLib(ctx, gametitle, username)
}
