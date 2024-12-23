package character_response

import "errors"

var (
	ErrInvalidCharacter = errors.New("invalid character")
	ErrInvalidRace      = errors.New("invalid race")
	ErrInvalidClass     = errors.New("invalid class")
)
