package dto

import (
	"deepfit/tools/ozzo_validation"
	validation "github.com/go-ozzo/ozzo-validation"
)

type VerifyPhoneNumberRequest struct {
	VerificationCode int `json:"verification_code"`
}

func (request *VerifyPhoneNumberRequest) Validate() error {
	return validation.ValidateStruct(request,
		validation.Field(&request.VerificationCode, ozzo_validation.VerificationCode...),
	)
}
