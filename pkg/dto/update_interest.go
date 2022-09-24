package dto

import (
	"deepfit/tools/ozzo_validation"
	validation "github.com/go-ozzo/ozzo-validation"
)

type UpdateInterestRequest struct {
	Interests []string `json:"interests"`
}

func (request *UpdateInterestRequest) Validate() error {
	return validation.ValidateStruct(request,
		validation.Field(&request.Interests, ozzo_validation.Interests...),
	)
}
