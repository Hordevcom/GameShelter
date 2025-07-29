package storage

import (
	"context"

	"github.com/Hordevcom/GameShelf/internal/config"
	"github.com/Hordevcom/GameShelf/internal/middleware/logging"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PGDB struct {
	DB     *pgxpool.Pool
	Conf   config.Config
	Logger *logging.Logger
}

func NewStorage(Conf config.Config, Logger *logging.Logger) *PGDB {
	db, err := pgxpool.New(context.Background(), Conf.DatabaseDsn)

	if err != nil {
		Logger.Error("Problem with connection to db: ", err)
		return nil
	}

	err = db.Ping(context.Background())
	if err != nil {
		Logger.Error("Problem with ping to db: ", err)
		return nil
	}

	return &PGDB{DB: db}
}
