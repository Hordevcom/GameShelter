package services

import (
	"context"

	"github.com/Hordevcom/GameShelf/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) AddUserToDB(ctx context.Context, user models.UserAuth) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.db.AddUserToDB(ctx, user.Username, string(hashedPassword))
}
