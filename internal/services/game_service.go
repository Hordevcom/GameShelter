package services

import (
	"context"

	"github.com/Hordevcom/GameShelf/internal/models"
)

type GameServerMaker interface {
	CheckGameExists(ctx context.Context, gamename string) (error, bool)
	AddGameToServer(ctx context.Context, game models.Game) error
}

type GameServerService struct {
	GameServerMaker
}

func (s *GameServerService) IsGameAlreadyExist(ctx context.Context, gamename string) (error, bool) {
	return s.CheckGameExists(ctx, gamename)
}

func (s *GameServerService) AddNewGame(ctx context.Context, game models.Game) error {
	return s.AddGameToServer(ctx, game)
}
