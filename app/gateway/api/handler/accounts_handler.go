package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/leonardo-gmuller/digital-bank-api/app/domain/erring"
	"github.com/leonardo-gmuller/digital-bank-api/app/domain/usecases"
	"github.com/leonardo-gmuller/digital-bank-api/app/gateway/api/handler/schema"
	"github.com/leonardo-gmuller/digital-bank-api/app/gateway/api/rest"
	"github.com/leonardo-gmuller/digital-bank-api/app/gateway/api/rest/response"
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
		var resp *response.Response

		list, err := h.useCase.ListAccount(req.Context())
		if err != nil {
			resp = response.InternalServerError(err)
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

			return
		}

		resp = response.OK(list.Accounts)
		rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck
	}
}

func (h *Handler) newAccount() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		var resp *response.Response

		input := &schema.RequestNewAccount{}

		err := json.NewDecoder(req.Body).Decode(&input)
		if err != nil {
			resp = response.BadRequest(err, err.Error()) //nolint:errcheck
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers)

			return
		}

		err = h.useCase.CreateAccount(req.Context(), usecases.InputCreateAccount{
			Name:    input.Name,
			Cpf:     input.Cpf,
			Secret:  input.Secret,
			Balance: input.Balance,
		})

		if err != nil {
			switch {
			case errors.Is(erring.ErrAccountCPFIsInvalid, err):
				resp = response.BadRequest(err, err.Error())
			case errors.Is(erring.ErrAccountExists, err):
				resp = response.BadRequest(err, err.Error())
			default:
				resp = response.InternalServerError(err)
			}
		} else {
			resp = response.Created("account created")
		}

		rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck
	}
}

func (h *Handler) getAccountBalanceByID() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		var resp *response.Response

		input := chi.URLParam(req, "id")

		id, err := strconv.Atoi(input)
		if err != nil {
			resp = response.InternalServerError(err)
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

			return
		}

		balance, err := h.useCase.GetBalanceByID(req.Context(), id)
		if err != nil {
			if errors.Is(erring.ErrAccountNotExists, err) {
				resp = response.BadRequest(err, err.Error())
			} else {
				resp = response.InternalServerError(err)
			}
		} else {
			resp = response.OK(balance)
		}

		rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck
	}
}
