package auth_handler

import (
	"github.com/gofiber/fiber/v2"

	"catchhotels/facade"
)

// Logout Destroy player session
//
//	@Summary        Destroy player session
//	@Description    Destroy session in cache and expire cookie
//	@Tags           Authentication
//	@Accept         json
//	@Produce        json
//	@Success        200  {boolean}  true
//	@Router         /auth/logout [delete]
func Logout() fiber.Handler {
	return func(c *fiber.Ctx) error {
		facade.Session(c).MustDestroy(c)
		return c.JSON(true)
	}
}
