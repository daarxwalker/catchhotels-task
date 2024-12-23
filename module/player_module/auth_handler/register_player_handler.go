package auth_handler

import (
	"github.com/gofiber/fiber/v2"

	"catchhotels/entity/player_entity"
	"catchhotels/facade"
	"catchhotels/module/player_module/auth_handler/auth_request"
	"catchhotels/module/player_module/auth_handler/auth_response"
	"catchhotels/service/dataverse_service"
)

// Register Create player account
//
//	@Summary        Create player account
//	@Description    Validate CIN in ARES
//	@Tags           Authentication
//	@Accept         json
//	@Produce        json
//	@Param          request body auth_request.RegisterPlayer true "Register request body"
//	@Success        200  {boolean}  true
//	@Router         /auth/register [post]
func Register() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body auth_request.RegisterPlayer
		if parseBodyErr := c.BodyParser(&body); parseBodyErr != nil {
			return parseBodyErr
		}
		if validateBodyErr := facade.Validate(c).Struct(body); validateBodyErr != nil {
			c.Status(fiber.StatusBadRequest)
			return validateBodyErr
		}
		aresResponse := facade.Ares(c).MustGet(body.Cin)
		if len(aresResponse.Cin) == 0 {
			c.Status(fiber.StatusBadRequest)
			return auth_response.ErrInvalidCin
		}
		var playersResponse dataverse_service.FindResponse[player_entity.Player]
		facade.Dataverse(c).MustFindWithEmail(player_entity.Table, body.Email, &playersResponse)
		if len(playersResponse.Value) > 0 {
			c.Status(fiber.StatusBadRequest)
			return auth_response.ErrPlayerExists
		}
		var createResponse fiber.Map
		facade.Dataverse(c).MustCreate(
			"cr568_players", player_entity.Player{
				FirstName: body.FirstName,
				LastName:  body.LastName,
				Phone:     body.Phone,
				Email:     body.Email,
				CIN:       body.Cin,
				VAT:       aresResponse.Vat,
				Address:   aresResponse.Address.Full,
			}, &createResponse,
		)
		return c.JSON(true)
	}
}
