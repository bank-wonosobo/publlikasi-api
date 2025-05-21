package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Validator interface {
	Validate(interface{}) error
}

type validatorImpl struct {
	validator *validator.Validate
}

func NewValidator() Validator {
	return &validatorImpl{
		validator: validator.New(),
	}
}

// Validate implements Validator.
func (v *validatorImpl) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			tag := err.Tag()
			errors = append(errors, fmt.Sprintf("%s: failed on the '%s' tag", field, tag))
		}
		return fmt.Errorf("%s", strings.Join(errors, ", "))
	}
	return nil
}
