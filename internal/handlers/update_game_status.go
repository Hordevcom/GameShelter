package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Hordevcom/GameShelf/internal/middleware/logging"
	"github.com/Hordevcom/GameShelf/internal/models"
)

type UserGameUpdater interface {
	UpdateGame(ctx context.Context, gameUpd models.UserGameUpdate, token string) error
}

func UpdateGameStatus(logger logging.Logger, ugu UserGameUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var UserGameUpd models.UserGameUpdate

		cookie, err := r.Cookie("token")
		if err != nil {
			logger.Error("problem with get cookie token: ", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		err = json.NewDecoder(r.Body).Decode(&UserGameUpd)

		if err != nil {
			logger.Error("problem with decode JSON: ", err)
			http.Error(w, "wrong JSON", http.StatusBadRequest)
			return
		}

		err = ugu.UpdateGame(r.Context(), UserGameUpd, cookie.Value)
		if err != nil {
			logger.Error("failed to update game status: ", err)
			http.Error(w, "failed to update game status", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
