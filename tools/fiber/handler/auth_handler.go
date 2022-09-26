package handler

import (
	"deepfit/constants"
	"deepfit/internal/user"
	"deepfit/pkg"
	"deepfit/pkg/dto"
	"deepfit/tools/jwt"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterHandler(c *fiber.Ctx) error {

	registerDto := new(dto.RegisterRequest)

	encodeError := json.Unmarshal(c.Body(), &registerDto)
	if encodeError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.Response{
			Status:  false,
			Message: constants.BAD_REQUEST,
			Data:    nil,
		})
	}

	validationError := registerDto.Validate()
	if validationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.Response{
			Status:  false,
			Message: constants.VALIDATION_ERROR,
			Data:    validationError,
		})
	}

	userObj, registerError := user.NewUserService().Register(*registerDto)
	if registerError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.Response{
			Status:  false,
			Message: registerError.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(pkg.Response{
		Status:  true,
		Message: constants.REGISTER_SUCCESS,
		Data: dto.RegisterResponse{
			Id:    userObj.ID,
			Token: jwt.New().SetUserId(userObj.ID).CreateToken().GetToken(),
		},
	})
}

func LoginHandler(c *fiber.Ctx) error {

	loginDto := new(dto.LoginRequest)

	encodeError := json.Unmarshal(c.Body(), &loginDto)
	if encodeError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.Response{
			Status:  false,
			Message: constants.BAD_REQUEST,
			Data:    nil,
		})
	}

	validationError := loginDto.Validate()
	if validationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.Response{
			Status:  false,
			Message: constants.VALIDATION_ERROR,
			Data:    validationError,
		})
	}

	userObj, loginError := user.NewUserService().Login(*loginDto)
	if loginError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.Response{
			Status:  false,
			Message: loginError.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(pkg.Response{
		Status:  true,
		Message: constants.LOGIN_SUCCESS,
		Data: dto.RegisterResponse{
			Id:    userObj.ID,
			Token: jwt.New().SetUserId(userObj.ID).CreateToken().GetToken(),
		},
	})
}

func VerifyPhoneNumberHandler(c *fiber.Ctx) error {

	verifyDto := new(dto.VerifyPhoneNumberRequest)

	encodeError := json.Unmarshal(c.Body(), &verifyDto)
	if encodeError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.Response{
			Status:  false,
			Message: constants.BAD_REQUEST,
			Data:    nil,
		})
	}

	validationError := verifyDto.Validate()
	if validationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.Response{
			Status:  false,
			Message: constants.VALIDATION_ERROR,
			Data:    validationError,
		})
	}

	userId := c.Locals("userId").(primitive.ObjectID)

	userObj, verifyError := user.NewUserService().VerifyPhoneNumber(*verifyDto, userId)
	if verifyError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.Response{
			Status:  false,
			Message: verifyError.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(pkg.Response{
		Status:  true,
		Message: constants.PHONE_VERIFICATION_SUCCESS,
		Data: dto.RegisterResponse{
			Id:    userObj.ID,
			Token: jwt.New().SetUserId(userObj.ID).CreateToken().GetToken(),
		},
	})
}

func ResendVerificationCodeHandler(c *fiber.Ctx) error {

	userId := c.Locals("userId").(primitive.ObjectID)

	_, resendError := user.NewUserService().ResendVerificationCode(userId)
	if resendError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.Response{
			Status:  false,
			Message: resendError.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(pkg.Response{
		Status:  true,
		Message: constants.VERIFICATION_CODE_RESENT,
		Data:    nil,
	})
}

func ChangePasswordHandler(c *fiber.Ctx) error {

	changePasswordDto := new(dto.ChangePasswordRequest)

	encodeError := json.Unmarshal(c.Body(), &changePasswordDto)
	if encodeError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.Response{
			Status:  false,
			Message: constants.BAD_REQUEST,
			Data:    nil,
		})
	}

	validationError := changePasswordDto.Validate()
	if validationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.Response{
			Status:  false,
			Message: constants.VALIDATION_ERROR,
			Data:    validationError,
		})
	}

	userId := c.Locals("userId").(primitive.ObjectID)

	userObj, changePasswordError := user.NewUserService().ChangePassword(*changePasswordDto, userId)
	if changePasswordError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.Response{
			Status:  false,
			Message: changePasswordError.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(pkg.Response{
		Status:  true,
		Message: constants.PASSWORD_CHANGE_SUCCESS,
		Data: dto.RegisterResponse{
			Id:    userObj.ID,
			Token: jwt.New().SetUserId(userObj.ID).CreateToken().GetToken(),
		},
	})
}
