package middleware

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		if recovered := recover(); recovered != nil {
			recoveredErr, ok := recovered.(error)
			if !ok {
				return c.Status(http.StatusInternalServerError).JSON(
					fiber.Map{
						"error":  http.StatusText(http.StatusInternalServerError),
						"status": http.StatusInternalServerError,
					},
				)
			}
			return createErrorResponse(c, recoveredErr)
		}
		return createErrorResponse(c, err)
	}
}

func createErrorResponse(c *fiber.Ctx, err error) error {
	statusCode := fiber.StatusInternalServerError
	var errorMessage string
	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		statusCode = fiberErr.Code
		errorMessage = fiberErr.Message
	}
	if len(errorMessage) == 0 {
		errorMessage = http.StatusText(statusCode)
	}
	return c.Status(statusCode).JSON(
		fiber.Map{
			"error":  errorMessage,
			"status": statusCode,
		},
	)
}
