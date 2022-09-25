package fiber

import (
	"deepfit/configs"
	"deepfit/constants"
	"deepfit/tools/fiber/handler"
	"deepfit/tools/fiber/middleware"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Router() {
	app := fiber.New()

	api := app.Group(constants.API_PREFIX)

	api.Post(constants.REGISTER, middleware.TokenCantGo, handler.RegisterHandler)
	api.Post(constants.LOGIN, middleware.TokenCantGo, handler.LoginHandler)
	api.Post(constants.VERIFY_PHONE_NUMBER, middleware.TokenCanGo, handler.VerifyPhoneNumberHandler)
	api.Put(constants.INTERESTS, middleware.TokenCantGo, handler.UpdateInterestsHandler)
	api.Put(constants.PROFILE_PHOTO, middleware.TokenCanGo, handler.UpdateProfilePhotoHandler)
	api.Put(constants.COVER_PHOTO, middleware.TokenCanGo, handler.UpdateCoverPhotoHandler)
	api.Put(constants.VERIFICATION_CODE, middleware.TokenCantGo, handler.ResendVerificationCodeHandler)
	api.Put(constants.PASSWORD, middleware.TokenCantGo, handler.ChangePasswordHandler)

	measurement := api.Group(constants.MEASUREMENT)

	measurement.Post(constants.EMPTY, middleware.TokenCanGo, handler.CreateMeasurementHandler)
	measurement.Put(constants.MEASUREMENT_ID, middleware.TokenCanGo, handler.UpdateMeasurementHandler)
	measurement.Delete(constants.MEASUREMENT_ID, middleware.TokenCanGo, handler.DeleteMeasurementHandler)
	measurement.Post(constants.MEASUREMENT_ID+constants.IMAGE, middleware.TokenCanGo, handler.AddImageToMeasurementHandler)
	measurement.Delete(constants.MEASUREMENT_ID+constants.IMAGE+constants.IMAGE_ID, middleware.TokenCanGo, handler.DeleteImageToMeasurementHandler)
	measurement.Put(constants.MEASUREMENT_ID+constants.IS_PUBLIC, middleware.TokenCanGo, handler.UpdateMeasurementIsPublicHandler)

	log.Fatal(app.Listen(":" + configs.PORT))
}
