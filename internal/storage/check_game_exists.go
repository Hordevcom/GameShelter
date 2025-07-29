package storage

import "context"

func (p *PGDB) CheckGameExists(ctx context.Context, gamename string) (error, bool) {
	var exist bool

	query := `SELECT EXISTS(SELECT 1 FROM games WHERE title = $1);`
	err := p.DB.QueryRow(ctx, query, gamename).Scan(&exist)

	if err != nil {
		return err, false
	}

	return nil, exist

}
