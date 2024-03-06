package authcontroller

import (
	"github.com/gofiber/fiber/v2"
	lib "kanlam/lib/server"
	signinusecase "kanlam/use-case/auth.use-case/sign-in.use-case"
)

func signIn(c *fiber.Ctx) error {
	signInUseCase := new(signinusecase.SignInUseCase)

	if parseErr := lib.ParseBody(c, signInUseCase); parseErr != nil {
		return lib.BuildResponse(c, nil, parseErr)
	}

	res, err := signInUseCase.SignIn()

	return lib.BuildResponse(c, res, err)
}
