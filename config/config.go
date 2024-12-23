package config

import (
	"github.com/spf13/viper"

	"catchhotels/config/app_config"
	"catchhotels/config/ares_config"
	"catchhotels/config/auth_config"
	"catchhotels/config/dataverse_config"
	"catchhotels/config/dragonfly_config"
	"catchhotels/config/dungeon_master_config"
)

const (
	Token = "config"
)

func Read() *viper.Viper {
	cfg := viper.New()
	app_config.Read(cfg)
	auth_config.Read(cfg)
	ares_config.Read(cfg)
	dataverse_config.Read(cfg)
	dragonfly_config.Read(cfg)
	dungeon_master_config.Read(cfg)
	return cfg
}
