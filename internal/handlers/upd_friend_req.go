package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Hordevcom/GameShelf/internal/middleware/logging"
	"github.com/Hordevcom/GameShelf/internal/models"
	"github.com/go-chi/render"
)

type FriendReqUpdater interface {
	UpdateFriendReqService(ctx context.Context, payload models.FriendRequestJSON, token string) error
}

func UpdateFriendReqHandler(logger logging.Logger, fru FriendReqUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var UserFrendsReq models.FriendRequestJSON

		cookie, err := r.Cookie("token")
		if err != nil {
			logger.Error("problem with get cookie token: ", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		err = json.NewDecoder(r.Body).Decode(&UserFrendsReq)

		if err != nil {
			logger.Error("problem with decode JSON: ", err)
			http.Error(w, "wrong JSON", http.StatusBadRequest)
			return
		}

		err = fru.UpdateFriendReqService(r.Context(), UserFrendsReq, cookie.Value)
		if err != nil {
			logger.Error("something went wrong request creation: ", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, map[string]string{"message": "Request was updated!"})
	}
}
