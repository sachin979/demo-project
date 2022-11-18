package config

import (
	"log"
	"todo/services"

	"github.com/gin-gonic/gin"
	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type DbConf struct {
	Host     string `env:"DB_HOST,required"`
	Port     int    `env:"DB_PORT,required"`
	Username string `env:"DB_USER,required"`
	Password string `env:"DB_PASS,required"`
	DbName   string `env:"DB_NAME,required"`
}

type Conf struct {
	Server serverConf
	Db     DbConf
}

type serverConf struct {
	Port string `env:"SERVER_PORT,required"`
}

type AppConfig struct {
	Dbs    *Dbs
	Router *gin.Engine
	Cfg    *Conf
}

type Dbs struct {
	DB *gorm.DB
}

type HandlerConfig struct {
	R           *gin.Engine
	TodoService services.ITodoService
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not loaded")
	}
}

func AppConfigF() *Conf {
	var c Conf

	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	return &c
}
