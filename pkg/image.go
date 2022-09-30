package pkg

import (
	"deepfit/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Image struct {
	Id   primitive.ObjectID `bson:"_id" json:"_id"`
	Name string             `bson:"name" json:"name"`
}

func NewImage(imageName, path string) Image {

	if imageName == "" {
		imageName = constants.DEFAULT_IMAGE
	}

	return Image{
		Id:   primitive.NewObjectID(),
		Name: imageName,
	}

}
