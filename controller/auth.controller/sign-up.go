package authcontroller

import (
	"github.com/gofiber/fiber/v2"
	lib "kanlam/lib/server"
	validatelib "kanlam/lib/validate"
	signupusecase "kanlam/use-case/auth.use-case/sign-up.use-case"
)

func SignUp(c *fiber.Ctx) error {
	signUp := new(signupusecase.SignUpUseCase)

	if err := c.BodyParser(signUp); err != nil {
		return lib.BuildResponse(c, nil, err)
	}

	if validateErr := validatelib.Validator.Struct(signUp); validateErr != nil {
		return lib.BuildResponse(c, nil, validateErr)
	}

	res, err := signUp.SignUp()

	return lib.BuildResponse(c, res, err)
}
