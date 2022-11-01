package handler

import (
	"deepfit/constants"
	"deepfit/internal/user"
	"deepfit/pkg"
	"deepfit/pkg/dto"
	"deepfit/tools/jwt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterHandler(c *fiber.Ctx) error {

	registerDto := new(dto.RegisterRequest)

	pkg.JSONEncoder(c.Body(), &registerDto)
	registerDto.Validate()

	repository := user.NewRepository()

	if repository.IsNicknameExists(registerDto.Nickname) {
		return c.Status(fiber.StatusBadRequest).JSON(
			pkg.NewResponse(false, constants.NICKNAME_ALREADY_EXISTS, nil),
		)
	}

	if repository.IsEmailExists(registerDto.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(
			pkg.NewResponse(false, constants.EMAIL_ALREADY_EXISTS, nil),
		)
	}

	userObj := user.NewUserService().Register(*registerDto)
	repository.Upsert(userObj)

	return c.Status(fiber.StatusOK).JSON(
		pkg.NewResponse(
			true,
			constants.REGISTER_SUCCESS,
			dto.RegisterResponse{
				Token: jwt.New().SetIsVerified(false).SetUserId(userObj.ID).CreateToken().GetToken(),
			},
		),
	)
}

func LoginHandler(c *fiber.Ctx) error {

	loginDto := new(dto.LoginRequest)

	pkg.JSONEncoder(c.Body(), &loginDto)
	loginDto.Validate()

	repository := user.NewRepository()

	userObj := repository.GetUserByEmail(loginDto.Email)
	userObj = user.NewUserService().Login(userObj, *loginDto)

	return c.Status(fiber.StatusOK).JSON(
		pkg.NewResponse(
			true,
			constants.LOGIN_SUCCESS,
			dto.RegisterResponse{
				Token: jwt.New().SetIsVerified(true).SetUserId(userObj.ID).CreateToken().GetToken(),
			},
		),
	)
}

func VerifyEmailHandler(c *fiber.Ctx) error {

	verifyDto := new(dto.VerifyEmailRequest)

	pkg.JSONEncoder(c.Body(), &verifyDto)
	verifyDto.Validate()

	userId := c.Locals("userId").(primitive.ObjectID)

	repository := user.NewRepository()
	userObj := repository.GetUserById(userId)

	userObj = user.NewUserService().VerifyEmail(userObj, *verifyDto)

	repository.Upsert(userObj)

	return c.Status(fiber.StatusOK).JSON(
		pkg.NewResponse(
			true,
			constants.EMAIL_VERIFICATION_SUCCESS,
			dto.RegisterResponse{
				Token: jwt.New().SetIsVerified(true).SetUserId(userObj.ID).CreateToken().GetToken(),
			},
		),
	)
}

func ForgotPasswordHandler(c *fiber.Ctx) error {

	forgotPasswordDto := new(dto.ForgotPasswordRequest)

	pkg.JSONEncoder(c.Body(), &forgotPasswordDto)
	forgotPasswordDto.Validate()

	repository := user.NewRepository()
	userObj := repository.GetUserByEmail(forgotPasswordDto.Email)
	userObj = user.NewUserService().ResendEmailVerificationCode(userObj)
	repository.Upsert(userObj)

	return c.Status(fiber.StatusOK).JSON(
		pkg.NewResponse(
			true,
			constants.FORGO_PASSWORD_EMAIL_SENT,
			dto.RegisterResponse{
				Token: jwt.New().SetIsVerified(false).SetUserId(userObj.ID).CreateToken().GetToken(),
			}),
	)
}

func VerifyPhoneNumberHandler(c *fiber.Ctx) error {

	verifyDto := new(dto.VerifyPhoneNumberRequest)

	pkg.JSONEncoder(c.Body(), &verifyDto)
	verifyDto.Validate()

	userId := c.Locals("userId").(primitive.ObjectID)

	repository := user.NewRepository()
	userObj := repository.GetUserById(userId)

	userObj = user.NewUserService().VerifyPhoneNumber(userObj, *verifyDto)

	repository.Upsert(userObj)

	return c.Status(fiber.StatusOK).JSON(
		pkg.NewResponse(
			true,
			constants.PHONE_VERIFICATION_SUCCESS,
			dto.RegisterResponse{
				Token: jwt.New().SetIsVerified(true).SetUserId(userObj.ID).CreateToken().GetToken(),
			},
		),
	)
}

func ResendPhoneVerificationCodeHandler(c *fiber.Ctx) error {

	userId := c.Locals("userId").(primitive.ObjectID)

	repository := user.NewRepository()
	userObj := repository.GetUserById(userId)
	userObj = user.NewUserService().ResendPhoneVerificationCode(userObj)
	repository.Upsert(userObj)

	return c.Status(fiber.StatusOK).JSON(
		pkg.NewResponse(true, constants.VERIFICATION_CODE_RESENT, nil),
	)
}

func ResendEmailVerificationCodeHandler(c *fiber.Ctx) error {

	userId := c.Locals("userId").(primitive.ObjectID)

	repository := user.NewRepository()
	userObj := repository.GetUserById(userId)
	userObj = user.NewUserService().ResendEmailVerificationCode(userObj)
	repository.Upsert(userObj)

	return c.Status(fiber.StatusOK).JSON(
		pkg.NewResponse(true, constants.VERIFICATION_CODE_RESENT, nil),
	)
}

func ChangePasswordHandler(c *fiber.Ctx) error {

	changePasswordDto := new(dto.ChangePasswordRequest)

	pkg.JSONEncoder(c.Body(), &changePasswordDto)
	changePasswordDto.Validate()

	userId := c.Locals("userId").(primitive.ObjectID)

	repository := user.NewRepository()

	userObj := repository.GetUserById(userId)
	userObj = user.NewUserService().ChangePassword(userObj, *changePasswordDto)
	repository.Upsert(userObj)

	return c.Status(fiber.StatusOK).JSON(
		pkg.NewResponse(
			true,
			constants.PASSWORD_CHANGE_SUCCESS,
			nil,
		),
	)
}
