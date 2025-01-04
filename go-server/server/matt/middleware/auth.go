package middleware

import (
	"net/http"

	"github.com/0xBradock/go-srvr/internal/user"
)

func AuthForRoles(next http.Handler, roles ...user.Role) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Get("Authorization")

		incomingRole := user.Customer

		authorized := false
		for _, role := range roles {
			if role == incomingRole {
				authorized = true
			}
		}

		if !authorized {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("unauthorized"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
