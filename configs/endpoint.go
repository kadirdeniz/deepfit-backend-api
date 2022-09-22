package configs

import "os"

var (
	API_VERSION_BASE_PATH = "/api/" + os.Getenv("API_VERSION")
)

const (
	REGISTER = "/register"
)
