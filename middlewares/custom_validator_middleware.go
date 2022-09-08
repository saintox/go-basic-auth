package middlewares

import "github.com/go-playground/validator/v10"

type Validator struct {
	Validator *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		Validator: validator.New(),
	}
}

func (v *Validator) Validate(i interface{}) error {
	if err := v.Validator.Struct(i); err != nil {
		return err
	}

	return nil
}