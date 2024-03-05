package taskcontroller

import (
	"github.com/gofiber/fiber/v2"
	lib "kanlam/lib/server"
)

func createTask(c *fiber.Ctx) error {
	return lib.BuildResponse(c, nil, nil)
}
