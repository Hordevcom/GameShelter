package app

import (
	"github.com/Hordevcom/GameShelf/internal/config"
	"github.com/Hordevcom/GameShelf/internal/handlers"
	"github.com/Hordevcom/GameShelf/internal/middleware/logging"
	"github.com/Hordevcom/GameShelf/internal/routes"
	"github.com/Hordevcom/GameShelf/internal/server"
	"github.com/Hordevcom/GameShelf/internal/services"
	"github.com/Hordevcom/GameShelf/internal/storage"
)

func Run() {
	logger := logging.NewLogger()
	config := config.NewConfig(logger)
	storage := storage.NewStorage(config, logger)
	service := services.NewService(*storage)
	handlers := handlers.NewHandler(*service, *logger)

	routes := routes.NewRouter(handlers, logger)

	server := server.NewServer(routes, config)

	logger.Info("Start server")
	server.ListenAndServe()
}
