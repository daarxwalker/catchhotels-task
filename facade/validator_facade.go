package facade

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"catchhotels/service/validator_service"
)

func Validate(c *fiber.Ctx) *validator.Validate {
	instance, instanceExists := c.Locals(validator_service.Token).(*validator_service.ValidatorService)
	if !instanceExists {
		panic(fmt.Sprintf("%s does not exist", validator_service.Token))
	}
	return instance.Validate()
}
