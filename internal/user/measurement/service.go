package measurement

import (
	"deepfit/configs"
	"deepfit/constants"
	"deepfit/pkg"
	"deepfit/pkg/dto"
	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MeasurementService struct{}

type IMeasurementService interface {
	Create(measurements []*Measurement, dto dto.MeasurementRequest) []*Measurement
	Update(measurements []*Measurement, dto dto.MeasurementRequest) []*Measurement
	Delete(measurements []*Measurement, measurementId primitive.ObjectID) []*Measurement
	AddImage(measurements []*Measurement, measurementId primitive.ObjectID, imageName string) []*Measurement
	DeleteImage(measurements []*Measurement, measurementId primitive.ObjectID, imageName string) []*Measurement
	UpdateIsPublic(measurements []*Measurement, measurementId primitive.ObjectID) []*Measurement
}

func NewMeasurementService() IMeasurementService {
	return &MeasurementService{}
}

func getMeasurement(measurements []*Measurement, measurementId primitive.ObjectID) *Measurement {
	measurementObj := funk.Find(measurements, func(measurement *Measurement) bool {
		return measurement.Id == measurementId && !measurement.Date.IsDeleted()
	})

	if measurementObj == nil {
		panic(
			pkg.NewError(constants.StatusBadRequest, constants.MEASUREMENT_NOT_FOUND, nil),
		)
	}

	return measurementObj.(*Measurement)
}

func (measurementService *MeasurementService) Create(measurements []*Measurement, dto dto.MeasurementRequest) []*Measurement {

	measurement := NewMeasurement(dto)

	measurements = append(measurements, measurement)

	return measurements
}

func (measurementService *MeasurementService) Update(measurements []*Measurement, dto dto.MeasurementRequest) []*Measurement {

	measurement := getMeasurement(measurements, dto.Id)

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

func (measurementService *MeasurementService) Delete(measurements []*Measurement, measurementId primitive.ObjectID) []*Measurement {
	getMeasurement(measurements, measurementId).Date.DeleteTime()
	return measurements
}

func (measurementService *MeasurementService) AddImage(measurements []*Measurement, measurementId primitive.ObjectID, imageName string) []*Measurement {

	measurement := getMeasurement(measurements, measurementId)

	if len(measurement.Images) >= configs.MAX_MEASUREMENT_IMAGE_COUNT {
		panic(pkg.NewError(constants.StatusBadRequest, constants.MAX_MEASUREMENT_IMAGE_COUNT, nil))
	}

	measurement.Images = append(measurement.Images, pkg.NewImage(imageName))
	measurement.Date.UpdateTime()

	return measurements
}

func (measurementService *MeasurementService) DeleteImage(measurements []*Measurement, measurementId primitive.ObjectID, imageName string) []*Measurement {

	measurement := getMeasurement(measurements, measurementId)
	if isContains := funk.Contains(measurement.Images, func(image pkg.Image) bool {
		return image.GetImageName() == imageName
	}); !isContains {
		panic(pkg.NewError(constants.StatusBadRequest, constants.IMAGE_NOT_FOUND, nil))
	}

	images := funk.Filter(measurement.Images, func(image pkg.Image) bool {
		return image.GetImageName() != imageName
	})

	measurement.SetImages(images.([]pkg.Image))
	measurement.Date.UpdateTime()

	return measurements
}

func (measurementService *MeasurementService) UpdateIsPublic(measurements []*Measurement, measurementId primitive.ObjectID) []*Measurement {

	measurementObj := funk.Find(measurements, func(measurement *Measurement) bool {
		return measurement.Id == measurementId && !measurement.Date.IsDeleted()
	})

	if measurementObj == nil {
		panic(
			pkg.NewError(constants.StatusBadRequest, constants.MEASUREMENT_NOT_FOUND, nil),
		)
	}

	measurement := measurementObj.(*Measurement)

	measurement.SetIsPublic(!measurement.IsPublic)
	measurement.Date.UpdateTime()

	return measurements
}
