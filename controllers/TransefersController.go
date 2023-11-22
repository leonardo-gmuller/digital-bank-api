package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/LeonardoMuller13/digital-bank-api/database"
	"github.com/LeonardoMuller13/digital-bank-api/models"
)

func GetTransfers(w http.ResponseWriter, r *http.Request) {
	var t []models.Transfer
	database.DB.Find(&t)
	json.NewEncoder(w).Encode(t)
}

func getAccountLogged(r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	tokenString = tokenString[len("Bearer "):]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
}
