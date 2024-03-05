package signupusecase

import (
	bcryptlib "kanlam/lib/bcrypt"
	uidlib "kanlam/lib/uid"
	userstorage "kanlam/storage/maindb/user.storage"
)

func (r *SignUpUseCase) createUser() error {
	hashedPassword, hashedErr := bcryptlib.BcryptLib.Hash(r.Password)
	if hashedErr != nil {
		return hashedErr
	}

	userId, uidErr := uidlib.UidLib.GenUid()
	if uidErr != nil {
		return uidErr
	}

	r.Id = userId
	r.HashedPassword = hashedPassword

	insertedUser := userstorage.User{
		Id:       userId,
		Username: r.Username,
		Password: hashedPassword,
	}

	if insertedErr := userstorage.UserStorage.CreateUser(insertedUser); insertedErr != nil {
		return insertedErr
	}

	return nil
}
