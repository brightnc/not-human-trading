package validators

import (
	validators "github.com/go-playground/validator/v10"
)

type Validator interface {
	ValidateStruct(v interface{}) error
}
type validator struct {
	validator *validators.Validate
}

func New() Validator {
	return &validator{
		validator: validators.New(),
	}
}

// ValidateStruct ...
// validate a struct that has tagged `validate`
func (vld *validator) ValidateStruct(v interface{}) error {
	return vld.validator.Struct(v)
}
