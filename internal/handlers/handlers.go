package handlers

import (
	"github.com/Hordevcom/GameShelf/internal/middleware/logging"
	"github.com/Hordevcom/GameShelf/internal/services"
)

type Handler struct {
	Services *services.Service
	Logger   *logging.Logger
}

func NewHandler(service services.Service, logger logging.Logger) *Handler {
	return &Handler{Services: &service, Logger: &logger}
}
