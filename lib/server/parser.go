package lib

import (
	"github.com/gofiber/fiber/v2"
	validatelib "kanlam/lib/validate"
)

func ParseBody(c *fiber.Ctx, useCase interface{}) error {
	if err := c.BodyParser(useCase); err != nil {
		return err
	}

	if validateErr := validatelib.Validator.Struct(useCase); validateErr != nil {
		return validateErr
	}

	return nil
}
