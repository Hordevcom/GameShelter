package storage

import (
	"context"

	"github.com/Hordevcom/GameShelf/internal/models"
)

func (p *PGDB) GetUserGames(ctx context.Context, username string) ([]models.UserGames, error) {
	var usGames []models.UserGames

	query := `SELECT game_title, game_status, game_store, updated_at FROM user_games
 WHERE username = $1`
	rows, err := p.DB.Query(ctx, query, username)
	if err != nil {
		return []models.UserGames{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var g models.UserGames

		err := rows.Scan(&g.GameTitle, &g.GameStatus, &g.GameStore, &g.AddedAt)
		if err != nil {
			p.Logger.Error("something went wrong with rows in DB: ", err)
			return nil, err
		}

		usGames = append(usGames, g)
	}

	return usGames, err
}
