package main

import "github.com/gofiber/fiber/v2"

func Router(app *fiber.App) {
	r := app.Group("/")

	r.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})
}
