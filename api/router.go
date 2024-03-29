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
	r.Get("/api/wishlists/:id", HandleRequireAuth, HandleGetWishlistById)
	r.Patch("/api/wishlists/:id", HandleRequireAuth, HandleUpdateWishlist)
	r.Delete("/api/wishlists/:id", HandleRequireAuth, HandleDeleteWishlist)

	r.Post("/api/wishlists/:id/products", HandleRequireAuth, HandleAddProduct)
	r.Patch("/api/wishlists/:wishlistId/products/:productId", HandleRequireAuth, HandleUpdateProduct)
	r.Delete("/api/wishlists/:wishlistId/products/:productId", HandleRequireAuth, HandleRemoveProduct)

	r.Post("/api/images", HandleUploadImages)

	r.Get("/api/users/:username", HandleGetUserByUsername)
	r.Get("/api/users/:username/wishlists", HandleGetUserWishlistsByUsername)
}
