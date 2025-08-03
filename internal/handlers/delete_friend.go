package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Hordevcom/GameShelf/internal/middleware/logging"
	"github.com/Hordevcom/GameShelf/internal/models"
	"github.com/go-chi/render"
)

type FriendRemover interface {
	DeleteFriendForUserService(ctx context.Context, token, friend string) error
}

func DeleteFriendHandler(logger logging.Logger, fr FriendRemover) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var Friend models.Friend

		cookie, err := r.Cookie("token")
		if err != nil {
			logger.Error("problem with get cookie token: ", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		err = json.NewDecoder(r.Body).Decode(&Friend)

		if err != nil {
			logger.Error("problem with decode JSON: ", err)
			http.Error(w, "wrong JSON", http.StatusBadRequest)
			return
		}

		err = fr.DeleteFriendForUserService(r.Context(), cookie.Value, Friend.Username)
		if err != nil {
			logger.Error("something went wrong delete friend: ", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, map[string]string{"message": "Friend was deleted!"})
	}
}
