package dto

import (
	"deepfit/tools/ozzo_validation"
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MeasurementRequest struct {
	Id       primitive.ObjectID `json:"_id"`
	Arm      int                `json:"arm"`
	Leg      int                `json:"leg"`
	Chest    int                `json:"chest"`
	Shoulder int                `json:"shoulder"`
	Waist    int                `json:"waist"`
	Hip      int                `json:"hip"`
	Height   int                `json:"height"`
	Weight   int                `json:"weight"`
	IsPublic bool               `json:"is-public"`
}

func (request *MeasurementRequest) Validate() {
	err := validation.ValidateStruct(request,
		validation.Field(&request.Arm, ozzo_validation.Arm...),
		validation.Field(&request.Chest, ozzo_validation.Chest...),
		validation.Field(&request.Leg, ozzo_validation.Leg...),
		validation.Field(&request.Hip, ozzo_validation.Hip...),
		validation.Field(&request.Shoulder, ozzo_validation.Shoulder...),
		validation.Field(&request.Waist, ozzo_validation.Waist...),
		validation.Field(&request.Weight, ozzo_validation.Weight...),
		validation.Field(&request.Height, ozzo_validation.Height...),
	)

	if err != nil {
		panic(
			validation.NewInternalError(err),
		)
	}
}
