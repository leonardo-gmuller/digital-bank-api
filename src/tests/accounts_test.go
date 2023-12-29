package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/LeonardoMuller13/digital-bank-api/src/controllers"
	"github.com/LeonardoMuller13/digital-bank-api/src/database"
	"github.com/LeonardoMuller13/digital-bank-api/src/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var ID int

func NewAccountMock() {
	account := models.Account{
		Name:    "Teste",
		Cpf:     "12345678901",
		Secret:  "123456",
		Balance: 1000,
	}
	database.DB.Create(&account)
	ID = int(account.ID)
}

func DeleteAccounts() {
	database.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Account{})
}

func TestGetAccountsHandler(t *testing.T) {
	database.ConnectDB()
	NewAccountMock()

	defer DeleteAccounts()

	req, _ := http.NewRequest("GET", "/accounts", nil)
	response := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/accounts", controllers.GetAccounts).Methods("GET")
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestNewAccountHandler(t *testing.T) {
	database.ConnectDB()
	account := models.Account{
		Name:    "Teste",
		Cpf:     "12345678901",
		Balance: 1000,
	}
	account.SetPassword("123456")
	data, err := json.Marshal(account)
	if err != nil {
		t.Fatal(err)
	}

	defer DeleteAccounts()

	req, _ := http.NewRequest("POST", "/accounts", bytes.NewBuffer(data))
	response := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/accounts", controllers.NewAccount).Methods("POST")
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)

	var accountResponse models.Account
	json.NewDecoder(response.Body).Decode(&accountResponse)
}

func TestGetAccountBalanceByIDHandler(t *testing.T) {
	database.ConnectDB()
	NewAccountMock()

	defer DeleteAccounts()

	req, _ := http.NewRequest("GET", "/accounts/"+strconv.Itoa(ID)+"/balance", nil)
	response := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/accounts/{account_id}/balance", controllers.GetAccountBalanceByID).Methods("GET")
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "1000\n", response.Body.String())
}

func TestNewAccountIfCpfIsNotValidHandler(t *testing.T) {
	database.ConnectDB()
	account := models.Account{
		Name:    "Teste",
		Cpf:     "02345678905",
		Secret:  "123456",
		Balance: 1000,
	}
	data, err := json.Marshal(account)
	if err != nil {
		t.Fatal(err)
	}
	defer DeleteAccounts()

	req, _ := http.NewRequest("POST", "/accounts", bytes.NewBuffer(data))
	response := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/accounts", controllers.NewAccount).Methods("POST")
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
	assert.Equal(t, "CPF is not valid.", response.Body.String())
}
func TestNewAccountIfAccountAlreadyExistsHandler(t *testing.T) {
	database.ConnectDB()
	account := models.Account{
		Name:    "Teste",
		Cpf:     "12345678901",
		Secret:  "123456",
		Balance: 1000,
	}
	data, err := json.Marshal(account)
	if err != nil {
		t.Fatal(err)
	}
	defer DeleteAccounts()

	req, _ := http.NewRequest("POST", "/accounts", bytes.NewBuffer(data))
	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/accounts", controllers.NewAccount).Methods("POST")
	router.ServeHTTP(response, req)

	req, _ = http.NewRequest("POST", "/accounts", bytes.NewBuffer(data))
	response = httptest.NewRecorder()
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
	assert.Equal(t, "Account already exists.", response.Body.String())
}

func TestGetAccountBalanceByIDIfNotFoundHandler(t *testing.T) {
	database.ConnectDB()
	NewAccountMock()

	defer DeleteAccounts()

	req, _ := http.NewRequest("GET", "/accounts/0/balance", nil)
	response := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/accounts/{account_id}/balance", controllers.GetAccountBalanceByID).Methods("GET")
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
	assert.Equal(t, "Account not found.", response.Body.String())
}
