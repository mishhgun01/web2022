package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	CertFile string
	KeyFile  string
	Host     string
	Port     string
	DbName   string
	DbHost   string
	DbUser   string
	DbPort   string
	DbPwd    string
}

func (cfg *Config) Fill() {
	err := godotenv.Load("/home/mishhgun01/go/src/web2022/internal/config/local.env")
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	cfg.CertFile = os.Getenv("CERT_FILE")
	cfg.KeyFile = os.Getenv("KEY_FILE")
	cfg.Host = os.Getenv("SERVER_HOST")
	cfg.Port = os.Getenv("SERVER_PORT")
	cfg.DbName = os.Getenv("DB_NAME")
	cfg.DbHost = os.Getenv("DB_HOST")
	cfg.DbUser = os.Getenv("DB_USER")
	cfg.DbPort = os.Getenv("DB_PORT")
	cfg.DbPwd = os.Getenv("DB_PASSWORD")
}
