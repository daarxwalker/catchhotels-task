package auth_handler

import (
	"github.com/gofiber/fiber/v2"

	"catchhotels/entity/player_entity"
	"catchhotels/facade"
	"catchhotels/module/player_module/auth_handler/auth_request"
	"catchhotels/module/player_module/auth_handler/auth_response"
	"catchhotels/service/dataverse_service"
	"catchhotels/service/session_service"
)

// Login Player authentication
//
//	@Summary        Player authentication
//	@Description    Login with email
//	@Tags           Authentication
//	@Accept         json
//	@Produce        json
//	@Param          request body auth_request.LoginPlayer true "Login request body"
//	@Success        200  {object}  session_service.Session
//	@Router         /auth/login [post]
func Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body auth_request.LoginPlayer
		if parseBodyErr := c.BodyParser(&body); parseBodyErr != nil {
			return parseBodyErr
		}
		if validateBodyErr := facade.Validate(c).Struct(body); validateBodyErr != nil {
			c.Status(fiber.StatusBadRequest)
			return validateBodyErr
		}
		var playersResponse dataverse_service.FindResponse[player_entity.Player]
		if findPlayersErr := facade.Dataverse(c).Find(
			player_entity.Table, &playersResponse,
		); findPlayersErr != nil || len(playersResponse.Value) == 0 {
			c.Status(fiber.StatusUnauthorized)
			return auth_response.ErrInvalidPlayer
		}
		var playerId string
		for _, player := range playersResponse.Value {
			if player.Email == body.Email {
				playerId = player.Id
				break
			}
		}
		newSession := session_service.Session{
			Id:        playerId,
			Email:     body.Email,
			IP:        c.Get(fiber.HeaderXForwardedFor),
			UserAgent: c.Get(fiber.HeaderUserAgent),
		}
		facade.Session(c).MustSet(c, newSession)
		return c.JSON(newSession)
	}
}
