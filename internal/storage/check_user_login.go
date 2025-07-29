package storage

import "context"

func (p *PGDB) CheckUserLogin(ctx context.Context, user string) bool {
	var username string
	query := `SELECT username FROM users WHERE username = $1`
	row := p.DB.QueryRow(ctx, query, user)
	row.Scan(&username)

	return username != ""
}
