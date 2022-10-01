package measurement

import (
	"deepfit/pkg"
	"deepfit/pkg/dto"
	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MeasurementService struct{}

type IMeasurementService interface {
	Create(measurements []Measurement, dto dto.MeasurementRequest) []Measurement
	Update(measurements []Measurement, dto dto.MeasurementRequest) []Measurement
	Delete(measurements []Measurement, measurementId primitive.ObjectID) []Measurement
	AddImage(measurements []Measurement, measurementId primitive.ObjectID, imageName string) []Measurement
	DeleteImage(measurements []Measurement, measurementId primitive.ObjectID, imageName string) []Measurement
	UpdateIsPublic(measurements []Measurement, measurementId primitive.ObjectID) []Measurement
}

func NewMeasurementService() IMeasurementService {
	return &MeasurementService{}
}

func (measurementService *MeasurementService) Create(measurements []Measurement, dto dto.MeasurementRequest) []Measurement {

	measurement := NewMeasurement(dto)

	measurements = append(measurements, *measurement)

	return measurements
}

func (measurementService *MeasurementService) Update(measurements []Measurement, dto dto.MeasurementRequest) []Measurement {

	measurement := funk.Find(measurements, func(measurement Measurement) bool {
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

	return measurements
}

func (measurementService *MeasurementService) Delete(measurements []Measurement, measurementId primitive.ObjectID) []Measurement {

	measurement := funk.Find(measurements, func(measurement Measurement) bool {
		return measurement.Id == measurementId
	}).(*Measurement)

	if measurement == nil {
		panic("Measurement not found")
	}

	measurement.Date.DeleteTime()

	return measurements
}

func (measurementService *MeasurementService) AddImage(measurements []Measurement, measurementId primitive.ObjectID, imageName string) []Measurement {

	measurement := funk.Find(measurements, func(measurement Measurement) bool {
		return measurement.Id == measurementId
	}).(*Measurement)

	if len(measurement.Images) >= 4 {
		panic("Maximum number of images reached")
	}

	if measurement == nil {
		panic("Measurement not found")
	}

	measurement.Images = append(measurement.Images, pkg.NewImage(imageName))
	measurement.Date.UpdateTime()

	return measurements
}

func (measurementService *MeasurementService) DeleteImage(measurements []Measurement, measurementId primitive.ObjectID, imageName string) []Measurement {

	measurement := funk.Find(measurements, func(measurement Measurement) bool {
		return measurement.Id == measurementId
	}).(*Measurement)

	if measurement == nil {
		panic("Measurement not found")
	}

	image := funk.Filter(measurement.Images, func(image pkg.Image) bool {
		return image.GetImageName() == imageName
	}).(*pkg.Image)

	if image == nil {
		panic("Image not found")
	}

	measurement.Date.UpdateTime()

	return measurements
}

func (measurementService *MeasurementService) UpdateIsPublic(measurements []Measurement, measurementId primitive.ObjectID) []Measurement {

	measurement := funk.Find(measurements, func(measurement Measurement) bool {
		return measurement.Id == measurementId
	}).(*Measurement)

	if measurement == nil {
		panic("Measurement not found")
	}

	measurement.SetIsPublic(!measurement.IsPublic)
	measurement.Date.UpdateTime()

	return measurements
}
