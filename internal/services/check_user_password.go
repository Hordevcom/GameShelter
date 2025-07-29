package services

import (
	"context"

	"github.com/Hordevcom/GameShelf/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) CheckUserPassword(ctx context.Context, user models.UserAuth) error {
	passFromDB, err := s.db.GetUserPassword(ctx, user.Username)
	if err != nil {
		return err
	}
	return bcrypt.CompareHashAndPassword([]byte(passFromDB), []byte(user.Password))
}
