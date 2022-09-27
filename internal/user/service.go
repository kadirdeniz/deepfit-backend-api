package user

import (
	"deepfit/constants"
	"deepfit/pkg"
	"deepfit/pkg/dto"
)

type UserService struct{}

type IUserService interface {
	Register(dto dto.RegisterRequest) *User
	Login(user *User, dto dto.LoginRequest) *User
	VerifyPhoneNumber(user *User, dto dto.VerifyPhoneNumberRequest) *User
	UpdateInterests(user *User, dto dto.UpdateInterestRequest) *User
	UpdateProfilePhoto(user *User, imageName string) *User
	UpdateCoverPhoto(user *User, imageName string) *User
	ResendVerificationCode(user *User) *User
	ChangePassword(user *User, dto dto.ChangePasswordRequest) *User
}

func NewUserService() IUserService {
	return &UserService{}
}

func (userService *UserService) Register(dto dto.RegisterRequest) *User {
	return NewUser(dto.Name, dto.Surname, dto.Nickname, dto.Phone, dto.Password)
}

func (userService *UserService) VerifyPhoneNumber(user *User, dto dto.VerifyPhoneNumberRequest) *User {

	if user.Phone.CheckVerificationCode(dto.VerificationCode) {
		panic(pkg.NewError(constants.StatusBadRequest, constants.INVALID_VERIFICATION_CODE, nil))
	}

	user.Phone.SetVerify()

	return user
}

func (userService *UserService) UpdateInterests(user *User, dto dto.UpdateInterestRequest) *User {

	user.Interests = dto.Interests

	return user
}

func (userService *UserService) UpdateProfilePhoto(user *User, imageName string) *User {
	return user.SetProfilePhoto(imageName)
}

func (userService *UserService) UpdateCoverPhoto(user *User, imageName string) *User {
	return user.SetCoverPhoto(imageName)
}

func (userService *UserService) Login(user *User, dto dto.LoginRequest) *User {

	if !user.ComparePasswords(dto.Password) {
		panic(pkg.NewError(constants.StatusBadRequest, constants.INVALID_PASSWORD, nil))
	}

	return user
}

func (userService *UserService) ResendVerificationCode(user *User) *User {
	user.Phone.SetVerificationCode()
	return user
}

func (userService *UserService) ChangePassword(user *User, dto dto.ChangePasswordRequest) *User {

	if !user.ComparePasswords(dto.OldPassword) {
		panic(pkg.NewError(constants.StatusBadRequest, constants.INVALID_PASSWORD, nil))
	}

	user.HashPassword(dto.NewPassword)

	return user
}
