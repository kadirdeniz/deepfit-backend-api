package constants

import "os"

var (
	API_PREFIX = "/api/" + os.Getenv("API_VERSION")
)

const (
	AUTH                = "auth/"
	REGISTER            = "/register"
	LOGIN               = "/login"
	VERIFY_PHONE_NUMBER = "/verify-phone-number"

	USER              = "/user"
	INTERESTS         = "/interests"
	PROFILE_PHOTO     = "/profile-photo"
	COVER_PHOTO       = "/cover-photo"
	VERIFICATION_CODE = "/verification-code"
	PASSWORD          = "/password"

	MEASUREMENT          = "/measurement"
	PARAM_MEASUREMENT_ID = "/:measurement_id"
	IMAGE                = "/image"
	PARAM_IMAGE_NAME     = "/:image_name"
	IS_PUBLIC            = "/is-public"

	EMPTY = "/"
)
