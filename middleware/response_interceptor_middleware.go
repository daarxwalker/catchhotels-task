package middleware

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func ResponseInterceptor() fiber.Handler {
	return func(c *fiber.Ctx) error {
		nextErr := c.Next()
		contentType := string(c.Response().Header.ContentType())
		if (contentType == fiber.MIMETextPlain && c.Response().StatusCode() >= fiber.StatusBadRequest) || nextErr != nil {
			err := string(c.Response().Body())
			if nextErr != nil {
				err = nextErr.Error()
			}
			modifiedBody, marshalModifiedBodyErr := json.Marshal(
				fiber.Map{
					"status": c.Response().StatusCode(),
					"error":  err,
				},
			)
			if marshalModifiedBodyErr != nil {
				return marshalModifiedBodyErr
			}
			c.Response().Header.Set("Content-Type", fiber.MIMEApplicationJSON)
			c.Response().SetBody(modifiedBody)
			return nil
		}
		if contentType == fiber.MIMEApplicationJSON {
			body := c.Response().Body()
			if len(body) == 0 {
				return nil
			}
			var data any
			if unmarshalResponseErr := json.Unmarshal(body, &data); unmarshalResponseErr != nil {
				return unmarshalResponseErr
			}
			modifiedBody, marshalModifiedBodyErr := json.Marshal(
				fiber.Map{
					"status": c.Response().StatusCode(),
					"result": data,
				},
			)
			if marshalModifiedBodyErr != nil {
				return marshalModifiedBodyErr
			}
			c.Response().SetBody(modifiedBody)
			return nil
		}
		return nil
	}
}
