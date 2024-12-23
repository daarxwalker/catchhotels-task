package character_handler

import (
	"github.com/gofiber/fiber/v2"

	"catchhotels/entity/character_entity"
	"catchhotels/entity/class_entity"
	"catchhotels/entity/race_entity"
	"catchhotels/facade"
	"catchhotels/module/player_module/character_handler/character_request"
	"catchhotels/module/player_module/character_handler/character_response"
)

// CreateCharacter Create a new player character
//
//	@Summary        Create a new player character
//	@Description    Stats are random with dice rolls and base stats.
//	@Tags           Player character
//	@Accept         json
//	@Produce        json
//	@Param          request body character_request.CreateCharacter true "New character request body"
//	@Success        200  {string}  "abc-def-efg-hij"
//	@Router         /player/character/create [post]
func CreateCharacter() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body character_request.CreateCharacter
		if parseBodyErr := c.BodyParser(&body); parseBodyErr != nil {
			return parseBodyErr
		}
		if validateBodyErr := facade.Validate(c).Struct(body); validateBodyErr != nil {
			c.Status(fiber.StatusBadRequest)
			return validateBodyErr
		}
		var race race_entity.Race
		var class class_entity.Class
		if findRaceErr := facade.Dataverse(c).FindOne(race_entity.Table, body.RaceId, &race); findRaceErr != nil {
			c.Status(fiber.StatusBadRequest)
			return character_response.ErrInvalidRace
		}
		if findClassErr := facade.Dataverse(c).FindOne(class_entity.Table, body.ClassId, &class); findClassErr != nil {
			c.Status(fiber.StatusBadRequest)
			return character_response.ErrInvalidClass
		}
		var character character_entity.Character
		session := facade.Session(c).MustGet(c)
		facade.Dataverse(c).MustCreate(
			character_entity.Table,
			character_entity.Character{
				PlayerId: session.Id,
				RaceId:   body.RaceId,
				ClassId:  body.ClassId,
				Name:     body.Name,
				Level:    1,
				Strength: race.Strength() + class.Strength() + facade.Dice(c).MustRoll(
					c, character_entity.AttributeStrength,
				),
				Dexterity: race.Dexterity() + class.Dexterity() + facade.Dice(c).MustRoll(
					c, character_entity.AttributeDexterity,
				),
				Intelligence: race.Intelligence() + class.Intelligence() + facade.Dice(c).MustRoll(
					c, character_entity.AttributeIntelligence,
				),
				Charisma: race.Charisma() + class.Charisma() + facade.Dice(c).MustRoll(
					c, character_entity.AttributeCharisma,
				),
			},
			&character,
		)
		return c.JSON(character.Id)
	}
}
