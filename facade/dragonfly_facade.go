package facade

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"catchhotels/service/dragonfly_service"
)

func Dragonfly(c *fiber.Ctx) *dragonfly_service.DragonflyService {
	instance, instanceExists := c.Locals(dragonfly_service.Token).(*dragonfly_service.DragonflyService)
	if !instanceExists {
		panic(fmt.Sprintf("%s does not exist", dragonfly_service.Token))
	}
	return instance
}
