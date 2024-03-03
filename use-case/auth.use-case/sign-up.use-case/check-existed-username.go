package signupusecase

import (
	"fmt"
	userstorage "kanlam/storage/maindb/user.storage"
)

func (r *SignUpUseCase) checkExistedUsername(username string) error {
	userStorage := new(userstorage.StructUserStorage)

	user, err := userStorage.GetByUsername(username)

	fmt.Println(user.Id)

	return err
}
