package services

import (
	"context"

	"github.com/Hordevcom/GameShelf/internal/middleware/auth"
)

type FriendsServiceGetter interface {
	GetFriendsForUserDB(ctx context.Context, username string) ([]string, error)
}

type FriendsServiceAdder interface {
	AddFriendForUserDB(ctx context.Context, username, friend string) error
}

type FriendsServiceRemover interface {
	DeleteFriendForUserDB(ctx context.Context, user, friend string) error
}

type Friends struct {
	FriendsServiceGetter
	FriendsServiceAdder
	FriendsServiceRemover
}

func (f *Friends) GetFriendsForUserService(ctx context.Context, token string) ([]string, error) {
	username := auth.GetUsername(token)

	return f.GetFriendsForUserDB(ctx, username)
}

func (f *Friends) AddFriendForUserService(ctx context.Context, token, friend string) error {
	username := auth.GetUsername(token)

	return f.AddFriendForUserDB(ctx, username, friend)
}

func (f *Friends) DeleteFriendForUserService(ctx context.Context, token, friend string) error {
	username := auth.GetUsername(token)

	return f.DeleteFriendForUserDB(ctx, username, friend)

}
