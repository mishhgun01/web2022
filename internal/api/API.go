package api

import (
	"web2022/internal/config"
	"web2022/internal/storage"

	"fmt"

	"github.com/gorilla/mux"
)

type API struct {
	R    *mux.Router
	db   *storage.Storage
	addr string
}

func New(cfg config.Config) (*API, error) {
	s, err := storage.New(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbUser,
		cfg.DbPwd,
		cfg.DbName,
	))
	if err != nil {
		return nil, err
	}

	router := mux.NewRouter()

	return &API{R: router, db: s, addr: fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)}, nil
}

func (api *API) FillEndpoints() {
	api.R.HandleFunc("/api/v1/ping")
	api.R.HandleFunc("/api/v1/users")
	api.R.HandleFunc("/api/v1/notes")
}
