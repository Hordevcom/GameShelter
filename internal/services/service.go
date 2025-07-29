package services

import "github.com/Hordevcom/GameShelf/internal/storage"

type Service struct {
	db storage.PGDB
}

func NewService(db storage.PGDB) *Service {
	return &Service{db: db}
}
