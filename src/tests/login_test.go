package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LeonardoMuller13/digital-bank-api/src/controllers"
	"github.com/LeonardoMuller13/digital-bank-api/src/database"
	"github.com/LeonardoMuller13/digital-bank-api/src/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type LoginRequest struct {
	Cpf      string
	Password string
}

type LoginResponse struct {
	Token string
}

func TestLogin(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)
	account := models.Account{
		Name:    "Teste",
		Cpf:     "12345678901",
		Balance: 1000,
	}
	account.SetPassword("123456")
	database.DB.Create(&account)

	login := LoginRequest{
		Cpf:      account.Cpf,
		Password: "123456",
	}
	data, err := json.Marshal(login)
	if err != nil {
		t.Fatal(err)
	}

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(data))
	response := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}
