package app

import (
	"github.com/Hordevcom/GameShelf/internal/config"
	"github.com/Hordevcom/GameShelf/internal/middleware/logging"
	"github.com/Hordevcom/GameShelf/internal/routes"
	"github.com/Hordevcom/GameShelf/internal/server"
	"github.com/Hordevcom/GameShelf/internal/storage"
)

func Run() {

	logger := logging.NewLogger()
	config := config.NewConfig(logger)
	Storages := storage.NewStorages(config, logger)

	routes := routes.NewRouter(logger, Storages)

	server := server.NewServer(routes, config)

	logger.Info("Start server")
	server.ListenAndServe()
}
