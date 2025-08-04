package storage

import (
	"context"

	"github.com/Hordevcom/GameShelf/internal/config"
	"github.com/Hordevcom/GameShelf/internal/middleware/logging"
	"github.com/Hordevcom/GameShelf/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserGamesStorage struct {
	DB     *pgxpool.Pool
	Conf   config.Config
	Logger *logging.Logger
}

func NewUserGameStorage(Conf config.Config, Logger *logging.Logger, db *pgxpool.Pool) *UserGamesStorage {
	return &UserGamesStorage{DB: db, Conf: Conf, Logger: Logger}
}

func (ug *UserGamesStorage) AddGameToUserDB(ctx context.Context, usergame models.UserGame) error {
	query := `INSERT INTO user_games (username, game_title, game_status, game_store)
				VALUES ($1, $2, $3, $4)`

	_, err := ug.DB.Exec(ctx, query, usergame.Username, usergame.GameTitle,
		usergame.GameStatus, usergame.GameStore)

	return err
}

func (ug *UserGamesStorage) CheckGameInUserLibDB(ctx context.Context, gametitle, username string) (error, bool) {
	var exist bool

	query := `SELECT EXISTS(SELECT 1 FROM user_games WHERE game_title = $1 AND username = $2);`
	err := ug.DB.QueryRow(ctx, query, gametitle, username).Scan(&exist)

	if err != nil {
		return err, false
	}

	return nil, exist
}

func (ug *UserGamesStorage) GetUserGamesDB(ctx context.Context, username string) ([]models.UserGames, error) {
	var usGames []models.UserGames

	query := `SELECT game_title, game_status, game_store, updated_at FROM user_games
 WHERE username = $1`
	rows, err := ug.DB.Query(ctx, query, username)
	if err != nil {
		return []models.UserGames{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var g models.UserGames

		err := rows.Scan(&g.GameTitle, &g.GameStatus, &g.GameStore, &g.AddedAt)
		if err != nil {
			ug.Logger.Error("something went wrong with rows in DB: ", err)
			return nil, err
		}

		usGames = append(usGames, g)
	}

	return usGames, err
}

func (ug *UserGamesStorage) UpdateUserGame(ctx context.Context, userGameUpd models.UserGameUpdate, username string) error {
	query := `UPDATE user_games 
              SET game_status = $1 
              WHERE username = $2 AND game_title = $3`

	_, err := ug.DB.Exec(ctx, query, userGameUpd.GameStatus, username, userGameUpd.GameTitle)

	return err
}
