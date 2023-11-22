package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/LeonardoMuller13/digital-bank-api/database"
	"github.com/LeonardoMuller13/digital-bank-api/models"
	"golang.org/x/crypto/bcrypt"
)

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	var ac []models.Account
	database.DB.Find(&ac)
	json.NewEncoder(w).Encode(ac)
}

func NewAccount(w http.ResponseWriter, r *http.Request) {
	var account models.Account
	json.NewDecoder(r.Body).Decode(&account)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(account.Secret), 8)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	account.Secret = string(hashedPassword)
	account.Balance = 0
	account.CreatedAt = time.Now()

	database.DB.Create(&account)
	json.NewEncoder(w).Encode(account)
}
