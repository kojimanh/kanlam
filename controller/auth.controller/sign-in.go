package authcontroller

import (
	"github.com/gofiber/fiber/v2"
	lib "kanlam/lib/server"
)

func signIn(c *fiber.Ctx) error {
	return lib.BuildResponse(c, "", nil)
}
