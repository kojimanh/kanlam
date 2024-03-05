package pingcontroller

import (
	"github.com/gofiber/fiber/v2"
)

func SetupPingController(app *fiber.App) {
	app.Get("/", ping)

	app.Get("/err", exampleError)
}
