package dto

import (
	"deepfit/tools/ozzo_validation"
	validation "github.com/go-ozzo/ozzo-validation"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (request *LoginRequest) Validate() {
	err := validation.ValidateStruct(request,
		validation.Field(&request.Email, ozzo_validation.Email...),
		validation.Field(&request.Password, ozzo_validation.Password...),
	)

	if err != nil {
		panic(
			validation.NewInternalError(err),
		)
	}
}
