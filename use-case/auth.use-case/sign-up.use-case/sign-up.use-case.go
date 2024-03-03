package signupusecase

import authdto "kanlam/entity/dto/auth.dto"

type SignUpUseCase struct {
	authdto.RegisterInputDto
}

func (r *SignUpUseCase) SignUp() (*authdto.RegisterOutputDto, error) {
	if validateErr := r.validateInput(); validateErr != nil {
		return nil, validateErr
	}

	if existedErr := r.checkExistedUsername(r.Username); existedErr != nil {
		return nil, existedErr
	}

	result := authdto.RegisterOutputDto{
		Username: r.Username,
	}

	return &result, nil
}
