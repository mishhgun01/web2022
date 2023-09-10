package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"web2022/internal/models"

	"golang.org/x/crypto/bcrypt"
)

// UsersHandler - обработчик запросов, связанных с сущностью пользователя.
func (api *API) UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		idString := r.URL.Query().Get("id")
		if idString == "" {
			http.Error(w, "Не передан идентификатор пользователя.", http.StatusBadRequest)
			return
		}

		id, _ := strconv.Atoi(idString)

		user, err := api.db.GetUserByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	case http.MethodPost:
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		pwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), 16)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user.Password = string(pwd)

		id, err := api.db.NewUser(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user.ID = id

		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	case http.MethodDelete:
		type request struct {
			id int
		}

		var req request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = api.db.DeleteUserByID(req.id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusUnauthorized)
	}
}
