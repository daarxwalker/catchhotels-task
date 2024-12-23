package player_handler

import (
	"github.com/gofiber/fiber/v2"

	"catchhotels/config/dungeon_master_config"
	"catchhotels/entity/character_entity"
	"catchhotels/entity/class_entity"
	"catchhotels/entity/player_entity"
	"catchhotels/entity/race_entity"
	"catchhotels/facade"
	"catchhotels/module/dungeon_master_module/player_handler/player_response"
	"catchhotels/service/dataverse_service"
)

// FindPlayer Find all players with grouped characters
//
//	@Summary        Find all players with grouped characters
//	@Description    Players are cached
//	@Tags           Dungeon Master
//	@Accept         json
//	@Produce        json
//	@Success        200  {array}  player_response.Player
//	@Router         /dungeon-master/player [get]
//	@Security       BearerAuth
func FindPlayer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		resultPlayers := make([]player_response.Player, 0)
		playersCacheKey := facade.Config(c).GetString(dungeon_master_config.PlayersCacheKey)
		playersCacheExists := facade.Dragonfly(c).Exists(c, playersCacheKey)
		if !playersCacheExists {
			var playersResponse dataverse_service.FindResponse[player_entity.Player]
			var charactersResponse dataverse_service.FindResponse[character_entity.Character]
			var racesResponse dataverse_service.FindResponse[race_entity.Race]
			var classesResponse dataverse_service.FindResponse[class_entity.Class]
			facade.Dataverse(c).MustFind(player_entity.Table, &playersResponse)
			facade.Dataverse(c).MustFind(character_entity.Table, &charactersResponse)
			facade.Dataverse(c).MustFind(race_entity.Table, &racesResponse)
			facade.Dataverse(c).MustFind(class_entity.Table, &classesResponse)
			resultPlayers = player_response.MapPlayers(
				playersResponse.Value,
				charactersResponse.Value,
				racesResponse.Value,
				classesResponse.Value,
			)
			facade.Dragonfly(c).MustSet(
				c,
				playersCacheKey,
				resultPlayers,
				facade.Config(c).GetDuration(dungeon_master_config.PlayersCacheDuration),
			)
		}
		if playersCacheExists {
			facade.Dragonfly(c).MustGet(
				c,
				playersCacheKey,
				&resultPlayers,
			)
		}
		return c.JSON(resultPlayers)
	}
}
