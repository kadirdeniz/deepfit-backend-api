package fiber

import (
	"deepfit/configs"
	"deepfit/constants"
	"deepfit/tools/fiber/handler"
	"deepfit/tools/fiber/middleware"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Router() {

	app := fiber.New(
		fiber.Config{
			ErrorHandler: handler.ErrorHandler,
		},
	)

	app.Use(recover.New())

	api := app.Group(constants.API_PREFIX)

	auth := api.Group(constants.AUTH)
	auth.Post(constants.REGISTER, middleware.TokenCantGo, handler.RegisterHandler)
	auth.Post(constants.LOGIN, middleware.TokenCantGo, handler.LoginHandler)
	auth.Post(constants.FORGOT_PASSWORD, middleware.TokenCantGo, handler.ForgotPasswordHandler)
	auth.Post(constants.VERIFY_EMAIL, middleware.TokenCanGo, middleware.VerifiedCantGo, handler.VerifyEmailHandler)

	user := api.Group(constants.USER)
	user.Put(constants.INTERESTS, middleware.TokenCanGo, middleware.VerifiedCanGo, handler.UpdateInterestsHandler)
	user.Put(constants.PROFILE_PHOTO, middleware.TokenCanGo, middleware.VerifiedCanGo, handler.UpdateProfilePhotoHandler)
	user.Put(constants.COVER_PHOTO, middleware.TokenCanGo, middleware.VerifiedCanGo, handler.UpdateCoverPhotoHandler)
	// user.Get(constants.PHONE_VERIFICATION_CODE, middleware.TokenCanGo, middleware.VerifiedCantGo, handler.ResendPhoneVerificationCodeHandler)
	user.Get(constants.EMAIL_VERIFICATION_CODE, middleware.TokenCanGo, middleware.VerifiedCantGo, handler.ResendEmailVerificationCodeHandler)
	user.Put(constants.PASSWORD, middleware.TokenCanGo, middleware.VerifiedCanGo, handler.ChangePasswordHandler)

	measurement := api.Group(constants.MEASUREMENT)
	measurement.Post(constants.EMPTY, middleware.TokenCanGo, middleware.VerifiedCanGo, handler.CreateMeasurementHandler)
	measurement.Put(constants.PARAM_MEASUREMENT_ID, middleware.TokenCanGo, middleware.VerifiedCanGo, handler.UpdateMeasurementHandler)
	measurement.Delete(constants.PARAM_MEASUREMENT_ID, middleware.TokenCanGo, middleware.VerifiedCanGo, handler.DeleteMeasurementHandler)
	measurement.Post(constants.PARAM_MEASUREMENT_ID+constants.IMAGE, middleware.TokenCanGo, middleware.VerifiedCanGo, handler.AddImageToMeasurementHandler)
	measurement.Delete(constants.PARAM_MEASUREMENT_ID+constants.IMAGE+constants.PARAM_IMAGE_NAME, middleware.TokenCanGo, middleware.VerifiedCanGo, handler.DeleteImageToMeasurementHandler)
	measurement.Put(constants.PARAM_MEASUREMENT_ID+constants.IS_PUBLIC, middleware.TokenCanGo, middleware.VerifiedCanGo, handler.UpdateMeasurementIsPublicHandler)

	log.Fatal(app.Listen(":" + configs.PORT))
}
