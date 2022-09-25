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
		SetWeight(dto.Weight)
}

func (measurement *Measurement) SetId(id primitive.ObjectID) *Measurement {
	measurement.Id = id
	measurement.Date.UpdateTime()

	return measurement
}

func (measurement *Measurement) SetRandomId() *Measurement {
	measurement.Id = primitive.NewObjectID()
	measurement.Date.UpdateTime()

	return measurement
}

func (measurement *Measurement) SetArm(arm int) *Measurement {
	measurement.Arm = arm
	measurement.Date.UpdateTime()

	return measurement
}

func (measurement *Measurement) SetLeg(leg int) *Measurement {
	measurement.Leg = leg
	measurement.Date.UpdateTime()

	return measurement
}

func (measurement *Measurement) SetChest(chest int) *Measurement {
	measurement.Chest = chest
	measurement.Date.UpdateTime()

	return measurement
}

func (measurement *Measurement) SetShoulder(shoulder int) *Measurement {
	measurement.Shoulder = shoulder
	measurement.Date.UpdateTime()

	return measurement
}

func (measurement *Measurement) SetWaist(waist int) *Measurement {
	measurement.Waist = waist
	measurement.Date.UpdateTime()

	return measurement
}

func (measurement *Measurement) SetHip(hip int) *Measurement {
	measurement.Hip = hip
	measurement.Date.UpdateTime()

	return measurement
}

func (measurement *Measurement) SetHeight(height int) *Measurement {
	measurement.Height = height
	measurement.Date.UpdateTime()

	return measurement
}

func (measurement *Measurement) SetWeight(weight int) *Measurement {
	measurement.Weight = weight
	measurement.Date.UpdateTime()

	return measurement
}
