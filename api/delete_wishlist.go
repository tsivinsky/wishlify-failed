package main

import "github.com/gofiber/fiber/v2"

func HandleDeleteWishlist(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)

	wishlistId, err := c.ParamsInt("id")
	if err != nil {
		return MakeApiError(400, err.Error())
	}

	var wishlist Wishlist
	if tx := Db.Model(&Wishlist{}).Where("id = ? AND user_id = ?", wishlistId, userId).Preload("User").Preload("Products").First(&wishlist).Delete(&wishlist); tx.Error != nil {
		return MakeApiError(500, tx.Error.Error())
	}

	return c.JSON(wishlist)
}
