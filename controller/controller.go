package controller

import (
	"github.com/gofiber/fiber/v2"
	authcontroller "kanlam/controller/auth.controller"
	pingcontroller "kanlam/controller/ping.controller"
)

func SetupController(app *fiber.App) {
	pingcontroller.SetupPingController(app)
	authcontroller.SetupAuthController(app)
}
