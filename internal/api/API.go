package api

import (
	"crypto/tls"
	"net/http"
	"web2022/internal/config"
	"web2022/internal/storage"

	"fmt"

	"github.com/gorilla/mux"
)

type API struct {
	R    *mux.Router
	db   *storage.Storage
	addr string

	TLSConfig *tls.Config
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
	cert, err := tls.LoadX509KeyPair(cfg.CertFile, cfg.KeyFile)
	if err != nil {
		return nil, err
	}

	return &API{
		R:    router,
		db:   s,
		addr: fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}, nil
}

func (api *API) FillEndpoints() {
	api.R.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello!")
	})
}

func (api *API) Serve() error {
	return http.ListenAndServeTLS(api.addr, "", "", api.R)
}
