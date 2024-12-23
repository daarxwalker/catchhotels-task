package auth_request

type RegisterPlayer struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Cin       string `json:"cin" validate:"required"`
}
