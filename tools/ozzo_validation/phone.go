package ozzo_validation

import (
	"deepfit/constants"
	validation "github.com/go-ozzo/ozzo-validation"
)

var Phone = []validation.Rule{
	validation.Required.Error(constants.REQUIRED_FIELD),
}
