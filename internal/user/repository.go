package user

type UserRepository struct{}

type IUserRepository interface {
}

func NewRepository() IUserRepository {
	return new(UserRepository)
}

func (repository *UserRepository) Upsert(user *User) error {
	return nil
}

func (repository *UserRepository) GetUserById(user *User) error {
	return nil
}

func (repository *UserRepository) IsPhoneNumberExists(user *User) error {
	return nil
}

func (repository *UserRepository) IsNicknameExists(user *User) error {
	return nil
}

func (repository *UserRepository) IsEmailExists(user *User) error {
	return nil
}
