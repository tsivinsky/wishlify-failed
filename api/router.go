package main

import "github.com/gofiber/fiber/v2"

func Router(app *fiber.App) {
	r := app.Group("/")

	r.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	r.Post("/api/auth/register", HandleRegisterUser)
	r.Post("/api/auth/login", HandleLoginUser)

	r.Get("/api/user", HandleRequireAuth, HandleGetUser)

	r.Get("/api/wishlists", HandleRequireAuth, HandleGetUserWishlists)
	r.Post("/api/wishlists", HandleRequireAuth, HandleCreateWishlist)
}
