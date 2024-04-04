package api

import (
	"net/http"

	"github.com/LeonardoMuller13/digital-bank-api/src/app/config"
	"github.com/LeonardoMuller13/digital-bank-api/src/app/domain/usecases"
	"github.com/LeonardoMuller13/digital-bank-api/src/app/gateway/api/handler"
	"github.com/LeonardoMuller13/digital-bank-api/src/app/gateway/api/middleware"
	"github.com/go-chi/chi/v5"
)

type API struct {
	Handler http.Handler
	cfg     config.Config
	useCase *usecases.UseCase
}

func BasicHandler() http.Handler {
	router := chi.NewMux()
	handler.RegisterHealthCheckRoute(router)

	return router
}

func New(cfg config.Config, useCase *usecases.UseCase) *API {
	api := &API{
		cfg:     cfg,
		useCase: useCase,
	}

	api.setupRouter()

	return api
}

func (api *API) setupRouter() {
	router := chi.NewRouter()

	if api.cfg.Development {
		router.Use(middleware.Logger)
	}

	router.Use(
		middleware.CORS,
	)

	api.registerRoutes(router)

	api.Handler = router
}
func (api *API) registerRoutes(router *chi.Mux) {
	handler.RegisterHealthCheckRoute(router)

	router.Route("/api/v2/digital-bank", func(publicRouter chi.Router) {
		handler.RegisterPublicRoutes(
			publicRouter,
			api.cfg,
			api.useCase,
		)
	})
}
