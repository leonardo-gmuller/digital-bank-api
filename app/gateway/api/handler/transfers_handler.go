package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"

	"github.com/LeonardoMuller13/digital-bank-api/app/domain/erring"
	"github.com/LeonardoMuller13/digital-bank-api/app/domain/usecases"
	"github.com/LeonardoMuller13/digital-bank-api/app/gateway/api/handler/schema"
	"github.com/LeonardoMuller13/digital-bank-api/app/gateway/api/middleware"
	"github.com/LeonardoMuller13/digital-bank-api/app/gateway/api/rest"
	"github.com/LeonardoMuller13/digital-bank-api/app/gateway/api/rest/response"
)

const (
	transfersPattern = "/transfers"
)

func (h *Handler) TransfersSetup(router chi.Router, tokenAuth *jwtauth.JWTAuth) {
	router.Route(transfersPattern, func(router chi.Router) {
		router.Use(
			jwtauth.Verifier(tokenAuth),
			jwtauth.Authenticator,
			middleware.AuthMiddleware,
		)
		router.Get("/", h.getTransfers())
		router.Post("/", h.newTransfer())
	})
}

func (h *Handler) getTransfers() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		var resp *response.Response

		list, err := h.useCase.ListUserTransfer(req.Context())
		if err != nil {
			resp = response.InternalServerError(err)
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

			return
		}

		resp = response.OK(list.Transfers)
		rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck
	}
}

func (h *Handler) newTransfer() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		var resp *response.Response

		input := &schema.RequestCreateTransfer{}

		err := json.NewDecoder(req.Body).Decode(&input)
		if err != nil {
			resp = response.BadRequest(err, err.Error()) //nolint:errcheck
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers)

			return
		}

		err = h.useCase.CreateTransfer(req.Context(), usecases.InputCreateTransfer{
			AccountDestinationCPF: input.AccountDestinationCPF,
			Amount:                input.Amount,
		})
		if err != nil {
			switch {
			case errors.Is(erring.ErrTransferAccountDestinationNotFound, err):
				resp = response.BadRequest(err, err.Error())
			case errors.Is(erring.ErrTransferUserNotFound, err):
				resp = response.BadRequest(err, err.Error())
			case errors.Is(erring.ErrTransferBalanceNotSufficient, err):
				resp = response.BadRequest(err, err.Error())
			default:
				resp = response.InternalServerError(err)
			}

			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

			return
		}

		resp = response.OK("Transfer successful")
		rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck
	}
}
