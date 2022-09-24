package user

import (
	"deepfit/constants"
	"deepfit/pkg/dto"
	"errors"
)

type UserService struct {
	repository IUserRepository
}

type IUserService interface {
	Register(dto dto.RegisterRequest) (*User, error)
}

func NewUserService() IUserService {
	return &UserService{
		repository: NewRepository(),
	}
}

func (userService *UserService) Register(dto dto.RegisterRequest) (*User, error) {

	user := NewUser(dto.Name, dto.Surname, dto.Nickname, dto.Phone, dto.Password)

	if !userService.repository.IsPhoneNumberExists(user) {
		return nil, errors.New(constants.PHONE_NUMBER_ALREADY_EXISTS)
	}

	if !userService.repository.IsNicknameExists(user) {
		return nil, errors.New(constants.NICKNAME_ALREADY_EXISTS)
	}

	if upsertError := userService.repository.Upsert(user); upsertError != nil {
		return nil, upsertError
	}

	return user, nil
}
