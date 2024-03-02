package pingcontroller

import (
	"github.com/gofiber/fiber/v2"
	lib "kanlam/lib/server"
	errorexampleusecase "kanlam/use-case/ping.use-case/error-example.use-case"
	pingusecase "kanlam/use-case/ping.use-case/ping.use-case"
)

func SetupPingController(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		ping := pingusecase.PingUseCase{}

		greeting, err := ping.Ping()

		return lib.BuildResponse(c, greeting, err)
	})

	app.Get("/err", func(c *fiber.Ctx) error {
		res, err := errorexampleusecase.ReturnError()

		return lib.BuildResponse(c, res, err)
	})
}
