package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Hordevcom/GameShelf/internal/middleware/auth"
	"github.com/Hordevcom/GameShelf/internal/models"
)

func (h *Handler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.UserAuth
		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			h.Logger.Error("wrong json: ", err)
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		err = h.Services.CheckUserPassword(r.Context(), user)

		if err != nil {
			h.Logger.Error("error with auth: ", err)
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
