package services

import (
	"context"

	"github.com/Hordevcom/GameShelf/internal/models"
)

func (s *Service) CheckUserLogin(ctx context.Context, user models.UserAuth) bool {
	return s.db.CheckUserLogin(ctx, user.Username)
}
