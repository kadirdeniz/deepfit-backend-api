package measurement

import (
	"deepfit/internal/user"
	"deepfit/pkg/dto"
	"github.com/thoas/go-funk"
)

type MeasurementService struct{}

type IMeasurementService interface {
	Create(dto dto.MeasurementRequest, user *user.User) *user.User
	Update(dto dto.MeasurementRequest, user *user.User) *user.User
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
	})

	if measurement == nil {
		panic("Measurement not found")
	}

	measurement.(*Measurement).
		SetArm(dto.Arm).
		SetLeg(dto.Leg).
		SetChest(dto.Chest).
		SetShoulder(dto.Shoulder).
		SetWaist(dto.Waist).
		SetHip(dto.Hip).
		SetHeight(dto.Height).
		SetWeight(dto.Weight)

	measurement.(*Measurement).Date.UpdateTime()
	user.Date.UpdateTime()

	return user
}
