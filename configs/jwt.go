package configs

import "os"

var (
	JWT_SECRET = os.Getenv("JWT_SECRET")
	JWT_TTL    = os.Getenv("JWT_TTL")
)
