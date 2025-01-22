package middleware

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/jwtauth"

	"github.com/leonardo-gmuller/digital-bank-api/app/domain/dto"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			http.Error(w, "failed to parse token from context", http.StatusUnauthorized)

			return
		}

		userID, ok := claims["sub"].(string)
		if !ok {
			http.Error(w, "invalid token claims: 'sub' not a string", http.StatusUnauthorized)

			return
		}

		userIDInt, err := strconv.Atoi(userID)
		if err != nil {
			http.Error(w, "invalid user ID format", http.StatusUnauthorized)

			return
		}

		u := new(dto.User)
		r = r.WithContext(context.WithValue(r.Context(), dto.UserKey, u))

		if u, ok := r.Context().Value(dto.UserKey).(*dto.User); ok {
			u.ID = userIDInt
		}

		next.ServeHTTP(w, r)
	})
}
