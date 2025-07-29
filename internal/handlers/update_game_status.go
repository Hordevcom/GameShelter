package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Hordevcom/GameShelf/internal/models"
)

func (h *Handler) UpdateGameStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var UserGameUpd models.UserGameUpdate

		cookie, err := r.Cookie("token")
		if err != nil {
			h.Logger.Error("problem with get cookie token: ", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		err = json.NewDecoder(r.Body).Decode(&UserGameUpd)

		if err != nil {
			h.Logger.Error("problem with decode JSON: ", err)
			http.Error(w, "wrong JSON", http.StatusBadRequest)
			return
		}

		err = h.Services.UpdateGame(r.Context(), UserGameUpd, cookie.Value)
		if err != nil {
			h.Logger.Error("failed to update game status: ", err)
			http.Error(w, "failed to update game status", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
