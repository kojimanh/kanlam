package signupusecase

import (
	"fmt"
	userstorage "kanlam/storage/maindb/user.storage"
)

func (r *SignUpUseCase) checkExistedUsername(username string) error {
	userStorage := new(userstorage.StructUserStorage)

	user, error := userStorage.GetByUsername(username)

	fmt.Println(user.Id)

	return error
}
