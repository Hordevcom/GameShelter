package main

import (
	"database/sql"

	"github.com/Hordevcom/GameShelf/internal/middleware/logging"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func main() {
	loger := logging.NewLogger()
	db, err := sql.Open("pgx", "postgres://postgres:1@localhost:5432/postgres")

	if err != nil {
		loger.Error("Failed to open DB: ", err)
		return
	}
	defer db.Close()

	if err := goose.Up(db, "../migrations"); err != nil {
		loger.Error("failed to apply migrations: ", err)
		return
	}

	loger.Info("Migrations runs successful")
}
