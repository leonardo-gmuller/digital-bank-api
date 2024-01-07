package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/LeonardoMuller13/digital-bank-api/src/config"
	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID interface{}
}

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
			if r.Context().Value(User{}) == nil {
				user := User{
					ID: claims["user"],
				}
				r = r.WithContext(context.WithValue(r.Context(), User{}, user))
			}
		}
		next.ServeHTTP(w, r)
	})
}

func verifyToken(tokenString string) (error, jwt.MapClaims) {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load configurations: %v", err)
	}
	secretKey := []byte(cfg.JwtSecretKey)
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
