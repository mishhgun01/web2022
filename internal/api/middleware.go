package api

import (
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

// HeadersMiddleware - мидлвэр для установки заголовков.
func (api *API) HeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/api/") {
			if r.Method == http.MethodOptions {

				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "*")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Content-Type", "application/json")
				return
			}

			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "application/json")
		}
		userName, _, ok := r.BasicAuth()
		if !ok {
			http.Error(w, errors.New("not authorized").Error(), http.StatusUnauthorized)
		}

		user, err := api.db.GetUserByName(userName)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		r.AddCookie(&http.Cookie{
			Name:  "lastPath",
			Value: user.LastPath,
		})

		next.ServeHTTP(w, r)
	})
}

// AuthMiddleware - мидлвэр для проведения http - аутентификации и авторизации.
func (api *API) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/api/") {
			if strings.Contains(r.URL.Path, "/api/v1/user") && (r.Method == http.MethodPost || r.Method == http.MethodOptions) {
				next.ServeHTTP(w, r)
				return
			}

			user, pass, ok := r.BasicAuth()
			pwd, err := api.db.GetUserPasswordByName(user)
			flag := verifyUserPassword(pwd, pass)
			if (strings.Contains(r.URL.Path, "/api/v1/ping") || strings.Contains(r.URL.Path, "/api/v1/user")) && r.Method == http.MethodGet {
				flag = bcrypt.CompareHashAndPassword([]byte(pwd), []byte(pass))
			}
			if !ok || err != nil || flag != nil {
				w.Header().Set("Notes-WWW-Authenticate", `Basic realm="api"`)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			if r.Method == http.MethodGet && strings.Contains(r.URL.Path, "/api/v1/user") {
				data, err := api.db.GetUserByName(user)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				json.NewEncoder(w).Encode(data)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

// verifyUserPassword - вспомогательная функци проверки совпадения паролей.
func verifyUserPassword(dbPassword, gotPassword string) error {
	if dbPassword == gotPassword {
		return nil
	}
	return errors.New("Wrong pwd")
}
