package signupusecase

import (
	authdto "kanlam/entity/dto/auth.dto"
)

type SignUpUseCase struct {
	authdto.SignUpInputDto
	Id             string
	HashedPassword string
}

func (r *SignUpUseCase) SignUp() (*authdto.SignUpOutputDto, error) {
	if validateErr := r.validateInput(); validateErr != nil {
		return nil, validateErr
	}

	if existedErr := r.checkExistedUsername(r.Username); existedErr != nil {
		return nil, existedErr
	}

	if insertErr := r.createUser(); insertErr != nil {
		return nil, insertErr
	}

	result := authdto.SignUpOutputDto{
		Id:       r.Id,
		Username: r.Username,
	}

	return &result, nil
}
