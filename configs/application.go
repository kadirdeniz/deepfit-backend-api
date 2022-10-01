package configs

import (
	"os"
	"strconv"
)

var (
	PORT                           = os.Getenv("PORT")
	API_VERSION                    = os.Getenv("API_VERSION")
	HOST                           = os.Getenv("HOST")
	MAX_MEASUREMENT_IMAGE_COUNT, _ = strconv.Atoi(os.Getenv("MAX_MEASUREMENT_IMAGE_COUNT"))
)
