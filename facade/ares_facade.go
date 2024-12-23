package facade

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"catchhotels/service/ares_service"
)

func Ares(c *fiber.Ctx) *ares_service.AresService {
	instance, instanceExists := c.Locals(ares_service.Token).(*ares_service.AresService)
	if !instanceExists {
		panic(fmt.Sprintf("%s does not exist", ares_service.Token))
	}
	return instance
}
