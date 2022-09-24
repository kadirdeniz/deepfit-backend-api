package handler

import (
	"deepfit/constants"
	"deepfit/internal/user"
	"deepfit/pkg/dto"
	"deepfit/tools/jwt"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func RegisterHandler(c *fiber.Ctx) error {

	registerDto := new(dto.RegisterRequest)

	encodeError := json.Unmarshal(c.Body(), &registerDto)
	if encodeError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.General{
			Status:  false,
			Message: constants.BAD_REQUEST,
			Data:    nil,
		})
	}

	validationError := registerDto.Validate()
	if validationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.General{
			Status:  false,
			Message: constants.VALIDATION_ERROR,
			Data:    validationError,
		})
	}

	userObj, registerError := user.NewUserService().Register(*registerDto)
	if registerError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.General{
			Status:  false,
			Message: registerError.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.General{
		Status:  true,
		Message: constants.REGISTER_SUCCESS,
		Data: dto.RegisterResponse{
			Id:    userObj.ID,
			Token: jwt.New().SetUserId(userObj.ID).CreateToken().GetToken(),
		},
	})
}

func VerifyPhoneNumberHandler(c *fiber.Ctx) error {

	verifyDto := new(dto.VerifyRequest)

	encodeError := json.Unmarshal(c.Body(), &verifyDto)
	if encodeError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.General{
			Status:  false,
			Message: constants.BAD_REQUEST,
			Data:    nil,
		})
	}

	validationError := verifyDto.Validate()
	if validationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.General{
			Status:  false,
			Message: constants.VALIDATION_ERROR,
			Data:    validationError,
		})
	}

	userObj, verifyError := user.NewUserService().Verify(*verifyDto)
	if verifyError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.General{
			Status:  false,
			Message: verifyError.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.General{
		Status:  true,
		Message: constants.VERIFY_SUCCESS,
		Data: dto.RegisterResponse{
			Id:    userObj.ID,
			Token: jwt.New().SetUserId(userObj.ID).CreateToken().GetToken(),
		},
	})
}
