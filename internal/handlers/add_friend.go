package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Hordevcom/GameShelf/internal/middleware/logging"
	"github.com/Hordevcom/GameShelf/internal/models"
	"github.com/go-chi/render"
)

type FriedsAdder interface {
	AddFriendForUserService(ctx context.Context, token, friend string) error
}

func AddFriendHandler(logger logging.Logger, fa FriedsAdder) http.HandlerFunc {
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

		err = fa.AddFriendForUserService(r.Context(), cookie.Value, Friend.Username)
		if err != nil {
			logger.Error("something went wrong request creation: ", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		render.JSON(w, r, map[string]string{"message": "Friend was added!"})
	}
}
