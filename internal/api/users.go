package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

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
		//TODO

	case http.MethodDelete:
		//TODO
	}
}
