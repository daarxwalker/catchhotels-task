package facade

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"catchhotels/service/dataverse_service"
)

func Dataverse(c *fiber.Ctx) *dataverse_service.DataverseService {
	instance, instanceExists := c.Locals(dataverse_service.Token).(*dataverse_service.DataverseService)
	if !instanceExists {
		panic(fmt.Sprintf("%s does not exist", dataverse_service.Token))
	}
	return instance
}
