package services

import (
	"context"

	"github.com/Hordevcom/GameShelf/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserAdder interface {
	AddUserToDB(ctx context.Context, username, password string) error
}

type UserAdd struct {
	UserAdder
}

func (u *UserAdd) AddUserToDB(ctx context.Context, user models.UserAuth) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return u.UserAdder.AddUserToDB(ctx, user.Username, string(hashedPassword))
}
