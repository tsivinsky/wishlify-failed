package main

import "github.com/gofiber/fiber/v2"

func HandleGetUserWishlists(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)

	var wishlists []Wishlist
	if tx := Db.Preload("User").Preload("Products").Find(&wishlists, "user_id = ?", userId); tx.Error != nil {
		return MakeApiError(500, tx.Error.Error())
	}

	return c.JSON(wishlists)
}
