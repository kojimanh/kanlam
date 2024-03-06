package authcontroller

import (
	"github.com/gofiber/fiber/v2"
	lib "kanlam/lib/server"
	signupusecase "kanlam/use-case/auth.use-case/sign-up.use-case"
)

func signUp(c *fiber.Ctx) error {
	signUpUseCase := new(signupusecase.SignUpUseCase)

	if validateErr := lib.ParseBody(c, signUpUseCase); validateErr != nil {
		return lib.BuildResponse(c, nil, validateErr)
	}

	res, err := signUpUseCase.SignUp()

	return lib.BuildResponse(c, res, err)
}
