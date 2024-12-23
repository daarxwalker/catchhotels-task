package race_entity

import "time"

type Race struct {
	CreatedAt        time.Time `json:"createdon"`
	Id               string    `json:"cr568_raceid"`
	Name             string    `json:"cr568_racename"`
	Description      string    `json:"cr568_description"`
	BaseStrength     *uint8    `json:"cr568_strengthbase"`
	BaseDexterity    *uint8    `json:"cr568_dexteritybase"`
	BaseIntelligence *uint8    `json:"cr568_intelligencebase"`
	BaseCharisma     *uint8    `json:"cr568_charismabase"`
	_                [4]byte
}

const (
	Table = "cr568_races"
)

func (r Race) Strength() uint8 {
	if r.BaseStrength == nil {
		return 0
	}
	return *r.BaseStrength
}

func (r Race) Dexterity() uint8 {
	if r.BaseDexterity == nil {
		return 0
	}
	return *r.BaseDexterity
}

func (r Race) Intelligence() uint8 {
	if r.BaseIntelligence == nil {
		return 0
	}
	return *r.BaseIntelligence
}

func (r Race) Charisma() uint8 {
	if r.BaseCharisma == nil {
		return 0
	}
	return *r.BaseCharisma
}
