package ares_config

import "github.com/spf13/viper"

const (
	Endpoint = "ares-endpoint"
)

func Read(cfg *viper.Viper) {
	cfg.Set(Endpoint, "https://ares.gov.cz/ekonomicke-subjekty-v-be/rest/ekonomicke-subjekty/")
}
