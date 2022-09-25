package user

import (
	"deepfit/constants"
	"deepfit/internal/user/measurement"
	"deepfit/pkg"
	"deepfit/tools/bcrypt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID        `bson:"_id" json:"_id"`
	Name         string                    `bson:"name" json:"name"`
	Surname      string                    `bson:"surname" json:"surname"`
	Nickname     string                    `bson:"nickname" json:"nickname"`
	Phone        Phone                     `bson:"phone" json:"phone"`
	Email        Email                     `bson:"email,omitempty" json:"email"`
	Password     string                    `bson:"password" json:"password"`
	Measurements []measurement.Measurement `bson:"measurements,omitempty" json:"measurements"`
	Interests    []string                  `bson:"interests,omitempty" json:"interests"`
	ProfilePhoto pkg.Image                 `bson:"profile_photo" json:"profile_photo"`
	CoverPhoto   pkg.Image                 `bson:"cover_photo" json:"cover_photo"`
	Date         pkg.Date                  `bson:"date" json:"date"`
}

func NewUser(name, surname, nickname, phone, password string) *User {
	return new(User).
		SetRandomId().
		SetName(name).
		SetSurname(surname).
		SetNickname(nickname).
		SetPhoneObj(phone).
		HashPassword(password).
		SetProfilePhoto("").
		SetCoverPhoto("")
}

func (user *User) SetId(id primitive.ObjectID) *User {
	user.ID = id
	return user
}

func (user *User) SetRandomId() *User {
	user.ID = primitive.NewObjectID()
	return user
}

func (user *User) SetName(name string) *User {
	user.Name = pkg.UpperCaseFirstLetters(name)
	user.Date.UpdateTime()
	return user
}

func (user *User) SetSurname(surname string) *User {
	user.Surname = pkg.UpperCaseFirstLetters(surname)
	user.Date.UpdateTime()
	return user
}

func (user *User) SetNickname(nickname string) *User {
	user.Nickname = pkg.LowerCaseReplaceEmpty(nickname)
	user.Date.UpdateTime()
	return user
}

func (user *User) SetPhone(phone string) *User {
	user.Phone.Phone = phone
	user.Date.UpdateTime()
	return user
}

func (user *User) SetPhoneObj(phone string) *User {
	user.Phone = NewPhone(phone)
	return user
}

func (user *User) SetEmail(email string) *User {
	user.Email.Email = pkg.LowerCaseReplaceEmpty(email)
	user.Date.UpdateTime()
	return user
}

func (user *User) SetEmailObj(email string) *User {
	user.Email = NewEmail(email)
	return user
}

func (user *User) HashPassword(password string) *User {
	user.Password = bcrypt.HashPassword(password)
	user.Date.UpdateTime()
	return user
}

func (user *User) ComparePasswords(password string) bool {
	return bcrypt.ComparePasswords(user.Password, password)
}

func (user *User) SetProfilePhoto(imageName string) *User {
	user.ProfilePhoto = pkg.NewImage(imageName, constants.PROFILE_PHOTO_PATH)
	user.Date.UpdateTime()
	return user
}

func (user *User) SetCoverPhoto(imageName string) *User {
	user.CoverPhoto = pkg.NewImage(imageName, constants.COVER_PHOTO_PATH)
	user.Date.UpdateTime()
	return user
}

func (user *User) SetMeasurements(measurements []measurement.Measurement) *User {
	user.Measurements = measurements
	return user
}
