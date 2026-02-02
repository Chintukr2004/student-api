package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Chintukr2004/student-api/internal/utils"
)

type contextKey string

const (
	UserIDKey contextKey = "userID"
	RoleKey   contextKey = "role"
)

func Auth(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "missing auth header", http.StatusUnauthorized)
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "invalid auth header", http.StatusUnauthorized)
				return
			}

			claims, err := utils.ParseToken(parts[1], secret)
			if err != nil {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
			ctx = context.WithValue(ctx, RoleKey, claims.Role)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
