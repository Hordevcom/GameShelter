package routes

import (
	"github.com/Hordevcom/GameShelf/internal/handlers"
	"github.com/Hordevcom/GameShelf/internal/middleware/auth"
	"github.com/Hordevcom/GameShelf/internal/middleware/logging"
	"github.com/go-chi/chi/v5"
)

func NewRouter(h *handlers.Handler, log *logging.Logger) *chi.Mux {
	router := chi.NewRouter()

	router.Use(log.WithLogging)

	router.Post("/api/user/register", h.UserRegister())
	router.Post("/api/user/login", h.Login())
	router.With(auth.AuthMiddleware).Post("/api/games", h.AddNewGame())
	router.With(auth.AuthMiddleware).Post("/api/users/games", h.AddGameToUser())
	router.With(auth.AuthMiddleware).Patch("/api/games/status", h.UpdateGameStatus())
	router.With(auth.AuthMiddleware).Get("/api/users/{username}/games", h.GetUserGames())

	return router
}
