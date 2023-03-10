package main

import "github.com/gofiber/fiber/v2"

func ErrorHandler(c *fiber.Ctx, err error) error {
	if e, ok := err.(*ApiError); ok {
		return c.Status(e.Code).JSON(fiber.Map{
			"message": e.Message,
		})
	}

	return c.Status(500).JSON(fiber.Map{
		"message": err.Error(),
	})
}
