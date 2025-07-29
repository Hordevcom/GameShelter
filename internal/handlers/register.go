package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Hordevcom/GameShelf/internal/middleware/auth"
	"github.com/Hordevcom/GameShelf/internal/models"
	"github.com/go-chi/render"
)

func (h *Handler) UserRegister() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.UserAuth
		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			h.Logger.Error("problem with decode JSON: ", err)
			http.Error(w, "wrong JSON", http.StatusBadRequest)
			return
		}

		loginExist := h.Services.CheckUserLogin(r.Context(), user)

		if loginExist {
			h.Logger.Error("problem with decode JSON: ", err)
			http.Error(w, "login already exist, try another", http.StatusBadRequest)
			return
		}

		err = h.Services.AddUserToDB(r.Context(), user)
		if err != nil {
			h.Logger.Error("problem with add user to db: ", err)
			http.Error(w, "server error", http.StatusInternalServerError)
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
		render.JSON(w, r, map[string]string{"message": "User registration complete!"})

	}
}
