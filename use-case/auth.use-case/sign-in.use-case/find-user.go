package signinusecase

import (
	"errors"
	userstorage "kanlam/storage/maindb/user.storage"
)

func (r *SignInUseCase) findUser() (*userstorage.User, error) {
	user, err := userstorage.UserStorage.GetByUsername(r.Username)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("User not existed")
	}

	return user, nil
}
