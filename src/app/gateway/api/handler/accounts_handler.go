package handler

import (
	"encoding/json"
	"net/http"

	"github.com/LeonardoMuller13/digital-bank-api/src/app/domain/dto"
	"github.com/LeonardoMuller13/digital-bank-api/src/app/gateway/api/rest"
	"github.com/LeonardoMuller13/digital-bank-api/src/app/gateway/api/rest/response"
	"github.com/go-chi/chi/v5"
)

const (
	accountsPattern = "/accounts"
)

func (h *Handler) AccountsSetup(router chi.Router) {
	router.Route(accountsPattern, func(r chi.Router) {
		r.Get("/", h.getAccounts())
		r.Post("/", h.newAccount())
		r.Get("/{id}/balance", h.getAccountBalanceByID())
	})
}

func (h *Handler) getAccounts() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		list, err := h.useCase.AccountsRepository.List(req.Context())
		var resp *response.Response
		if err != nil {
			resp = response.InternalServerError(err)
		} else {
			resp = response.OK(list)
		}
		rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers)
	}
}

func (h *Handler) newAccount() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		input := &dto.RequestNewAccount{}
		json.NewDecoder(req.Body).Decode(&input)
		err := h.useCase.AccountsRepository.Create(req.Context(), *input)
		var resp *response.Response
		if err != nil {
			resp = response.InternalServerError(err)
		} else {
			resp = response.OK("Cadastro feito com sucesso")
		}
		rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers)
	}
}

func (h *Handler) getAccountBalanceByID() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		input := chi.URLParam(req, "id")
		json.NewDecoder(req.Body).Decode(&input)
		balance, err := h.useCase.AccountsRepository.GetAccountBalanceByID(req.Context(), input)
		var resp *response.Response
		if err != nil {
			resp = response.InternalServerError(err)
		} else {
			resp = response.OK(balance)
		}
		rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers)
	}
}
