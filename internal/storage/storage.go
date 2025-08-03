package storage

import (
	"context"

	"github.com/Hordevcom/GameShelf/internal/config"
	"github.com/Hordevcom/GameShelf/internal/middleware/logging"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storages struct {
	UserStorage
	GameStorage
	UserGamesStorage
	FriendReqStorage
	FriendsStorage
}

func NewConnectionToDB(Conf config.Config, Logger *logging.Logger) (*pgxpool.Pool, error) {
	db, err := pgxpool.New(context.Background(), Conf.DatabaseDsn)

	if err != nil {
		Logger.Error("Problem with connection to db: ", err)
		return nil, err
	}

	err = db.Ping(context.Background())
	if err != nil {
		Logger.Error("Problem with ping to db: ", err)
		return nil, err
	}

	return db, nil
}

func NewStorages(Conf config.Config, Logger *logging.Logger) *Storages {
	db, err := NewConnectionToDB(Conf, Logger)
	if err != nil {
		Logger.Error("Problem with connection to db: ", err)
		return nil
	}
	UserStorage := NewUserStorage(Conf, Logger, db)
	GameStorage := NewGameStorage(Conf, Logger, db)
	UserGamesStorage := NewUserGameStorage(Conf, Logger, db)
	FriendReqStorage := NewFriendReqStorage(Conf, Logger, db)
	FriendsStorage := NewFriendsStorage(Conf, Logger, db)

	return &Storages{
		UserStorage:      *UserStorage,
		GameStorage:      *GameStorage,
		UserGamesStorage: *UserGamesStorage,
		FriendReqStorage: *FriendReqStorage,
		FriendsStorage:   *FriendsStorage,
	}
}
