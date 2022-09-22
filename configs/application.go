package configs

import "os"

var (
	PORT        = os.Getenv("PORT")
	API_VERSION = os.Getenv("API_VERSION")
	HOST        = os.Getenv("HOST")
)
