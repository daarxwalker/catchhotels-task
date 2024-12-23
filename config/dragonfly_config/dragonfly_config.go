package dragonfly_config

import "github.com/spf13/viper"

const (
	Uri      = "dragonfly-uri"
	Password = "dragonfly-password"
	DB       = "dragonfly-db"
)

func Read(v *viper.Viper) {
	v.Set(Uri, "catchhotels-dragonfly:6379")
	v.Set(Password, "")
	v.Set(DB, 0)
}
