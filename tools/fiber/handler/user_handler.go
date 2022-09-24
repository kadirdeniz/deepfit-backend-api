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

func UpdateProfilePhotoHandler(c *fiber.Ctx) error {
	userId := c.Locals("userId").(primitive.ObjectID)

	photo, photoError := c.FormFile("profile_photo")
	if photoError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.General{
			Status:  false,
			Message: constants.BAD_REQUEST,
			Data:    nil,
		})
	}

	if updateProfilePhotoError := user.NewUserService().UpdateProfilePhoto(photo.Filename, userId); updateProfilePhotoError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.General{
			Status:  false,
			Message: updateProfilePhotoError.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.General{
		Status:  true,
		Message: constants.UPDATE_PROFILE_PHOTO_SUCCESS,
		Data:    nil,
	})
}

func UpdateCoverPhotoHandler(c *fiber.Ctx) error {
	userId := c.Locals("userId").(primitive.ObjectID)

	photo, photoError := c.FormFile("cover_photo")
	if photoError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.General{
			Status:  false,
			Message: constants.BAD_REQUEST,
			Data:    nil,
		})
	}

	if updateCoverPhotoError := user.NewUserService().UpdateCoverPhoto(photo.Filename, userId); updateCoverPhotoError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.General{
			Status:  false,
			Message: updateCoverPhotoError.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.General{
		Status:  true,
		Message: constants.UPDATE_COVER_PHOTO_SUCCESS,
		Data:    nil,
	})
}
