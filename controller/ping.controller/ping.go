package pingcontroller

import (
	"github.com/gofiber/fiber/v2"
	lib "kanlam/lib/server"
	pingusecase "kanlam/use-case/ping.use-case/ping.use-case"
)

func ping(c *fiber.Ctx) error {
	ping := pingusecase.PingUseCase{}

	greeting, err := ping.Ping()

	return lib.BuildResponse(c, greeting, err)
}
