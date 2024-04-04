package handler

import (
	"net/http"

	"github.com/LeonardoMuller13/digital-bank-api/src/app/config"
	"github.com/LeonardoMuller13/digital-bank-api/src/app/domain/usecases"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	cfg     config.Config
	useCase *usecases.UseCase
}

func New(cfg config.Config, useCase *usecases.UseCase) Handler {
	return Handler{
		cfg:     cfg,
		useCase: useCase,
	}
}

func RegisterHealthCheckRoute(router chi.Router) {
	router.Get("/healthcheck", func(rw http.ResponseWriter, _ *http.Request) {
		rw.WriteHeader(http.StatusOK)
	})
}

func RegisterPublicRoutes(
	router chi.Router,
	cfg config.Config,
	useCase *usecases.UseCase,
) {
	handler := New(cfg, useCase)
	handler.LoginSetup(router)
	handler.AccountsSetup(router)
	handler.TransfersSetup(router)
}
