package ozzo_validation

import (
	"deepfit/constants"
	validation "github.com/go-ozzo/ozzo-validation"
)

var Interests = []validation.Rule{
	validation.Required.Error(constants.REQUIRED_FIELD),
	validation.Length(5, 25).Error(constants.WRONG_LENGHT),
}
