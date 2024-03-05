package signupusecase

import (
	"errors"
	userstorage "kanlam/storage/maindb/user.storage"
)

func (r *SignUpUseCase) checkExistedUsername(username string) error {
	user, err := userstorage.UserStorage.GetByUsername(username)

	if err == nil && user == nil {
		return nil
	}

	if err != nil {
		return errors.New("Failed while check existed")
	}

	if user.Id != "" {
		return errors.New("User existed")
	}

	return err
}
