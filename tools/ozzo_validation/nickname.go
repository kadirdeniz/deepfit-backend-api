package ozzo_validation

import (
	"deepfit/constants"
	"deepfit/pkg"
	validation "github.com/go-ozzo/ozzo-validation"
	"regexp"
)

var Nickname = []validation.Rule{
	validation.Required.Error(constants.REQUIRED_FIELD),
	validation.Length(1, 25).Error(constants.WRONG_LENGHT),
	validation.Match(regexp.MustCompile(pkg.AlphaNumericRegex)).Error(constants.WRONG_FORMAT),
}
