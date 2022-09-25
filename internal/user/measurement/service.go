package measurement

import (
	"deepfit/internal/user"
	"deepfit/pkg/dto"
)

type MeasurementService struct {
	repository IMeasurementRepository
}

type IMeasurementService interface {
	Create(dto dto.MeasurementRequest, user *user.User) *user.User
}

func NewMeasurementService() IMeasurementService {
	return &MeasurementService{
		repository: NewMeasurementRepository(),
	}
}

func (measurementService *MeasurementService) Create(dto dto.MeasurementRequest, user *user.User) *user.User {

	measurement := NewMeasurement(dto)

	user.Measurements = append(user.Measurements, measurement)

	return user
}
