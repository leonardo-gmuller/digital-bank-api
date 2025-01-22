package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/leonardo-gmuller/digital-bank-api/app/domain/erring"
	"github.com/leonardo-gmuller/digital-bank-api/app/domain/usecases"
	"github.com/leonardo-gmuller/digital-bank-api/app/gateway/api/handler/schema"
	"github.com/leonardo-gmuller/digital-bank-api/app/gateway/api/rest"
	"github.com/leonardo-gmuller/digital-bank-api/app/gateway/api/rest/response"
)

const (
	loginPattern = "/auth"
)

func (h *Handler) LoginSetup(router chi.Router) {
	router.Route(loginPattern, func(r chi.Router) {
		r.Post("/", h.login())
	})
}

func (h *Handler) login() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		var resp *response.Response

		creds := &schema.AuthRequest{}

		err := json.NewDecoder(req.Body).Decode(creds)
		if err != nil {
			resp = response.InternalServerError(err)
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

			return
		}

		input := usecases.AuthInput{
			Cpf:    creds.Cpf,
			Secret: creds.Password,
		}

		output, err := h.useCase.Auth(req.Context(), input, h.cfg.JWT.TokenAuth, h.cfg.JWT.ExpiresIn)
		if err != nil {
			switch {
			case errors.Is(err, erring.ErrLoginUserNotFound):
				resp = response.NotFound(err, err.Error(), "")
				rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

				return
			case errors.Is(err, erring.ErrLoginUnauthorized):
				resp = response.Unauthorized(err.Error())
				rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

				return
			case errors.Is(err, erring.ErrLoginTokenNotCreated):
				resp = response.InternalServerError(errors.New("internal error"))
				rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

				return
			}
		}

		resp = response.OK(&schema.AuthResponse{AccessToken: output.AccessToken})
		rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck
	}
}
