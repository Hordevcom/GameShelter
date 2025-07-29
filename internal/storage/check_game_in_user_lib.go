package storage

import "context"

func (p *PGDB) CheckGameInUserLib(ctx context.Context, gametitle, username string) (error, bool) {
	var exist bool

	query := `SELECT EXISTS(SELECT 1 FROM user_games WHERE game_title = $1 AND username = $2);`
	err := p.DB.QueryRow(ctx, query, gametitle, username).Scan(&exist)

	if err != nil {
		return err, false
	}

	return nil, exist
}
