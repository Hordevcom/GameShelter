package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Hordevcom/GameShelf/internal/models"
	"github.com/go-chi/render"
)

func (h *Handler) AddGameToUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var UserGames models.UserGameJSON

		cookie, err := r.Cookie("token")
		if err != nil {
			h.Logger.Error("problem with get cookie token: ", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		err = json.NewDecoder(r.Body).Decode(&UserGames)

		if err != nil {
			h.Logger.Error("problem with decode JSON: ", err)
			http.Error(w, "wrong JSON", http.StatusBadRequest)
			return
		}

		err, exist := h.Services.IsGameAlreadyExist(r.Context(), UserGames.GameTitle)
		if err != nil {
			h.Logger.Error("problem with finding game: ", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		if !exist {
			h.Logger.Error("try to add non exist game: ", err)
			http.Error(w, "Game not exist in library", http.StatusBadRequest)
			return
		}

		err, exist = h.Services.CheckGameInUserLib(r.Context(), UserGames.GameTitle, cookie.Value)
		if err != nil {
			h.Logger.Error("problem with checking game: ", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		if exist {
			h.Logger.Error("try to add game that already in user lib: ")
			http.Error(w, "Game already in library", http.StatusBadRequest)
			return
		}

		err = h.Services.AddGameToUser(r.Context(), UserGames, cookie.Value)
		if err != nil {
			h.Logger.Error("problem with add game to user: ", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, map[string]string{"message": "Game added to user library"})
	}
}
