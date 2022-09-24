package measurement

import (
	"deepfit/pkg"
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
