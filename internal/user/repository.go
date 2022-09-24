package user

import (
	"deepfit/configs"
	"deepfit/constants"
	"deepfit/tools/mongodb"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type UserRepository struct{}

type IUserRepository interface {
	Upsert(user *User) error
	GetUserById(userId primitive.ObjectID) (*User, error)
	IsPhoneNumberExists(user *User) bool
	IsNicknameExists(user *User) bool
	IsEmailExists(user *User) bool
}

func NewRepository() IUserRepository {
	return new(UserRepository)
}

func (repository *UserRepository) Upsert(user *User) error {
	userCollection := mongodb.GetCollection(configs.USER_COLLECTION)
	opts := options.Update().SetUpsert(true)

	if _, updateErr := userCollection.UpdateByID(mongodb.CTX, user.ID, bson.M{"$set": user}, opts); updateErr != nil {
		log.Fatal(updateErr)
		return errors.New(constants.DATABASE_OPERATION_ERROR)
	}

	return nil
}

func (repository *UserRepository) GetUserById(userId primitive.ObjectID) (*User, error) {

	var userObj User
	userCollection := mongodb.GetCollection(configs.USER_COLLECTION)

	response := userCollection.FindOne(mongodb.CTX, bson.M{"_id": userId})

	if response.Err() != nil {
		if response.Err() == mongo.ErrNoDocuments {
			return nil, errors.New(constants.USER_NOT_FOUND)
		}
		log.Fatal(response.Err())
		return nil, errors.New(constants.DATABASE_OPERATION_ERROR)
	}

	if err := response.Decode(&userObj); err != nil {
		log.Fatal(err)
		return nil, errors.New(constants.DATABASE_OPERATION_ERROR)
	}

	return &userObj, nil
}

func (repository *UserRepository) IsPhoneNumberExists(user *User) bool {
	filter := bson.M{"user.phone.phone": user.Nickname}

	userCollection := mongodb.GetCollection(configs.USER_COLLECTION)

	count, err := userCollection.CountDocuments(mongodb.CTX, filter)
	if err != nil {
		log.Fatal(err)
		panic(errors.New(constants.DATABASE_OPERATION_ERROR))
	}

	return count > 0
}

func (repository *UserRepository) IsNicknameExists(user *User) bool {
	filter := bson.M{"user.nickname": user.Nickname}

	userCollection := mongodb.GetCollection(configs.USER_COLLECTION)

	count, err := userCollection.CountDocuments(mongodb.CTX, filter)
	if err != nil {
		log.Fatal(err)
		panic(errors.New(constants.DATABASE_OPERATION_ERROR))
	}

	return count > 0
}

func (repository *UserRepository) IsEmailExists(user *User) bool {
	filter := bson.M{"user.email.email": user.Nickname}

	userCollection := mongodb.GetCollection(configs.USER_COLLECTION)

	count, err := userCollection.CountDocuments(mongodb.CTX, filter)
	if err != nil {
		log.Fatal(err)
		panic(errors.New(constants.DATABASE_OPERATION_ERROR))
	}

	return count > 0
}
