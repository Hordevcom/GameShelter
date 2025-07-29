package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) GetUserGames() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")

		games, err := h.Services.GetUserGames(r.Context(), username)

		if err != nil {
			h.Logger.Error("something went wrong with getting game: ", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(games)
		if err != nil {
			h.Logger.Error("something went wrong with encode json: ", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}
	}
}
