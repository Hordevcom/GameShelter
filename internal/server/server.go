package server

import (
	"net/http"

	"github.com/Hordevcom/GameShelf/internal/config"
	"github.com/go-chi/chi/v5"
)

func NewServer(router *chi.Mux, config config.Config) http.Server {
	return http.Server{
		Addr:    config.ServerAdress,
		Handler: router,
	}
}
