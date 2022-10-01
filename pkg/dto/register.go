package dto

import (
	"deepfit/constants"
	"deepfit/tools/ozzo_validation"
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisterResponse struct {
	Id    primitive.ObjectID `json:"id"`
	Token string             `json:"token"`
}

type RegisterRequest struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Nickname   string `json:"nickname"`
	Phone      string `json:"phone"`
	Password   string `json:"password"`
	RePassword string `json:"re-password"`
}

func (request *RegisterRequest) Validate() {
	err := validation.ValidateStruct(request,
		validation.Field(&request.Name, ozzo_validation.Name...),
		validation.Field(&request.Surname, ozzo_validation.Name...),
		validation.Field(&request.Nickname, ozzo_validation.Nickname...),
		validation.Field(&request.Phone, ozzo_validation.Phone...),
		validation.Field(&request.Password, ozzo_validation.Password...),
		validation.Field(&request.RePassword, validation.In(request.Password).Error(constants.PASSWORD_DOESNT_MATCH)),
	)
	if err != nil {
		panic(
			validation.NewInternalError(err),
		)
	}
}
