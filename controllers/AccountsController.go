package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/LeonardoMuller13/digital-bank-api/database"
	"github.com/LeonardoMuller13/digital-bank-api/models"
	"github.com/gorilla/mux"
)

type RequestNewAccount struct {
	Name   string
	Cpf    string
	Secret string
}

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	var ac []models.Account
	database.DB.Find(&ac)
	json.NewEncoder(w).Encode(ac)
}

func NewAccount(w http.ResponseWriter, r *http.Request) {
	var account models.Account
	var creds RequestNewAccount
	json.NewDecoder(r.Body).Decode(&creds)

	account.Name = creds.Name
	account.Cpf = creds.Cpf
	account.SetPassword(creds.Secret)
	database.DB.Create(&account)
	json.NewEncoder(w).Encode(account)
}

func GetAccountBalanceByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["account_id"]
	var a models.Account
	database.DB.First(&a, id)
	json.NewEncoder(w).Encode(a.GetBalance())
}
