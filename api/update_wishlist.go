package main

import "github.com/gofiber/fiber/v2"

type UpdateWishlistBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func HandleUpdateWishlist(c *fiber.Ctx) error {
	wishlistId, err := c.ParamsInt("id")
	if err != nil {
		return MakeApiError(500, err.Error())
	}

	var body UpdateWishlistBody

	if err := c.BodyParser(&body); err != nil {
		return MakeApiError(400, err.Error())
	}

	var wishlist Wishlist
	if tx := Db.Preload("User").Preload("Products").First(&wishlist, wishlistId); tx.Error != nil {
		return MakeApiError(404, err.Error())
	}

	if body.Name != "" {
		newDisplayName := GenerateWishlistDisplayName(body.Name)
		if tx := Db.Where("display_name = ?", newDisplayName).First(&Wishlist{}); tx.RowsAffected > 0 {
			return MakeApiValidationError("name", "Name already taken")
		}

		wishlist.Name = body.Name
		wishlist.DisplayName = newDisplayName
	}

	if body.Description != "" {
		wishlist.Description = body.Description
	}

	if tx := Db.Save(&wishlist); tx.Error != nil {
		return MakeApiError(500, err.Error())
	}

	return c.JSON(wishlist)
}
