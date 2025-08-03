package services

import "github.com/Hordevcom/GameShelf/internal/storage"

type Services struct {
	AuthService      AuthService
	UserAdder        UserAdd
	GameAdder        GameServerService
	UserGameAdder    UserAddGame
	UserGamesFetcher UserGamesGet
	UserGameUpdater  UserGameDbUpd
	FriendRequest    FriendRequest
	Friends          Friends
}

func NewServices(storages *storage.Storages) *Services {
	return &Services{
		AuthService:      AuthService{UserChecker: &storages.UserStorage},
		UserAdder:        UserAdd{UserAdder: &storages.UserStorage},
		GameAdder:        GameServerService{GameServerMaker: &storages.GameStorage},
		UserGameAdder:    UserAddGame{UserGameAdderDBAdder: &storages.UserGamesStorage},
		UserGamesFetcher: UserGamesGet{UserGameDBGetter: &storages.UserGamesStorage},
		UserGameUpdater:  UserGameDbUpd{UserGameDBUpdater: &storages.UserGamesStorage},
		FriendRequest: FriendRequest{
			FriendReqDBAdder:   &storages.FriendReqStorage,
			FriendReqDBUpdater: &storages.FriendReqStorage,
			FriendReqDBDeleter: &storages.FriendReqStorage,
			FrindReqDBGetter:   &storages.FriendReqStorage},
		Friends: Friends{
			FriendsServiceGetter:  &storages.FriendsStorage,
			FriendsServiceAdder:   &storages.FriendsStorage,
			FriendsServiceRemover: &storages.FriendsStorage,
		},
	}
}
