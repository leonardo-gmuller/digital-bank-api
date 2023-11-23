package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/LeonardoMuller13/digital-bank-api/database"
	"github.com/LeonardoMuller13/digital-bank-api/models"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("my_secret_key")

func ProtectedHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Missing authorization header")
			return
		}
		tokenString = tokenString[len("Bearer "):]

		err, claims := verifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Invalid token")
			return
		}
		if claims != nil {
			if r.Context().Value("account") == nil {
				account := &models.Account{}
				result := database.DB.First(&account, "cpf = ?", claims["username"])
				if result.Error != nil {
					// If there is an issue with the database, return a 500 error
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				r = r.WithContext(context.WithValue(r.Context(), "account", account))
			}
		}
		next.ServeHTTP(w, r)
	})
}

func verifyToken(tokenString string) (error, jwt.MapClaims) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return err, nil
	}

	if !token.Valid {
		return fmt.Errorf("Invalid token"), nil
	}
	claims := token.Claims.(jwt.MapClaims)
	return nil, claims
}
