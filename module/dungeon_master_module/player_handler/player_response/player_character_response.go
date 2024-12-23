package player_response

type Character struct {
	Name         string `json:"name"`
	Class        string `json:"class"`
	Race         string `json:"race"`
	Level        uint8  `json:"level"`
	Strength     uint8  `json:"strength"`
	Dexterity    uint8  `json:"dexterity"`
	Intelligence uint8  `json:"intelligence"`
	Charisma     uint8  `json:"charisma"`
}
