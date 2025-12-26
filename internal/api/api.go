package api

import (
	"log/slog"

	"github.com/coelhoedudev/gobit/internal/service"
	"github.com/go-chi/chi/v5"
)

type Api struct {
	Router      *chi.Mux
	UserService *service.UserService
	Logger      *slog.Logger
}
