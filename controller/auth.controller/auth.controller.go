package authcontroller

import (
	"github.com/gofiber/fiber/v2"
)

func SetupAuthController(app *fiber.App) {
	app.Post("/auth/sign-up", signUp)

	app.Post("/auth/sign-in", signIn)
}
