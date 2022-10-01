package measurement

import (
	"deepfit/pkg"
	"deepfit/pkg/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Measurement struct {
	Id       primitive.ObjectID `bson:"_id" json:"_id"`
	Arm      int                `bson:"arm,omitempty" json:"arm"`
	Leg      int                `bson:"leg,omitempty" json:"leg"`
	Chest    int                `bson:"chest,omitempty" json:"chest"`
	Shoulder int                `bson:"shoulder,omitempty" json:"shoulder"`
	Waist    int                `bson:"waist,omitempty" json:"waist"`
	Hip      int                `bson:"hip,omitempty" json:"hip"`
	Height   int                `bson:"height,omitempty" json:"height"`
	Weight   int                `bson:"weight,omitempty" json:"weight"`
	IsPublic bool               `bson:"is-public,omitempty" json:"is-public"`
	Images   []pkg.Image        `bson:"images,omitempty" json:"images"`
	Date     pkg.Date           `bson:"time" json:"time"`
}

func NewMeasurement(dto dto.MeasurementRequest) *Measurement {

	return new(Measurement).
		SetRandomId().
		SetArm(dto.Arm).
		SetLeg(dto.Leg).
		SetChest(dto.Chest).
		SetShoulder(dto.Shoulder).
		SetWaist(dto.Waist).
		SetHip(dto.Hip).
		SetHeight(dto.Height).
		SetWeight(dto.Weight).
		SetIsPublic(false).
		SetDate()

}

func (measurement *Measurement) SetId(id primitive.ObjectID) *Measurement {
	measurement.Id = id

	return measurement
}

func (measurement *Measurement) SetRandomId() *Measurement {
	measurement.Id = primitive.NewObjectID()

	return measurement
}

func (measurement *Measurement) SetArm(arm int) *Measurement {
	measurement.Arm = arm

	return measurement
}

func (measurement *Measurement) SetLeg(leg int) *Measurement {
	measurement.Leg = leg

	return measurement
}

func (measurement *Measurement) SetChest(chest int) *Measurement {
	measurement.Chest = chest

	return measurement
}

func (measurement *Measurement) SetShoulder(shoulder int) *Measurement {
	measurement.Shoulder = shoulder

	return measurement
}

func (measurement *Measurement) SetWaist(waist int) *Measurement {
	measurement.Waist = waist

	return measurement
}

func (measurement *Measurement) SetHip(hip int) *Measurement {
	measurement.Hip = hip

	return measurement
}

func (measurement *Measurement) SetIsPublic(isPublic bool) *Measurement {
	measurement.IsPublic = isPublic

	return measurement
}

func (measurement *Measurement) SetHeight(height int) *Measurement {
	measurement.Height = height

	return measurement
}

func (measurement *Measurement) SetWeight(weight int) *Measurement {
	measurement.Weight = weight

	return measurement
}

func (measurement *Measurement) SetDate() *Measurement {
	measurement.Date = *pkg.New()

	return measurement
}

func (measurement *Measurement) SetImages(images []pkg.Image) *Measurement {
	measurement.Images = images

	return measurement
}
