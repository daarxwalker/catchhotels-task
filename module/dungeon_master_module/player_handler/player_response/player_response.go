package player_response

import (
	"github.com/samber/lo"

	"catchhotels/entity/character_entity"
	"catchhotels/entity/class_entity"
	"catchhotels/entity/player_entity"
	"catchhotels/entity/race_entity"
)

type Player struct {
	Id         string      `json:"id"`
	Email      string      `json:"email"`
	Characters []Character `json:"characters"`
}

func MapPlayers(
	players []player_entity.Player,
	characters []character_entity.Character,
	races []race_entity.Race,
	classes []class_entity.Class,
) []Player {
	return lo.Map(
		players,
		func(player player_entity.Player, index int) Player {
			return Player{
				Id:    player.Id,
				Email: player.Email,
				Characters: lo.Map(
					lo.Filter(
						characters,
						func(character character_entity.Character, _ int) bool {
							return character.PlayerId == player.Id
						},
					),
					func(character character_entity.Character, index int) Character {
						return Character{
							Name:         character.Name,
							Class:        findClassName(character.ClassId, classes),
							Race:         findRaceName(character.RaceId, races),
							Level:        character.Level,
							Strength:     character.Strength,
							Dexterity:    character.Dexterity,
							Intelligence: character.Intelligence,
							Charisma:     character.Charisma,
						}
					},
				),
			}
		},
	)
}

func findRaceName(raceId string, races []race_entity.Race) string {
	race, raceExists := lo.Find(
		races, func(race race_entity.Race) bool {
			return race.Id == raceId
		},
	)
	if !raceExists {
		return ""
	}
	return race.Name
}

func findClassName(classId string, classes []class_entity.Class) string {
	class, classExists := lo.Find(
		classes, func(race class_entity.Class) bool {
			return race.Id == classId
		},
	)
	if !classExists {
		return ""
	}
	return class.Name
}
