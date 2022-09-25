package handler

import (
	"deepfit/constants"
	"deepfit/internal/user"
	"deepfit/internal/user/measurement"
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
