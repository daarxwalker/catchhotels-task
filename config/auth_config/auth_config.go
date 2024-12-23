package auth_config

import (
	"time"

	"github.com/spf13/viper"
)

const (
	SessionExpiration = "auth-session-expiration"
)

func Read(v *viper.Viper) {
	v.Set(SessionExpiration, 30*24*time.Hour)
}
