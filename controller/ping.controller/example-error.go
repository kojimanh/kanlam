package pingcontroller

import (
	"github.com/gofiber/fiber/v2"
	lib "kanlam/lib/server"
	errorexampleusecase "kanlam/use-case/ping.use-case/error-example.use-case"
)

func exampleError(c *fiber.Ctx) error {
	res, err := errorexampleusecase.ReturnError()

	return lib.BuildResponse(c, res, err)
}
