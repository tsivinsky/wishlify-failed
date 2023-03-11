package main

import "github.com/gofiber/fiber/v2"

func HandleGetUserByUsername(c *fiber.Ctx) error {
	username := c.Params("username")

	var user *User
	if tx := Db.First(&user, "username = ?", username); tx.Error != nil {
		return MakeApiError(404, tx.Error.Error())
	}

	return c.JSON(user)
}

func HandleGetUserWishlistsByUsername(c *fiber.Ctx) error {
	username := c.Params("username")

	var user *User
	if tx := Db.First(&user, "username = ?", username); tx.Error != nil {
		return MakeApiError(404, tx.Error.Error())
	}

	var wishlists []*Wishlist
	if tx := Db.Preload("User").Preload("Products").Find(&wishlists, "user_id = ?", user.ID); tx.Error != nil {
		return MakeApiError(404, tx.Error.Error())
	}

	return c.JSON(wishlists)
}
