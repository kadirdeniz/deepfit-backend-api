package handler

import (
	"deepfit/constants"
	"deepfit/internal/user"
	"deepfit/internal/user/measurement"
	"deepfit/pkg"
	"deepfit/pkg/dto"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateMeasurementHandler(c *fiber.Ctx) error {

	measurementDto := new(dto.MeasurementRequest)

	encodeError := json.Unmarshal(c.Body(), &measurementDto)
	if encodeError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.General{
			Status:  false,
			Message: constants.BAD_REQUEST,
			Data:    nil,
		})
	}

	validationError := measurementDto.Validate()
	if validationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.General{
			Status:  false,
			Message: constants.VALIDATION_ERROR,
			Data:    validationError,
		})
	}

	userId := c.Locals("userId").(primitive.ObjectID)

	userObj, getUserErr := user.NewRepository().GetUserById(userId)
	if getUserErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.General{
			Status:  false,
			Message: getUserErr.Error(),
			Data:    nil,
		})
	}

	userObj = measurement.NewMeasurementService().Create(*measurementDto, userObj)

	if upsertErr := user.NewRepository().Upsert(userObj); upsertErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.General{
			Status:  false,
			Message: upsertErr.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.General{
		Status:  true,
		Message: constants.CREATE_MEASUREMENT_SUCCESS,
		Data:    nil,
	})
}

func UpdateMeasurementHandler(c *fiber.Ctx) error {

	measurementDto := new(dto.MeasurementRequest)

	encodeError := json.Unmarshal(c.Body(), &measurementDto)
	if encodeError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.General{
			Status:  false,
			Message: constants.BAD_REQUEST,
			Data:    nil,
		})
	}

	validationError := measurementDto.Validate()
	if validationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.General{
			Status:  false,
			Message: constants.VALIDATION_ERROR,
			Data:    validationError,
		})
	}

	userId := c.Locals("userId").(primitive.ObjectID)
	measurementDto.Id, _ = primitive.ObjectIDFromHex(c.Params("measurement_id"))

	userObj, getUserErr := user.NewRepository().GetUserById(userId)
	if getUserErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.General{
			Status:  false,
			Message: getUserErr.Error(),
			Data:    nil,
		})
	}

	userObj = measurement.NewMeasurementService().Update(*measurementDto, userObj)

	if upsertErr := user.NewRepository().Upsert(userObj); upsertErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.General{
			Status:  false,
			Message: upsertErr.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.General{
		Status:  true,
		Message: constants.UPDATE_MEASUREMENT_SUCCESS,
		Data:    nil,
	})
}

func DeleteMeasurementHandler(c *fiber.Ctx) error {

	userId := c.Locals("userId").(primitive.ObjectID)
	measurementId, _ := primitive.ObjectIDFromHex(c.Params("measurement_id"))

	userObj, getUserErr := user.NewRepository().GetUserById(userId)
	if getUserErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.General{
			Status:  false,
			Message: getUserErr.Error(),
			Data:    nil,
		})
	}

	userObj = measurement.NewMeasurementService().Delete(userObj, measurementId)

	if upsertErr := user.NewRepository().Upsert(userObj); upsertErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.General{
			Status:  false,
			Message: upsertErr.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.General{
		Status:  true,
		Message: constants.DELETE_MEASUREMENT_SUCCESS,
		Data:    nil,
	})

}

func AddImageToMeasurementHandler(c *fiber.Ctx) error {

	userId := c.Locals("userId").(primitive.ObjectID)
	measurementId, _ := primitive.ObjectIDFromHex(c.Params("measurement_id"))

	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.General{
			Status:  false,
			Message: constants.BAD_REQUEST,
			Data:    nil,
		})
	}

	userObj, getUserErr := user.NewRepository().GetUserById(userId)
	if getUserErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.General{
			Status:  false,
			Message: getUserErr.Error(),
			Data:    nil,
		})
	}

	userObj = measurement.NewMeasurementService().AddImage(userObj, measurementId, pkg.HashImageName(file.Filename))

	if upsertErr := user.NewRepository().Upsert(userObj); upsertErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.General{
			Status:  false,
			Message: upsertErr.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.General{
		Status:  true,
		Message: constants.ADD_IMAGE_TO_MEASUREMENT_SUCCESS,
		Data:    nil,
	})
}
