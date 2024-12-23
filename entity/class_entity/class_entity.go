package class_entity

import "time"

type Class struct {
	CreatedAt        time.Time `json:"createdon"`
	Id               string    `json:"cr568_classid"`
	Name             string    `json:"cr568_classname"`
	Description      string    `json:"cr568_description"`
	BaseStrength     *uint8    `json:"cr568_strengthbase"`
	BaseDexterity    *uint8    `json:"cr568_dexteritybase"`
	BaseIntelligence *uint8    `json:"cr568_intelligencebase"`
	BaseCharisma     *uint8    `json:"cr568_charismabase"`
	_                [4]byte
}

const (
	Table = "cr568_classes"
)

func (c Class) Strength() uint8 {
	if c.BaseStrength == nil {
		return 0
	}
	return *c.BaseStrength
}

func (c Class) Dexterity() uint8 {
	if c.BaseDexterity == nil {
		return 0
	}
	return *c.BaseDexterity
}

func (c Class) Intelligence() uint8 {
	if c.BaseIntelligence == nil {
		return 0
	}
	return *c.BaseIntelligence
}

func (c Class) Charisma() uint8 {
	if c.BaseCharisma == nil {
		return 0
	}
	return *c.BaseCharisma
}
