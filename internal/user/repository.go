package user

import (
	"deepfit/constants"
	"deepfit/pkg"
	"deepfit/tools/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct{}

type IUserRepository interface {
	Upsert(user *User)
	GetUserById(userId primitive.ObjectID) *User
	GetUserByPhone(phone string) *User
	GetUserByEmail(email string) *User
	IsPhoneNumberExists(phone string) bool
	IsNicknameExists(nickname string) bool
	IsEmailExists(email string) bool
}

func NewRepository() IUserRepository {
	return new(UserRepository)
}

func (repository *UserRepository) Upsert(user *User) {
	opts := options.Update().SetUpsert(true)

	if _, updateErr := mongodb.UserCollection.UpdateByID(mongodb.CTX, user.ID, bson.M{"$set": user}, opts); updateErr != nil {
		panic(
			pkg.NewError(constants.StatusInternalServerError, constants.DATABASE_OPERATION_ERROR, updateErr),
		)
	}
}

func (repository *UserRepository) GetUserById(userId primitive.ObjectID) *User {

	var userObj User

	response := mongodb.UserCollection.FindOne(mongodb.CTX, bson.M{"_id": userId})
	if response.Err() != nil {
		if response.Err() == mongo.ErrNoDocuments {
			panic(pkg.NewError(constants.StatusNotFound, constants.USER_NOT_FOUND, nil))
		}
		panic(pkg.NewError(constants.StatusInternalServerError, constants.DATABASE_OPERATION_ERROR, response.Err()))
	}

	if err := response.Decode(&userObj); err != nil {
		panic(pkg.NewError(constants.StatusInternalServerError, constants.DATABASE_OPERATION_ERROR, err))
	}

	return &userObj
}

func (repository *UserRepository) GetUserByPhone(phone string) *User {
	var userObj User

	response := mongodb.UserCollection.FindOne(mongodb.CTX, bson.D{{"phone.phone", phone}, {"email.is_verified", true}})
	if response.Err() != nil {
		if response.Err() == mongo.ErrNoDocuments {
			panic(pkg.NewError(constants.StatusNotFound, constants.PHONE_NUMBER_NOT_FOUND, nil))
		}
		panic(pkg.NewError(constants.StatusInternalServerError, constants.DATABASE_OPERATION_ERROR, response.Err()))
	}

	if err := response.Decode(&userObj); err != nil {
		panic(pkg.NewError(constants.StatusInternalServerError, constants.DATABASE_OPERATION_ERROR, err))
	}

	return &userObj
}

func (repository *UserRepository) GetUserByEmail(email string) *User {
	var userObj User

	response := mongodb.UserCollection.FindOne(mongodb.CTX, bson.D{{"email.email", email}, {"email.is_verified", true}})
	if response.Err() != nil {
		if response.Err() == mongo.ErrNoDocuments {
			panic(pkg.NewError(constants.StatusNotFound, constants.EMAIL_NOT_FOUND, nil))
		}
		panic(pkg.NewError(constants.StatusInternalServerError, constants.DATABASE_OPERATION_ERROR, response.Err()))
	}

	if err := response.Decode(&userObj); err != nil {
		panic(pkg.NewError(constants.StatusInternalServerError, constants.DATABASE_OPERATION_ERROR, err))
	}

	return &userObj
}

func (repository *UserRepository) IsPhoneNumberExists(phone string) bool {
	filter := bson.D{{"phone.phone", phone}, {"email.is_verified", true}}

	count, err := mongodb.UserCollection.CountDocuments(mongodb.CTX, filter)
	if err != nil {
		panic(pkg.NewError(constants.StatusInternalServerError, constants.DATABASE_OPERATION_ERROR, err))
	}

	return count > 0
}

func (repository *UserRepository) IsNicknameExists(nickname string) bool {
	filter := bson.D{{"nickname", nickname}, {"email.is_verified", true}}

	count, err := mongodb.UserCollection.CountDocuments(mongodb.CTX, filter)
	if err != nil {
		panic(pkg.NewError(constants.StatusInternalServerError, constants.DATABASE_OPERATION_ERROR, err))
	}

	return count > 0
}

func (repository *UserRepository) IsEmailExists(email string) bool {
	filter := bson.D{{"email.email", email}, {"email.is_verified", true}}

	count, err := mongodb.UserCollection.CountDocuments(mongodb.CTX, filter)
	if err != nil {
		panic(pkg.NewError(constants.StatusInternalServerError, constants.DATABASE_OPERATION_ERROR, err))
	}

	return count > 0
}
