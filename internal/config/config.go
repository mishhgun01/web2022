package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Config - структура конфига приложения.
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
	WebDir   string
}

const projectDirName = "web2022/internal/config" // change to relevant project name

// Fill - заполнение полей конфига из системных переменных.
func (cfg *Config) Fill() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	configDir := "/internal/config/"
	webDir := "/webapp/"
	godotenv.Load(wd + configDir + "local.env")
	cfg.CertFile = wd + configDir + os.Getenv("CERT_FILE")
	cfg.KeyFile = wd + configDir + os.Getenv("KEY_FILE")
	cfg.Host = os.Getenv("SERVER_HOST")
	cfg.Port = os.Getenv("SERVER_PORT")
	cfg.DbName = os.Getenv("DB_NAME")
	cfg.DbHost = os.Getenv("DB_HOST")
	cfg.DbUser = os.Getenv("DB_USER")
	cfg.DbPort = os.Getenv("DB_PORT")
	cfg.DbPwd = os.Getenv("DB_PASSWORD")
	cfg.WebDir = wd + webDir + os.Getenv("WEB_DIR")
	log.Println(cfg.WebDir)
}
