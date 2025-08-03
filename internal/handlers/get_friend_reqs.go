package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Hordevcom/GameShelf/internal/middleware/logging"
	"github.com/Hordevcom/GameShelf/internal/models"
)

type FriendReqGetter interface {
	GetListOfFriendReqService(ctx context.Context, token string) ([]models.FriendRequestDB, error)
}

func GetListOfFriendReqHandler(logger logging.Logger, frg FriendReqGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			logger.Error("problem with get cookie token: ", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		friends, err := frg.GetListOfFriendReqService(r.Context(), cookie.Value)

		if err != nil {
			logger.Error("something went wrong with getting game: ", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(friends)
		if err != nil {
			logger.Error("something went wrong with encode json: ", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}
	}
}
