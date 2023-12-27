package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LeonardoMuller13/digital-bank-api/src/database"
	"github.com/LeonardoMuller13/digital-bank-api/src/models"
	"github.com/bitly/go-simplejson"
)

type RequestTransfer struct {
	Account_Destination_CPF string
	Amount                  int
}

func GetTransfers(w http.ResponseWriter, r *http.Request) {
	var t []models.Transfer
	accountOrigin := r.Context().Value("account").(*models.Account)
	result := database.DB.Find(&t, "account_origin_id = ?", accountOrigin.ID)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Account not found.")
		return
	}
	json.NewEncoder(w).Encode(t)
	return
}

func NewTransfer(w http.ResponseWriter, r *http.Request) {
	fail := func(status int, err string) {
		w.WriteHeader(status)
		fmt.Fprint(w, err)
		return
	}
	tx := database.DB.Begin()

	defer tx.Rollback()

	newRequest := &RequestTransfer{}
	json.NewDecoder(r.Body).Decode(&newRequest)
	accountOrigin := r.Context().Value("account").(*models.Account)

	var accountDestination models.Account
	result := database.DB.First(&accountDestination, "cpf = ?", newRequest.Account_Destination_CPF)
	if result.Error != nil {
		fail(http.StatusInternalServerError, "Account destination not found.")
		return
	}

	err := accountOrigin.Transfer(newRequest.Amount, &accountDestination)
	if err != nil {
		fail(http.StatusInternalServerError, err.Error())
		return
	}

	tx.Save(&accountOrigin)
	tx.Save(&accountDestination)

	var newTransfer models.Transfer
	newTransfer.AccountDestinationId = int(accountDestination.ID)
	newTransfer.AccountOriginId = int(accountOrigin.ID)
	newTransfer.Amount = newRequest.Amount
	tx.Create(&newTransfer)

	w.WriteHeader(http.StatusOK)
	response := simplejson.New()
	response.Set("message", "Transfer completed successfully!")
	payload, err := response.MarshalJSON()
	if err != nil {
		fail(http.StatusInternalServerError, "Internal error.")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)

	tx.Commit()
	return
}
