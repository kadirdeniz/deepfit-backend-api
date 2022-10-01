package dto

import (
	"deepfit/tools/ozzo_validation"
	validation "github.com/go-ozzo/ozzo-validation"
)

type ResendVerificationCodeRequest struct {
	PhoneNumber string `json:"phone_number"`
}

func (request *ResendVerificationCodeRequest) Validate() {
	err := validation.ValidateStruct(request,
		validation.Field(&request.PhoneNumber, ozzo_validation.Phone...),
	)

	if err != nil {
		panic(
			validation.NewInternalError(err),
		)
	}
}
