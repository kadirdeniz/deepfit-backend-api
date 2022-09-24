package ozzo_validation

import (
	"deepfit/constants"
	validation "github.com/go-ozzo/ozzo-validation"
)

var VerificationCode = []validation.Rule{
	validation.Required.Error(constants.REQUIRED_FIELD),
	validation.Min(10000).Error(constants.WRONG_LENGHT),
	validation.Max(99999).Error(constants.WRONG_LENGHT),
}
