package signinusecase

import authdto "kanlam/entity/dto/auth.dto"

type SignInUseCase struct {
	authdto.SignInInputDto
}

func (r *SignInUseCase) SignIn() (*authdto.SignInOutputDto, error) {
	result := new(authdto.SignInOutputDto)

	return result, nil
}
