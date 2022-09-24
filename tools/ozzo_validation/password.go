package ozzo_validation

import (
	"deepfit/constants"
	validation "github.com/go-ozzo/ozzo-validation"
)

var Password = []validation.Rule{
	validation.Required.Error(constants.REQUIRED_FIELD),
	validation.Length(8, 16).Error(constants.WRONG_LENGHT),
}
