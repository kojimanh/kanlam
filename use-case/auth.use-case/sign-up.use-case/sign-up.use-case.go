package signupusecase

import (
	authdto "kanlam/entity/dto/auth.dto"
)

type SignUpUseCase struct {
	authdto.RegisterInputDto
	Id             string
	HashedPassword string
}

func (r *SignUpUseCase) SignUp() (*authdto.RegisterOutputDto, error) {
	if validateErr := r.validateInput(); validateErr != nil {
		return nil, validateErr
	}

	if existedErr := r.checkExistedUsername(r.Username); existedErr != nil {
		return nil, existedErr
	}

	if insertErr := r.createUser(); insertErr != nil {
		return nil, insertErr
	}

	result := authdto.RegisterOutputDto{
		Id:       r.Id,
		Username: r.Username,
	}

	return &result, nil
}
