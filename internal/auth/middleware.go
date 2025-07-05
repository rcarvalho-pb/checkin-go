package auth

import (
	"context"
	"log"
	"net/http"
)

type contextKey string

const UserIDKey = contextKey("userID")
const UserNameKey = contextKey("userName")
const UserEmailKey = contextKey("userEmail")
const UserRoleKey = contextKey("userRole")

func (a *AuthHandler) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := a.ValidateToken(r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		claims, err := a.ParseJwtTokenWithClaims(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), UserIDKey, claims.ID)
		ctx = context.WithValue(ctx, UserNameKey, claims.Name)
		ctx = context.WithValue(ctx, UserEmailKey, claims.Email)
		ctx = context.WithValue(ctx, UserRoleKey, claims.Role)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func (a *AuthHandler) LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Host, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
