package validator

import (
	validator "github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func ValidatorInit() {
	Validate = validator.New()
}
