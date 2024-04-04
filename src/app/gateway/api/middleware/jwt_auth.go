package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/LeonardoMuller13/digital-bank-api/src/app/domain/dto"
	"github.com/golang-jwt/jwt/v5"
)

func ProtectedHandler(jwtSecret string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-type", "application/json")
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, "Missing authorization header")
				return
			}
			tokenString = tokenString[len("Bearer "):]

			err, claims := verifyToken(tokenString, jwtSecret)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, "Invalid token")
				return
			}
			if claims != nil {
				if r.Context().Value(dto.User{}) == nil {
					user := dto.User{
						ID: claims["user"],
					}
					r = r.WithContext(context.WithValue(r.Context(), dto.User{}, user))
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}

func verifyToken(tokenString string, jwtSecret string) (error, jwt.MapClaims) {
	secretKey := []byte(jwtSecret)
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
