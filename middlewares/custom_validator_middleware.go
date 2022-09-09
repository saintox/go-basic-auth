package middlewares

import "github.com/go-playground/validator/v10"

type CustomValidator struct {
	Validator *validator.Validate
}

func NewValidator() *CustomValidator {
	return &CustomValidator{
		Validator: validator.New(),
	}
}

func (v *CustomValidator) Validate(i interface{}) error {
	if err := v.Validator.Struct(i); err != nil {
		return err
	}

	return nil
}
