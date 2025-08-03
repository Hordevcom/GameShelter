package services

import (
	"context"
	"errors"

	"github.com/Hordevcom/GameShelf/internal/middleware/auth"
	"github.com/Hordevcom/GameShelf/internal/models"
)

type FriendReqDBAdder interface {
	CreateFriendReqDB(ctx context.Context, payload models.FriendRequestJSON) error
}

type FriendReqDBUpdater interface {
	UpdateFriendReqDB(ctx context.Context, payload models.FriendRequestJSON) error
}

type FriendReqDBDeleter interface {
	DeleteFriendReqDB(ctx context.Context, payload models.FriendRequestJSON) error
}

type FrindReqDBGetter interface {
	GetListOfReqDB(ctx context.Context, username string) ([]models.FriendRequestDB, error)
}

type FriendRequest struct {
	FriendReqDBAdder
	FriendReqDBUpdater
	FriendReqDBDeleter
	FrindReqDBGetter
}

func (f *FriendRequest) CreateFriendReqService(ctx context.Context, payload models.FriendRequestJSON, token string) error {
	username := auth.GetUsername(token)
	if payload.Sender != username {
		return errors.New("user and sender are different")
	}

	return f.CreateFriendReqDB(ctx, payload)
}

func (f *FriendRequest) UpdateFriendReqService(ctx context.Context, payload models.FriendRequestJSON, token string) error {
	username := auth.GetUsername(token)
	if payload.Sender != username {
		return errors.New("user and sender are different")
	}

	return f.UpdateFriendReqDB(ctx, payload)
}

func (f *FriendRequest) DeleteFriendReqService(ctx context.Context, payload models.FriendRequestJSON, token string) error {
	username := auth.GetUsername(token)
	if payload.Sender != username {
		return errors.New("user and sender are different")
	}

	return f.DeleteFriendReqDB(ctx, payload)
}

func (f *FriendRequest) GetListOfFriendReqService(ctx context.Context, token string) ([]models.FriendRequestDB, error) {
	username := auth.GetUsername(token)

	return f.GetListOfReqDB(ctx, username)
}
