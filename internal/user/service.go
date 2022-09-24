package user

import (
	"deepfit/constants"
	"deepfit/pkg/dto"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	repository IUserRepository
}

type IUserService interface {
	Register(dto dto.RegisterRequest) (*User, error)
	Login(dto dto.LoginRequest) (*User, error)
	VerifyPhoneNumber(dto dto.VerifyPhoneNumberRequest, userId primitive.ObjectID) (*User, error)
	UpdateInterests(dto dto.UpdateInterestRequest, userId primitive.ObjectID) error
	UpdateProfilePhoto(imageName string, userId primitive.ObjectID) error
	UpdateCoverPhoto(imageName string, userId primitive.ObjectID) error
	ResendVerificationCode(id primitive.ObjectID) (*User, error)
}

func NewUserService() IUserService {
	return &UserService{
		repository: NewRepository(),
	}
}

func (userService *UserService) Register(dto dto.RegisterRequest) (*User, error) {

	if !userService.repository.IsPhoneNumberExists(dto.Phone) {
		return nil, errors.New(constants.PHONE_NUMBER_ALREADY_EXISTS)
	}

	if !userService.repository.IsNicknameExists(dto.Nickname) {
		return nil, errors.New(constants.NICKNAME_ALREADY_EXISTS)
	}

	user := NewUser(dto.Name, dto.Surname, dto.Nickname, dto.Phone, dto.Password)

	if upsertError := userService.repository.Upsert(user); upsertError != nil {
		return nil, upsertError
	}

	return user, nil
}

func (userService *UserService) VerifyPhoneNumber(dto dto.VerifyPhoneNumberRequest, userId primitive.ObjectID) (*User, error) {

	user, getUserErr := userService.repository.GetUserById(userId)
	if getUserErr != nil {
		return nil, getUserErr
	}

	if user.Phone.CheckVerificationCode(dto.VerificationCode) {
		return nil, errors.New(constants.INVALID_VERIFICATION_CODE)
	}

	user.Phone.SetVerify()

	if upsertError := userService.repository.Upsert(user); upsertError != nil {
		return nil, upsertError
	}

	return user, nil
}

func (userService *UserService) UpdateInterests(dto dto.UpdateInterestRequest, userId primitive.ObjectID) error {

	user, getUserErr := userService.repository.GetUserById(userId)
	if getUserErr != nil {
		return getUserErr
	}

	user.Interests = dto.Interests

	if upsertError := userService.repository.Upsert(user); upsertError != nil {
		return upsertError
	}

	return nil
}

func (userService *UserService) UpdateProfilePhoto(imageName string, userId primitive.ObjectID) error {

	user, getUserErr := userService.repository.GetUserById(userId)
	if getUserErr != nil {
		return getUserErr
	}

	user.SetProfilePhoto(imageName)

	if upsertError := userService.repository.Upsert(user); upsertError != nil {
		return upsertError
	}

	return nil
}

func (userService *UserService) UpdateCoverPhoto(imageName string, userId primitive.ObjectID) error {

	user, getUserErr := userService.repository.GetUserById(userId)
	if getUserErr != nil {
		return getUserErr
	}

	user.SetCoverPhoto(imageName)

	if upsertError := userService.repository.Upsert(user); upsertError != nil {
		return upsertError
	}

	return nil
}

func (userService *UserService) Login(dto dto.LoginRequest) (*User, error) {

	user, getUserErr := userService.repository.GetUserByPhone(dto.Phone)
	if getUserErr != nil {
		return nil, getUserErr
	}

	if !user.ComparePasswords(dto.Password) {
		return nil, errors.New(constants.INVALID_PASSWORD)
	}

	return user, nil
}

func (userService *UserService) ResendVerificationCode(id primitive.ObjectID) (*User, error) {

	user, getUserErr := userService.repository.GetUserById(id)
	if getUserErr != nil {
		return nil, getUserErr
	}

	user.Phone.SetVerificationCode()

	if upsertError := userService.repository.Upsert(user); upsertError != nil {
		return nil, upsertError
	}

	return user, nil
}
