package constants

import "os"

var (
	API_PREFIX = "/api/" + os.Getenv("API_VERSION")
)

const (
	REGISTER            = "/register"
	LOGIN               = "/login"
	VERIFY_PHONE_NUMBER = "/verify-phone-number"
	INTERESTS           = "/interests"
	PROFILE_PHOTO       = "/profile-photo"
	COVER_PHOTO         = "/cover-photo"
	VERIFICATION_CODE   = "/verification-code"
	PASSWORD            = "/password"

	MEASUREMENT    = "/measurement"
	MEASUREMENT_ID = "/:measurement_id"
	IMAGE          = "/image"
	IMAGE_ID       = "/:image_id"
	IS_PUBLIC      = "/is-public"

	EMPTY = "/"
)
