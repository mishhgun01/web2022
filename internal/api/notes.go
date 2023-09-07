package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"web2022/internal/models"
)

// NotesCRUDHandler - обработчик запросов, связанных с сущностью заметки.
func (api *API) NotesCRUDHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		userIDString := r.URL.Query().Get("user_id")
		if userIDString == "" {
			http.Error(w, "Не передан идентификатор пользователя.", http.StatusBadRequest)
			return
		}

		userID, _ := strconv.Atoi(userIDString)

		notes, err := api.db.GetNotesByUserID(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(notes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	case http.MethodPost:
		var note models.Note
		err := json.NewDecoder(r.Body).Decode(&note)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		id, err := api.db.NewNote(note)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	case http.MethodPatch:
		var note models.Note
		err := json.NewDecoder(r.Body).Decode(&note)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		err = api.db.UpdateNote(note)
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
