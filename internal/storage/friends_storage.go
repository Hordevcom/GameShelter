package storage

import (
	"context"

	"github.com/Hordevcom/GameShelf/internal/config"
	"github.com/Hordevcom/GameShelf/internal/middleware/logging"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FriendsStorage struct {
	DB     *pgxpool.Pool
	Conf   config.Config
	Logger *logging.Logger
}

func NewFriendsStorage(Conf config.Config, Logger *logging.Logger, db *pgxpool.Pool) *FriendsStorage {

	return &FriendsStorage{DB: db, Conf: Conf, Logger: Logger}
}

func (f *FriendsStorage) GetFriendsForUserDB(ctx context.Context, username string) ([]string, error) {
	var friends []string

	query := `SELECT friend FROM friends
 WHERE username = $1`
	rows, err := f.DB.Query(ctx, query, username)
	if err != nil {
		return []string{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var g string

		err := rows.Scan(&g)
		if err != nil {
			f.Logger.Error("something went wrong with rows in DB: ", err)
			return nil, err
		}

		friends = append(friends, g)
	}

	return friends, err
}

func (f *FriendsStorage) AddFriendForUserDB(ctx context.Context, username, friend string) error {
	query := `INSERT INTO friends (username, friend)
				VALUES ($1, $2)`

	_, err := f.DB.Exec(ctx, query, username, friend)

	return err
}

func (f *FriendsStorage) DeleteFriendForUserDB(ctx context.Context, user, friend string) error {
	query := `DELETE FROM friends 
              WHERE username = $1 AND friend = $2`

	_, err := f.DB.Exec(ctx, query, user, friend)

	return err
}
