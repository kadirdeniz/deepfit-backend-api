package ozzo_validation

import (
	"deepfit/constants"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

var Surname = []validation.Rule{
	validation.Required.Error(constants.REQUIRED_FIELD),
	validation.Length(1, 25).Error(constants.WRONG_LENGHT),
	is.Alpha.Error(constants.WRONG_FORMAT),
}
