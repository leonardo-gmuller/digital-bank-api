package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LeonardoMuller13/digital-bank-api/src/controllers"
	"github.com/LeonardoMuller13/digital-bank-api/src/database"
	"github.com/LeonardoMuller13/digital-bank-api/src/middleware"
	"github.com/LeonardoMuller13/digital-bank-api/src/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var passwordDefault = "123456"

type RequestTransfer struct {
	Account_Destination_CPF string
	Amount                  int
}

type ResponseNewTransfer struct {
	Message string
}

func login(account models.Account) (string, error) {
	login := LoginRequest{
		Cpf:      account.Cpf,
		Password: passwordDefault,
	}
	data, err := json.Marshal(login)
	if err != nil {
		return "", err
	}

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(data))
	response := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.ServeHTTP(response, req)

	var loginResponse LoginResponse
	json.NewDecoder(response.Body).Decode(&loginResponse)

	return loginResponse.Token, nil
}

func TestGetTransfersHandler(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)
	account := models.Account{
		Name:    "Teste",
		Cpf:     "12345678901",
		Balance: 1000,
	}
	account.SetPassword(passwordDefault)
	database.DB.Create(&account)

	token, err := login(account)
	if err != nil {
		t.Fatal(err)
	}

	req, _ := http.NewRequest("GET", "/transfers", nil)
	req.Header.Add("Authorization", "Bearer "+token)
	response := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/transfers", controllers.GetTransfers).Methods("GET")
	router.Use(middleware.ProtectedHandler)
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func newTransfer() (httptest.ResponseRecorder, error) {
	accountOrigin := models.Account{
		Name:    "Teste 1",
		Cpf:     "12345678901",
		Balance: 1000,
	}
	accountOrigin.SetPassword("123456")
	database.DB.Create(&accountOrigin)

	accountDest := models.Account{
		Name:    "Teste 2",
		Cpf:     "12345678902",
		Balance: 1000,
	}
	accountDest.SetPassword("123456")
	database.DB.Create(&accountDest)

	defer DeleteAccounts()

	token, err := login(accountOrigin)
	if err != nil {
		return httptest.ResponseRecorder{}, err
	}

	newTransfer := RequestTransfer{
		Account_Destination_CPF: accountDest.Cpf,
		Amount:                  50,
	}
	data, err := json.Marshal(newTransfer)
	if err != nil {
		return httptest.ResponseRecorder{}, err
	}

	req, _ := http.NewRequest("POST", "/transfers", bytes.NewBuffer(data))
	req.Header.Add("Authorization", "Bearer "+token)
	response := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/transfers", controllers.NewTransfer).Methods("POST")
	router.Use(middleware.ProtectedHandler)
	router.ServeHTTP(response, req)

	return *response, nil
}

func TestNewTransfersHandler(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)
	response, err := newTransfer()
	if err != nil {
		t.Fatal(err)
	}

	var newTransferResponse ResponseNewTransfer
	json.NewDecoder(response.Body).Decode(&newTransferResponse)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Transfer completed successfully!", newTransferResponse.Message)
}
