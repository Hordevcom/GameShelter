package storage

import (
	"context"

	"github.com/Hordevcom/GameShelf/internal/models"
)

func (p *PGDB) AddGameToServer(ctx context.Context, game models.Game) error {
	query := `INSERT INTO games (title, genre)
				VALUES ($1, $2)`

	_, err := p.DB.Exec(ctx, query, game.Title, game.Genre)

	return err
}
