package measurement

import (
	"deepfit/constants"
	"deepfit/internal/user"
	"deepfit/pkg"
	"deepfit/pkg/dto"
	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MeasurementService struct{}

type IMeasurementService interface {
	Create(dto dto.MeasurementRequest, user *user.User) *user.User
	Update(dto dto.MeasurementRequest, user *user.User) *user.User
	Delete(user *user.User, measurementId primitive.ObjectID) *user.User
	AddImage(user *user.User, measurementId primitive.ObjectID, imageName string) *user.User
}

func NewMeasurementService() IMeasurementService {
	return &MeasurementService{}
}

func (measurementService *MeasurementService) Create(dto dto.MeasurementRequest, user *user.User) *user.User {

	measurement := NewMeasurement(dto)

	user.Measurements = append(user.Measurements, *measurement)
	user.Date.UpdateTime()

	return user
}

func (measurementService *MeasurementService) Update(dto dto.MeasurementRequest, user *user.User) *user.User {

	measurement := funk.Find(user.Measurements, func(measurement Measurement) bool {
		return measurement.Id == dto.Id
	}).(*Measurement)

	if measurement == nil {
		panic("Measurement not found")
	}

	measurement.
		SetArm(dto.Arm).
		SetLeg(dto.Leg).
		SetChest(dto.Chest).
		SetShoulder(dto.Shoulder).
		SetWaist(dto.Waist).
		SetHip(dto.Hip).
		SetHeight(dto.Height).
		SetWeight(dto.Weight)

	measurement.Date.UpdateTime()
	user.Date.UpdateTime()

	return user
}

func (measurementService *MeasurementService) Delete(user *user.User, measurementId primitive.ObjectID) *user.User {

	measurement := funk.Find(user.Measurements, func(measurement Measurement) bool {
		return measurement.Id == measurementId
	}).(*Measurement)

	if measurement == nil {
		panic("Measurement not found")
	}

	measurement.Date.DeleteTime()
	user.Date.UpdateTime()

	return user
}

func (measurementService *MeasurementService) AddImage(user *user.User, measurementId primitive.ObjectID, imageName string) *user.User {

	measurement := funk.Find(user.Measurements, func(measurement Measurement) bool {
		return measurement.Id == measurementId
	}).(*Measurement)

	if len(measurement.Images) >= 4 {
		panic("Maximum number of images reached")
	}

	if measurement == nil {
		panic("Measurement not found")
	}

	measurement.Images = append(measurement.Images, pkg.NewImage(imageName, constants.MEASUREMENT_IMAGE_PATH))
	measurement.Date.UpdateTime()
	user.Date.UpdateTime()

	return user

}
