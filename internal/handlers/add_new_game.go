package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Hordevcom/GameShelf/internal/middleware/logging"
	"github.com/Hordevcom/GameShelf/internal/models"
	"github.com/go-chi/render"
)

type GameStorageAdder interface {
	IsGameAlreadyExist(ctx context.Context, gamename string) (error, bool)
	AddNewGame(ctx context.Context, game models.Game) error
}

func AddNewGame(logger logging.Logger, gsa GameStorageAdder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var game models.Game
		err := json.NewDecoder(r.Body).Decode(&game)
		if err != nil {
			logger.Error("wrong json: ", err)
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		err, exist := gsa.IsGameAlreadyExist(r.Context(), game.Title)

		if err != nil {
			logger.Error("something went wrong with check game exist: ", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		if exist {
			http.Error(w, "Game already added", http.StatusBadRequest)
			return
		}

		err = gsa.AddNewGame(r.Context(), game)

		if err != nil {
			logger.Error("something went wrong with check game exist: ", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		render.JSON(w, r, map[string]string{"message": "Game added to server library"})

	}
}
