package services

import "context"

func (s *Service) IsGameAlreadyExist(ctx context.Context, gamename string) (error, bool) {
	return s.db.CheckGameExists(ctx, gamename)
}
