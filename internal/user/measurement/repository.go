package measurement

type MeasurementRepository struct{}

type IMeasurementRepository interface{}

func NewMeasurementRepository() *MeasurementRepository {
	return &MeasurementRepository{}
}
