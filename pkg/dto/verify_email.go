package dto

import (
	"deepfit/tools/ozzo_validation"
	validation "github.com/go-ozzo/ozzo-validation"
)

type VerifyEmailRequest struct {
	VerificationCode int `json:"verification_code"`
}

func (request *VerifyEmailRequest) Validate() {
	err := validation.ValidateStruct(request,
		validation.Field(&request.VerificationCode, ozzo_validation.VerificationCode...),
	)

	if err != nil {
		panic(
			validation.NewInternalError(err),
		)
	}
}
