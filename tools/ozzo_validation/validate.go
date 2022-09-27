package ozzo_validation

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Validation interface {
	Validate() error
}

func Validate(dto interface{}) {
	if validationErr := dto.(Validation).Validate(); validationErr != nil {
		panic(validation.NewInternalError(validationErr))
	}
}
