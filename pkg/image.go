package pkg

import (
	"deepfit/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Image struct {
	Id        primitive.ObjectID `bson:"_id" json:"_id"`
	Thumbnail string             `bson:"thumbnail" json:"thumbnail"`
	Original  string             `bson:"original" json:"original"`
	Large     string             `bson:"large" json:"large"`
}

func NewImage(imageName, path string) Image {

	if imageName == "" {
		imageName = constants.DEFAULT_IMAGE
	}

	return Image{
		Id:        primitive.NewObjectID(),
		Thumbnail: path + constants.THUMBNAIL + imageName,
		Original:  path + constants.ORIGINAL + imageName,
		Large:     path + constants.LARGE + imageName,
	}

}
