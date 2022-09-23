package pkg

import (
	"deepfit/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

type Image struct {
	Id        primitive.ObjectID `bson:"_id" json:"_id"`
	Thumbnail string             `bson:"thumbnail" json:"thumbnail"`
	Original  string             `bson:"original" json:"original"`
	Large     string             `bson:"large" json:"large"`
}

func (image *Image) IsDefaultProfilePhoto() bool {
	return image.GetImageName() == constants.DEFAULT_PROFILE_PHOTO
}

func (image *Image) GetImageName() string {
	splittedImage := strings.Split(image.Thumbnail, "/")
	return splittedImage[len(splittedImage)-1]
}

func (image *Image) GetImageThumbnail() string {
	return image.Thumbnail
}

func (image *Image) GetImageOriginal() string {
	return image.Original
}

func (image *Image) GetImageLarge() string {
	return image.Large
}

func (image *Image) GetImageId() primitive.ObjectID {
	return image.Id
}

func (image *Image) SetImageThumbnail(thumbnail string) {
	image.Thumbnail = thumbnail
}

func (image *Image) SetImageOriginal(original string) {
	image.Original = original
}

func (image *Image) SetImageLarge(large string) {
	image.Large = large
}

func (image *Image) SetImageId(imageId primitive.ObjectID) {
	image.Id = imageId
}
