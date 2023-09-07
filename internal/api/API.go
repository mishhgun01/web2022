package api

import (
	"log"
	"net/http"
	"web2022/internal/config"
	"web2022/internal/storage"

	"fmt"

	"github.com/gorilla/mux"
)

type API struct {
	R  *mux.Router
	db *storage.Storage
	serveOptions
}

type serveOptions struct {
	addr string
	cert string
	key  string
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

	return &API{
		R:  router,
		db: s,
		serveOptions: serveOptions{
			addr: fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
			cert: cfg.CertFile,
			key:  cfg.KeyFile,
		},
	}, nil
}

func (api *API) FillEndpoints() {
	api.R.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello!")
	})
}

func (api *API) Serve() error {
	log.Println("ping")
	return http.ListenAndServeTLS(api.addr, api.cert, api.key, api.R)
}
