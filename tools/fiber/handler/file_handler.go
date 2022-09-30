package handler

import (
	"deepfit/constants"
	"deepfit/pkg"
	"github.com/gofiber/fiber/v2"
	"mime/multipart"
)

type fiberCTX fiber.Ctx

func ImageFileHandler(c *fiber.Ctx, image_path string) *multipart.FileHeader {
	photo, photoError := c.FormFile(image_path)
	if photoError != nil {
		panic(
			pkg.NewError(constants.StatusBadRequest, constants.BAD_REQUEST, photoError),
		)
	}

	return photo
}
