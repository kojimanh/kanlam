package signupusecase

import (
	"strings"
)

func (r *SignUpUseCase) validateInput() error {
	username := strings.Trim(r.Username, " ")
	username = strings.ToLower(username)

	r.Username = username

	return nil
}
