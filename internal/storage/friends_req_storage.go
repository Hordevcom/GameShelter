package storage

import (
	"context"

	"github.com/Hordevcom/GameShelf/internal/config"
	"github.com/Hordevcom/GameShelf/internal/middleware/logging"
	"github.com/Hordevcom/GameShelf/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FriendReqStorage struct {
	DB     *pgxpool.Pool
	Conf   config.Config
	Logger *logging.Logger
}

func NewFriendReqStorage(Conf config.Config, Logger *logging.Logger, db *pgxpool.Pool) *FriendReqStorage {

	return &FriendReqStorage{DB: db, Conf: Conf, Logger: Logger}
}

func (f *FriendReqStorage) CreateFriendReqDB(ctx context.Context, payload models.FriendRequestJSON) error {
	query := `INSERT INTO friend_requests (sender, receiver, req_status)
				VALUES ($1, $2, $3)`

	_, err := f.DB.Exec(ctx, query, payload.Sender, payload.Receiver, payload.RequestStatus)

	return err
}

func (f *FriendReqStorage) UpdateFriendReqDB(ctx context.Context, payload models.FriendRequestJSON) error {
	query := `UPDATE friend_requests 
              SET req_status = $1 
              WHERE sender = $2 AND receiver = $3`

	_, err := f.DB.Exec(ctx, query, payload.RequestStatus, payload.Sender, payload.Receiver)

	return err
}

func (f *FriendReqStorage) DeleteFriendReqDB(ctx context.Context, payload models.FriendRequestJSON) error {
	query := `DELETE FROM friend_requests 
              WHERE sender = $1 AND receiver = $2`

	_, err := f.DB.Exec(ctx, query, payload.Sender, payload.Receiver)

	return err
}

func (f *FriendReqStorage) GetListOfReqDB(ctx context.Context, username string) ([]models.FriendRequestDB, error) {
	var requests []models.FriendRequestDB

	query := `SELECT sender, receiver, req_status, created_at FROM friend_requests
 WHERE sender = $1`
	rows, err := f.DB.Query(ctx, query, username)
	if err != nil {
		return []models.FriendRequestDB{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var g models.FriendRequestDB

		err := rows.Scan(&g.Sender, &g.Receiver, &g.RequestStatus, &g.CretedAt)
		if err != nil {
			f.Logger.Error("something went wrong with rows in DB: ", err)
			return nil, err
		}

		requests = append(requests, g)
	}

	return requests, err
}
