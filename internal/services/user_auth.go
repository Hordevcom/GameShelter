package services

import (
	"context"

	"github.com/Hordevcom/GameShelf/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserChecker interface {
	GetUserPassword(ctx context.Context, username string) (string, error)
	CheckUserLogin(ctx context.Context, user string) bool
}

type AuthService struct {
	UserChecker
}

func (a *AuthService) CheckUserPassword(ctx context.Context, user models.UserAuth) error {
	usPass, err := a.UserChecker.GetUserPassword(ctx, user.Username)
	if err != nil {
		return err
	}
	return bcrypt.CompareHashAndPassword([]byte(usPass), []byte(user.Password))
}

func (a *AuthService) CheckUserLoginService(ctx context.Context, user models.UserAuth) bool {
	return a.UserChecker.CheckUserLogin(ctx, user.Username)
}
