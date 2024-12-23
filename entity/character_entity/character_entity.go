package character_entity

import "time"

type Character struct {
	CreatedAt    time.Time `json:"createdon"`
	Id           string    `json:"cr568_characterid"`
	PlayerId     string    `json:"_cr568_player_value"`
	RaceId       string    `json:"_cr568_raceid_value"`
	ClassId      string    `json:"_cr568_classid_value"`
	Name         string    `json:"cr568_charactername"`
	Level        uint8     `json:"cr568_level"`
	Strength     uint8     `json:"cr568_strength"`
	Dexterity    uint8     `json:"cr568_dexterity"`
	Intelligence uint8     `json:"cr568_intelligence"`
	Charisma     uint8     `json:"cr568_charisma"`
}

const (
	Table = "cr568_characters"
)

const (
	AttributeStrength     = "strength"
	AttributeDexterity    = "dexterity"
	AttributeIntelligence = "intelligence"
	AttributeCharisma     = "charisma"
)
