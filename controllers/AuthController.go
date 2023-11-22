package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/LeonardoMuller13/digital-bank-api/database"
	"github.com/LeonardoMuller13/digital-bank-api/models"
	"github.com/bitly/go-simplejson"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("my_secret_key")

type Credentials struct {
	Password string `json:"password", db:"password"`
	Cpf      string `json:"cpf", db:"cpf"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	account := &models.Account{}
	result := database.DB.First(&account, "cpf = ?", creds.Cpf)
	if result.Error != nil {
		// If there is an issue with the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(account.Secret), []byte(creds.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println(creds.Password)
		fmt.Errorf("Nao autorizado")
		return
	}
	tokenString, err := createToken(creds.Cpf)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Errorf("No username found")
	}
	w.WriteHeader(http.StatusOK)

	response := simplejson.New()
	response.Set("token", tokenString)

	payload, err := response.MarshalJSON()
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
	return
}

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
