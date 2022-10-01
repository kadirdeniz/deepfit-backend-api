package dto

import (
	"deepfit/tools/ozzo_validation"
	validation "github.com/go-ozzo/ozzo-validation"
)

type LoginRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (request *LoginRequest) Validate() {
	err := validation.ValidateStruct(request,
		validation.Field(&request.Phone, ozzo_validation.Phone...),
		validation.Field(&request.Password, ozzo_validation.Password...),
	)

	if err != nil {
		panic(
			validation.NewInternalError(err),
		)
	}
}
