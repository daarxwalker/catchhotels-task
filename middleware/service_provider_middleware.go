package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"

	"catchhotels/config"
	"catchhotels/service/ares_service"
	"catchhotels/service/dataverse_service"
	"catchhotels/service/dice_service"
	"catchhotels/service/dragonfly_service"
	"catchhotels/service/session_service"
	"catchhotels/service/validator_service"
)

func ServiceProvider(cfg *viper.Viper) fiber.Handler {
	aresService := ares_service.New(cfg)
	dataverseService, createDataverseErr := dataverse_service.New(cfg)
	if createDataverseErr != nil {
		log.Fatalln(createDataverseErr)
	}
	dragonflyService := dragonfly_service.New(cfg)
	sessionService := session_service.New(cfg, dragonflyService)
	diceService := dice_service.New(dataverseService, sessionService)
	validatorService := validator_service.New()
	return func(c *fiber.Ctx) error {
		c.Locals(config.Token, cfg)
		c.Locals(ares_service.Token, aresService)
		c.Locals(dataverse_service.Token, dataverseService)
		c.Locals(dice_service.Token, diceService)
		c.Locals(dragonfly_service.Token, dragonflyService)
		c.Locals(session_service.Token, sessionService)
		c.Locals(validator_service.Token, validatorService)
		return c.Next()
	}
}
