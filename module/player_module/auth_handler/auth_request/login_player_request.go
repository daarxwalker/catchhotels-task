package auth_request

type LoginPlayer struct {
	Email string `json:"email" validate:"required"`
}
