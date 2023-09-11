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

	currentWorkDirectory, _ := os.Getwd()
	serverWorkDir := currentWorkDirectory + "/internal/"

	err := godotenv.Load(serverWorkDir + `config/local.env`)
	if err != nil {
		log.Println(err.Error())
		panic(err.Error())
	}
	cfg.CertFile = serverWorkDir + "config/" + os.Getenv("CERT_FILE")
	cfg.KeyFile = serverWorkDir + "config/" + os.Getenv("KEY_FILE")
	cfg.Host = os.Getenv("SERVER_HOST")
	cfg.Port = os.Getenv("SERVER_PORT")
	cfg.DbName = os.Getenv("DB_NAME")
	cfg.DbHost = os.Getenv("DB_HOST")
	cfg.DbUser = os.Getenv("DB_USER")
	cfg.DbPort = os.Getenv("DB_PORT")
	cfg.DbPwd = os.Getenv("DB_PASSWORD")
	cfg.WebDir = currentWorkDirectory + "/webapp/" + os.Getenv("WEB_DIR")
}
