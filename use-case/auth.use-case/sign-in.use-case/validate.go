package signinusecase

import "strings"

func (r *SignInUseCase) validateInput() {
	username := strings.Trim(r.Username, " ")
	username = strings.ToLower(username)

	r.Username = username
}
