package app

import (
	"web2022/internal/api"
	"web2022/internal/config"

	"github.com/joho/godotenv"
)

func StartServer() error {
	err := godotenv.Load("local.env")
	if err != nil {
		panic(err)
	}
	var cfg config.Config
	cfg.Fill()

	api, err := api.New(cfg)
	if err != nil {
		return err
	}

	api.FillEndpoints()

	return api.Serve()
}
