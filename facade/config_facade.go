package facade

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"

	"catchhotels/config"
)

func Config(c *fiber.Ctx) *viper.Viper {
	instance, instanceExists := c.Locals(config.Token).(*viper.Viper)
	if !instanceExists {
		panic(fmt.Sprintf("%s does not exist", config.Token))
	}
	return instance
}
