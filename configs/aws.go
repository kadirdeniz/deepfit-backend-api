package configs

import "os"

var (
	S3_BUCKET_BASE_URL = os.Getenv("S3_BUCKET_BASE_URL")
)
