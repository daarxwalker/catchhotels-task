package middleware

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"

	"catchhotels/config/dungeon_master_config"
)

func DungeonMaster(cfg *viper.Viper) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if !strings.HasSuffix(c.Get("Authorization"), cfg.GetString(dungeon_master_config.Secret)) {
			return c.JSON(
				fiber.Map{
					"error":  http.StatusText(http.StatusUnauthorized),
					"status": http.StatusUnauthorized,
				},
			)
		}
		return c.Next()
	}
}
