package config

import (
	"os"
)

type Api struct {
	Port           string   `json:"port" mapstructure:"APP_PORT"`
	Env            string   `json:"env" mapstructure:"ENV"`
	JWTSecretKey   string   `json:"-" mapstructure:"JWT_SECRET_KEY"`
	JWTExpiredTime string   `json:"jwt_expired_time" mapstructure:"JWT_EXPIRED_TIME"`
	Database       Database `json:"database"`
}

type Database struct {
	Host     string `json:"host" mapstructure:"DATABASE_HOST"`
	Port     string `json:"port" mapstructure:"DATABASE_PORT"`
	Username string `json:"username" mapstructure:"DATABASE_USERNAME"`
	Password string `json:"password" mapstructure:"DATABASE_PASSWORD"`
	Schema   string `json:"shcema" mapstructure:"DATABASE_SCHEMA"`
	Loc      string `json:"loc" mapstructure:"DATABASE_LOC"`
}

func LoadConfingAPI() *Api {
	api := &Api{}

	api.Port = os.Getenv("APP_PORT")
	api.JWTSecretKey = os.Getenv("JWT_SECRET_KEY")
	api.JWTExpiredTime = os.Getenv("JWT_SECRET_KEY")

	return api
}

func LoadConfigDatabase() *Database {
	database := &Database{}

	database.Host = os.Getenv("DATABASE_HOST")
	database.Port = os.Getenv("DATABASE_PORT")
	database.Username = os.Getenv("DATABASE_USERNAME")
	database.Password = os.Getenv("DATABASE_PASSWORD")
	database.Schema = os.Getenv("DATABASE_SCHEMA")
	database.Loc = os.Getenv("DATABASE_LOC")

	return database
}
