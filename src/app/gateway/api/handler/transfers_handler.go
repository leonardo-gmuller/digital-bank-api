package handler

import (
	"encoding/json"
	"net/http"

	"github.com/LeonardoMuller13/digital-bank-api/src/app/domain/dto"
	"github.com/LeonardoMuller13/digital-bank-api/src/app/gateway/api/middleware"
	"github.com/LeonardoMuller13/digital-bank-api/src/app/gateway/api/rest"
	"github.com/LeonardoMuller13/digital-bank-api/src/app/gateway/api/rest/response"
	"github.com/go-chi/chi/v5"
)

const (
	transfersPattern = "/transfers"
)

func (h *Handler) TransfersSetup(router chi.Router) {
	router = router.With(middleware.ProtectedHandler(h.cfg.JwtSecretKey))
	router.Route(transfersPattern, func(r chi.Router) {
		r.Get("/", h.getTransfers())
		r.Post("/", h.newTransfer())
	})
}

func (h *Handler) getTransfers() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		list, err := h.useCase.TransfersRepository.List(req.Context())
		var resp *response.Response
		if err != nil {
			resp = response.InternalServerError(err)
		} else {
			resp = response.OK(list)
		}
		rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers)
	}
}

func (h *Handler) newTransfer() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		input := &dto.RequestTransfer{}
		json.NewDecoder(req.Body).Decode(&input)
		err := h.useCase.TransfersRepository.Create(req.Context(), *input)
		var resp *response.Response
		if err != nil {
			resp = response.InternalServerError(err)
		} else {
			resp = response.OK("Transfer successfull")
		}
		rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers)
	}
}
