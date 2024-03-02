package server

import (
	"github.com/gofiber/fiber/v2"
	"kanlam/controller"
	"kanlam/lib/config"
)

func StartServer() {
	app := fiber.New()

	serverConfig := config.GetServerConfig()
	serverHost := serverConfig.Server

	controller.SetupController(app)

	app.Listen(serverHost)
}
