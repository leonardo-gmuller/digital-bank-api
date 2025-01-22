package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/LeonardoMuller13/digital-bank-api/app/config"
	"github.com/LeonardoMuller13/digital-bank-api/app/domain/usecases"
	"github.com/LeonardoMuller13/digital-bank-api/app/gateway/api/handler"
	"github.com/LeonardoMuller13/digital-bank-api/app/gateway/api/middleware"
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

	api.registerRoutes(router)

	api.Handler = router
}

func (api *API) registerRoutes(router *chi.Mux) {
	handler.RegisterHealthCheckRoute(router)

	router.Route("/api/v2", func(r chi.Router) {
		r.Use(middleware.Logger, middleware.RealIP, middleware.ContentTypeJSON)
		handler.RegisterRoutes(
			r,
			api.cfg,
			api.useCase,
		)
	})
}
