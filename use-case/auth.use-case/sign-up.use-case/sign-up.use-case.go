package signupusecase

import authdto "kanlam/entity/dto/auth.dto"

type SignUpUseCase struct {
	authdto.RegisterInputDto
}

func (r SignUpUseCase) SignUp() (string, error) {
	return r.Username, nil
}
