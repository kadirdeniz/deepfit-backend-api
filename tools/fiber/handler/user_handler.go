package handler

import (
	"deepfit/constants"
	"deepfit/internal/user"
	"deepfit/pkg/dto"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateInterestsHandler(c *fiber.Ctx) error {
	updateInterestsDto := new(dto.UpdateInterestRequest)

	encodeError := json.Unmarshal(c.Body(), &updateInterestsDto)
	if encodeError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.General{
			Status:  false,
			Message: constants.BAD_REQUEST,
			Data:    nil,
		})
	}

	validationError := updateInterestsDto.Validate()
	if validationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.General{
			Status:  false,
			Message: constants.VALIDATION_ERROR,
			Data:    validationError,
		})
	}

	userId := c.Locals("userId").(primitive.ObjectID)

	if updateInterestsError := user.NewUserService().UpdateInterests(*updateInterestsDto, userId); updateInterestsError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.General{
			Status:  false,
			Message: updateInterestsError.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.General{
		Status:  true,
		Message: constants.UPDATE_INTERESTS_SUCCESS,
		Data:    nil,
	})

}
