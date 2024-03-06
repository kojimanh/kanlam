package signinusecase

import (
	authdto "kanlam/entity/dto/auth.dto"
	userstorage "kanlam/storage/maindb/user.storage"
)

type SignInUseCase struct {
	authdto.SignInInputDto
	user *userstorage.User
}

func (r *SignInUseCase) SignIn() (*authdto.SignInOutputDto, error) {
	r.validateInput()

	user, findErr := r.findUser()
	if findErr != nil {
		return nil, findErr
	}

	r.user = user

	passwordError := r.checkUserPassword()
	if passwordError != nil {
		return nil, passwordError
	}

	token := r.signToken()

	result := new(authdto.SignInOutputDto)

	result.Username = r.user.Username
	result.Id = r.user.Id
	result.RefreshToken = token.RefreshToken
	result.AccessToken = token.AccessToken

	return result, nil
}
