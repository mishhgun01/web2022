package api

import (
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func (api *API) HeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/api/") {

			user, pass, ok := r.BasicAuth()
			pwd, err := api.db.GetUserPasswordByName(user)
			if !ok || err != nil || verifyUserPassword(pwd, pass) != nil {
				w.Header().Set("Notes-WWW-Authenticate", `Basic realm="api"`)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "*")

			if r.Method == http.MethodOptions {
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "*")

				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func verifyUserPassword(dbPassword, gotPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(gotPassword))
}
