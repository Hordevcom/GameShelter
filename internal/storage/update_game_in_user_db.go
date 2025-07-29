package storage

import (
	"context"

	"github.com/Hordevcom/GameShelf/internal/models"
)

func (p *PGDB) UodateUserGame(ctx context.Context, userGameUpd models.UserGameUpdate, username string) error {
	query := `UPDATE user_games 
              SET game_status = $1 
              WHERE username = $2 AND game_title = $3`

	_, err := p.DB.Exec(ctx, query, userGameUpd.GameStatus, username, userGameUpd.GameTitle)

	return err
}
