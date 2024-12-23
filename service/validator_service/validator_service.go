package validator_service

import (
	"github.com/go-playground/validator/v10"
)

type ValidatorService struct {
	validate *validator.Validate
}

const (
	Token = "validator_service"
)

func New() *ValidatorService {
	return &ValidatorService{
		validate: validator.New(),
	}
}

func (s *ValidatorService) Validate() *validator.Validate {
	return s.validate
}
