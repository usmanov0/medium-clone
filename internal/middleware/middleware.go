package middleware

import (
	"context"
	"example.com/my-medium-clone/internal/common/jwt"
	"net/http"
	"strings"
)

func AuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")
		if parts := strings.Split(bearerToken, " "); len(parts) == 2 && parts[0] == "Bearer" {
			tokenStr := parts[1]

			claims, err := jwt.VerifyToken(tokenStr)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "UserId", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
	})
}
