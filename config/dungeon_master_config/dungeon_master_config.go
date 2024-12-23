package dungeon_master_config

import (
	"time"

	"github.com/spf13/viper"
)

const (
	Secret               = "dungeon-master-secret"
	PlayersCacheKey      = "dungeon-master-players-cache-key"
	PlayersCacheDuration = "dungeon-master-players-cache-duration"
)

func Read(cfg *viper.Viper) {
	cfg.Set(PlayersCacheKey, "dungeon-master-players")
	cfg.Set(PlayersCacheDuration, 15*time.Minute)
	cfg.Set(Secret, "SvawVmY^}9i5L&kx:XiaM2[Ol!PME")
}
