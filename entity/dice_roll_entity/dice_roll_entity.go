package dice_roll_entity

import "time"

type DiceRoll struct {
	CreatedAt   time.Time `json:"createdon"`
	RollAt      time.Time `json:"cr568_rolldate"`
	Id          string    `json:"cr568_dicerollid"`
	CharacterId string    `json:"_cr568_characterid_value"`
	Name        string    `json:"cr568_diceroll1"`
	Result      uint8     `json:"cr568_rollresult"`
}

const (
	Table = "cr568_dicerolls"
)
