package taskcontroller

import "github.com/gofiber/fiber/v2"

func SetupTaskController(app *fiber.App) {
	app.Post("/task", createTask)
}
