package storage

import (
	"context"

	"github.com/Hordevcom/GameShelf/internal/config"
	"github.com/Hordevcom/GameShelf/internal/middleware/logging"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserStorage struct {
	DB     *pgxpool.Pool
	Conf   config.Config
	Logger *logging.Logger
}

func NewUserStorage(Conf config.Config, Logger *logging.Logger, db *pgxpool.Pool) *UserStorage {

	return &UserStorage{DB: db, Conf: Conf, Logger: Logger}
}

func (u *UserStorage) CheckUserLogin(ctx context.Context, user string) bool {
	var username string
	query := `SELECT username FROM users WHERE username = $1`
	row := u.DB.QueryRow(ctx, query, user)
	row.Scan(&username)

	return username != ""
}

func (u *UserStorage) GetUserPassword(ctx context.Context, username string) (string, error) {
	var userPassword string

	query := `SELECT password_hash FROM users WHERE username = $1`
	err := u.DB.QueryRow(ctx, query, username).Scan(&userPassword)

	return userPassword, err
}

func (u *UserStorage) AddUserToDB(ctx context.Context, username, password string) error {
	var user string
	query := `INSERT INTO users (username, password_hash)
				VALUES ($1, $2) ON CONFLICT (username) DO NOTHING
				RETURNING username`
	err := u.DB.QueryRow(ctx, query, username, password).Scan(&user)

	return err
}
