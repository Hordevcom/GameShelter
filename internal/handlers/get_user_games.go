package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Hordevcom/GameShelf/internal/middleware/logging"
	"github.com/Hordevcom/GameShelf/internal/models"
	"github.com/go-chi/chi/v5"
)

type UserGameGetter interface {
	GetUserGames(ctx context.Context, username string) ([]models.UserGames, error)
}

func GetUserGames(logger logging.Logger, ugg UserGameGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")

		games, err := ugg.GetUserGames(r.Context(), username)

		if err != nil {
			logger.Error("something went wrong with getting game: ", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(games)
		if err != nil {
			logger.Error("something went wrong with encode json: ", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}
	}
}
