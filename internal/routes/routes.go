package routes

import (
	"github.com/Hordevcom/GameShelf/internal/handlers"
	"github.com/Hordevcom/GameShelf/internal/middleware/auth"
	"github.com/Hordevcom/GameShelf/internal/middleware/logging"
	"github.com/Hordevcom/GameShelf/internal/services"
	"github.com/Hordevcom/GameShelf/internal/storage"
	"github.com/go-chi/chi/v5"
)

func NewRouter(log *logging.Logger, storages *storage.Storages) *chi.Mux {
	router := chi.NewRouter()

	appServices := services.NewServices(storages)

	router.Use(log.WithLogging)

	router.Post("/api/user/register", handlers.UserRegister(&appServices.AuthService, *log, &appServices.UserAdder))
	router.Post("/api/user/login", handlers.Login(&appServices.AuthService, *log))

	router.With(auth.AuthMiddleware).
		Post("/api/games", handlers.AddNewGame(*log, &appServices.GameAdder))

	router.With(auth.AuthMiddleware).
		Post("/api/users/games", handlers.AddGameToUser(*log, &appServices.GameAdder, &appServices.UserGameAdder))

	router.With(auth.AuthMiddleware).
		Patch("/api/games/status", handlers.UpdateGameStatus(*log, &appServices.UserGameUpdater))

	router.With(auth.AuthMiddleware).
		Get("/api/users/{username}/games", handlers.GetUserGames(*log, &appServices.UserGamesFetcher))

	return router
}
