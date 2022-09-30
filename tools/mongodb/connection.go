package mongodb

import (
	"context"
	"deepfit/configs"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var CTX = context.TODO()

var MongoClient *mongo.Client
var Database = MongoClient.Database(configs.MONGO_INITDB_DATABASE)
var UserCollection = Database.Collection(configs.USER_COLLECTION)

func CreateConnection() {
	client, err := mongo.Connect(CTX, options.Client().ApplyURI(configs.MongoDBConnectionURI))
	if err != nil {
		panic(err)
	}

	ping(client)

	MongoClient = client

	fmt.Println("Connected to MongoDB!")
}

func GetCollection(collectionName string) *mongo.Collection {
	return MongoClient.Database(configs.MONGO_INITDB_DATABASE).Collection(collectionName)
}

func DisconnectMongo(client *mongo.Client) {
	if err := client.Disconnect(CTX); err != nil {
		panic(err)
	}
}

func DropCollection(collectionName string) {
	collection := GetCollection(collectionName)

	if err := collection.Drop(CTX); err != nil {
		panic(err)
	}
}

func ping(client *mongo.Client) {
	if err := client.Ping(CTX, readpref.Primary()); err != nil {
		panic(err)
	}
}
