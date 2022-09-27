package handler

import (
	"deepfit/constants"
	"deepfit/internal/user"
	"deepfit/internal/user/measurement"
	"deepfit/pkg"
	"deepfit/pkg/dto"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateMeasurementHandler(c *fiber.Ctx) error {

	measurementDto := new(dto.MeasurementRequest)

	pkg.JSONEncoder(c.Body(), &measurementDto)
	measurementDto.Validate()

	userId := c.Locals("userId").(primitive.ObjectID)

	repository := user.NewRepository()
	userObj := repository.GetUserById(userId)

	user.NewRepository().Upsert(
		userObj.SetMeasurements(
			measurement.NewMeasurementService().Create(userObj.Measurements, *measurementDto),
		),
	)

	return c.Status(fiber.StatusOK).JSON(
		pkg.NewResponse(true, constants.CREATE_MEASUREMENT_SUCCESS, nil),
	)
}

func UpdateMeasurementHandler(c *fiber.Ctx) error {

	measurementDto := new(dto.MeasurementRequest)

	pkg.JSONEncoder(c.Body(), &measurementDto)
	measurementDto.Validate()

	userId := c.Locals("userId").(primitive.ObjectID)
	measurementDto.Id, _ = primitive.ObjectIDFromHex(c.Params("measurement_id"))

	// TODO: Validate measurementDto
	userObj, getUserErr := user.NewRepository().GetUserById(userId)
	if getUserErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.Response{
			Status:  false,
			Message: getUserErr.Error(),
			Data:    nil,
		})
	}

	user.NewRepository().Upsert(
		userObj.SetMeasurements(
			measurement.NewMeasurementService().Update(userObj.Measurements, *measurementDto),
		),
	)

	return c.Status(fiber.StatusOK).JSON(
		pkg.NewResponse(true, constants.UPDATE_MEASUREMENT_SUCCESS, nil),
	)
}

func DeleteMeasurementHandler(c *fiber.Ctx) error {

	userId := c.Locals("userId").(primitive.ObjectID)
	measurementId, _ := primitive.ObjectIDFromHex(c.Params("measurement_id"))

	userObj, getUserErr := user.NewRepository().GetUserById(userId)
	if getUserErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.Response{
			Status:  false,
			Message: getUserErr.Error(),
			Data:    nil,
		})
	}

	user.NewRepository().Upsert(
		userObj.SetMeasurements(
			measurement.NewMeasurementService().Delete(userObj.Measurements, measurementId),
		),
	)

	return c.Status(fiber.StatusOK).JSON(
		pkg.NewResponse(true, constants.DELETE_MEASUREMENT_SUCCESS, nil),
	)

}

func AddImageToMeasurementHandler(c *fiber.Ctx) error {

	userId := c.Locals("userId").(primitive.ObjectID)
	measurementId, _ := primitive.ObjectIDFromHex(c.Params("measurement_id"))

	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.Response{
			Status:  false,
			Message: constants.BAD_REQUEST,
			Data:    nil,
		})
	}

	userObj, getUserErr := user.NewRepository().GetUserById(userId)
	if getUserErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.Response{
			Status:  false,
			Message: getUserErr.Error(),
			Data:    nil,
		})
	}

	user.NewRepository().Upsert(
		userObj.SetMeasurements(
			measurement.NewMeasurementService().AddImage(userObj.Measurements, measurementId, pkg.HashImageName(file.Filename)),
		),
	)

	return c.Status(fiber.StatusOK).JSON(
		pkg.NewResponse(true, constants.ADD_IMAGE_TO_MEASUREMENT_SUCCESS, nil),
	)

}

func DeleteImageToMeasurementHandler(c *fiber.Ctx) error {

	userId := c.Locals("userId").(primitive.ObjectID)
	measurementId, _ := primitive.ObjectIDFromHex(c.Params("measurement_id"))
	imageId, _ := primitive.ObjectIDFromHex(c.Params("image_id"))

	userObj, getUserErr := user.NewRepository().GetUserById(userId)
	if getUserErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.Response{
			Status:  false,
			Message: getUserErr.Error(),
			Data:    nil,
		})
	}

	user.NewRepository().Upsert(
		userObj.SetMeasurements(
			measurement.NewMeasurementService().DeleteImage(userObj.Measurements, measurementId, imageId),
		),
	)

	return c.Status(fiber.StatusOK).JSON(
		pkg.NewResponse(true, constants.DELETE_IMAGE_TO_MEASUREMENT_SUCCESS, nil),
	)
}

func UpdateMeasurementIsPublicHandler(c *fiber.Ctx) error {

	userId := c.Locals("userId").(primitive.ObjectID)
	measurementId, _ := primitive.ObjectIDFromHex(c.Params("measurement_id"))

	userObj, getUserErr := user.NewRepository().GetUserById(userId)
	if getUserErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.Response{
			Status:  false,
			Message: getUserErr.Error(),
			Data:    nil,
		})
	}

	user.NewRepository().Upsert(
		userObj.SetMeasurements(
			measurement.NewMeasurementService().UpdateIsPublic(userObj.Measurements, measurementId),
		),
	)

	return c.Status(fiber.StatusOK).JSON(
		pkg.NewResponse(true, constants.UPDATE_MEASUREMENT_IS_PUBLIC_SUCCESS, nil),
	)
}
