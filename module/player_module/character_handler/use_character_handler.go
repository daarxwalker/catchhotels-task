package character_handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"catchhotels/entity/character_entity"
	"catchhotels/facade"
	"catchhotels/module/player_module/character_handler/character_response"
)

// UseCharacter Choose and play with character
//
//	@Summary        Choose and play with character
//	@Description    CharacterId is stored in session
//	@Tags           Player character
//	@Accept         json
//	@Produce        json
//	@Param          characterId path string true "Character ID"
//	@Success        200  {boolean}  true
//	@Router         /player/character/use/:characterId [post]
func UseCharacter() fiber.Handler {
	return func(c *fiber.Ctx) error {
		characterId := c.Params("characterId")
		if len(characterId) == 0 {
			c.Status(http.StatusBadRequest)
			return character_response.ErrInvalidCharacter
		}
		var character character_entity.Character
		if findCharacterErr := facade.Dataverse(c).FindOne(
			character_entity.Table, characterId, &character,
		); findCharacterErr != nil {
			c.Status(http.StatusBadRequest)
			return character_response.ErrInvalidCharacter
		}
		session := facade.Session(c).MustGet(c)
		session.CharacterId = characterId
		facade.Session(c).MustSet(c, session)
		return c.JSON(true)
	}
}
