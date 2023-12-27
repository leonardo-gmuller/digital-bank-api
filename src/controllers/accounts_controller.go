package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LeonardoMuller13/digital-bank-api/src/database"
	"github.com/LeonardoMuller13/digital-bank-api/src/helpers"
	"github.com/LeonardoMuller13/digital-bank-api/src/models"
	"github.com/gorilla/mux"
)

type RequestNewAccount struct {
	Name   string
	Cpf    string
	Secret string
}

type ResponseAccount struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Cpf  string `json:"cpf"`
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
	if !helpers.CPFIsValid(creds.Cpf) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "CPF is not valid.")
		return
	}
	result := database.DB.First(&account, "cpf = ?", creds.Cpf)
	if result.Error == nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Account already exists.")
		return
	}
	account.Name = creds.Name
	account.Cpf = creds.Cpf
	account.Balance = 1000 //For tests, init account with 1000
	account.SetPassword(creds.Secret)
	database.DB.Create(&account)
	json.NewEncoder(w).Encode(account)
}

func GetAccountBalanceByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["account_id"]
	var a models.Account
	result := database.DB.First(&a, id)
	if result.Error == nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Account not found.")
		return
	}
	json.NewEncoder(w).Encode(a.Balance)
}
