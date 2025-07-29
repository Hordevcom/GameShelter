package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Hordevcom/GameShelf/internal/models"
	"github.com/go-chi/render"
)

func (h *Handler) AddNewGame() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var game models.Game
		err := json.NewDecoder(r.Body).Decode(&game)
		if err != nil {
			h.Logger.Error("wrong json: ", err)
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		err, exist := h.Services.IsGameAlreadyExist(r.Context(), game.Title)

		if err != nil {
			h.Logger.Error("something went wrong with check game exist: ", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		if exist {
			http.Error(w, "Game already added", http.StatusBadRequest)
			return
		}

		err = h.Services.AddNewGame(r.Context(), game)

		if err != nil {
			h.Logger.Error("something went wrong with check game exist: ", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		render.JSON(w, r, map[string]string{"message": "Game added to server library"})

	}
}
