package app

import (
	"github.com/Hordevcom/GameShelf/internal/config"
	"github.com/Hordevcom/GameShelf/internal/middleware/logging"
	"github.com/Hordevcom/GameShelf/internal/routes"
	"github.com/Hordevcom/GameShelf/internal/server"
	"github.com/Hordevcom/GameShelf/internal/services"
	"github.com/Hordevcom/GameShelf/internal/storage"
)

func Run() {

	logger := logging.NewLogger()
	config := config.NewConfig(logger)
	Storages := storage.NewStorages(config, logger)
	appServices := services.NewServices(Storages)

	routes := routes.NewRouter(logger, appServices)

	serverAPI := server.NewServer(routes, config)

	go func() {
		logger.Info("Start grpc-server")
		if err := server.StartGRPCServer(appServices, ":50051"); err != nil {
			logger.Fatalf("failed to start gRPC server: %v", err)
		}
	}()

	logger.Info("Start server")
	serverAPI.ListenAndServe()
}
