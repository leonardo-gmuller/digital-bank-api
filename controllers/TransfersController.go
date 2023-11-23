package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/LeonardoMuller13/digital-bank-api/database"
	"github.com/LeonardoMuller13/digital-bank-api/models"
	"github.com/bitly/go-simplejson"
)

type RequestTransfer struct {
	Account_Destination_CPF string
	Amount                  int
}

func GetTransfers(w http.ResponseWriter, r *http.Request) {
	accountOrigin := r.Context().Value("account").(*models.Account)
	var t []models.Transfer
	result := database.DB.Find(&t, "account_origin_id = ?", accountOrigin.ID)
	if result.Error != nil {
		// If there is an issue with the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(t)
	return
}

func NewTransfer(w http.ResponseWriter, r *http.Request) {
	newRequest := &RequestTransfer{}
	json.NewDecoder(r.Body).Decode(&newRequest)
	accountOrigin := r.Context().Value("account").(*models.Account)
	var accountDestination models.Account
	result := database.DB.First(&accountDestination, "cpf = ?", newRequest.Account_Destination_CPF)
	if result.Error != nil {
		// If there is an issue with the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err := accountOrigin.Transfer(newRequest.Amount, &accountDestination)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	database.DB.Save(&accountOrigin)
	database.DB.Save(&accountDestination)
	var newTransfer models.Transfer
	newTransfer.Account_destination_ID = int(accountDestination.ID)
	newTransfer.Account_origin_ID = int(accountOrigin.ID)
	newTransfer.Amount = newRequest.Amount
	database.DB.Create(&newTransfer)

	w.WriteHeader(http.StatusOK)
	response := simplejson.New()
	response.Set("message", "Transferencia feita com sucesso!")
	payload, err := response.MarshalJSON()
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
	return
}
