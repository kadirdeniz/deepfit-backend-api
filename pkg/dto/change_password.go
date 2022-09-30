package dto

import (
	"deepfit/constants"
	"deepfit/tools/ozzo_validation"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
)

type ChangePasswordRequest struct {
	OldPassword        string `json:"old_password"`
	NewPassword        string `json:"new_password"`
	ConfirmNewPassword string `json:"confirm_new_password"`
}

func (request *ChangePasswordRequest) Validate() {
	err := validation.ValidateStruct(request,
		validation.Field(&request.OldPassword, ozzo_validation.Password...),
		validation.Field(&request.NewPassword, ozzo_validation.Password...),
		validation.Field(&request.ConfirmNewPassword, ozzo_validation.Password...),
		validation.Field(&request.NewPassword, validation.By(func(value interface{}) error {
			if request.NewPassword != request.ConfirmNewPassword {
				return errors.New(constants.PASSWORDS_NOT_MATCH)
			}
			return nil
		})),
	)

	if err != nil {
		panic(
			validation.NewInternalError(err),
		)
	}
}
