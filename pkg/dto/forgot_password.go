package dto

import (
	"deepfit/tools/ozzo_validation"
	validation "github.com/go-ozzo/ozzo-validation"
)

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

func (request *ForgotPasswordRequest) Validate() {
	err := validation.ValidateStruct(request,
		validation.Field(&request.Email, ozzo_validation.Email...),
	)

	if err != nil {
		panic(
			validation.NewInternalError(err),
		)
	}
}
