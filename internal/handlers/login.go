package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Hordevcom/GameShelf/internal/middleware/auth"
	"github.com/Hordevcom/GameShelf/internal/middleware/logging"
	"github.com/Hordevcom/GameShelf/internal/models"
)

type UserAuth interface {
	CheckUserPassword(ctx context.Context, user models.UserAuth) error
	CheckUserLoginService(ctx context.Context, user models.UserAuth) bool
}

func Login(userAuth UserAuth, logger logging.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.UserAuth
		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			logger.Error("wrong json: ", err)
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		err = userAuth.CheckUserPassword(r.Context(), user)

		if err != nil {
			logger.Error("error with auth: ", err)
			http.Error(w, "Wrong login or password", http.StatusUnauthorized)
			return
		}
		token, _ := auth.BuildJWTString(user.Username)
		cookie := &http.Cookie{
			Name:     "token",
			Value:    token,
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, cookie)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

	}
}
