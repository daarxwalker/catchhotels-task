package auth_response

import "errors"

var (
	ErrPlayerExists  = errors.New("player already exists")
	ErrInvalidPlayer = errors.New("invalid player")
	ErrInvalidCin    = errors.New("invalid cin")
)
