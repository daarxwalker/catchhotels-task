package middleware

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"catchhotels/facade"
)

func AccessGuard() fiber.Handler {
	return func(c *fiber.Ctx) error {
		sessionService := facade.Session(c)
		session, getSessionErr := sessionService.Get(c)
		if !session.Exists() || getSessionErr != nil {
			c.Status(http.StatusUnauthorized)
			return errors.New("session not found")
		}
		ip := c.Get(fiber.HeaderXForwardedFor)
		userAgent := c.Get(fiber.HeaderUserAgent)
		if !session.CompareIP(ip) || !session.CompareUserAgent(userAgent) {
			c.Status(http.StatusForbidden)
			return errors.New("invalid ip or user agent")
		}
		if renewSessionErr := sessionService.Renew(c); renewSessionErr != nil {
			c.Status(http.StatusInternalServerError)
			return errors.New("renew session error")
		}
		return c.Next()
	}
}
