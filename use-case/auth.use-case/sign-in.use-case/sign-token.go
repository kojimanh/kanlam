package signinusecase

import authdto "kanlam/entity/dto/auth.dto"

func (r *SignInUseCase) signToken() authdto.Token {
	return authdto.Token{
		AccessToken:  "access",
		RefreshToken: "refresh",
	}
}
