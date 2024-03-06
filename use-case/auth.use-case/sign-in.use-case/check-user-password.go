package signinusecase

import (
	"errors"
	bcryptlib "kanlam/lib/bcrypt"
)

func (r *SignInUseCase) checkUserPassword() error {
	isPasswordCorrect := bcryptlib.BcryptLib.Compare(r.user.Password, r.Password)

	if !isPasswordCorrect {
		return errors.New("Password is not correct")
	}

	return nil
}
