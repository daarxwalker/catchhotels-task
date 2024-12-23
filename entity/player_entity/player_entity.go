package player_entity

import "time"

type Player struct {
	CreatedAt time.Time `json:"createdon,omitempty"`
	Id        string    `json:"cr568_playerid,omitempty"`
	FirstName string    `json:"cr568_firstname"`
	LastName  string    `json:"cr568_lastname"`
	Phone     string    `json:"cr568_phone"`
	Email     string    `json:"cr568_email"`
	CIN       string    `json:"cr568_ic"`
	VAT       string    `json:"cr568_vatid"`
	Address   string    `json:"cr568_address"`
}

const (
	Table = "cr568_players"
)
