package facade

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"catchhotels/service/dice_service"
)

func Dice(c *fiber.Ctx) *dice_service.DiceService {
	instance, instanceExists := c.Locals(dice_service.Token).(*dice_service.DiceService)
	if !instanceExists {
		panic(fmt.Sprintf("%s does not exist", dice_service.Token))
	}
	return instance
}
