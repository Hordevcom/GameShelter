package storage

import (
	"context"

	"github.com/Hordevcom/GameShelf/internal/models"
)

func (p *PGDB) AddGameToUser(ctx context.Context, usergame models.UserGame) error {
	query := `INSERT INTO user_games (username, game_title, game_status, game_store)
				VALUES ($1, $2, $3, $4)`

	_, err := p.DB.Exec(ctx, query, usergame.Username, usergame.GameTitle,
		usergame.GameStatus, usergame.GameStore)

	return err
}
