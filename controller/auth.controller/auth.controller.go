package authcontroller

import (
	"github.com/gofiber/fiber/v2"
)

func SetupAuthController(app *fiber.App) {
	app.Post("/auth/register", SignUp)
}
