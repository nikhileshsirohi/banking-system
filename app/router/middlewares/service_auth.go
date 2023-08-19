package middlewares

import (
	"encoding/base64"
	"net/http"
	"strings"
)

const (
	username = "banking_system"
	password = "banking_system_123"
)

// Middleware function for Basic Authentication
func BasicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Parse and decode the Authorization header
		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || authParts[0] != "Basic" {
			http.Error(w, "Invalid authorization header", http.StatusBadRequest)
			return
		}

		authValue, err := base64.StdEncoding.DecodeString(authParts[1])
		if err != nil {
			http.Error(w, "Invalid authorization header", http.StatusBadRequest)
			return
		}

		credentials := strings.SplitN(string(authValue), ":", 2)
		if len(credentials) != 2 || credentials[0] != username || credentials[1] != password {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
