package app

import (
	"web2022/internal/api"
	"web2022/internal/config"
)

func StartServer() error {
	var cfg config.Config
	cfg.Fill()

	api, err := api.New(cfg)
	if err != nil {
		return err
	}

	api.FillEndpoints()

	return api.Serve()
}
