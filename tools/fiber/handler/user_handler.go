package handler

import (
	"deepfit/constants"
	"deepfit/internal/user"
	"deepfit/pkg"
	"deepfit/pkg/dto"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateInterestsHandler(c *fiber.Ctx) error {
	updateInterestsDto := new(dto.UpdateInterestRequest)

	pkg.JSONEncoder(c.Body(), &updateInterestsDto)
	updateInterestsDto.Validate()

	userId := c.Locals("userId").(primitive.ObjectID)

	repository := user.NewRepository()

	repository.Upsert(
		user.NewUserService().UpdateInterests(
			repository.GetUserById(userId),
			*updateInterestsDto,
		),
	)

	return c.Status(fiber.StatusOK).JSON(
		pkg.NewResponse(true, constants.UPDATE_INTERESTS_SUCCESS, nil),
	)
}

func UpdateProfilePhotoHandler(c *fiber.Ctx) error {
	userId := c.Locals("userId").(primitive.ObjectID)

	photo := ImageFileHandler(c, "profile_photo")

	repository := user.NewRepository()
	userObj := repository.GetUserById(userId)

	repository.Upsert(
		user.NewUserService().UpdateProfilePhoto(userObj, photo.Filename),
	)

	return c.Status(fiber.StatusOK).JSON(
		pkg.NewResponse(true, constants.UPDATE_PROFILE_PHOTO_SUCCESS, nil),
	)
}

func UpdateCoverPhotoHandler(c *fiber.Ctx) error {
	userId := c.Locals("userId").(primitive.ObjectID)

	photo := ImageFileHandler(c, "cover_photo")

	repository := user.NewRepository()
	userObj := repository.GetUserById(userId)

	repository.Upsert(
		user.NewUserService().UpdateCoverPhoto(userObj, photo.Filename),
	)

	return c.Status(fiber.StatusOK).JSON(
		pkg.NewResponse(true, constants.UPDATE_COVER_PHOTO_SUCCESS, nil),
	)
}
