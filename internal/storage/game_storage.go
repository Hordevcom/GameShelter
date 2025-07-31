package storage

import (
	"context"

	"github.com/Hordevcom/GameShelf/internal/config"
	"github.com/Hordevcom/GameShelf/internal/middleware/logging"
	"github.com/Hordevcom/GameShelf/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type GameStorage struct {
	DB     *pgxpool.Pool
	Conf   config.Config
	Logger *logging.Logger
}

func NewGameStorage(Conf config.Config, Logger *logging.Logger, db *pgxpool.Pool) *GameStorage {
	return &GameStorage{DB: db, Conf: Conf, Logger: Logger}
}

func (g *GameStorage) CheckGameExists(ctx context.Context, gamename string) (error, bool) {
	var exist bool

	query := `SELECT EXISTS(SELECT 1 FROM games WHERE title = $1);`
	err := g.DB.QueryRow(ctx, query, gamename).Scan(&exist)

	if err != nil {
		return err, false
	}

	return nil, exist
}

func (g *GameStorage) AddGameToServer(ctx context.Context, game models.Game) error {
	query := `INSERT INTO games (title, genre)
				VALUES ($1, $2)`

	_, err := g.DB.Exec(ctx, query, game.Title, game.Genre)

	return err
}
