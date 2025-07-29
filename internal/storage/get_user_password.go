package storage

import "context"

func (p *PGDB) GetUserPassword(ctx context.Context, username string) (string, error) {
	var userPassword string

	query := `SELECT password_hash FROM users WHERE username = $1`
	err := p.DB.QueryRow(ctx, query, username).Scan(&userPassword)

	return userPassword, err
}
