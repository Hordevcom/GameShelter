package server

import (
	"fmt"
	"log"
	"net"

	"github.com/Hordevcom/GameShelf/internal/handlers"
	"github.com/Hordevcom/GameShelf/internal/services"
	pb "github.com/Hordevcom/GameShelf/pkg/pb"
	"google.golang.org/grpc"
)

func StartGRPCServer(services *services.Services, port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	server := grpc.NewServer()
	pb.RegisterGameShelterServiceServer(server, &handlers.GameShelterGRPCServer{Services: services})

	log.Printf("gRPC server listening on port %d", port)
	return server.Serve(lis)
}
