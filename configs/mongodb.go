package configs

import (
	"fmt"
	"os"
)

var (
	MONGO_INITDB_ROOT_USERNAME = os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	MONGO_INITDB_ROOT_PASSWORD = os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	MONGO_INITDB_HOST          = os.Getenv("MONGO_INITDB_HOST")
	MONGO_INITDB_PORT          = os.Getenv("MONGO_INITDB_PORT")
	MONGO_INITDB_DATABASE      = os.Getenv("MONGO_INITDB_DATABASE")

	MongoDBConnectionURI = fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=admin",
		MONGO_INITDB_ROOT_USERNAME,
		MONGO_INITDB_ROOT_PASSWORD,
		MONGO_INITDB_HOST,
		MONGO_INITDB_PORT,
	)
)

const (
	USER_COLLECTION = "users"
)
