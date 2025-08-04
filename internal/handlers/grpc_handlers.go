package handlers

import (
	"context"

	"github.com/Hordevcom/GameShelf/internal/models"
	"github.com/Hordevcom/GameShelf/internal/services"
	pb "github.com/Hordevcom/GameShelf/pkg/pb"
)

type GameShelterGRPCServer struct {
	pb.UnimplementedGameShelterServiceServer
	Services *services.Services
}

func (s *GameShelterGRPCServer) AddNewGame(ctx context.Context, req *pb.NewGameRequest) (*pb.NewGameResponce, error) {
	err := s.Services.GameAdder.AddNewGame(ctx, models.Game{
		Title: req.Title,
		Genre: req.Genre,
	})

	if err != nil {
		return nil, err
	}

	return &pb.NewGameResponce{
		Message: "Game added to server library",
	}, nil
}

func (s *GameShelterGRPCServer) GetUserGames(ctx context.Context, req *pb.UserGamesRequest) (*pb.UserGamesResponce, error) {
	userGames, err := s.Services.UserGamesFetcher.GetUserGames(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	var games []*pb.UserGame
	for _, g := range userGames {
		games = append(games, &pb.UserGame{
			GameTitle:  g.GameTitle,
			GameStatus: g.GameStatus,
			GameStore:  g.GameStore,
		})
	}

	return &pb.UserGamesResponce{Games: games}, nil
}

func (s *GameShelterGRPCServer) CreateFriendRequest(ctx context.Context, req *pb.FriendReqRequest) (*pb.FriendReqResponce, error) {
	err := s.Services.FriendRequest.CreateFriendReqService(ctx, models.FriendRequestJSON{
		Sender:        req.Sender,
		Receiver:      req.Receiver,
		RequestStatus: req.Status,
	}, req.Token)

	if err != nil {
		return nil, err
	}

	return &pb.FriendReqResponce{Message: "Request has been created!"}, nil
}
