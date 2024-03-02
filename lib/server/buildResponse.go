package lib

import (
	"github.com/gofiber/fiber/v2"
	entity "kanlam/entity/dto"
)

func BuildResponse(c *fiber.Ctx, data interface{}, e error) error {
	var errorMessage = ""

	if e != nil {
		errorMessage = e.Error()
	}

	response := entity.ResponseDto{
		Code:  "OK",
		Data:  data,
		Error: errorMessage,
	}

	return c.JSON(response)
}
