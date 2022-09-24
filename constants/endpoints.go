package constants

import "os"

var (
	API_PREFIX = "/api/" + os.Getenv("API_VERSION")
)

const (
	REGISTER            = "/register"
	VERIFY_PHONE_NUMBER = "/verify-phone-number"
	INTERESTS           = "/interests"
)
