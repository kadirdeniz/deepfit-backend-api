package ozzo_validation

import (
	"deepfit/constants"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

var Email = []validation.Rule{
	validation.Required.Error(constants.REQUIRED_FIELD),
	validation.Length(1, 50).Error(constants.WRONG_LENGHT),
	is.Email.Error(constants.WRONG_FORMAT),
}
