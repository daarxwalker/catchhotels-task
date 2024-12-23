package facade

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"catchhotels/service/session_service"
)

func Session(c *fiber.Ctx) *session_service.SessionService {
	instance, instanceExists := c.Locals(session_service.Token).(*session_service.SessionService)
	if !instanceExists {
		panic(fmt.Sprintf("%s does not exist", session_service.Token))
	}
	return instance
}
