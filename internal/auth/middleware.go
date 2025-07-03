package auth

import "net/http"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := ValidateToken(r); err == nil {
			next.ServeHTTP(w, r)
		}
	})
}
