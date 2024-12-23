package app_config

import (
	"os"

	"github.com/spf13/viper"
)

const (
	Env  = "app-env"
	Name = "app-name"
	Port = "app-port"
)

const (
	EnvDevelopment = "development"
)

func Read(cfg *viper.Viper) {
	cfg.Set(Env, os.Getenv("APP_ENV"))
	cfg.Set(Name, os.Getenv("APP_NAME"))
	cfg.Set(Port, os.Getenv("APP_PORT"))
}
