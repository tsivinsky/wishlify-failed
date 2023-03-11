package main

import "github.com/gofiber/fiber/v2"

func HandleGetWishlistById(c *fiber.Ctx) error {
	wishlistId, err := c.ParamsInt("id")
	if err != nil {
		return MakeApiError(400, err.Error())
	}

	var wishlist *Wishlist
	if tx := Db.Preload("User").Preload("Products").First(&wishlist, wishlistId); tx.Error != nil {
		return MakeApiError(404, tx.Error.Error())
	}

	return c.JSON(wishlist)
}
