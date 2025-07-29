package storage

import (
	"context"
)

func (p *PGDB) AddUserToDB(ctx context.Context, username, password string) error {
	var user string
	query := `INSERT INTO users (username, password_hash)
				VALUES ($1, $2) ON CONFLICT (username) DO NOTHING
				RETURNING username`
	err := p.DB.QueryRow(ctx, query, username, password).Scan(&user)

	return err
}
