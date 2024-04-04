package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/LeonardoMuller13/digital-bank-api/src/app/gateway/api/rest"
	"github.com/LeonardoMuller13/digital-bank-api/src/app/gateway/api/rest/response"
	"github.com/bitly/go-simplejson"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Password string `json:"password"`
	Cpf      string `json:"cpf"`
}

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
		creds := &Credentials{}
		err := json.NewDecoder(req.Body).Decode(creds)
		var resp *response.Response
		if err != nil {
			resp = response.InternalServerError(err)
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers)
			return
		}

		account, err := h.useCase.AccountsRepository.GetAccountByCpf(req.Context(), creds.Cpf)
		if err != nil {
			resp = response.InternalServerError(fmt.Errorf("username not found"))
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers)
			return
		}
		fmt.Println(account.Secret)
		fmt.Println(creds.Password)
		if err = bcrypt.CompareHashAndPassword([]byte(account.Secret), []byte(creds.Password)); err != nil {
			resp = response.Unauthorized()
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers)
			return
		}
		tokenString, err := createToken(account.ID, h.cfg.JwtSecretKey)
		if err != nil {
			resp = response.InternalServerError(fmt.Errorf("internal error"))
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers)
			return
		}

		payload := simplejson.New()
		payload.Set("token", tokenString)
		resp = response.OK(payload)
		rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers)
	}
}

func createToken(user uint, jwtSecret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user": user,
			"exp":  time.Now().Add(time.Hour * 24).Unix(),
		})

	jwtKey := []byte(jwtSecret)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
