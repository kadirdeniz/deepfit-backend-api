package pkg

import (
	"deepfit/constants"
)

type Image string

func NewImage(imageName string) Image {

	if imageName == "" {
		imageName = constants.DEFAULT_IMAGE
	}

	return Image(imageName)
}

func (image *Image) GetImageName() string {
	return string(*image)
}
